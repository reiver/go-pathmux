package pathmux

import (
	"github.com/reiver/go-pathmatch"

	"net/http"
)

// Example
//
//	err := mux.HandlePath(handler, "/v1/users/{user_id}/wallets/{wallet_name}")
func (receiver *Mux) HandlePattern(handler http.Handler, pattern string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	var compiledPattern pathmatch.Pattern
	if err := pathmatch.CompileTo(&compiledPattern, pattern); nil != err {
		return err
	}

	var patternHandler internalPatternHandler = internalPatternHandler{
		Pattern: compiledPattern,
		Handler: handler,
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.patternHandlers = append(receiver.patternHandlers, patternHandler)

	return nil
}

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
