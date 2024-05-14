package route

import (
	"net/http"
)

type Route struct {
	HTTPMethod string
	Path       string
	Handler    http.HandlerFunc
}
