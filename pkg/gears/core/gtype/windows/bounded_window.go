package windows

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/gtime"
)

type BoundedWindow interface {
	Window

	Interval() gtime.Interval
}

func NewBoundedWindow(i gtime.Interval) BoundedWindow {
	if !i.IsBounded() {
		panic(fmt.Errorf("unbounded interval"))
	}

	bw := &boundedWindow{
		i: i,
	}

	return bw
}

type boundedWindow struct {
	i gtime.Interval
}

func (bw *boundedWindow) String() string {
	return fmt.Sprintf("BoundedWindow: [%s]", bw.i)
}

func (bw *boundedWindow) MaxTimestamp() time.Time {
	return bw.i.EndPoint()
}

func (bw *boundedWindow) Equals(w Window) bool {
	x, ok := w.(*boundedWindow)
	if !ok {
		return false
	}

	return bw.Interval().Equals(x.Interval())
}

func (bw *boundedWindow) Interval() gtime.Interval {
	return bw.i
}
