package service

import (
	"time"

	"github.com/BreakDown-CS/go-fiber-api/apperror"
	"github.com/BreakDown-CS/go-fiber-api/config"
	"github.com/BreakDown-CS/go-fiber-api/repository"
	"github.com/BreakDown-CS/go-fiber-api/store"
	"github.com/golang-jwt/jwt/v5"
)

func Login(username, password string) (string, string, error) {

	user, dbPassword, err := repository.FindByUsername(username)
	if err != nil {
		return "", "", apperror.New(401, "invalid username or password")
	}

	// TODO: Access Token
	accessExp := time.Now().Add(15 * time.Minute)

	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     accessExp.Unix(),
	}

	accessToken, _ := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		accessClaims,
	).SignedString(config.JwtSecret)

	// TODO: Refresh Token
	refreshExp := time.Now().Add(7 * 24 * time.Hour)

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     refreshExp.Unix(),
	}

	refreshToken, _ := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		refreshClaims,
	).SignedString(config.JwtSecret)

	store.Save(accessToken, user.ID, refreshExp)

	if password != dbPassword {
		return "", "", apperror.New(401, "invalid username or password")
	}

	return accessToken, refreshToken, nil
}

func Logout(token string) {
	store.Add(token)
}

func Refresh(refreshToken string) (string, error) {
	rt, ok := store.Get(refreshToken)

	if !ok {
		return "", apperror.New(401, "invalid refresh token")
	}

	if time.Now().After(rt.Expire) {
		store.Delete(refreshToken)
		return "", apperror.New(401, "refresh token expired")
	}

	claims := jwt.MapClaims{
		"user_id": rt.UserID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	newAccess, _ := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	).SignedString(config.JwtSecret)

	return newAccess, nil
}
