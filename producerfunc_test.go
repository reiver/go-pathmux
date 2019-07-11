package pathmux_test

import (
	"github.com/reiver/go-pathmux"

	"testing"
)

func TestProducerFuncAsProducer(t *testing.T) {

	var producer pathmux.Producer = pathmux.ProducerFunc(nil) // THIS IS THE LINE THAT ACTUALLY MATTERS!

	if nil == producer {
		t.Error("This should never happen.")
	}
}
