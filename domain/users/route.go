package users

import (
	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type Handlers interface {
	GetUsers() httprouter.Handle
}

type routes struct {
	handlers Handlers
}

func (r *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/users",
			Handler:    r.handlers.GetUsers(),
		},
	}
}

func NewRoutes(handlers Handlers) *routes {
	return &routes{handlers: handlers}
}
