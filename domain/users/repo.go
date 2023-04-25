package users

import (
	"context"

	"github.com/chris-ramon/golang-scaffolding/db"
	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type repo struct {
}

func (r *repo) FindUsers(ctx context.Context) ([]*userTypes.User, error) {
	return nil, nil
}

func NewRepo(db db.DB) *repo {
	return &repo{}
}
