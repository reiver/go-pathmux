package pathmux_test

import (
	"fmt"
	"net/http"
	"strings"
)

type nopHandler struct {
	Name string
}

func (nopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Nothing here.
}

type nopProducer struct {
	Name string
}

func (receiver nopProducer) Produce(a ...string) http.Handler {
	var s string = fmt.Sprintf("%s [/] %s", receiver.Name, strings.Join(a, ", "))

	return nopHandler{
		Name: s,
	}
}
