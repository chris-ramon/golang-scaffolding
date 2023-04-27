package auth

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
	"github.com/julienschmidt/httprouter"
)

type serviceMock struct {
}

func (s *serviceMock) CurrentUser(jwtToken string) (*types.CurrentUser, error) {
	return nil, nil
}

func (s *serviceMock) AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
	return nil, nil
}

func TestGetPing(t *testing.T) {
	srvMock := &serviceMock{}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.GetPing()(w, req, params)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "ok" {
		t.Fatalf("expected: ok, got: %v", body)
	}
}
