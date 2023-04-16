package auth

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type Service interface {
	CurrentUser() (types.CurrentUser, error)
}

type handlers struct {
	service Service
}

func (h *handlers) GetPing() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	}
}

func (h *handlers) GetCurrentUser() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		u, err := h.service.CurrentUser()
		if err != nil {
			log.Printf("failed to find current user: %v", err)
			http.Error(w, "failed to find current user", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(u.Username))
	}
}

func NewHandlers(service Service) *handlers {
	return &handlers{service: service}
}
