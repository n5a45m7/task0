package storage

import (
	"app"
	"errors"
)

var (
	ErrStorAccCrUserNotExist error = errors.New("AccountCreator Error: User does not exist")
)

type AccountCreator interface {
	Create(dto CreateAccountDTO) (app.Account, error)
}

type AccountReceiver interface {
	GetByUser(userID int) ([]app.Account, error)
}

type CreateAccountDTO struct {
	CustomerID int
}