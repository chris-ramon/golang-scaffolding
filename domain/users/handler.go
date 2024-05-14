package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/admin-golang/admin/dataloader"
	"github.com/chris-ramon/golang-scaffolding/domain/users/mappers"
	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type handlers struct {
	srv Service
}

func (h *handlers) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.srv.FindUsers(r.Context())

		if err != nil {
			http.Error(w, "failed to find users", http.StatusInternalServerError)
			return
		}

		usersAPI := mappers.UsersFromTypeToAPI(users)

		response := dataloader.Response{
			Data: usersAPI,
			Meta: dataloader.Meta{
				Headers: []string{"ID", "Username"},
				Components: map[string]string{
					"id":       "text",
					"username": "text",
				},
				Pagination: dataloader.Pagination{
					TotalCount:  len(usersAPI),
					PerPage:     10,
					CurrentPage: 0,
					RowsPerPage: []int{10, 25, 50},
				},
			},
		}

		resp, err := json.Marshal(response)
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
