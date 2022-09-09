package services

import (
	"context"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/repositories"
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/dto"
	"InnoTaxi/internal/pkg/model"
)

type IUserService interface {
	RegisterUser(ctx context.Context, user dto.User) (int, error)
	LogInUser(ctx context.Context, request dto.LogInUserRequest) (string, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, request dto.ChangeUserRequest) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error)
	LogOutUser(ctx context.Context, token string) error
}

type UserService struct {
	userRepository  repositories.IUserRepository
	cacheRepository repositories.ICacheRepository
	tokenForger     auth.TokenForger
	hasher          auth.IHasher
	config          *configs.Config
}

func NewUserService(userRepository repositories.IUserRepository, cacheRepository repositories.ICacheRepository, tokenForger auth.TokenForger, hasher auth.IHasher, config *configs.Config) *UserService {
	return &UserService{userRepository: userRepository, cacheRepository: cacheRepository, tokenForger: tokenForger, hasher: hasher, config: config}
}

func (s *UserService) RegisterUser(ctx context.Context, user dto.User) (int, error) {
	if err := s.userRepository.DoesPhoneExist(ctx, user.PhoneNumber); err != nil {
		return -1, err
	}

	hashedPassword, err := s.hasher.HashPassword(user.Password)
	if err != nil {
		return -1, err
	}

	newUser := model.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    hashedPassword,
	}

	id, err := s.userRepository.AddUser(ctx, newUser)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *UserService) LogInUser(ctx context.Context, request dto.LogInUserRequest) (string, error) {
	user, err := s.userRepository.GetUserByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return "", err
	}

	if err = s.hasher.CheckPassword(user.Password, request.Password); err != nil {
		return "", err
	}

	token, err := s.tokenForger.Encode(
		auth.TokenClaims{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
		*s.config)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, request dto.ChangeUserRequest) error {
	user, err := s.userRepository.GetUserByPhone(ctx, request.PhoneNumber)
	if err != nil {
		return err
	}

	if err = s.userRepository.DoesPhoneExist(ctx, request.NewPhoneNumber); err != nil {
		return err
	}

	if err = s.userRepository.UpdateUser(ctx, request, user.Id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error) {
	user, err := s.userRepository.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) LogOutUser(ctx context.Context, token string) error {
	if err := s.cacheRepository.PutToken(ctx, token); err != nil {
		return err
	}
	return nil
}
