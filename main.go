package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql/testutil"
	"github.com/graphql-go/handler"

	"github.com/chris-ramon/golang-scaffolding/config"
)

func main() {
	conf := config.New(8080)
	h := handler.New(&handler.Config{
		Schema:   &testutil.StarWarsSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	log.Printf("server running on port :%d", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}
