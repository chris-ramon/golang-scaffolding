package ctxutil

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type AuthHeaderNameType string

const AuthHeaderName AuthHeaderNameType = "Authorization"

func WithAuthHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, AuthHeaderName, header.Get(string(AuthHeaderName)))
}

func AuthHeaderValueFromCtx(ctx context.Context) (string, error) {
	authorizationWithBearer := strings.Split(ctx.Value(AuthHeaderName).(string), " ")
	if len(authorizationWithBearer) != 2 {
		return "", errors.New("invalid authorization header value")
	}
	return authorizationWithBearer[1], nil
}
