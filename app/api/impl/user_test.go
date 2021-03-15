package impl

import (
	"app"
	"app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAPI(t *testing.T) {
	storageUser := &_storageUserReceiverStub{}
	storageAcc := &_storageAccReceiverrStub{}
	storageTx := &_storageTxReceiverStub{}

	userAPI := NewUserAPI(
		storageUser,
		storageAcc,
		storageTx,
	)

	_, err := userAPI.GetInfo(1)
	assert.NoError(t, err)

}

type _storageUserReceiverStub struct {
	returnErr error
}

func (s *_storageUserReceiverStub) GetUser(userID int) (app.User, bool, error) {
	if s.returnErr != nil {
		return app.User{}, false, s.returnErr
	}
	return app.User{
		ID: userID,
	}, true, nil
}

type _storageAccReceiverrStub struct {
	returnErr error
}

func (s *_storageAccReceiverrStub) GetByUser(userID int) ([]app.Account, error) {
	if s.returnErr != nil {
		return []app.Account{}, s.returnErr
	}
	return []app.Account{
		app.Account{
			ID:         1,
			CustomerID: userID,
		},
	}, nil
}

type _storageTxReceiverStub struct {
	returnErr error
	amount    float64
	accAmount float64
}

func (s *_storageTxReceiverStub) GetByAccount(accountID int, filters ...storage.FilterApplyItem) ([]app.Transaction, error) {
	if s.returnErr != nil {
		return []app.Transaction{}, s.returnErr
	}
	return []app.Transaction{
		app.Transaction{
			ID:            1,
			AccountID:     accountID,
			Amount:        s.amount,
			AccountAmount: s.accAmount,
		},
	}, nil
}

func (s *_storageTxReceiverStub) GetLastAccountTx(accountID int) (app.Transaction, bool, error) {
	txs, err := s.GetByAccount(accountID)
	if err != nil {
		return app.Transaction{}, false, err
	}
	return txs[0], true, nil
}
