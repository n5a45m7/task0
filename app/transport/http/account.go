package http

import (
	"app/api"
	"encoding/json"
	"io/ioutil"
	nhttp "net/http"
)

type AccountCreateHandler struct {
	apiCreator api.AccountCreator
}

func NewAccountCreateHandler(apiCreator api.AccountCreator) Handler {
	return &AccountCreateHandler{
		apiCreator: apiCreator,
	}
}

func (h *AccountCreateHandler) Handler(w nhttp.ResponseWriter, r *nhttp.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		nhttp.Error(w, err.Error(), nhttp.StatusInternalServerError)
		return
	}

	// Unmarshal
	var req accCreateRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		nhttp.Error(w, err.Error(), 500)
		return
	}

	// Create account
	account, err := h.apiCreator.Create(
		api.AccountCreateRequest{
			CustomerID:    req.CustomerID,
			InitialCredit: req.InitialCredit,
		},
	)
	if err != nil {
		switch err {
		case api.ErrAPIAccCrUserNotExist:
			nhttp.Error(w, err.Error(), nhttp.StatusNotFound)
		default:
			nhttp.Error(w, err.Error(), nhttp.StatusBadRequest)
		}
		return
	}

	resp := accCreateOkResponse{
		ID: account.ID,
	}
	b, _ = json.Marshal(resp)
	w.Header().Add("content-type", "application/json")
	w.Write(b)
}

type accCreateRequest struct {
	CustomerID    int     `json:"customerID"`
	InitialCredit float64 `json:"initialCredit"`
}

type accCreateOkResponse struct {
	ID int `json:"id"`
}
