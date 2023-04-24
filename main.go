package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/db"
	"github.com/chris-ramon/golang-scaffolding/domain/admin"
	"github.com/chris-ramon/golang-scaffolding/domain/auth"
	"github.com/chris-ramon/golang-scaffolding/domain/gql"
	"github.com/chris-ramon/golang-scaffolding/domain/users"
	"github.com/chris-ramon/golang-scaffolding/pkg/route"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	conf := config.New()
	dbConf := config.NewDBConfig()

	db, err := db.New(dbConf)
	if err != nil {
		handleErr(err)
	}

	router := httprouter.New()

	usersRepo := users.NewRepo(db)
	usersService := users.NewService(usersRepo)
	usersHandlers, err := users.NewHandlers(usersService)
	if err != nil {
		handleErr(err)
	}
	usersRoutes := users.NewRoutes(usersHandlers)

	authService := auth.NewService()
	authHandlers := auth.NewHandlers(authService)
	authRoutes := auth.NewRoutes(authHandlers)

	gqlHandlers, err := gql.NewHandlers(authService)
	if err != nil {
		handleErr(err)
	}

	gqlRoutes := gql.NewRoutes(gqlHandlers)

	adminHandlers, err := admin.NewHandlers()
	if err != nil {
		handleErr(err)
	}
	adminRoutes := admin.NewRoutes(adminHandlers)

	routes := []route.Route{}
	routes = append(routes, authRoutes.All()...)
	routes = append(routes, gqlRoutes.All()...)
	routes = append(routes, adminRoutes.All()...)
	routes = append(routes, usersRoutes.All()...)

	for _, r := range routes {
		router.Handle(r.HTTPMethod, r.Path, r.Handler)
	}

	log.Printf("server running on port :%s", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router))
}
