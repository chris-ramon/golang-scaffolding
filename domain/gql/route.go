package gql

import (
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/chris-ramon/golang-scaffolding/pkg/ctxutil"
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
			Handler: func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(ctxutil.WithAuthHeader(r.Context(), r.Header))
				ro.handlers.GetGraphQL().ServeHTTP(w, r)
			},
		},
		route.Route{
			HTTPMethod: "POST",
			Path:       "/graphql",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(ctxutil.WithAuthHeader(r.Context(), r.Header))
				ro.handlers.PostGraphQL().ServeHTTP(w, r)
			},
		},
	}
}

func NewRoutes(handlers Handlers) *routes {
	return &routes{handlers: handlers}
}
