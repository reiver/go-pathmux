package pathmux

import (
	"net/http"
)

// Handler returns the http.Handler that would deal with ‘path’.
func (receiver *Mux) Handler(path string) http.Handler {

	var handler http.Handler

	if nil == handler {
		handler = receiver.pathHandler(path)
	}

	if nil == handler {
		handler = receiver.patternHandler(path)
	}

	return handler
}

func (receiver *Mux) pathHandler(path string) http.Handler {
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

func (receiver *Mux) patternHandler(path string) http.Handler {
	if nil == receiver {
		return nil
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	var handler http.Handler
	for _, patternHandler := range receiver.patternHandlers {
		var producer Producer = patternHandler.Producer
		if nil == producer {
			continue
		}

		var matched bool
		var matches []string
		{
			var err error

			matched, err = patternHandler.Pattern.FindAndLoad(path, &matches)
			if nil != err {
//@TODO
				continue
			}
		}
		if !matched {
			continue
		}

		handler = producer.Produce(matches...)
		break
	}

	return handler
}
