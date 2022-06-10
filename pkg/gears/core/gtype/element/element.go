package element

import (
	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/gtime"
)

// Element is the core basic data unit abstraction.
type Element interface {
	Type() Type
	Key() common.Key
	Timestamp() gtime.Instant
}
