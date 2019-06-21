package pathmux

import (
	"net/http"

	"testing"
)

func TestMuxAsHandler(t *testing.T) {

	var handler http.Handler = new(Mux) // THIS IS THE LINE THAT ACTUALLY MATTERS

	if nil == handler {
		t.Errorf("This should not happen.")
	}
}

