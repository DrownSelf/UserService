package repositories

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/DrownSelf/UserService/internal/appErrors"
	configs "github.com/DrownSelf/UserService/internal/config"
	"github.com/DrownSelf/UserService/internal/entities"
	"github.com/DrownSelf/UserService/internal/migrations"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user entities.User) (int, error)
	UpdateUser(ctx context.Context, user entities.User) error
	GetUserById(ctx context.Context, id int) (entities.User, error)
	GetUserByPhone(ctx context.Context, phone string) (entities.User, error)
	DoesPhoneExist(ctx context.Context, phone string) error
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, id int) error
	RelateRating(ctx context.Context, id int) error
	GetAverageRating(ctx context.Context, id int) (float64, error)
	AppendRating(ctx context.Context, id int, rating float64) error
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

func (r *UserRepository) AddUser(ctx context.Context, user entities.User) (int, error) {
	var id int
	createdTime := time.Now().UTC()
	query := `insert into users ("name", "phoneNumber", "email", "password", "created_at", "is_deleted") values($1, $2, $3, $4, $5, false) RETURNING id`
	err := r.db.QueryRowContext(ctx, query,
		user.Name, user.PhoneNumber, user.Email, user.Password, createdTime).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, user entities.User) error {
	updateTime := time.Now().UTC()
	query := `update users set "name"=$1, "phoneNumber"=$2, "email" = $3, "updated_at"=$4 where "id" = $5`

	err := r.db.QueryRowContext(ctx, query, user.Name, user.PhoneNumber,
		user.Email, updateTime, user.Id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (entities.User, error) {
	var user entities.User

	query := `select users.id, users."phoneNumber", users.email, users.name, ratings.average_rating from users 
                                inner join ratings on users.id = ratings.id
                                where users.is_deleted = false and users.id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.PhoneNumber, &user.Email, &user.Name, &user.Rating)
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
	query := `select id from users where "phoneNumber" = $1 and "is_deleted" = false`
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return appErrors.ErrUserExists
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (entities.User, error) {
	var user entities.User

	query := `select * from users where "phoneNumber" = $1 and "is_deleted" = false`
	row := r.db.QueryRowContext(ctx, query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted)
	if err != nil {
		if err == sql.ErrNoRows {
			err = appErrors.ErrUserDoesntExist
		}
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	rows, err := r.db.QueryContext(ctx, `select * from users where is_deleted = false`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsDeleted)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.db.Exec(`update users set "is_deleted"= true where "id" = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) RelateRating(ctx context.Context, id int) error {
	_, err := r.db.Exec(`insert into ratings("id", "average_rating") values($1, 0)`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AppendRating(ctx context.Context, id int, rating float64) error {
	_, err := r.db.Exec(`call addrating($1, $2)`, id, rating)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAverageRating(ctx context.Context, id int) (float64, error) {
	var averageRating float64
	query := `select average_rating from ratings where id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&averageRating)

	if err != nil {
		return -1, err
	}
	return averageRating, nil
}
