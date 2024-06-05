package fields

import "fmt"

// ContinuationFlag is the continuation definition
//
// Specified in Doc 9871 / D.2.4.1
type ContinuationFlag byte

const (
	// NoNextRegister indicates that this is the last register.
	NoNextRegister ContinuationFlag = 0
	// NextRegisterPresent indicates that subsequent register shall be extracted.
	NextRegisterPresent ContinuationFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (cf ContinuationFlag) ToString() string {

	switch cf {
	case NoNextRegister:
		return "0 - last register"
	case NextRegisterPresent:
		return "1 - subsequent register shall be extracted"
	default:
		return fmt.Sprintf("%v - Unknown code", cf)
	}
}

// ReadContinuationFlag reads the ContinuationFlag from a 56 bits data field
func ReadContinuationFlag(data []byte) ContinuationFlag {
	bits := (data[1] & 0x80) >> 7
	return ContinuationFlag(bits)
}
