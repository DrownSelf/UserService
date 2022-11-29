package services

import (
	"context"

	pb "github.com/DrownSelf/OrderService/pkg/grpc"

	"github.com/DrownSelf/UserService/internal/auth"
	configs "github.com/DrownSelf/UserService/internal/config"
	"github.com/DrownSelf/UserService/internal/entities"
	"github.com/DrownSelf/UserService/internal/repositories"
)

type IUserService interface {
	RegisterUser(ctx context.Context, user entities.User) (int, error)
	LogInUser(ctx context.Context, phoneNumber string, password string) (string, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, updatedUser entities.User, oldPhoneNumber string) error
	GetUser(ctx context.Context, phoneNumber string) (entities.User, error)
	LogOutUser(ctx context.Context, token string) error
	UpdateUserRating(ctx context.Context, phoneNumber string, rating float64) error
	MakeOrder(ctx context.Context, gottenUser entities.User, from string, to string, taxiType string) (*pb.UserRideResponse, error)
	RateRideFromUser(ctx context.Context, rating int32, id string) error
}

type UserService struct {
	userRepository  repositories.IUserRepository
	cacheRepository repositories.ICacheRepository
	orderClient     pb.OrderServiceClient
	tokenForger     auth.TokenForger
	hasher          auth.IHasher
	config          *configs.Config
}

func NewUserService(userRepository repositories.IUserRepository, client pb.OrderServiceClient, cacheRepository repositories.ICacheRepository, tokenForger auth.TokenForger, hasher auth.IHasher, config *configs.Config) *UserService {
	return &UserService{userRepository: userRepository, orderClient: client, cacheRepository: cacheRepository, tokenForger: tokenForger, hasher: hasher, config: config}
}

func (s *UserService) RegisterUser(ctx context.Context, user entities.User) (int, error) {
	if err := s.userRepository.DoesPhoneExist(ctx, user.PhoneNumber); err != nil {
		return -1, err
	}

	hashedPassword, err := s.hasher.HashPassword(user.Password)
	if err != nil {
		return -1, err
	}

	newUser := entities.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    hashedPassword,
	}

	id, err := s.userRepository.AddUser(ctx, newUser)
	if err != nil {
		return -1, err
	}

	err = s.userRepository.RelateRating(ctx, id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *UserService) LogInUser(ctx context.Context, phoneNumber string, password string) (string, error) {
	user, err := s.userRepository.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		return "", err
	}

	if err = s.hasher.CheckPassword(user.Password, password); err != nil {
		return "", err
	}

	token, err := s.tokenForger.Encode(
		auth.TokenClaims{
			Id:          user.Id,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
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

func (s *UserService) UpdateUser(ctx context.Context, updatedUser entities.User, oldPhoneNumber string) error {
	user, err := s.userRepository.GetUserByPhone(ctx, oldPhoneNumber)
	if err != nil {
		return err
	}

	updatedUser.Id = user.Id
	if err = s.userRepository.DoesPhoneExist(ctx, updatedUser.PhoneNumber); err != nil {
		return err
	}

	if err = s.userRepository.UpdateUser(ctx, updatedUser); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(ctx context.Context, phoneNumber string) (entities.User, error) {
	user, err := s.userRepository.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		return entities.User{}, err
	}

	userWithRating, err := s.userRepository.GetUserById(ctx, user.Id)
	if err != nil {
		return entities.User{}, err
	}

	return userWithRating, nil
}

func (s *UserService) LogOutUser(ctx context.Context, token string) error {
	if err := s.cacheRepository.PutToken(ctx, token); err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUserRating(ctx context.Context, phoneNumber string, rating float64) error {
	order, err := s.userRepository.GetUserByPhone(ctx, phoneNumber)
	if err != nil {
		return err
	}

	if err = s.userRepository.AppendRating(ctx, order.Id, rating); err != nil {
		return err
	}
	return nil
}

func (s *UserService) MakeOrder(ctx context.Context, gottenUser entities.User, from string, to string, taxiType string) (*pb.UserRideResponse, error) {
	response, err := s.orderClient.MakeOrder(ctx, &pb.OrderTaxiRequest{
		User: &pb.User{
			PhoneNumber: gottenUser.PhoneNumber,
			Email:       gottenUser.Email,
			Name:        gottenUser.Name,
		},
		From:     from,
		To:       to,
		TaxiType: taxiType,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *UserService) RateRideFromUser(ctx context.Context, rating int32, id string) error {
	_, err := s.orderClient.RateRideFromUser(ctx, &pb.RateDriverFromUser{
		Rating:  rating,
		OrderId: id,
	})

	if err != nil {
		return err
	}
	return nil
}
