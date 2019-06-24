package pathmux_test

import (
	"net/http"
)

type nopHandler struct {
	Name string
}

func (nopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Nothing here.
}
