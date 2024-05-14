package auth

import (
	"net/http"
	"testing"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type handlersMock struct {
	getPing        http.HandlerFunc
	getCurrentUser http.HandlerFunc
	postSignIn     http.HandlerFunc
}

func (h *handlersMock) GetPing() http.HandlerFunc        { return h.getPing }
func (h *handlersMock) GetCurrentUser() http.HandlerFunc { return h.getCurrentUser }
func (h *handlersMock) PostSignIn() http.HandlerFunc     { return h.postSignIn }

func TestRoutesAll(t *testing.T) {
	h := &handlersMock{
		getPing: func(w http.ResponseWriter, r *http.Request) {
		},
		getCurrentUser: func(w http.ResponseWriter, r *http.Request) {
		},
		postSignIn: func(w http.ResponseWriter, r *http.Request) {
		},
	}

	routes := NewRoutes(h)

	expectedRoutes := []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/auth/ping",
			Handler:    h.getPing,
		},
		route.Route{
			HTTPMethod: "GET",
			Path:       "/auth/current-user",
			Handler:    h.getCurrentUser,
		},
		route.Route{
			HTTPMethod: "POST",
			Path:       "/auth/sign-in",
			Handler:    h.postSignIn,
		},
	}

	for idx, actualRoute := range routes.All() {
		if actualRoute.HTTPMethod != expectedRoutes[idx].HTTPMethod {
			t.Fatalf("expected: %v, got: %v", expectedRoutes[idx].HTTPMethod, actualRoute.HTTPMethod)
		}
	}
}
