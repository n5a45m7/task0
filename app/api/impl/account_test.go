package impl

import (
	"app"
	"app/api"
	"app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountAPI(t *testing.T) {
	storageUsr := &_storageUserReceiverStub{}
	storageAcc := &_storageAccCreatorStub{}
	storageTx := &_storageTxCreatorStub{}
	accountAPI := NewAccountAPI(
		storageUsr,
		storageAcc,
		storageTx,
	)

	_, err := accountAPI.Create(api.AccountCreateRequest{
		CustomerID:    1,
		InitialCredit: 0,
	})
	assert.NoError(t, err)

	_, err = accountAPI.Create(api.AccountCreateRequest{
		CustomerID:    1,
		InitialCredit: 100,
	})
	assert.NoError(t, err)

	_, err = accountAPI.Create(api.AccountCreateRequest{
		CustomerID:    1,
		InitialCredit: -100,
	})
	assert.Error(t, err)
	assert.Equal(t, api.ErrAPIAccCrNegativeAmount, err)

	storageAcc.returnErr = storage.ErrStorAccCrUserNotExist
	_, err = accountAPI.Create(api.AccountCreateRequest{
		CustomerID: 1,
	})
	assert.Error(t, err)
	assert.Equal(t, api.ErrAPIAccCrUserNotExist, err)
	storageAcc.returnErr = nil

	storageTx.returnErr = storage.ErrStorTrCrNegativeAccountAmount
	_, err = accountAPI.Create(api.AccountCreateRequest{
		CustomerID:    1,
		InitialCredit: 100,
	})
	assert.Error(t, err)
	assert.Equal(t, api.ErrAPIAccCrNegativeAmount, err)
	storageTx.returnErr = nil
}

type _storageAccCreatorStub struct {
	returnErr error
}

func (s *_storageAccCreatorStub) Create(dto storage.CreateAccountDTO) (app.Account, error) {
	if s.returnErr != nil {
		return app.Account{}, s.returnErr
	}
	return app.Account{
		ID:         1,
		CustomerID: dto.CustomerID,
	}, nil
}

type _storageTxCreatorStub struct {
	returnErr error
}

func (s *_storageTxCreatorStub) Create(dto storage.CreateTransactionDTO) (app.Transaction, error) {
	if s.returnErr != nil {
		return app.Transaction{}, s.returnErr
	}
	return app.Transaction{
		ID:            1,
		AccountID:     dto.AccountID,
		Amount:        dto.Amount,
		AccountAmount: dto.Amount,
	}, nil
}
