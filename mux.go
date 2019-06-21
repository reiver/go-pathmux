package pathmux

import (
	"github.com/reiver/go-pathmatch"

	"net/http"
	"sync"
)

// Mux registers (a number of) handlers (i.e., http.Handler) to run (a number of) for a number of paths (which can include patters).
type Mux struct {
	mutex sync.RWMutex
	pathHandlers map[string]http.Handler
	patternHandlers []internalPatternHandler
}

type internalPatternHandler struct {
	Pattern pathmatch.Pattern
	Handler http.Handler
}
