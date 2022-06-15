package windowed

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
	"github.com/buck119br/gears/pkg/gears/core/gtype/windows"
)

type Values interface {
	element.Element

	Window() windows.Window
	Panes() []windows.Pane

	Set(windows.Pane, element.Buffer)
	Get(windows.Pane) (element.Buffer, error)
}

func NewValues(w windows.Window, t time.Time) Values {
	v := new(values)
	v.k = common.NewKey(nil)

	v.t = t

	v.w = w
	v.p = make(map[string]windows.Pane)
	v.e = make(map[string]element.Buffer)

	return v
}

type values struct {
	k common.Key

	t time.Time

	w windows.Window
	p map[string]windows.Pane
	e map[string]element.Buffer
}

func (v *values) Type() element.Type {
	return element.WindowedValues
}

func (v *values) Key() common.Key {
	return v.k
}

func (v *values) Timestamp() time.Time {
	return v.t
}

func (v *values) Window() windows.Window {
	return v.w
}

func (v *values) Panes() []windows.Pane {
	ps := make([]windows.Pane, 0, len(v.p))
	for _, p := range v.p {
		ps = append(ps, p)
	}

	return ps
}

func (v *values) Set(p windows.Pane, b element.Buffer) {
	v.p[p.String()] = p
	v.e[p.String()] = b
}

func (v *values) Get(p windows.Pane) (element.Buffer, error) {
	e, ok := v.e[p.String()]
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	return e, nil
}
