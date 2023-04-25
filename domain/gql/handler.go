package gql

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/chris-ramon/golang-scaffolding/domain/gql/schema"
	"github.com/chris-ramon/golang-scaffolding/domain/internal/services"
)

type handlers struct {
	gqlHandler *handler.Handler
}

func (h *handlers) PostGraphQL() *handler.Handler {
	return h.gqlHandler
}

func (h *handlers) GetGraphQL() *handler.Handler {
	return h.gqlHandler
}

func NewHandlers(authService services.AuthService, userService services.UserService) (*handlers, error) {
	appSchema, err := schema.New()
	if err != nil {
		return nil, err
	}

	rootObjectFn := func(ctx context.Context, r *http.Request) map[string]interface{} {
		rootObject := map[string]interface{}{
			"services": &services.Services{
				AuthService: authService,
				UserService: userService,
			},
		}
		return rootObject
	}

	h := handler.New(&handler.Config{
		Schema:       &appSchema,
		Pretty:       true,
		Playground:   true,
		RootObjectFn: rootObjectFn,
	})

	return &handlers{
		gqlHandler: h,
	}, nil
}
