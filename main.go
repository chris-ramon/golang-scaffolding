package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/domain/auth"
	"github.com/chris-ramon/golang-scaffolding/domain/gql"
	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	conf := config.New(8080)

	router := httprouter.New()

	authService := auth.NewService()
	authHandlers := auth.NewHandlers(authService)
	authRoutes := auth.NewRoutes(authHandlers)

	gqlHandlers, err := gql.NewHandlers(authService)
	if err != nil {
		handleErr(err)
	}

	gqlRoutes := gql.NewRoutes(gqlHandlers)

	routes := []route.Route{}
	routes = append(routes, authRoutes.All()...)
	routes = append(routes, gqlRoutes.All()...)

	for _, r := range routes {
		router.Handle(r.HTTPMethod, r.Path, r.Handler)
	}

	log.Printf("server running on port :%d", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}
