package auth

import (
	"context"
	"crypto/rsa"
	_ "embed"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
)

//go:embed app.rsa
var appRsa []byte // openssl genrsa -out app.rsa 2048

//go:embed app.rsa.pub
var appRsaPub []byte // openssl rsa -in app.rsa -pubout > app.rsa.pub

type service struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

type customClaims struct {
	Data map[string]string `json:"data"`
	jwt.RegisteredClaims
}

func (s *service) CurrentUser(jwtToken string) (*types.CurrentUser, error) {
	parsedJWTToken, err := jwt.ParseWithClaims(jwtToken, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return s.verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims := parsedJWTToken.Claims.(*customClaims)

	return &types.CurrentUser{
		Username: claims.Data["username"],
	}, nil
}

func (s *service) AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	data := map[string]string{
		"username": username,
	}

	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = customClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
		},
	}

	jwtToken, err := t.SignedString(s.signKey)
	if err != nil {
		return nil, err
	}

	return &types.CurrentUser{
		Username: username,
		JWT:      jwtToken,
	}, nil
}

func (s *service) FindUsers(ctx context.Context) ([]*types.CurrentUser, error) {
	return nil, nil
}

func NewService() (*service, error) {
	signBytes := appRsa

	sKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	verifyBytes := appRsaPub

	vKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return &service{
		signKey:   sKey,
		verifyKey: vKey,
	}, nil
}
