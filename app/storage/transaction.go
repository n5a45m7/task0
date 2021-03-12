package storage

import (
	"app"
	"errors"
)

var (
	ErrStorTrCrNegativeAccountAmount error = errors.New("TransactionCreator Error: AccountAmount is negative")
)


type TransactionCreator interface {
	Create(dto CreateTransactionDTO) (app.Transction, error)
}

type TransactionbReceiver interface {
	GetByAccount(accountID int, filters... FilterApplyItem) ([]app.Transction, error)
}

type CreateTransactionDTO struct {
	AccountID int
	Amount float64
	AccountAmount float64 
}

// Transaction filtration options

func WithLimitOffset(limit int, offset int) FilterApplyItem {
	return func(f *filters) {
		f.limit = limit
		f.offset = offset
	}
}

func WithDir(dir DIR) FilterApplyItem {
	return func(f *filters) {
		f.dir = dir
	}
}

type FilterApplyItem func(f *filters)

type filters struct {
	limit int
	offset int
	dir DIR
}

var trFilDef filters = filters {
	limit: 10,
	offset: 0,
	dir: DIR_DESC,
}