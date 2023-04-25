package users

import (
	"context"

	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type service struct {
	repo Repo
}

func (s *service) FindUsers(ctx context.Context) ([]*userTypes.User, error) {
	return s.repo.FindUsers(ctx)
}

func NewService(repo Repo) *service {
	return &service{repo: repo}
}

type Repo interface {
	FindUsers(ctx context.Context) ([]*userTypes.User, error)
}
