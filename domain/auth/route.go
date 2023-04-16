package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type routes struct {
}

func (r *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/auth/ping",
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Write([]byte("ok"))
			},
		},
	}
}

func NewRoutes() *routes {
	return &routes{}
}
