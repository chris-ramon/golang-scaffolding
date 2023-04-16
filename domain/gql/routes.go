package gql

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/domain/gql/schema"
	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

type routes struct {
	handler *handler.Handler
}

func (ro *routes) All() []route.Route {
	return []route.Route{
		route.Route{
			HTTPMethod: "GET",
			Path:       "/graphql",
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				ro.handler.ServeHTTP(w, r)
			},
		},
		route.Route{
			HTTPMethod: "POST",
			Path:       "/graphql",
			Handler: func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				ro.handler.ServeHTTP(w, r)
			},
		},
	}
}

func NewRoutes() *routes {
	_schema, err := schema.New()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &_schema,
		Pretty:     true,
		Playground: true,
	})

	return &routes{handler: h}
}
