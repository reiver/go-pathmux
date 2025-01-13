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
	producerPatternHandlers []internalProducerPatternHandler
	directoryHandlers map[string]http.Handler
}

type internalProducerPatternHandler struct {
	Pattern pathmatch.Pattern
	Producer Producer
}
