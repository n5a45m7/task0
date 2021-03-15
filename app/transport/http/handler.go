package http

import "net/http"

// Handler --
type Handler interface {
	Handler(writer http.ResponseWriter, req *http.Request)
}
