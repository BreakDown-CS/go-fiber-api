package service

import (
	"github.com/BreakDown-CS/go-fiber-api/model"
	"github.com/BreakDown-CS/go-fiber-api/repository"
)

func ChackHealth() string {
	return "OK"
}

func GetUsers() ([]model.User, error) {
	return repository.FindAllUsers()
}
