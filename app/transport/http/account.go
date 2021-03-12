package http

import (
	nhttp "net/http"
	"app/api"
)

type AccountCreateHandler struct {
	apiCreator api.AccountCreator
}

func NewAccountCreateHandler(apiCreator api.AccountCreator) Handler {
	return &AccountCreateHandler{
		apiCreator: apiCreator,
	}
}

func (h *AccountCreateHandler) Handler(writer nhttp.ResponseWriter, req *nhttp.Request) {
}
