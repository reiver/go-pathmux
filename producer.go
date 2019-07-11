package pathmux

import (
	"net/http"
)

type Producer interface {
	Produce(...string) http.Handler
}
