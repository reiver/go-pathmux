package pathmux

import (
	"net/http"
)

type PatternHandler interface {
	ServeHTTP(http.ResponseWriter, *ParameterizedRequest)
}

type PatternHandlerFunc func(http.ResponseWriter, *ParameterizedRequest)

var _ PatternHandler = PatternHandlerFunc(nil)

func (fn PatternHandlerFunc) ServeHTTP(rw http.ResponseWriter, pr *ParameterizedRequest) {
	fn(rw, pr)
}
