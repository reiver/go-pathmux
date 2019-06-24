package pathmux

import (
	"net/http"
)

// Handler returns the http.Handler that would deal with ‘path’.
func (receiver *Mux) Handler(path string) http.Handler {

	var handler http.Handler

	if nil == handler {
		handler = receiver.PathHandler(path)
	}

	if nil == handler {
		handler = receiver.PatternHandler(path)
	}

	return handler
}

// PathHandler exists for optimization, and debugging purposes.
//
// In most cases you should instead use pathmux.Mux.Handler().
func (receiver *Mux) PathHandler(path string) http.Handler {
	if nil == receiver {
		return nil
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	var handler http.Handler
	{
		var found bool

		handler, found = receiver.pathHandlers[path]
		if !found {
			return nil
		}
	}

	return handler
}

// PatternHandler exists for optimization, and debugging purposes.
//
// In most cases you should instead use pathmux.Mux.Handler().
func (receiver *Mux) PatternHandler(path string) http.Handler {
	if nil == receiver {
		return nil
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	var handler http.Handler
	for _, patternHandler := range receiver.patternHandlers {
		matched, err := patternHandler.Pattern.Match(path)
		if nil != err {
//@TODO
			continue
		}
		if !matched {
			continue
		}

		handler = patternHandler.Handler
		if nil == handler {
			return nil
		}
		break
	}
	if nil == handler {
		return nil
	}

	return handler
}
