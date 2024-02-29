package pathmux

import (
	"errors"
)

var (
	errBadPath       = errors.New("pathmux: bad path")
	errFound         = errors.New("pathmux: found")
	errNilHandler    = errors.New("pathmux: nil handler")
	errNilProducer   = errors.New("pathmux: nil producer")
	errNilReceiver   = errors.New("pathmux: nil receiver")
)
