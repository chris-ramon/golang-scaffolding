package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

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

func (h *handlers) GetPing() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	}
}

func (h *handlers) GetCurrentUser() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r = r.WithContext(ctxutil.WithAuthHeader(r.Context(), r.Header))
		authorization, err := ctxutil.AuthHeaderValueFromCtx(r.Context())
		if err != nil {
			log.Printf("failed to get authorization header: %v", err)
			http.Error(w, "failed to get authorization header", http.StatusInternalServerError)
			return
		}

		u, err := h.service.CurrentUser(r.Context(), authorization)
		if err != nil {
			log.Printf("failed to find current user: %v", err)
			http.Error(w, "failed to find current user", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(u.Username))
	}
}

func (h *handlers) PostSignIn() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
