package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handlers struct {
}

func (h *handlers) GetPing() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	}
}

func (h *handlers) GetCurrentUser() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	}
}

func NewHandlers() *handlers {
	return &handlers{}
}
