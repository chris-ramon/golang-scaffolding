package auth

import (
	"context"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type service struct {
	jwt JWT
}

func (s *service) CurrentUser(ctx context.Context, jwtToken string) (*types.CurrentUser, error) {
	data, err := s.jwt.Validate(ctx, jwtToken)
	if err != nil {
		return nil, err
	}

	return &types.CurrentUser{
		Username: data["username"],
	}, nil
}

func (s *service) AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
	data := map[string]string{
		"username": username,
	}

	jwtToken, err := s.jwt.Generate(ctx, data)
	if err != nil {
		return nil, err
	}

	return &types.CurrentUser{
		Username: username,
		JWT:      *jwtToken,
	}, nil
}

func NewService(jwt JWT) (*service, error) {
	return &service{jwt: jwt}, nil
}

type JWT interface {
	Generate(ctx context.Context, data map[string]string) (*string, error)
	Validate(ctx context.Context, jwtToken string) (map[string]string, error)
}
