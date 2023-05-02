package auth

import (
	"context"
	"testing"
)

type jwtMock struct {
}

func (j *jwtMock) Generate(ctx context.Context, data map[string]string) (*string, error) {
	jwtToken := ""
	return &jwtToken, nil
}

func (j *jwtMock) Validate(ctx context.Context, jwtToken string) (map[string]string, error) {
	return map[string]string{}, nil
}

func TestCurrentUser(t *testing.T) {
	type testCase struct {
		name string
	}

	testCases := []testCase{
		testCase{
			name: "success",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			srv, err := NewService(&jwtMock{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			authedUser, err := srv.AuthUser(context.Background(), "user", "pwd")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			currentUser, err := srv.CurrentUser(authedUser.JWT)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if currentUser == nil {
				t.Fatalf("unexpected nil value")
			}
		})
	}
}
