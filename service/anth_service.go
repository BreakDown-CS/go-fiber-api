package service

import (
	"time"

	"github.com/BreakDown-CS/go-fiber-api/apperror"
	"github.com/BreakDown-CS/go-fiber-api/config"
	"github.com/BreakDown-CS/go-fiber-api/repository"
	"github.com/golang-jwt/jwt/v5"
)

func Login(username, password string) (string, error) {

	user, dbPassword, err := repository.FindByUsername(username)
	if err != nil {
		return "", apperror.New(401, "invalid username or password")
	}

	if password != dbPassword {
		return "", apperror.New(401, "invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return "", apperror.New(500, "internal server error")
	}

	return signedToken, nil
}
