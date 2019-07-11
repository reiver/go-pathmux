package pathmux

import (
	"net/http"
)

// ServeHTTP makes *pathmux.Mux fits the http.Handler interface, and makes pathmux.Mux
// a type of "middleware".
func (receiver *Mux) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == receiver {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	url := request.URL
	if nil == url {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	path := url.Path

	var handler http.Handler = receiver.Handler(path)
	if nil == handler {
		http.NotFound(responseWriter, request)
		return
	}

	handler.ServeHTTP(responseWriter, request)
	return
}
