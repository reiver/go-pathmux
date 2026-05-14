package pathmux

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	errBadPath       = erorr.Error("pathmux: bad path")
	errFound         = erorr.Error("pathmux: found")
	errNilHandler    = erorr.Error("pathmux: nil handler")
	errNilProducer   = erorr.Error("pathmux: nil producer")
	errNilReceiver   = erorr.Error("pathmux: nil receiver")
)
