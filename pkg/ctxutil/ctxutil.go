package ctxutil

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const authHeaderName = "Authorization"

func WithAuthHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, authHeaderName, header.Get(authHeaderName))
}

func AuthHeaderValueFromCtx(ctx context.Context) (string, error) {
	authorizationWithBearer := strings.Split(ctx.Value(authHeaderName).(string), " ")
	if len(authorizationWithBearer) != 2 {
		return "", errors.New("invalid authorization header value")
	}
	return authorizationWithBearer[1], nil
}
