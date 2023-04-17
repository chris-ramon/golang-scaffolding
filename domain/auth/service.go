package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

type service struct {
}

type customClaims struct {
	Data map[string]string `json:"data"`
	jwt.RegisteredClaims
}

func (s *service) CurrentUser() (types.CurrentUser, error) {
	return types.CurrentUser{
		Username: "test user",
	}, nil
}

func (s *service) AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(3 * time.Minute))
	data := map[string]string{}

	claims := customClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("local-signing-secret"))
	if err != nil {
		return nil, err
	}

	return &types.CurrentUser{
		Username: "test user",
		JWT:      token,
	}, nil
}

func NewService() *service {
	return &service{}
}
