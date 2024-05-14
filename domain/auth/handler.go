package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
	"github.com/chris-ramon/golang-scaffolding/pkg/ctxutil"
)

type Service interface {
	CurrentUser(ctx context.Context, jwtToken string) (*types.CurrentUser, error)
	AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error)
}

type handlers struct {
	service Service
}

func (h *handlers) GetPing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}
}

func (h *handlers) GetCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ctxutil.WithAuthHeader(r.Context(), r.Header))
		jwtToken, err := ctxutil.AuthHeaderValueFromCtx(r.Context())
		if err != nil {
			log.Printf("failed to get authorization header: %v", err)
			http.Error(w, "failed to get authorization header", http.StatusInternalServerError)
			return
		}

		u, err := h.service.CurrentUser(r.Context(), jwtToken)
		if err != nil {
			log.Printf("failed to find current user: %v", err)
			http.Error(w, "failed to find current user", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(u.Username))
	}
}

func (h *handlers) PostSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("failed to read request body: %v", err)
			http.Error(w, "failed to read request body", http.StatusInternalServerError)
			return
		}

		var reqBodyJSON struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.Unmarshal(b, &reqBodyJSON); err != nil {
			log.Printf("failed to json unmarshal request body: %v", err)
			http.Error(w, "failed to json unmarshal request body", http.StatusInternalServerError)
			return
		}

		u, err := h.service.AuthUser(r.Context(), reqBodyJSON.Email, reqBodyJSON.Password)
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
