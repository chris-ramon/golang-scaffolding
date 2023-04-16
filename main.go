package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/gql/schema"
)

func main() {
	conf := config.New(8080)
	_schema, err := schema.New()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &_schema,
		Pretty:     true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Printf("server running on port :%d", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
