package pathmux

import (
	"errors"
)

var (
	errFound         = errors.New("pathmux: Found")
	errNilHandler    = errors.New("pathmux: Nil Handler")
	errNilProducer   = errors.New("pathmux: Nil Producer")
	errNilReceiver   = errors.New("pathmux: Nil Receiver")
)
