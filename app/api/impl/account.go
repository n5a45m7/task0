package impl

import (
	"app"
	"app/api"
	"app/storage"
)

type accountAPI struct {
	userSt storage.UserReceiver
	accSt  storage.AccountCreator
	txSt   storage.TransactionCreator
}

func NewAccountAPI(
	userSt storage.UserReceiver,
	accSt storage.AccountCreator,
	txSt storage.TransactionCreator,
) api.AccountCreator {
	return &accountAPI{
		userSt: userSt,
		accSt:  accSt,
		txSt:   txSt,
	}
}

func (a *accountAPI) Create(request api.AccountCreateRequest) (app.Account, error) {
	if _, ok, _ := a.userSt.GetUser(request.CustomerID); !ok {
		return app.Account{}, api.ErrAPIUserUserNotFound
	}
	acc, err := a.accSt.Create(
		storage.CreateAccountDTO{
			CustomerID: request.CustomerID,
		},
	)

	switch err {
	case nil:
		// do nothing
	case storage.ErrStorAccCrUserNotExist:
		return acc, api.ErrAPIAccCrUserNotExist
	default:
		return acc, err
	}

	if request.InitialCredit < 0 {
		return acc, api.ErrAPIAccCrNegativeAmount
	}
	if request.InitialCredit == 0 {
		return acc, nil
	}

	_, err = a.txSt.Create(storage.CreateTransactionDTO{
		AccountID: acc.ID,
		Amount:    request.InitialCredit,
	})
	switch err {
	case nil:
		// do nothing
	case storage.ErrStorTrCrNegativeAccountAmount:
		return acc, api.ErrAPIAccCrNegativeAmount
	default:
		return acc, err
	}

	return acc, nil
}
