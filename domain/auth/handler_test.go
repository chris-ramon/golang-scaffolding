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

type testReaderError int

func (testReaderError) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
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
		testReaderError(0),
	)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	h.PostSignIn()(w, req, params)

	expectedBody := "failed to read request body"
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("expected: %v, got: %v", expectedBody, w.Body.String())
	}
}

func TestPostSignIn_JSONUnmarshal(t *testing.T) {
	type testCase struct {
		name           string
		srvMock        *serviceMock
		request        *http.Request
		responseWriter *httptest.ResponseRecorder
		params         httprouter.Params
		handler        func() httprouter.Handle
		expectedBody   string
	}

	h := &handlers{}

	testCases := []testCase{
		{
			name: "test json unmarshal error",
			srvMock: &serviceMock{
				authUser: func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
					return &types.CurrentUser{
						Username: "test user",
					}, nil
				},
			},
			request: httptest.NewRequest(
				"POST",
				"/auth/sign-in",
				bytes.NewBuffer([]byte(`{invalid}`)),
			),
			responseWriter: httptest.NewRecorder(),
			params:         httprouter.Params{},
			handler:        h.PostSignIn,
			expectedBody:   "failed to json unmarshal request body",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			h.service = testCase.srvMock
			testCase.handler()(testCase.responseWriter, testCase.request, testCase.params)
			if !strings.Contains(testCase.responseWriter.Body.String(), testCase.expectedBody) {
				t.Fatalf("expected: %v, got: %v", testCase.expectedBody, testCase.responseWriter.Body.String())
			}
		})
	}
}
