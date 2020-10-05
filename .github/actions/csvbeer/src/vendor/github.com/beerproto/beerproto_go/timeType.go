package beerproto_go

import "time"

// Seconds returns the duration as a floating point number of seconds.
func (t *TimeType) Seconds() int64 {
	switch t.Unit {
	case TimeType_SEC:
		return t.Value
	case TimeType_MIN:
		return t.Value * 60
	case TimeType_HR:
		return t.Value * 60 * 60
	case TimeType_DAY:
		return t.Value * 60 * 60 * 24
	case TimeType_WEEK:
		return t.Value * 60 * 60 * 24 * 7
	}

	return 0
}

// Milliseconds returns the duration as an integer millisecond count.
func (t *TimeType) Milliseconds() int64 {
	return t.Seconds() * int64(time.Millisecond)
}
