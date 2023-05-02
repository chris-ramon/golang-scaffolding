package jwt

import (
	"context"
	"crypto/rsa"
	_ "embed"
	"time"

	jwtV5 "github.com/golang-jwt/jwt/v5"
)

//go:embed app.rsa
var appRsa []byte // openssl genrsa -out app.rsa 2048

//go:embed app.rsa.pub
var appRsaPub []byte // openssl rsa -in app.rsa -pubout > app.rsa.pub

type customClaims struct {
	Data map[string]string `json:"data"`
	jwtV5.RegisteredClaims
}

type jwt struct {
	defaultExpireTimeInHours time.Duration
	signKey                  *rsa.PrivateKey
	verifyKey                *rsa.PublicKey
}

func (j *jwt) Generate(ctx context.Context, data map[string]string) (*string, error) {
	expiresAt := jwtV5.NewNumericDate(time.Now().Add(j.defaultExpireTimeInHours * time.Hour))

	t := jwtV5.New(jwtV5.GetSigningMethod("RS256"))
	t.Claims = customClaims{
		data,
		jwtV5.RegisteredClaims{
			ExpiresAt: expiresAt,
		},
	}

	jwtToken, err := t.SignedString(j.signKey)
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func (j *jwt) Validate(ctx context.Context, jwtToken string) (map[string]string, error) {
	parsedJWTToken, err := jwtV5.ParseWithClaims(jwtToken, &customClaims{}, func(t *jwtV5.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims := parsedJWTToken.Claims.(*customClaims)

	return claims.Data, nil
}

func NewJWT() (*jwt, error) {
	signBytes := appRsa

	sKey, err := jwtV5.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	verifyBytes := appRsaPub

	vKey, err := jwtV5.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return &jwt{
		defaultExpireTimeInHours: 24,
		signKey:                  sKey,
		verifyKey:                vKey,
	}, nil
}
