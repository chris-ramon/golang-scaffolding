// Package api defines types for external communication.
package api

type CurrentUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"jwt"`
}
