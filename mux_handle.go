package pathmux

import (
	"net/http"

	pkgpath "github.com/reiver/go-path"
	"github.com/reiver/go-pathmatch"
)

// HandlePattern registers a producer (i.e., pathmux.Producer) that will create
// handlers (i.e., http.Handler) to handle patterns.
//
// Example
//
//	err := mux.HandlePath(producer, "/v1/users/{user_id}/wallets/{wallet_name}")
func (receiver *Mux) HandlePattern(producer Producer, pattern string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == producer {
		return errNilProducer
	}

	var compiledPattern pathmatch.Pattern
	if err := pathmatch.CompileTo(&compiledPattern, pattern); nil != err {
		return err
	}

	var patternHandler internalPatternHandler = internalPatternHandler{
		Pattern: compiledPattern,
		Producer: producer,
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.patternHandlers = append(receiver.patternHandlers, patternHandler)

	return nil
}

// HandlePath registers a handler (i.e., http.Handler) to handle paths.
//
// Example
//
//	err := mux.HandlePath(handler, "/v1/help")
func (receiver *Mux) HandlePath(handler http.Handler, path string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.pathHandlers {
		receiver.pathHandlers = map[string]http.Handler{}
	}

	if _, found := receiver.pathHandlers[path]; found {
		return errFound
	}

	receiver.pathHandlers[path] = handler

	return nil
}

// HandleDirectory registers a handler (i.e., http.Handler) to handle a directory and everything under it.
//
// Example
//
//	err := mux.HandleDirectory(handler, "/wiki")
//
// Note that if you register "/wiki", then it will also (usually) receive requests for
// "/wiki/", "/wiki/something", "/wiki/image.jpeg", "/wiki/once/twice/thrice/fource", etc.
//
// The exceptions being:
// (1) if there is already a path-handler,
// (2) if there is already a pattern-handler, and
// (3) if there is a lower directory-handler.
//
// Also, not that HandleDirectory will remove trailing-slashes in your path.
// So this:
//
//	err := mux.HandleDirectory(handler, "/wiki/")
//
// ... will actually be registered as:
//
//	err := mux.HandleDirectory(handler, "/wiki")
func (receiver *Mux) HandleDirectory(handler http.Handler, path string) error {
	if nil == receiver{
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	path = pkgpath.Canonical("" + path)
	path = pkgpath.RemoveTrailingSeparators(path)
	if len(path) <= 0 || '/' != path[0] {
		return errBadPath
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.directoryHandlers {
		receiver.directoryHandlers = map[string]http.Handler{}
	}

	if _, found := receiver.directoryHandlers[path]; found {
		return errFound
	}

	receiver.directoryHandlers[path] = handler

	return nil
}
