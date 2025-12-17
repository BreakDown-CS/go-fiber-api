package repository

import (
	"errors"

	"github.com/BreakDown-CS/go-fiber-api/model"
)

func FindAllUsers() ([]model.User, error) {
	// users := []model.User{
	// 	{ID: 1, Name: "ARM"},
	// 	{ID: 2, Name: "John"},
	// }

	// return users, nil
	// fmt.Println("user")
	return nil, errors.New("db down")
}
