package watermark

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type Watermark interface {
	fmt.Stringer
	element.Element
}

func NewWatermark(i time.Time) Watermark {
	w := &watermark{
		i: i,
	}
	return w
}

type watermark struct {
	i time.Time
}

func (wm *watermark) String() string {
	return fmt.Sprintf("%s", common.AnyToString(wm.i))
}

func (wm *watermark) Type() element.Type {
	return element.Watermark
}

func (wm *watermark) Key() common.Key {
	return nil
}

func (wm *watermark) Timestamp() time.Time {
	return wm.i
}
