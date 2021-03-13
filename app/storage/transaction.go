package storage

import (
	"app"
	"errors"
)

var (
	ErrStorTrCrNegativeAccountAmount error = errors.New("TransactionCreator Error: AccountAmount is negative")
)


type TransactionCreator interface {
	Create(dto CreateTransactionDTO) (app.Transaction, error)
}

type TransactionReceiver interface {
	GetByAccount(accountID int, filters... FilterApplyItem) ([]app.Transaction, error)
	GetLastAccountTx(accountID int) (app.Transaction, bool, error)
}

type CreateTransactionDTO struct {
	AccountID int
	Amount float64
}

// Transaction filtration options

func WithOffsetLimit(offset int, limit int) FilterApplyItem {
	return func(f *Filters) {
		f.Limit = limit
		f.Offset = offset
	}
}

func WithDir(dir DIR) FilterApplyItem {
	return func(f *Filters) {
		f.Dir = dir
	}
}

type FilterApplyItem func(f *Filters)

type Filters struct {
	Limit int
	Offset int
	Dir DIR
}

func GetDefFilters() Filters {
	return Filters {
		Limit: 10,
		Offset: 0,
		Dir: DIR_DESC,
	}
}