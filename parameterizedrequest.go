package pathmux

import (
	"net/http"
)

type ParameterizedRequest struct {
	httpRequest    *http.Request
	parameterNames []string
	parameters     []string
}

// ParameterizedRequest appends a parameter inside of a [ParameterizedRequest],
// and appends the index of the new parameter.
//
// The name passed to AppendParameter is the same name passed to [ParameterByName].
// The index returned by AppendParameter is the same index passed to [ParameterByIndex].
//
// You would probably only use this if you are writing a tests.
func (receiver *ParameterizedRequest) AppendParameter(name string, value string) int {
	if nil == receiver {
		panic(errNilReceiver)
	}

	if nil == receiver.parameterNames {
		receiver.parameterNames = []string{}
	}
	if nil == receiver.parameters {
		receiver.parameters = []string{}
	}

	receiver.parameterNames = append(receiver.parameterNames, name)
	receiver.parameters     = append(receiver.parameters, value)

	return len(receiver.parameters) - 1
}

// HTTPRequest returns the internal [http.Request] that [ParameterizedRequest] wraps.
func (receiver ParameterizedRequest) HTTPRequest() *http.Request {
	return receiver.httpRequest
}

// SetHTTPRequest sets the [http.Request] that is inside of a [ParameterizedRequest].
//
// You would probably only use this if you are writing a tests.
func (receiver *ParameterizedRequest) SetHTTPRequest(req *http.Request) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.httpRequest = req
}

func (receiver ParameterizedRequest) ParameterByIndex(index int) (string, bool) {
	if index < 0 {
		var nada string
		return nada, false
	}
	if len(receiver.parameters) <= index {
		var nada string
		return nada, false
	}

	return receiver.parameters[index], true
}

func (receiver ParameterizedRequest) ParameterByName(name string) (string, bool) {
	index, found := parameterNameIndex(name, receiver.parameterNames...)
	if !found {
		var nada string
		return nada, false
	}
	if index < 0 || len(receiver.parameterNames) <= index {
		var nada string
		return nada, false
	}

	return receiver.ParameterByIndex(index)
}
