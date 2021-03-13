package service

import (
	"app"
	"errors"
)

var (
	ErrServTrCrNegativeAccountAmount error = errors.New("TransactionCreator Error: AccountAmount is negative")
)


type TransactionCreator interface {
	Create(dto CreateTransactionDTO) (app.Transaction, error)
}

type TransactionbReceiver interface {
	// Returns last limit transactions from offset related to accountID
	GetByAccount(accountID int, limit int, offset int) ([]app.Transaction, error)
}

type CreateTransactionDTO struct {
	AccountID int
	Amount float64 
}