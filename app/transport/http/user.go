package http

import (
	"app/api"
	"encoding/json"
	nhttp "net/http"
	"strconv"
)

type UserGetInfoHandler struct {
	apiReceiver api.UserGetInfo
}

func NewUserGetInfoHandler(apiReceiver api.UserGetInfo) Handler {
	return &UserGetInfoHandler{
		apiReceiver: apiReceiver,
	}
}

func (h *UserGetInfoHandler) Handler(w nhttp.ResponseWriter, r *nhttp.Request) {
	userIDStr := r.URL.Query().Get("userID")
	if userIDStr == "" {
		nhttp.Error(w, "User not found", nhttp.StatusNotFound)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		nhttp.Error(w, err.Error(), nhttp.StatusBadRequest)
		return
	}

	userInfo, err := h.apiReceiver.GetInfo(userID)
	if err != nil {
		switch err {
		case api.ErrAPIUserUserNotFound:
			nhttp.Error(w, err.Error(), nhttp.StatusNotFound)
		default:
			nhttp.Error(w, err.Error(), nhttp.StatusBadRequest)
		}
		return
	}

	resp := userInfoResponse{
		ID:      userInfo.Udata.ID,
		Name:    userInfo.Udata.Name,
		Surname: userInfo.Udata.Surname,
	}
	for _, acc := range userInfo.Accounts {
		accResp := userInfoAccResponse{
			ID:      acc.AData.ID,
			Balance: acc.Balance,
		}
		for _, tx := range acc.Transactions {
			txResp := userInfoTxResponse{
				ID:        tx.ID,
				Amount:    tx.Amount,
				AccAmount: tx.AccountAmount,
			}
			accResp.Txs = append(accResp.Txs, txResp)
		}
		resp.Accounts = append(resp.Accounts, accResp)
	}
	b, _ := json.Marshal(resp)
	w.Write(b)
}

type userInfoResponse struct {
	ID       int                   `json:"id"`
	Name     string                `json:"name"`
	Surname  string                `json:"surname"`
	Accounts []userInfoAccResponse `json:"accounts"`
}

type userInfoAccResponse struct {
	ID      int                  `json:"id"`
	Balance float64              `json:"balance"`
	Txs     []userInfoTxResponse `json:"txs"`
}

type userInfoTxResponse struct {
	ID        int     `json:"id"`
	Amount    float64 `json:"amount"`
	AccAmount float64 `json:"accAmount"`
}
