package http

import (
	"app/api"
	nhttp "net/http"
)

type UserGetInfoHandler struct {
	apiReceiver api.UserGetInfo
}

func NewUserGetInfoHandler(apiReceiver api.UserGetInfo) Handler {
	return &UserGetInfoHandler{
		apiReceiver: apiReceiver,
	}
}

func (h *UserGetInfoHandler) Handler(writer nhttp.ResponseWriter, req *nhttp.Request) {
}