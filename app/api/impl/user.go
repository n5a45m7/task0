package impl

import (
	"app/api"
	"app/storage"
)

type userAPI struct {
	userSt storage.UserReceiver
	accSt storage.AccountReceiver
	txSt storage.TransactionReceiver
}

func NewUserAPI(
	userSt storage.UserReceiver,
	accSt storage.AccountReceiver,
	txSt storage.TransactionReceiver,
) api.UserGetInfo {
	return &userAPI{
		userSt: userSt,
		accSt: accSt,
		txSt: txSt,
	}
}

func (a *userAPI) GetInfo(userID int) (api.UserInfo, error) {
	user, ok, err := a.userSt.GetUser(userID)
	if !ok {
		return api.UserInfo{}, api.ErrAPIUserUserNotFound
	}

	accounts, err := a.accSt.GetByUser(user.ID)
	if err != nil {
		return api.UserInfo{}, err
	}

	result := api.UserInfo{
		Udata: user,
	}
	for _, acc := range accounts {
		txs, _ := a.txSt.GetByAccount(
			acc.ID,
			storage.WithOffsetLimit(0, 10),
			storage.WithDir(storage.DIR_DESC),
		)
		balance := 0.0
		if len(txs) > 0 {
			balance = txs[0].AccountAmount
		}
		result.Accounts = append(
			result.Accounts, 
			api.UserInfoAccount{
				AData: acc,
				Transactions: txs,
				Balance: balance,
			},
		)		
	}
	
	return result, nil
}