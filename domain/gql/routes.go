package gql

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type Handlers interface {
	PostGraphQL() *handler.Handler
	GetGraphQL() *handler.Handler
}

type routes struct {
	handlers Handlers
}

func (ro *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/graphql",
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				ro.handlers.PostGraphQL().ServeHTTP(w, r)
			},
		},
		route.Route{
			HTTPMethod: "POST",
			Path:       "/graphql",
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				ro.handlers.GetGraphQL().ServeHTTP(w, r)
			},
		},
	}
}

func NewRoutes(handlers Handlers) *routes {
	return &routes{handlers: handlers}
}
