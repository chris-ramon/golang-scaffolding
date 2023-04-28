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

type testReaderError int

func (testReaderError) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
}

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
	type testCase struct {
		name               string
		srvMock            *serviceMock
		request            *http.Request
		responseWriter     *httptest.ResponseRecorder
		params             httprouter.Params
		handler            func() httprouter.Handle
		header             http.Header
		expectedBody       string
		expectedStatusCode uint
	}

	h := &handlers{}
	header := map[string][]string{
		"Authorization": []string{"Bearer Test-JWT-Token"},
	}

	testCases := []testCase{
		{
			name: "success",
			srvMock: &serviceMock{
				currentUser: func(jwtToken string) (*types.CurrentUser, error) {
					return &types.CurrentUser{
						Username: "test user",
					}, nil
				},
			},
			request:            httptest.NewRequest("GET", "/auth/current-user", nil),
			responseWriter:     httptest.NewRecorder(),
			params:             httprouter.Params{},
			handler:            h.GetCurrentUser,
			header:             header,
			expectedBody:       "test user",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "auth header error",
			srvMock:            &serviceMock{},
			request:            httptest.NewRequest("GET", "/auth/current-user", nil),
			responseWriter:     httptest.NewRecorder(),
			params:             httprouter.Params{},
			handler:            h.GetCurrentUser,
			header:             map[string][]string{},
			expectedBody:       "failed to get authorization header",
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "current user error",
			srvMock: &serviceMock{
				currentUser: func(jwtToken string) (*types.CurrentUser, error) {
					return nil, errors.New("test error")
				},
			},
			request:            httptest.NewRequest("GET", "/auth/current-user", nil),
			responseWriter:     httptest.NewRecorder(),
			params:             httprouter.Params{},
			handler:            h.GetCurrentUser,
			header:             header,
			expectedBody:       "failed to find current user",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			h.service = testCase.srvMock

			for k, v := range testCase.header {
				if len(v) > 0 {
					testCase.request.Header.Set(k, v[0])
				}
			}

			testCase.handler()(testCase.responseWriter, testCase.request, testCase.params)

			if !strings.Contains(testCase.responseWriter.Body.String(), testCase.expectedBody) {
				t.Fatalf("expected: %v, got: %v", testCase.expectedBody, testCase.responseWriter.Body.String())
			}
		})
	}
}

func TestPostSignIn(t *testing.T) {
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
			name: "success",
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
				bytes.NewBuffer([]byte(`{"email":"test@test.com","password":"test-pwd"}`)),
			),
			responseWriter: httptest.NewRecorder(),
			params:         httprouter.Params{},
			handler:        h.PostSignIn,
			expectedBody:   "test user",
		},
		{
			name: "test body error",
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
				testReaderError(0),
			),
			responseWriter: httptest.NewRecorder(),
			params:         httprouter.Params{},
			handler:        h.PostSignIn,
			expectedBody:   "failed to read request body",
		},
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
		{
			name: "test auth user error",
			srvMock: &serviceMock{
				authUser: func(ctx context.Context, username string, pwd string) (*types.CurrentUser, error) {
					return nil, errors.New("test error")
				},
			},
			request: httptest.NewRequest(
				"POST",
				"/auth/sign-in",
				bytes.NewBuffer([]byte(`{"email":"test@test.com","password":"test-pwd"}`)),
			),
			responseWriter: httptest.NewRecorder(),
			params:         httprouter.Params{},
			handler:        h.PostSignIn,
			expectedBody:   "failed to find current user",
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

func TestNewHandlers(t *testing.T) {
	srvMock := &serviceMock{}
	if handlers := NewHandlers(srvMock); handlers == nil {
		t.Fatalf("unexpected nil value")
	}
}
