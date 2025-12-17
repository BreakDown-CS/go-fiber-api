package repository

import "github.com/BreakDown-CS/go-fiber-api/model"

func FindAllUsers() ([]model.User, error) {
	users := []model.User{
		{ID: 1, Name: "ARM"},
		{ID: 2, Name: "John"},
	}

	return users, nil
}
