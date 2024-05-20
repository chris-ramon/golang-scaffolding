package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/db"
	"github.com/chris-ramon/golang-scaffolding/domain/admin"
	"github.com/chris-ramon/golang-scaffolding/domain/auth"
	"github.com/chris-ramon/golang-scaffolding/domain/gql"
	"github.com/chris-ramon/golang-scaffolding/domain/users"
	"github.com/chris-ramon/golang-scaffolding/pkg/jwt"
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

	if err := db.Migrate(); err != nil {
		handleErr(err)
	} else {
		log.Println("successfully run migrations")
	}

	router := http.NewServeMux()

	usersRepo := users.NewRepo(db)
	usersService := users.NewService(usersRepo)
	usersHandlers, err := users.NewHandlers(usersService)
	if err != nil {
		handleErr(err)
	}
	usersRoutes := users.NewRoutes(usersHandlers)

	jwt, err := jwt.NewJWT()
	if err != nil {
		handleErr(err)
	}

	authService, err := auth.NewService(jwt)
	if err != nil {
		handleErr(err)
	}
	authHandlers := auth.NewHandlers(authService)
	authRoutes := auth.NewRoutes(authHandlers)

	gqlHandlers, err := gql.NewHandlers(authService, usersService)
	if err != nil {
		handleErr(err)
	}

	gqlRoutes := gql.NewRoutes(gqlHandlers)

	adminHandlers, err := admin.NewHandlers()
	if err != nil {
		handleErr(err)
	}
	adminRoutes := admin.NewRoutes(adminHandlers)

	routes := slices.Concat(
		authRoutes.All(),
		gqlRoutes.All(),
		adminRoutes.All(),
		usersRoutes.All(),
	)

	for _, r := range routes {
		router.HandleFunc(fmt.Sprintf("%s %s", r.HTTPMethod, r.Path), r.Handler)
	}

	log.Printf("server running on port :%s", conf.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router))
}
