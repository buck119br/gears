package watermark

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrNilWatermark = errors.New("watermark must not be nil")
)

const (
	Default EstimatorType = iota
	Manual
	TimestampObserving
)

type EstimatorType int

func (t EstimatorType) String() string {
	switch t {
	case Default:
		return "Default"
	case Manual:
		return "Manual"
	case TimestampObserving:
		return "TimestampObserving"
	default:
		panic(fmt.Errorf("invalid watermark estimator type: [%d]", t))
	}
}

// Estimator is used for estimating output watermarks of a splittable DoFn.
type Estimator interface {
	Type() EstimatorType

	// Current returns estimated output
	// This method must return monotonically increasing watermarks across instances that are constructed from prior state.
	Current() Watermark

	// GetState gets current state of the Estimator instance,
	// which can be used to recreate the Estimator when processing the restriction.
	// The internal state of the estimator must not be mutated by this method.
	// The state returned must not be mutated.
	GetState() Watermark
}

func NewWallTime(w Watermark) Estimator {
	if w == nil {
		panic(ErrNilWatermark)
	}

	wt := &wallTime{
		t: Default,
		w: w,
	}

	return wt
}

// wallTime is an Estimator that tracks wall time.
// Note that this watermark estimator expects wall times of all machines performing the processing to be close to each other.
// Any machine with a wall clock that is far in the past may cause the pipeline to perform poorly while a watermark far in the future may cause records to be marked as late.
type wallTime struct {
	t EstimatorType
	w Watermark
}

func (wt *wallTime) Type() EstimatorType {
	return wt.t
}

func (wt *wallTime) Current() Watermark {
	now := time.Now()
	if now.After(wt.w.Timestamp()) {
		wt.w = NewWatermark(now)
	}
	return wt.w
}

func (wt *wallTime) GetState() Watermark {
	return wt.w
}

// ManualEstimator is an Estimator which is controlled manually.
type ManualEstimator interface {
	Estimator
	// Set sets a timestamp before or at the timestamps of all future elements produced.
	Set(time.Time)
}

func NewManual(w Watermark) ManualEstimator {
	if w == nil {
		panic(ErrNilWatermark)
	}

	m := &manual{
		t: Manual,
		w: w,
	}

	return m
}

type manual struct {
	t       EstimatorType
	w       Watermark
	lastSet Watermark
}

func (m *manual) Type() EstimatorType {
	return m.t
}

func (m *manual) Current() Watermark {
	if m.lastSet != nil && m.lastSet.Timestamp().After(m.w.Timestamp()) {
		m.w = m.lastSet
	}

	return m.w
}

func (m *manual) GetState() Watermark {
	return m.w
}

func (m *manual) Set(i time.Time) {
	m.lastSet = NewWatermark(i)
}

// TimestampObservingEstimator is an Estimator that observes the timestamps of all records output.
type TimestampObservingEstimator interface {
	Estimator
	// Observe updates watermark estimate with the latest output timestamp.
	Observe(time.Time)
}

func NewMonotonicallyIncreasing(w Watermark) TimestampObservingEstimator {
	if w == nil {
		panic(ErrNilWatermark)
	}

	mi := &monotonicallyIncreasing{
		t: TimestampObserving,
		w: w,
	}

	return mi
}

// monotonicallyIncreasing is an Estimator that observes timestamps of records output reporting the timestamp of the last element seen as the current
// Note that this watermark estimator expects output timestamps in monotonically increasing order.
// If they are not, then the watermark will advance based upon the last observed timestamp as long as it is greater than any previously reported
type monotonicallyIncreasing struct {
	t            EstimatorType
	w            Watermark
	lastObserved Watermark
}

func (mi *monotonicallyIncreasing) Type() EstimatorType {
	return mi.t
}

func (mi *monotonicallyIncreasing) Current() Watermark {
	if mi.lastObserved != nil && mi.lastObserved.Timestamp().After(mi.w.Timestamp()) {
		mi.w = mi.lastObserved
	}

	return mi.w
}

func (mi *monotonicallyIncreasing) GetState() Watermark {
	return mi.w
}

func (mi *monotonicallyIncreasing) Observe(i time.Time) {
	mi.lastObserved = NewWatermark(i)
}
