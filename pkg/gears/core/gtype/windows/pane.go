package windows

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

type Pane interface {
	fmt.Stringer

	Window() Window
	Key() common.Key
}

func NewPane(w Window, key common.Key) Pane {
	p := &pane{
		w: w,
		k: key,
	}

	return p
}

type pane struct {
	w Window
	k common.Key
}

func (p *pane) String() string {
	return fmt.Sprintf("Window: [%s], Key: [%s]", p.w, p.k)
}

func (p *pane) Window() Window {
	return p.w
}

func (p *pane) Key() common.Key {
	return p.k
}
