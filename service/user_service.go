package service

import (
	"github.com/BreakDown-CS/go-fiber-api/apperror"
	"github.com/BreakDown-CS/go-fiber-api/model"
	"github.com/BreakDown-CS/go-fiber-api/repository"
)

func ChackHealth() string {
	return "OK"
}

func GetUsers() ([]model.User, error) {

	users, err := repository.FindAllUsers()
	if err != nil {
		return nil, apperror.New(500, "database error")
	}

	return users, nil
}
