package windows

import (
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/gtime"
)

type GlobalWindow interface {
	Window
}

func NewGlobalWindow() GlobalWindow {
	gw := &globalWindow{}
	return gw
}

type globalWindow struct{}

func (gw *globalWindow) String() string {
	return "GlobalWindow"
}

func (gw *globalWindow) MaxTimestamp() time.Time {
	return gtime.Max
}

func (gw *globalWindow) Equals(w Window) bool {
	_, ok := w.(*globalWindow)
	if !ok {
		return false
	}

	return true
}
