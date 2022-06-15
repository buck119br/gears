package element

import (
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

// Element is the core basic data unit abstraction.
type Element interface {
	Type() Type
	Key() common.Key
	Timestamp() time.Time
}
