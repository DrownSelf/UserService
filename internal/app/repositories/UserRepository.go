package repositories

import (
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
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
	connectionString := configs.MakeConnectionString(*config)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *model.User) (int, error) {
	var id int
	db := r.db

	query := `insert into users ("name", "phoneNumber", "email", "password") values($1, $2, $3, $4) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}

func (r *UserRepository) ChangeUser(ctx context.Context, user *model.User) error {
	db := r.db
	query := `update users set("name" = $1, "phoneNumber" = $2, "email" = $3, "password" = $4) where "id" = $5`

	row := db.QueryRowContext(ctx, query, user.Name, user.PhoneNumber, user.Email,
		user.Password, user.Id)

	if row.Err() != nil {
		return row.Err()
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
		return nil, errors.New("unable to read data")
	}
}

func (r *UserRepository) DoesNumberExists(ctx context.Context, phone string) (*model.User, error) {
	db := r.db
	var user model.User

	query := `select * from users where "phoneNumber" = $1`
	row := db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
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
