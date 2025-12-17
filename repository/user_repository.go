package repository

import (
	"errors"

	"github.com/BreakDown-CS/go-fiber-api/model"
)

func FindAllUsers() ([]model.User, error) {
	users := []model.User{
		{ID: 1, Name: "ARM"},
		{ID: 2, Name: "John"},
	}

	return users, nil
}

func FindByUsername(username string) (*model.User, string, error) {
	if username == "arm" {
		return &model.User{
			ID:   1,
			Name: "ARM",
		}, "1234", nil // password mock
	}

	return nil, "", errors.New("user not found")
}
