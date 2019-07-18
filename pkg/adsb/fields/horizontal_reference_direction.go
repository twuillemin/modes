package fields

import "fmt"

// HorizontalReferenceDirection is the Horizontal Reference Direction definition
//
// Specified in Doc 9871 / B.2.3.10.13
type HorizontalReferenceDirection byte

const (
	// HRDTrueNorth indicates True North
	HRDTrueNorth HorizontalReferenceDirection = 0
	// HRDMagneticNorth indicates Magnetic North
	HRDMagneticNorth HorizontalReferenceDirection = 1
)

// ToString returns a basic, but readable, representation of the field
func (direction HorizontalReferenceDirection) ToString() string {

	switch direction {
	case HRDTrueNorth:
		return "0 - true north"
	case HRDMagneticNorth:
		return "1 - magnetic north"
	default:
		return fmt.Sprintf("%v - Unknown code", direction)
	}
}

// ReadHorizontalReferenceDirection reads the HorizontalReferenceDirection from a 56 bits data field
func ReadHorizontalReferenceDirection(data []byte) HorizontalReferenceDirection {
	bits := (data[6] & 0x04) >> 2
	return HorizontalReferenceDirection(bits)
}
