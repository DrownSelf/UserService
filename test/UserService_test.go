package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"InnoTaxi/internal/app/appErrors"
	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/app/services"
	"InnoTaxi/internal/pkg/DTO"
	"InnoTaxi/internal/pkg/configs"
	"InnoTaxi/internal/pkg/model"
	test "InnoTaxi/test/mocks"
)

func TestServiceRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	userRepo := test.NewMockIUserRepository(ctrl)
	cacheRepo := test.NewMockICacheRepository(ctrl)
	hasher := test.NewMockIHasher(ctrl)
	tokenForger := auth.NewJwt("Silent Secret")
	config := configs.Config{Secret: "SilentSecret"}

	type hashBehaviour func(r *test.MockIHasher, password string)
	type checkBehaviour func(r *test.MockIUserRepository, user DTO.User, ctx context.Context)
	type addBehaviour func(r *test.MockIUserRepository, user model.User, ctx context.Context)

	testTable := []struct {
		ctx            context.Context
		incomingData   DTO.User
		newUser        model.User
		expectedIndex  int
		checkBehaviour checkBehaviour
		hashBehaviour  hashBehaviour
		addBehaviour   addBehaviour
	}{
		{
			ctx,
			DTO.User{
				Name:        "Walter White",
				Password:    "8004355",
				Email:       "strigelskiy.petr@gmail.com",
				PhoneNumber: "+375447505544",
			},
			model.User{
				Name:        "Walter White",
				Email:       "strigelskiy.petr@gmail.com",
				PhoneNumber: "+375447505544",
			},
			-1,
			func(r *test.MockIUserRepository, user DTO.User, ctx context.Context) {
				r.EXPECT().DoesPhoneExist(ctx, user.PhoneNumber).Return(appErrors.ErrUserExists)
			},
			func(r *test.MockIHasher, password string) {

			},
			func(r *test.MockIUserRepository, user model.User, ctx context.Context) {

			},
		},
		{
			ctx,
			DTO.User{
				Name:        "Michael Scoffild",
				Password:    "3608216",
				Email:       "goodddman@gmail.com",
				PhoneNumber: "+375296509109",
			},
			model.User{
				Name:        "Michael Scoffild",
				Email:       "goodddman@gmail.com",
				PhoneNumber: "+375296509109",
			},
			2,
			func(r *test.MockIUserRepository, user DTO.User, ctx context.Context) {
				r.EXPECT().DoesPhoneExist(ctx, user.PhoneNumber).Return(nil)
			},
			func(r *test.MockIHasher, password string) {
				r.EXPECT().HashPassword(password).Return("", nil)
			},
			func(r *test.MockIUserRepository, user model.User, ctx context.Context) {
				r.EXPECT().AddUser(ctx, user).Return(2, nil)
			},
		},
		{
			ctx,
			DTO.User{
				Name:        "Drown",
				Password:    "24578643",
				Email:       "aceRainbow@mail.ru",
				PhoneNumber: "+375335478908",
			},
			model.User{
				Name:        "Drown",
				Email:       "aceRainbow@mail.ru",
				PhoneNumber: "+375335478908",
			},
			-1,
			func(r *test.MockIUserRepository, user DTO.User, ctx context.Context) {
				r.EXPECT().DoesPhoneExist(ctx, user.PhoneNumber).Return(nil)
			},
			func(r *test.MockIHasher, password string) {
				r.EXPECT().HashPassword(password).Return("", nil)
			},
			func(r *test.MockIUserRepository, user model.User, ctx context.Context) {
				r.EXPECT().AddUser(ctx, user).Return(-1, pq.ErrChannelNotOpen)
			},
		},
		{
			ctx,
			DTO.User{
				Name:        "Jacobs",
				Password:    "21570243",
				Email:       "lastexample@mail.ru",
				PhoneNumber: "+375336606644",
			},
			model.User{
				Name:        "Jacobs",
				Email:       "lastexample@mail.ru",
				PhoneNumber: "",
			},
			-1,
			func(r *test.MockIUserRepository, user DTO.User, ctx context.Context) {
				r.EXPECT().DoesPhoneExist(ctx, user.PhoneNumber).Return(nil)
			},
			func(r *test.MockIHasher, password string) {
				r.EXPECT().HashPassword(password).Return("", bcrypt.ErrMismatchedHashAndPassword)
			},
			func(r *test.MockIUserRepository, user model.User, ctx context.Context) {

			},
		},
	}
	for _, testcase := range testTable {
		t.Run("test", func(t *testing.T) {

			testcase.addBehaviour(userRepo, testcase.newUser, ctx)
			testcase.hashBehaviour(hasher, testcase.incomingData.Password)
			testcase.checkBehaviour(userRepo, testcase.incomingData, ctx)

			var service services.IUserService = services.NewUserService(userRepo, cacheRepo, tokenForger, hasher, &config)
			index, err := service.RegisterUser(ctx, testcase.incomingData)
			fmt.Println(err)
			assert.Equal(t, testcase.expectedIndex, index)
		})
	}
}

func TestLogInUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	userRepo := test.NewMockIUserRepository(ctrl)
	cacheRepo := test.NewMockICacheRepository(ctrl)
	hasher := test.NewMockIHasher(ctrl)
	tokenForger := test.NewMockTokenForger(ctrl)
	config := configs.Config{Secret: "SilentSecret"}

	type getUserBehaviour func(r *test.MockIUserRepository, ctx context.Context, phoneNumber string)
	type checkPasswordBehaviour func(h *test.MockIHasher, userPassword string, providedPassword string)
	type encodeTokenBehaviour func(t *test.MockTokenForger, claims auth.TokenClaims, config configs.Config)

	testTable := []struct {
		ctx                    context.Context
		incomingData           DTO.LogInUserRequest
		gottenUser             model.User
		expectedToken          string
		getUserBehaviour       getUserBehaviour
		checkPasswordBehaviour checkPasswordBehaviour
		encodeTokenBehaviour   encodeTokenBehaviour
	}{
		{
			ctx,
			DTO.LogInUserRequest{"+375447505544", ""},
			model.User{
				Name:        "Jacobs",
				Password:    "",
				Email:       "example@mail.ru",
				PhoneNumber: "+375447505544",
			},
			"",
			func(r *test.MockIUserRepository, ctx context.Context, phoneNumber string) {
				r.EXPECT().GetUserByPhone(ctx, phoneNumber).Return(model.User{}, appErrors.ErrUserDoesntExist)
			},
			func(h *test.MockIHasher, userPassword string, providedPassword string) {

			},
			func(t *test.MockTokenForger, claims auth.TokenClaims, config configs.Config) {

			},
		},
		{
			ctx,
			DTO.LogInUserRequest{"+375296509109", ""},
			model.User{
				Name:        "Peter",
				Password:    "",
				Email:       "rubicon@gmail.com",
				PhoneNumber: "+375296509109",
			},
			"cool token",
			func(r *test.MockIUserRepository, ctx context.Context, phoneNumber string) {
				r.EXPECT().GetUserByPhone(ctx, phoneNumber).Return(model.User{}, nil)
			},
			func(h *test.MockIHasher, userPassword string, providedPassword string) {
				h.EXPECT().CheckPassword(userPassword, providedPassword).Return(nil)
			},
			func(t *test.MockTokenForger, claims auth.TokenClaims, config configs.Config) {
				t.EXPECT().Encode(claims, config).Return("cool token", nil)
			},
		},
		{
			ctx,
			DTO.LogInUserRequest{"+3733304109", ""},
			model.User{
				Name:        "Drake",
				Password:    "",
				Email:       "changeable@gmail.com",
				PhoneNumber: "+3733304109",
			},
			"",
			func(r *test.MockIUserRepository, ctx context.Context, phoneNumber string) {
				r.EXPECT().GetUserByPhone(ctx, phoneNumber).Return(model.User{}, nil)
			},
			func(h *test.MockIHasher, userPassword string, providedPassword string) {
				h.EXPECT().CheckPassword(userPassword, providedPassword).Return(appErrors.ErrWrongPassword)
			},
			func(t *test.MockTokenForger, claims auth.TokenClaims, config configs.Config) {

			},
		},
		{
			ctx,
			DTO.LogInUserRequest{"+3754404109", ""},
			model.User{
				Name:        "Jackson",
				Password:    "",
				Email:       "summer@gmail.com",
				PhoneNumber: "+3754404109",
			},
			"",
			func(r *test.MockIUserRepository, ctx context.Context, phoneNumber string) {
				r.EXPECT().GetUserByPhone(ctx, phoneNumber).Return(model.User{}, nil)
			},
			func(h *test.MockIHasher, userPassword string, providedPassword string) {
				h.EXPECT().CheckPassword(userPassword, providedPassword).Return(nil)
			},
			func(t *test.MockTokenForger, claims auth.TokenClaims, config configs.Config) {
				t.EXPECT().Encode(claims, config).Return("", jwt.ErrTokenInvalidClaims)
			},
		},
	}
	for _, testcase := range testTable {
		t.Run("test", func(t *testing.T) {
			testcase.getUserBehaviour(userRepo, ctx, testcase.incomingData.PhoneNumber)
			testcase.checkPasswordBehaviour(hasher, testcase.incomingData.Password, testcase.gottenUser.Password)
			testcase.encodeTokenBehaviour(tokenForger, auth.TokenClaims{}, config)

			var service services.IUserService = services.NewUserService(userRepo, cacheRepo, tokenForger, hasher, &config)
			token, err := service.LogInUser(ctx, testcase.incomingData)
			fmt.Println(err)
			assert.Equal(t, testcase.expectedToken, token)
		})
	}
}
