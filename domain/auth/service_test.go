package auth

import (
	"context"
	"errors"
	"testing"
)

type jwtMock struct {
	generate func(ctx context.Context, data map[string]string) (*string, error)
	validate func(ctx context.Context, jwtToken string) (map[string]string, error)
}

func (j *jwtMock) Generate(ctx context.Context, data map[string]string) (*string, error) {
	return j.generate(ctx, data)
}

func (j *jwtMock) Validate(ctx context.Context, jwtToken string) (map[string]string, error) {
	return j.validate(ctx, jwtToken)
}

func TestCurrentUser(t *testing.T) {
	type testCase struct {
		name                     string
		generate                 func(ctx context.Context, data map[string]string) (*string, error)
		validate                 func(ctx context.Context, jwtToken string) (map[string]string, error)
		currentUserExpectedError error
		authUserExpectedError    error
	}

	testError := errors.New("test error")

	testCases := []testCase{
		testCase{
			name: "success",
			generate: func(ctx context.Context, data map[string]string) (*string, error) {
				jwtToken := ""
				return &jwtToken, nil
			},
			validate: func(ctx context.Context, jwtToken string) (map[string]string, error) {
				return nil, nil
			},
		},
		testCase{
			name: "validate error",
			generate: func(ctx context.Context, data map[string]string) (*string, error) {
				jwtToken := ""
				return &jwtToken, nil
			},
			validate: func(ctx context.Context, jwtToken string) (map[string]string, error) {
				return nil, testError
			},
			currentUserExpectedError: testError,
		},
		testCase{
			name: "auth user error",
			generate: func(ctx context.Context, data map[string]string) (*string, error) {
				return nil, testError
			},
			validate: func(ctx context.Context, jwtToken string) (map[string]string, error) {
				return nil, nil
			},
			authUserExpectedError: testError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			srv, err := NewService(&jwtMock{
				generate: testCase.generate,
				validate: testCase.validate,
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			authedUser, err := srv.AuthUser(context.Background(), "user", "pwd")
			if testCase.authUserExpectedError != nil {
				if err != testCase.authUserExpectedError {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if authedUser != nil {
				currentUser, err := srv.CurrentUser(context.Background(), authedUser.JWT)
				if testCase.currentUserExpectedError != nil {
					if err != testCase.currentUserExpectedError {
						t.Fatalf("unexpected error: %v", err)
					}
				} else {
					if currentUser == nil {
						t.Fatalf("unexpected nil value")
					}
				}
			}
		})
	}
}
