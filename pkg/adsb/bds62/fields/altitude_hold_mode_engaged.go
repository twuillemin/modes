package fields

import "fmt"

// AltitudeHoldModeEngaged is the Altitude Hold Mode Engaged definition
//
// Specified in Doc 9871 / C.2.3.9.15
type AltitudeHoldModeEngaged byte

const (
	// AHENotEngaged indicates that altitude hold mode is not active
	AHENotEngaged AltitudeHoldModeEngaged = 0
	// AHEngaged indicates that altitude hold mode is active
	AHEngaged AltitudeHoldModeEngaged = 1
)

// ToString returns a basic, but readable, representation of the field
func (status AltitudeHoldModeEngaged) ToString() string {

	switch status {
	case AHENotEngaged:
		return "0 - altitude hold mode is not active"
	case AHEngaged:
		return "1 - altitude hold mode is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadAltitudeHoldModeEngaged reads the AltitudeHoldModeEngaged from a 56 bits data field
func ReadAltitudeHoldModeEngaged(data []byte) AltitudeHoldModeEngaged {
	bits := (data[6] & 0x40) >> 6
	return AltitudeHoldModeEngaged(bits)
}
