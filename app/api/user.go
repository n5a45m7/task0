package api

import (
	"app"
	"errors"
)

var (
	ErrAPIUserUserNotFound error = errors.New("UserGetInfo Error: user not found")
)

type UserGetInfo interface {
	GetInfo(userID int) (UserInfo, error)
}

type UserInfo struct {
	Udata    app.User
	Accounts []UserInfoAccount
}

type UserInfoAccount struct {
	AData        app.Account
	Balance      float64
	Transactions []app.Transaction
}
