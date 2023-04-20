package services

import (
	"context"
	authTypes "github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type AuthService interface {
	CurrentUser(jwtToken string) (*authTypes.CurrentUser, error)
	AuthUser(ctx context.Context, username string, pwd string) (*authTypes.CurrentUser, error)
}

type Services struct {
	AuthService AuthService
}
