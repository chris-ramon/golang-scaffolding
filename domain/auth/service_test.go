package auth

import (
	"context"
	"testing"
)

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
			srv, err := NewService()
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
