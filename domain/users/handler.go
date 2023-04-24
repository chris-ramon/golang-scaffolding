package users

import (
	"context"
	"net/http"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
	"github.com/julienschmidt/httprouter"
)

type handlers struct {
}

func (h *handlers) GetUsers() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	}
}

func NewHandlers(usersService Service) (*handlers, error) {
	return &handlers{}, nil
}

type Service interface {
	FindUsers(ctx context.Context) (*[]types.CurrentUser, error)
}
