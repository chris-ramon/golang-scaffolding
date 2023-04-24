package users

import (
	"context"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type service struct {
}

func (s *service) FindUsers(ctx context.Context) (*[]types.CurrentUser, error) {
	return nil, nil
}

func NewService(repo Repo) *service {
	return &service{}
}

type Repo interface {
	FindUsers(ctx context.Context) (*[]types.CurrentUser, error)
}
