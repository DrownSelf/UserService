package repositories

import (
	"InnoTaxi/internal/pkg/model"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
	DB *sql.DB
}

func (repository *UserRepository) AddUser(ctx context.Context, user *model.User) (int, error) {
	var id int
	db := repository.DB

	query := `insert into users ("name", "phoneNumber", "email", "password") values($1, $2, $3, $4) RETURNING id`
	err := db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}

func (repository *UserRepository) ChangeUser(ctx context.Context, user *model.User) error {
	db := repository.DB
	query := `update users set("name" = $1, "phoneNumber" = $2, "email" = $3, "password" = $4) where "id" = $5`

	row := db.QueryRowContext(ctx, query, user.Name, user.PhoneNumber, user.Email,
		user.Password, user.Id)

	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (repository *UserRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	db := repository.DB
	var user model.User

	query := `select * from users where "id" = $1`
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("There isn't value which you found.")
		return nil, nil
	case nil:
		return &user, nil
	default:
		log.Fatalf("unable to read row. %v", err)
	}
	return nil, err
}

func (repository *UserRepository) DoesNumberExists(ctx context.Context, phone string) (*model.User, error) {
	db := repository.DB
	var user model.User

	query := `select * from users where "phoneNumber" = $1`
	row := db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repository *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	db := repository.DB
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

func (repository *UserRepository) DeleteUser(context context.Context, id int) error {
	db := repository.DB
	res, err := db.Exec(`delete from users where "id" = $1`, id)

	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	fmt.Println("affectedrows: $1", rowsAffected)
	return nil
}
