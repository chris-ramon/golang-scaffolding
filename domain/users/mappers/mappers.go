package mappers

import (
	"github.com/chris-ramon/golang-scaffolding/domain/users/api"
	"github.com/chris-ramon/golang-scaffolding/domain/users/types"
)

func UsersFromTypeToAPI(users []*types.User) []api.User {
	apiUsers := []api.User{}

	for _, user := range users {
		apiUsers = append(apiUsers, api.User{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	return apiUsers
}
