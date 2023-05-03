package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/domain/users/mappers"
	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type handlers struct {
	srv Service
}

func (h *handlers) GetUsers() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		users, err := h.srv.FindUsers(r.Context())

		if err != nil {
			http.Error(w, "failed to find users", http.StatusInternalServerError)
			return
		}

		usersAPI := mappers.UsersFromTypeToAPI(users)

		resp, err := json.Marshal(usersAPI)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func NewHandlers(usersService Service) (*handlers, error) {
	return &handlers{srv: usersService}, nil
}

type Service interface {
	FindUsers(ctx context.Context) ([]*userTypes.User, error)
}
