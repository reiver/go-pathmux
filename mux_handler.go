package pathmux

import (
	"net/http"

	pkgpath "github.com/reiver/go-path"
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

	if nil == handler {
		handler = receiver.directoryHandler(path)
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
	for _, patternHandler := range receiver.producerPatternHandlers {
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

func (receiver *Mux) directoryHandler(path string) http.Handler {
	if nil == receiver {
		return nil
	}

	path = pkgpath.Canonical("/" + path)
	path = pkgpath.RemoveTrailingSeparators(path)

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	var handler http.Handler
	for {
		var found bool

		handler, found = receiver.directoryHandlers[path]
		if found {
			break
		}

		if "/" == path {
			break
		}
		path = pkgpath.Parent(path)
		path = pkgpath.RemoveTrailingSeparators(path)
	}

	return handler
}
