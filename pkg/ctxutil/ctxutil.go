package ctxutil

import (
	"context"
	"net/http"
)

func WithAuthHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, "Authorization", header.Get("Authorization"))
}
