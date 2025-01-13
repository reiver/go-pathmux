package pathmux

import (
	"net/http"
)

type ParameterizedRequest struct {
	httpRequest *http.Request
	parameterNames[]string
	parameters []string
}

func (receiver ParameterizedRequest) HTTPRequest() *http.Request {
	return receiver.httpRequest
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
