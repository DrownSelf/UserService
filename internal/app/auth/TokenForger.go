package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"InnoTaxi/internal/app/errors"
	"InnoTaxi/internal/pkg/configs"
)

type TokenForger interface {
	Encode(name string, email string, config configs.Config) (string, error)
	Decode(cipher string) error
}

type JWTClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTForger struct {
	secret string
}

func NewJwt(secret string) *JWTForger {
	return &JWTForger{secret: secret}
}

func (forger *JWTForger) Encode(name string, email string, config configs.Config) (string, error) {
	secret := []byte(forger.secret)
	expirationTime := time.Now().Add(config.ExpTime).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = expirationTime

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (forger *JWTForger) Decode(cipher string) error {
	_, err := jwt.Parse(cipher,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.ErrInvalidToken
			}
			return []byte(forger.secret), nil
		})

	if err != nil {
		return errors.ErrInvalidToken
	}

	return nil
}
