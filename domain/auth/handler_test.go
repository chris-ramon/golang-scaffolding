package auth

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/types"
	"github.com/julienschmidt/httprouter"
)

type serviceMock struct {
	currentUser func(jwtToken string) (*types.CurrentUser, error)
	authUser    func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error)
}

func (s *serviceMock) CurrentUser(jwtToken string) (*types.CurrentUser, error) {
	return s.currentUser(jwtToken)
}

func (s *serviceMock) AuthUser(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
	return s.authUser(ctx, username, pwd)
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

func TestGetCurrentUser(t *testing.T) {
	srvMock := &serviceMock{
		currentUser: func(jwtToken string) (*types.CurrentUser, error) {
			return &types.CurrentUser{
				Username: "test user",
			}, nil
		},
		authUser: func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
			return nil, nil
		},
	}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest("GET", "/auth/current-user", nil)
	req.Header.Set("Authorization", "Bearer Test-JWT-Token")
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.GetCurrentUser()(w, req, params)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "test user" {
		t.Fatalf("expected: ok, got: %s", body)
	}
}

func TestGetCurrentUser_AuthHeaderError(t *testing.T) {
	srvMock := &serviceMock{}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest("GET", "/auth/current-user", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.GetCurrentUser()(w, req, params)

	expectedBody := "failed to get authorization header"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("expected: %v, got: %v", expectedBody, w.Body.String())
	}

	expectedStatusCode := http.StatusInternalServerError
	if w.Code != expectedStatusCode {
		t.Fatalf("expected: %v, got: %v", expectedStatusCode, w.Code)
	}
}

func TestGetCurrentUser_CurrentUserError(t *testing.T) {
	srvMock := &serviceMock{
		currentUser: func(jwtToken string) (*types.CurrentUser, error) {
			return nil, errors.New("test error")
		},
	}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest("GET", "/auth/current-user", nil)
	req.Header.Set("Authorization", "Bearer Test-JWT-Token")
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.GetCurrentUser()(w, req, params)

	expectedBody := "failed to find current user"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("expected: %v, got: %v", expectedBody, w.Body.String())
	}

	expectedStatusCode := http.StatusInternalServerError
	if w.Code != expectedStatusCode {
		t.Fatalf("expected: %v, got: %v", expectedStatusCode, w.Code)
	}
}

func TestPostSignIn(t *testing.T) {
	srvMock := &serviceMock{
		authUser: func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
			return &types.CurrentUser{
				Username: "test user",
			}, nil
		},
	}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest(
		"POST",
		"/auth/sign-in",
		bytes.NewBuffer([]byte(`{"email":"test@test.com","password":"test-pwd"}`)),
	)
	req.Header.Set("Authorization", "Bearer Test-JWT-Token")
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.PostSignIn()(w, req, params)

	body, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != "test user" {
		t.Fatalf("expected: ok, got: %s", body)
	}
}

func TestPostSignIn_BodyError(t *testing.T) {
	srvMock := &serviceMock{
		authUser: func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
			return &types.CurrentUser{
				Username: "test user",
			}, nil
		},
	}

	h := &handlers{
		service: srvMock,
	}

	req := httptest.NewRequest(
		"POST",
		"/auth/sign-in",
		bytes.NewBuffer([]byte(`{invalid}`)),
	)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.PostSignIn()(w, req, params)

	expectedBody := "failed to json unmarshal request body"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("expected: %v, got: %v", expectedBody, w.Body.String())
	}
}
