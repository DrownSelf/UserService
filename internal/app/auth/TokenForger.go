package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"InnoTaxi/internal/pkg/configs"
)

type TokenForger interface {
	Encode(tokenClaims TokenClaims, config configs.Config) (string, error)
	Decode(cipher string) error
}

type JWTForger struct {
	secret string
}

type TokenClaims struct {
	Id    int
	Name  string
	Email string
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
			return []byte(forger.secret), nil
		})
	if err != nil {
		return err
	}
	return nil
}
