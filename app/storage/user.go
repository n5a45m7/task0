package storage

import (
	"app"
)

type UserCreator interface {
	Create(dto CreateUserDTO) (app.User, error)
}

type UserReceiver interface {
	GetUser(userID int) (app.User, bool, error)
}

type CreateUserDTO struct {
	Name string
	Surname string
}