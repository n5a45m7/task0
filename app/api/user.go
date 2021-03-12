package api

import "app"

type UserGetInfo interface {
	GetInfo(userID int) (*UserInfo, error)
}

type UserInfo struct {
	Udata app.User
	Accounts []UserInfoAccount
}

type UserInfoAccount struct {
	AData app.Account
	Transactions []app.Transction
}