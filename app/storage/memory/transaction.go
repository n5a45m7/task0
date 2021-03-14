package memory

import (
	"app"
	"app/storage"
	"sync"
)

type transactionStorage struct {
	// [accountID] => []transactions ASC order
	data map[int][]app.Transaction
	// id auto increment
	idInc int

	mu sync.RWMutex
}

func NewTransactionStorage() interface {
	storage.TransactionCreator
	storage.TransactionReceiver
} {
	return &transactionStorage{
		data:  make(map[int][]app.Transaction),
		idInc: 1,
	}
}

func (s *transactionStorage) Create(dto storage.CreateTransactionDTO) (app.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// create future transaction
	newTx := app.Transaction{
		ID:            s.idInc,
		AccountID:     dto.AccountID,
		Amount:        dto.Amount,
		AccountAmount: 0,
	}

	// get list of transactions for accountID
	list, ok := s.data[newTx.AccountID]
	if !ok {
		// create slice for transactions
		list = make([]app.Transaction, 0)
		s.data[newTx.AccountID] = list
	}
	if len(list) > 0 {
		// set current account amount
		newTx.AccountAmount = list[len(list)-1].AccountAmount
	}

	// apply transaction
	newTx.AccountAmount += newTx.Amount

	if newTx.AccountAmount < 0 {
		// this is error, do not save transaction
		return newTx, storage.ErrStorTrCrNegativeAccountAmount
	}

	// everything ok

	// increment auto increment
	s.idInc++
	// save tx to storage
	s.data[newTx.AccountID] = append(s.data[newTx.AccountID], newTx)

	return newTx, nil
}

func (s *transactionStorage) GetByAccount(accountID int, filters ...storage.FilterApplyItem) ([]app.Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filterOptions := storage.GetDefFilters()
	for _, apply := range filters {
		apply(&filterOptions)
	}

	txs, ok := s.data[accountID]
	if !ok || len(txs) <= 0 {
		// no transactions for this account
		return nil, nil
	}

	// si = start index, ei = end index to receive data from slice, dep on filtration order
	si, ei := 0, 0
	// sr = should revert, determine if result should be returned in descending order
	sr := false
	switch filterOptions.Dir {
	case storage.DIR_ASC:
		si, ei = filterOptions.Offset, filterOptions.Offset+filterOptions.Limit-1
		if ei >= len(txs) {
			ei = len(txs) - 1
		}
	case storage.DIR_DESC:
		si, ei = len(txs)-filterOptions.Offset-filterOptions.Limit, len(txs)-1-filterOptions.Offset
		if si < 0 {
			si = 0
		}
		sr = true
	}

	// incorect limit offset, return emprt txs list
	if si > ei {
		return nil, nil
	}

	result := make([]app.Transaction, ei-si+1)
	copy(result, txs[si:ei+1])
	if sr {
		reverse := func(arr []app.Transaction) {
			for i := 0; i < len(arr)/2; i++ {
				j := len(arr) - i - 1
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		reverse(result)
	}

	return result, nil
}

func (s *transactionStorage) GetLastAccountTx(accountID int) (app.Transaction, bool, error) {
	txs, err := s.GetByAccount(accountID, storage.WithOffsetLimit(0, 1))
	if err != nil {
		return app.Transaction{}, false, err
	}

	if len(txs) <= 0 {
		return app.Transaction{}, false, nil
	}

	return txs[0], true, nil
}
