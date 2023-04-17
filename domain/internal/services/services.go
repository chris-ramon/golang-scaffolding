package services

import (
	authTypes "github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type AuthService interface {
	CurrentUser() (authTypes.CurrentUser, error)
}

type Services struct {
	AuthService AuthService
}
