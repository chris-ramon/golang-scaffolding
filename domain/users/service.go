package users

import (
	"context"

	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type service struct {
}

func (s *service) FindUsers(ctx context.Context) ([]*userTypes.User, error) {
	return nil, nil
}

func NewService(repo Repo) *service {
	return &service{}
}

type Repo interface {
	FindUsers(ctx context.Context) ([]*userTypes.User, error)
}
