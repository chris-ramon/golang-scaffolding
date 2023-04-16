package route

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	HTTPMethod string
	Path       string
	Handler    httprouter.Handle
}
