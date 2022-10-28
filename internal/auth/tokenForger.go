package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	configs "github.com/DrownSelf/UserService/internal/config"
)

type TokenForger interface {
	Encode(tokenClaims TokenClaims, config configs.Config) (string, error)
	Decode(cipher string) (TokenClaims, error)
}

type JWTForger struct {
	secret string
}

type TokenClaims struct {
	Id          int
	Name        string
	Email       string
	PhoneNumber string
}

func NewJwt(secret string) *JWTForger {
	return &JWTForger{secret: secret}
}

func (forger *JWTForger) Encode(tokenClaims TokenClaims, config configs.Config) (string, error) {
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

func (forger *JWTForger) Decode(cipher string) (TokenClaims, error) {
	token, err := jwt.Parse(cipher,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(forger.secret), nil
		})
	if err != nil {
		return TokenClaims{}, err
	}
	claims := token.Claims.(jwt.MapClaims)

	return TokenClaims{
		Id:          int(claims["id"].(float64)),
		PhoneNumber: fmt.Sprint(claims["phoneNumber"]),
		Email:       fmt.Sprint(claims["email"]),
		Name:        fmt.Sprint(claims["name"]),
	}, nil
}
