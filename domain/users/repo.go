package users

import (
	"context"

	"github.com/chris-ramon/golang-scaffolding/db"
	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type repo struct {
}

func (r *repo) FindUsers(ctx context.Context) (*[]types.CurrentUser, error) {
	return nil, nil
}

func NewRepo(db *db.DB) *repo {
	return &repo{}
}
