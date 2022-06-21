package gears

import (
	"github.com/buck119br/gears/pkg/gears/core/gtype/source"
)

type DoFunction func(inputs, outputs []source.Source, args ...interface{})
