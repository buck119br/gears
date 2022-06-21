package gears

import (
	"sync"

	"github.com/buck119br/gears/pkg/gears/core/gfunc"
	"github.com/buck119br/gears/pkg/gears/core/graph"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
	"github.com/buck119br/gears/pkg/gears/core/gtype/event"
	"github.com/buck119br/gears/pkg/gears/core/gtype/kv"
	"github.com/buck119br/gears/pkg/gears/core/gtype/row"
	"github.com/buck119br/gears/pkg/gears/core/gtype/source"
	"github.com/buck119br/gears/pkg/gears/core/gtype/watermark"
	"github.com/buck119br/gears/pkg/gears/core/gtype/windowed"
	"github.com/buck119br/gears/pkg/gears/log"
)

func Print(df Dataflow, inputs []Dataset) {
	fn := gfunc.NewFunction(doPrint)
	n := df.Graph().AddNode(graph.Write, fn).WithConcurrency(1)
	for _, input := range inputs {
		ib := graph.NewInbound(input.Edge())
		n.AddIn(ib)
	}
}

func doPrint(inputs, outputs []source.Source, args ...any) {
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)

		go func(i source.Source) {
			defer wg.Done()

			for i.HasNext() {
				i.Next()
				elm := i.Current()

				switch elm.Type() {
				case element.KV:
					v := elm.(kv.KV[any])
					log.Infof("kv: [%s]", v)
				case element.Event:
					v := elm.(event.Event)
					log.Infof("event: [%s]", v)
				case element.Row:
					v := elm.(row.Row)
					log.Infof("row: [%s]", v)
				case element.WindowedValues:
					wv := elm.(windowed.Values)
					for _, p := range wv.Panes() {
						b, _ := wv.Get(p)
						log.Infof("pane: [%s] values: [%v]", p, b.GetAll())
					}
				case element.Watermark:
					v := elm.(watermark.Watermark)
					log.Infof("watermark: [%s]", v)
				default:
					log.Errorf("invalid element type: [%s]", elm.Type())
				}
			}
		}(input)
	}

	wg.Wait()
}
