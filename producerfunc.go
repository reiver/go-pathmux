package pathmux

import (
	"net/http"
)

type ProducerFunc func( ...string)http.Handler

func (fn ProducerFunc) Produce(a ...string) http.Handler {
	return fn(a...)
}
