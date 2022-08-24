package repositories

import (
	"context"
	"database/sql"
	"time"

	"InnoTaxi/internal/app/errors"
	"InnoTaxi/internal/app/migrations"
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"

	_ "github.com/lib/pq"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user model.User) (int, error)
	ChangePassword(ctx context.Context, newPassword string, id int) error
	GetUserById(ctx context.Context, id int) (model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (model.User, error)
	DoesPhoneExists(ctx context.Context, phone string) error
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

	if err = migrations.Up(db); err != nil {
		if err = db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

func (r *UserRepository) DestroyRepository(ctx context.Context) error {
	quit := make(chan error)
	go func() {
		quit <- r.db.Close()
	}()

	select {
	case err := <-quit:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (r *UserRepository) AddUser(ctx context.Context, user model.User) (int, error) {
	var id int
	createdTime := time.Now().UTC()
	query := `insert into users ("name", "phoneNumber", "email", "password", "created_at") values($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password, createdTime).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, err
}

func (r *UserRepository) ChangePassword(ctx context.Context, newPassword string, id int) error {
	updateTime := time.Now().UTC()
	query := `update users set "password"=$1, "updated_at"=$2 where "id" = $3`

	err := r.db.QueryRowContext(ctx, query, newPassword, updateTime, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query := `select * from users where "id" = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)

	switch err {
	case sql.ErrNoRows:
		return user, errors.ErrUserDoesntExist
	case nil:
		return user, nil
	default:
		return user, err
	}
}

func (r *UserRepository) DoesPhoneExists(ctx context.Context, phone string) error {
	var user model.User

	query := `select * from users where "phoneNumber" = $1`
	row := r.db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return errors.ErrUserExists
	default:
		return err
	}
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User

	query := `select * from users where "phoneNumber" = $1`
	row := r.db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return user, errors.ErrUserDoesntExist
	case nil:
		return user, err
	default:
		return user, err
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	rows, err := r.db.QueryContext(ctx, `select * from users`)

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
	_, err := r.db.Exec(`delete from users where "id" = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
