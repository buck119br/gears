package source

import (
	"testing"
	"time"

	. "gopkg.in/check.v1"

	"github.com/buck119br/gears/pkg/gears/core/gtype/watermark"
)

func TestSource(t *testing.T) { TestingT(t) }

type SourceSuite struct {
}

var _ = Suite(&SourceSuite{})

func (ss *SourceSuite) TestLocalSource(c *C) {
	s, err := NewBuilder().WithOption(NewOption().WithMode(Local)).Build()
	c.Assert(err, Equals, nil)

	go func() {
		for s.HasNext() {
			s.Next()
			elm := s.Current()
			c.Logf("source consumer received element: [%s]", elm)
		}
		c.Logf("source consumer will be closed")
	}()

	err = s.Emit(watermark.NewWatermark(time.Now()))
	c.Assert(err, Equals, nil)

	c.Logf("source will be closed")
	err = s.Close()
	c.Assert(err, Equals, nil)
	c.Logf("source is closed")

	err = s.Emit(watermark.NewWatermark(time.Now()))
	c.Assert(err, Equals, ErrClosed)

	time.Sleep(100 * time.Millisecond)
}
