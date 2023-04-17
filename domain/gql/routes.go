package gql

import (
	"context"
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
	appSchema, err := schema.New()
	if err != nil {
		log.Fatal(err)
	}

	rootObjectFn := func(ctx context.Context, r *http.Request) map[string]interface{} {
		rootObject := map[string]interface{}{
			"services": &schema.Services{},
		}
		return rootObject
	}

	h := handler.New(&handler.Config{
		Schema:       &appSchema,
		Pretty:       true,
		Playground:   true,
		RootObjectFn: rootObjectFn,
	})

	return &routes{handler: h}
}
