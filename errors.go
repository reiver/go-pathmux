package pathmux

import (
	"errors"
)

var (
	errFound         = errors.New("pathmux: Found")
	errNilHandler    = errors.New("pathmux: Nil Handler")
	errNilReceiver   = errors.New("pathmux: Nil Receiver")
)
