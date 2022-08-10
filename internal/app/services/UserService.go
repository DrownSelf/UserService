package services

import (
	auth2 "InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/repositories"
	DTO2 "InnoTaxi/internal/pkg/DTO"
	"InnoTaxi/internal/pkg/model"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type IUserService interface {
	RegisterUser(ctx context.Context, request DTO2.CreateUserRequest) (int, error)
	LogInUser(ctx context.Context, request DTO2.LogInUserRequest) (*DTO2.GetUserResponse, error)
}

type UserService struct {
	repository  repositories.IUserRepository
	tokenForger auth2.TokenForger
	hasher      auth2.IHasher
	validator   *validator.Validate
}

func New(repository repositories.IUserRepository,
	tokenForger auth2.TokenForger,
	hasher auth2.IHasher,
	validate *validator.Validate) *UserService {
	return &UserService{repository: repository, tokenForger: tokenForger, hasher: hasher, validator: validate}
}

func (service *UserService) RegisterUser(ctx context.Context, request DTO2.CreateUserRequest) (int, error) {
	repository := service.repository
	hashedPassword, err := service.hasher.HashPassword(request.Password)
	if err != nil {
		return -1, err
	}

	err = service.validator.Struct(request)
	if err != nil {
		return -1, err
	}

	check, err := repository.DoesNumberExists(ctx, request.PhoneNumber)
	if check == nil {
		return -1, err
	}

	var newUser model.User = model.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    hashedPassword,
	}

	id, err := repository.AddUser(ctx, &newUser)

	if err != nil {
		return -1, err
	}

	return id, err
}

func (service *UserService) LogInUser(ctx context.Context, request DTO2.LogInUserRequest) (*DTO2.GetUserResponse, error) {
	repository := service.repository

	err := service.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	user, err := repository.DoesNumberExists(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	err = service.hasher.CheckPassword(user.Password, request.Password)
	if err != nil {
		return nil, err
	}

	token, err := service.tokenForger.Encode(user.Name, user.Email)
	if err != nil {
		fmt.Println("omg")
		return nil, err
	}

	return &DTO2.GetUserResponse{user.Id, token}, nil
}
