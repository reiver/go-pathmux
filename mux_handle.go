package pathmux

import (
	"github.com/reiver/go-pathmatch"

	"net/http"
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
