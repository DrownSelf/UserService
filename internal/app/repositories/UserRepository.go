package repositories

import (
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"
	"context"
	"database/sql"
	"time"

	_ "InnoTaxi/internal/app/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user *model.User) (int, error)
	ChangeUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int) (*model.User, error)
	DoesNumberExists(ctx context.Context, phone string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(config *configs.Config) (*UserRepository, error) {
	db, err := sql.Open("postgres", config.PgSource)
	if err != nil {
		return nil, err
	}

	goose.SetDialect("postgres")
	if config.Reset {
		if err = goose.Down(db, "."); err != nil {
			return nil, err
		}
	}

	if err = goose.Up(db, "."); err != nil {
		db.Close()
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

func (r *UserRepository) DestroyRepository() error {
	err := goose.Down(r.db, ".")
	err = r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *model.User) (int, error) {
	var id int
	db := r.db
	createdTime := time.Now().UTC()
	query := `insert into users ("name", "phoneNumber", "email", "password", "created_at") values($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password, createdTime).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, err
}

func (r *UserRepository) ChangeUser(ctx context.Context, user *model.User) error {
	var id int
	db := r.db
	updateTime := time.Now().UTC()
	query := `update users set("name" = $1, "phoneNumber" = $2, "email" = $3, "password" = $4, "updated_at" = $5) where "id" = $5`

	err := db.QueryRowContext(ctx, query, user.Name, user.PhoneNumber, user.Email,
		user.Password, user.Id, updateTime).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	db := r.db
	var user model.User

	query := `select * from users where "id" = $1`
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

func (r *UserRepository) DoesNumberExists(ctx context.Context, phone string) (*model.User, error) {
	db := r.db
	var user model.User

	query := `select * from users where "phoneNumber" = $1`
	row := db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	db := r.db
	var users []model.User
	rows, err := db.QueryContext(ctx, `select * from users`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepository) DeleteUser(context context.Context, id int) error {
	db := r.db
	_, err := db.Exec(`delete from users where "id" = $1`, id)

	if err != nil {
		return err
	}
	return nil
}
