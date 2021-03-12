package api

import (
	"app"
	"errors"
)

var (
	ErrAPIAccCrUserNotExist error = errors.New("AccountCreator Error: User does not exist")
	ErrAPIAccCrNegativeAmount error = errors.New("AccountCreator Error: Initial credit is negative")
)

type AccountCreator interface {
	Create(request AccountCreateRequest) (*app.Account, error)
}

type AccountCreateRequest struct {
	CustomerID int
	InitialCredit float64
}