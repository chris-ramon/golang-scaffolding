package users

import (
	"context"
	"fmt"

	"github.com/chris-ramon/golang-scaffolding/db"
	"github.com/chris-ramon/golang-scaffolding/db/models"
	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

type repo struct {
	db db.DB
}

func (r *repo) FindUsers(ctx context.Context) ([]*userTypes.User, error) {
	args := models.ListUsersParams{
		Limit:  10,
		Offset: 0,
	}

	users, err := r.db.Queries().ListUsers(ctx, args)

	if err != nil {
		return nil, err
	}

	var result []*userTypes.User

	for _, user := range users {
		result = append(result, &userTypes.User{
			ID:       fmt.Sprintf("%d", user.ID),
			Username: user.Username,
		})
	}

	return result, nil
}

func NewRepo(db db.DB) *repo {
	return &repo{db: db}
}
