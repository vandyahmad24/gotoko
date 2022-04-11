package auth

import (
	"errors"

	"vandyahmad/gotoko/config"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(config.GetEnvVariable("JWT_SECRET"))

func NewServiceAuth() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	toke, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return SECRET_KEY, nil
	})
	if err != nil {
		return toke, err
	}
	return toke, nil
}
