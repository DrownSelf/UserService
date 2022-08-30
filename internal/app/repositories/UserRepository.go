package repositories

import (
	"context"
	"database/sql"
	"time"

	"InnoTaxi/internal/app/appErrors"
	"InnoTaxi/internal/app/migrations"
	"InnoTaxi/internal/pkg/DTO"
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"

	_ "github.com/lib/pq"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user model.User) (int, error)
	ChangeUser(ctx context.Context, request DTO.ChangeUserRequest, id int) error
	GetUserById(ctx context.Context, id int) (model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (model.User, error)
	DoesPhoneExist(ctx context.Context, phone string) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	DeleteUser(context context.Context, id int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(config *configs.Config) (*UserRepository, error) {
	db, err := sql.Open("postgres", config.PgSource)
	if err != nil {
		return nil, err
	}

	if err = migrations.Up(db); err != nil {
		if err := db.Close(); err != nil {
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
	query := `insert into users ("name", "phoneNumber", "email", "password", "created_at", "is_deleted", "rating") values($1, $2, $3, $4, $5, false, 0) RETURNING id`
	err := r.db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password, createdTime).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, err
}

func (r *UserRepository) ChangeUser(ctx context.Context, request DTO.ChangeUserRequest, id int) error {
	updateTime := time.Now().UTC()
	query := `update users set "name"=$1, "phoneNumber"=$2, "email" = $3, "updated_at"=$4 where "id" = $5`

	err := r.db.QueryRowContext(ctx, query, request.NewName, request.NewPhoneNumber,
		request.NewEmail, updateTime, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query := `select * from users where "id" = $1 and "is_deleted" = false`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			err = appErrors.ErrUserDoesntExist
		}
		return user, err
	}
	return user, nil
}

func (r *UserRepository) DoesPhoneExist(ctx context.Context, phone string) error {
	var id int
	query := `select id from users where "phoneNumber" = $1 and is_deleted = false`
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return appErrors.ErrUserExists
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User

	query := `select * from users where "phoneNumber" = $1 and is_deleted = false`
	row := r.db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			err = appErrors.ErrUserDoesntExist
		}
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	rows, err := r.db.QueryContext(ctx, `select * from users where is_deleted = false`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Rating, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepository) DeleteUser(context context.Context, id int) error {
	_, err := r.db.Exec(`update users set "is_deleted"= true where "id" = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
