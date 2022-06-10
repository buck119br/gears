package gtime

import (
	"errors"
)

var (
	InfLength int64 = -1
)

var (
	ErrInvalidEndPoint = errors.New("end point is invalid, less than start point")
)
