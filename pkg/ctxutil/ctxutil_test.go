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

func TestAuthHeaderValueFromCtx(t *testing.T) {
	testAuthHeaderValue := "Test-JWT-Token"
	testAuthHeaderValueWithBearer := "Bearer " + testAuthHeaderValue

	ctx := context.WithValue(context.Background(), AuthHeaderName, testAuthHeaderValueWithBearer)

	actualCtxValue, err := AuthHeaderValueFromCtx(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actualCtxValue != testAuthHeaderValue {
		t.Fatalf("expected: %v, got: %v", testAuthHeaderValue, actualCtxValue)
	}
}

func TestAuthHeaderValueFromCtx_Error(t *testing.T) {
	testAuthHeaderValue := "Test-JWT-Token"

	ctx := context.WithValue(context.Background(), AuthHeaderName, testAuthHeaderValue)

	actualCtxValue, err := AuthHeaderValueFromCtx(ctx)
	if err == nil {
		t.Fatalf("expected error got nil")
	}
	if actualCtxValue != "" {
		t.Fatalf("expected empty string value got: %v", actualCtxValue)
	}
}
