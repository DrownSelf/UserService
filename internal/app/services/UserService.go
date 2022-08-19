package services

import (
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/errors"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/pkg/DTO"
	"InnoTaxi/internal/pkg/model"
	"context"
)

type IUserService interface {
	RegisterUser(ctx context.Context, request DTO.CreateUserRequest) (int, error)
	LogInUser(ctx context.Context, request DTO.LogInUserRequest) (*DTO.GetUserResponse, error)
}

type UserService struct {
	repository  repositories.IUserRepository
	tokenForger auth.TokenForger
	hasher      auth.IHasher
}

func New(repository repositories.IUserRepository,
	tokenForger auth.TokenForger,
	hasher auth.IHasher) *UserService {
	return &UserService{repository: repository, tokenForger: tokenForger, hasher: hasher}
}

func (s *UserService) RegisterUser(ctx context.Context, request DTO.CreateUserRequest) (int, error) {
	repository := s.repository
	hashedPassword, err := s.hasher.HashPassword(request.Password)
	if err != nil {
		return -1, err
	}

	user, err := repository.DoesNumberExists(ctx, request.PhoneNumber)
	if err != nil {
		return -1, err
	}
	if user != nil {
		return -1, errors.ErrUserExists
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

	return id, nil
}

func (s *UserService) LogInUser(ctx context.Context, request DTO.LogInUserRequest) (*DTO.GetUserResponse, error) {
	repository := s.repository

	user, err := repository.DoesNumberExists(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserDoesntExist
	}

	err = s.hasher.CheckPassword(user.Password, request.Password)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenForger.Encode(user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	return &DTO.GetUserResponse{user.Id, token}, nil
}
