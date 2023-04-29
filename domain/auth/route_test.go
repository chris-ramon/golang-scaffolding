package auth

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type handlersMock struct {
	getPing        httprouter.Handle
	getCurrentUser httprouter.Handle
	postSignIn     httprouter.Handle
}

func (h *handlersMock) GetPing() httprouter.Handle        { return h.getPing }
func (h *handlersMock) GetCurrentUser() httprouter.Handle { return h.getCurrentUser }
func (h *handlersMock) PostSignIn() httprouter.Handle     { return h.postSignIn }

func TestRoutesAll(t *testing.T) {
	h := &handlersMock{
		getPing: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		},
		getCurrentUser: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		},
		postSignIn: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
