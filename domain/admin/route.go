package admin

import (
	"net/http"

	"github.com/chris-ramon/golang-scaffolding/pkg/ctxutil"
	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type Handlers interface {
	GetAdmin() http.Handler
}

type routes struct {
	handlers Handlers
}

func (ro *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/admin",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				r = r.WithContext(ctxutil.WithAuthHeader(r.Context(), r.Header))
				ro.handlers.GetAdmin().ServeHTTP(w, r)
			},
		},
	}
}

func NewRoutes(handlers Handlers) *routes {
	return &routes{handlers: handlers}
}
