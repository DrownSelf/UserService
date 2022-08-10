package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenForger interface {
	Encode(name string, email string) (string, error)
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

func (forger *JWTForger) Encode(name string, email string) (string, error) {
	secret := []byte(forger.secret)
	expirationTime := time.Now().Add(1 * time.Hour).Unix()
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
	token, err := jwt.ParseWithClaims(
		cipher,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(forger.secret), nil
		})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
