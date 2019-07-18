package fields

import "fmt"

// Time is the time synchronization definition
//
// Specified in Doc 9871 / A.2.3.2.2
type Time byte

const (
	// TNotSynchronized indicates that the time is not synchronized to UTC
	TNotSynchronized Time = 0
	// TSynchronizedUTC indicates that the time is synchronized to UTC
	TSynchronizedUTC Time = 1
)

// ToString returns a basic, but readable, representation of the field
func (time Time) ToString() string {

	switch time {
	case TNotSynchronized:
		return "0 - not synchronized to UTC"
	case TSynchronizedUTC:
		return "1 - synchronized to UTC"
	default:
		return fmt.Sprintf("%v - Unknown code", time)
	}
}

// ReadTime reads the Time from a 56 bits data field
func ReadTime(data []byte) Time {
	bits := (data[2] & 0x08) >> 3
	return Time(bits)
}
