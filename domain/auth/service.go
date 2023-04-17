package auth

import (
	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type service struct {
}

func (s *service) CurrentUser() (types.CurrentUser, error) {
	return types.CurrentUser{
		Username: "test user",
	}, nil
}

func NewService() *service {
	return &service{}
}
