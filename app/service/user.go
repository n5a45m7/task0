package service

import (
	"app"
)

type UserReceiver interface {
	GetUser(userID int) (app.User, error)
}