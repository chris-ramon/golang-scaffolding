package auth

import (
	"net/http"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type Handlers interface {
	GetPing() http.HandlerFunc
	GetCurrentUser() http.HandlerFunc
	PostSignIn() http.HandlerFunc
}

type routes struct {
	handlers Handlers
}

func (r *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/auth/ping",
			Handler:    r.handlers.GetPing(),
		},
		route.Route{
			HTTPMethod: "GET",
			Path:       "/auth/current-user",
			Handler:    r.handlers.GetCurrentUser(),
		},
		route.Route{
			HTTPMethod: "POST",
			Path:       "/auth/sign-in",
			Handler:    r.handlers.PostSignIn(),
		},
	}
}

func NewRoutes(handlers Handlers) *routes {
	return &routes{handlers: handlers}
}
