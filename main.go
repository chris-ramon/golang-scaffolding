package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/domain/auth"
	"github.com/chris-ramon/golang-scaffolding/domain/gql"
)

func main() {
	conf := config.New(8080)

	router := httprouter.New()

	authService := auth.NewService()

	handlers := auth.NewHandlers(authService)

	authRoutes := auth.NewRoutes(handlers)
	for _, r := range authRoutes.All() {
		router.Handle(r.HTTPMethod, r.Path, r.Handler)
	}

	gqlRoutes := gql.NewRoutes()
	for _, r := range gqlRoutes.All() {
		router.Handle(r.HTTPMethod, r.Path, r.Handler)
	}

	log.Printf("server running on port :%d", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}
