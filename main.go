package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/domain/auth"
	"github.com/chris-ramon/golang-scaffolding/gql/schema"
)

func main() {
	conf := config.New(8080)
	_schema, err := schema.New()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	authRoutes := auth.NewRoutes()
	for _, r := range authRoutes.All() {
		router.Handle(r.HTTPMethod, r.Path, r.Handler)
	}

	h := handler.New(&handler.Config{
		Schema:     &_schema,
		Pretty:     true,
		Playground: true,
	})

	router.Handle("GET", "/graphql", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
	})
	router.Handle("POST", "/graphql", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
	})

	log.Printf("server running on port :%d", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}
