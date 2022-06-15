package windows

import (
	"fmt"
	"time"
)

// Window is a grouping of elements into finite buckets.
// Window has a maximum timestamp which means that, at some point, all elements that go into one window will have arrived.
type Window interface {
	fmt.Stringer

	MaxTimestamp() time.Time

	Equals(Window) bool
}
