package services

import (
	"context"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/pkg/DTO"
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"
)

type IUserService interface {
	RegisterUser(ctx context.Context, request DTO.CreateUserRequest) (int, error)
	LogInUser(ctx context.Context, request DTO.LogInUserRequest) (string, error)
	DeleteUser(ctx context.Context, id int) error
	ChangeUserPassword(ctx context.Context, request DTO.ChangeUserPassword) error
}

type UserService struct {
	repository  repositories.IUserRepository
	tokenForger auth.TokenForger
	hasher      auth.IHasher
	config      *configs.Config
}

func NewUserService(repository repositories.IUserRepository,
	tokenForger auth.TokenForger,
	hasher auth.IHasher,
	config *configs.Config) *UserService {
	return &UserService{repository: repository, tokenForger: tokenForger, hasher: hasher, config: config}
}

func (s *UserService) RegisterUser(ctx context.Context, request DTO.CreateUserRequest) (int, error) {
	hashedPassword, err := s.hasher.HashPassword(request.Password)
	if err != nil {
		return -1, err
	}

	err = s.repository.DoesPhoneExists(ctx, request.PhoneNumber)
	if err != nil {
		return -1, err
	}

	newUser := model.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    hashedPassword,
	}

	id, err := s.repository.AddUser(ctx, newUser)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *UserService) LogInUser(ctx context.Context, request DTO.LogInUserRequest) (string, error) {
	user, err := s.repository.GetUserByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return "", err
	}

	err = s.hasher.CheckPassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}

	token, err := s.tokenForger.Encode(user.Name, user.Email, user.Id, *(s.config))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ChangeUserPassword(ctx context.Context, request DTO.ChangeUserPassword) error {
	user, err := s.repository.GetUserByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return err
	}

	err = s.hasher.CheckPassword(user.Password, request.Password)
	if err != nil {
		return err
	}

	newPassord, err := s.hasher.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	err = s.repository.ChangePassword(ctx, newPassord, user.Id)
	if err != nil {
		return err
	}
	return nil
}
