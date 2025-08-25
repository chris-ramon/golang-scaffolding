package users

import (
	"context"
	"encoding/json"
	"time"

	userTypes "github.com/chris-ramon/golang-scaffolding/domain/users/types"
	cachePkg "github.com/chris-ramon/golang-scaffolding/pkg/cache"
)

const findUsersKey string = "findUsersKey"

type service struct {
	repo  Repo
	cache *cachePkg.Cache[string, string]
}

func (s *service) FindUsers(ctx context.Context) ([]*userTypes.User, error) {
	items, _ := s.cache.Get(findUsersKey)
	if items != nil {
		var result []*userTypes.User
		if err := json.Unmarshal([]byte(*items), &result); err != nil {
			return nil, err
		}

		return result, nil
	}

	users, err := s.repo.FindUsers(ctx)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, nil
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	ttl := cachePkg.TTL(10 * time.Minute)
	s.cache.Set(findUsersKey, string(usersJSON), &ttl)

	return users, nil
}

func NewService(repo Repo, cache *cachePkg.Cache[string, string]) *service {
	return &service{repo: repo, cache: cache}
}

type Repo interface {
	FindUsers(ctx context.Context) ([]*userTypes.User, error)
}
