package ctxutil

import (
	"context"
	"net/http"
	"testing"
)

func TestWithAuthHeader(t *testing.T) {
	ctx := context.Background()

	testAuthHeaderValue := "Bearer Test-JWT-Token"

	headers := http.Header{
		"Authorization": []string{testAuthHeaderValue},
	}

	ctxWithAuthHeader := WithAuthHeader(ctx, headers)

	actualCtxValue := ctxWithAuthHeader.Value(AuthHeaderName)

	if actualCtxValue != testAuthHeaderValue {
		t.Fatalf("expected: %v, got: %v", testAuthHeaderValue, actualCtxValue)
	}
}
