package auth

import (
	"context"
	"crypto/rsa"
	"log"
	"os"
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

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	handleError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	signBytes, err := os.ReadFile("./domain/auth/app.rsa") // openssl genrsa -out app.rsa 2048
	handleError(err)

	sKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	handleError(err)
	signKey = sKey

	verifyBytes, err := os.ReadFile("./domain/auth/app.rsa.pub") // openssl rsa -in app.rsa -pubout > app.rsa.pub
	handleError(err)

	vKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	handleError(err)
	verifyKey = vKey
}

func (s *service) CurrentUser(jwtToken string) (*types.CurrentUser, error) {
	parsedJWTToken, err := jwt.ParseWithClaims(jwtToken, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return verifyKey, nil
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

	jwtToken, err := t.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	return &types.CurrentUser{
		Username: username,
		JWT:      jwtToken,
	}, nil
}

func NewService() *service {
	return &service{}
}
