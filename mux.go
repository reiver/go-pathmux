package pathmux

import (
	"github.com/reiver/go-pathmatch"

	"net/http"
	"sync"
)

// Mux registers (a number of) handlers (i.e., http.Handler), and producers (i.e., pathmux.Producer)
// to handle (a number of) paths, and patterns.
type Mux struct {
	mutex sync.RWMutex
	pathHandlers map[string]http.Handler
	patternHandlers []internalPatternHandler
	producerPatternHandlers []internalProducerPatternHandler
	directoryHandlers map[string]http.Handler
}

type internalPatternHandler struct {
	Pattern pathmatch.Pattern
	Handler PatternHandler
}

func (receiver *internalPatternHandler) HTTPHandler(keys []string, values []string) http.Handler {

	serveHTTP := func(responseWriter http.ResponseWriter, request *http.Request) {
		var parameterizedRequest = ParameterizedRequest{
			httpRequest: request,
			parameterNames: keys,
			parameters: values,
		}

		receiver.Handler.ServeHTTP(responseWriter, &parameterizedRequest)
	}

	return http.HandlerFunc(serveHTTP)
}

type internalProducerPatternHandler struct {
	Pattern pathmatch.Pattern
	Producer Producer
}
