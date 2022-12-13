package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/DrownSelf/UserService/internal/appErrors"
	configs "github.com/DrownSelf/UserService/internal/config"
)

type TokenAuth interface {
	Encode(tokenClaims TokenClaims, config configs.Config) (string, error)
	Decode(cipher string) (TokenClaims, error)
}

type JWTAuth struct {
	secret string
}

type TokenClaims struct {
	Id          int
	Name        string
	Email       string
	PhoneNumber string
}

func NewJwt(secret string) *JWTAuth {
	return &JWTAuth{secret: secret}
}

func (forger *JWTAuth) Encode(tokenClaims TokenClaims, config configs.Config) (string, error) {
	secret := []byte(forger.secret)
	expirationTime := time.Now().Add(config.ExpTime).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = tokenClaims.Id
	claims["name"] = tokenClaims.Name
	claims["email"] = tokenClaims.Email
	claims["phoneNumber"] = tokenClaims.PhoneNumber
	claims["exp"] = expirationTime

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (forger *JWTAuth) Decode(cipher string) (TokenClaims, error) {
	token, err := jwt.Parse(cipher,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(forger.secret), nil
		})
	if err != nil {
		return TokenClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return TokenClaims{}, appErrors.ErrInvalidToken
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return TokenClaims{}, appErrors.ErrInvalidToken
	}

	return TokenClaims{
		Id:          int(id),
		PhoneNumber: fmt.Sprint(claims["phoneNumber"]),
		Email:       fmt.Sprint(claims["email"]),
		Name:        fmt.Sprint(claims["name"]),
	}, nil
}
