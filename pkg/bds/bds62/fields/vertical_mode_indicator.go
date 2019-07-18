package fields

import "fmt"

// VerticalModeIndicator is the Vertical Mode Indicator definition
//
// Specified in Doc 9871 / B.2.3.9.6
type VerticalModeIndicator byte

const (
	// VMIUnknown indicates unknown mode or information unavailable
	VMIUnknown VerticalModeIndicator = 0
	// VMIAcquiringMode indicates "Acquiring" Mode
	VMIAcquiringMode VerticalModeIndicator = 1
	// VMICapturingMode indicates "Capturing" or "Maintaining" Mode
	VMICapturingMode VerticalModeIndicator = 2
	// VMIReserved3 is reserved
	VMIReserved3 VerticalModeIndicator = 3
)

// ToString returns a basic, but readable, representation of the field
func (indicator VerticalModeIndicator) ToString() string {

	switch indicator {
	case VMIUnknown:
		return "0 - unknown mode or information unavailable"
	case VMIAcquiringMode:
		return "1 - \"Acquiring\" Mode"
	case VMICapturingMode:
		return "2 - \"Capturing\" or \"Maintaining\" Mode"
	case VMIReserved3:
		return "3 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", indicator)
	}
}

// ReadVerticalModeIndicator reads the VerticalModeIndicator from a 56 bits data field
func ReadVerticalModeIndicator(data []byte) VerticalModeIndicator {
	bits := (data[1] & 0x06) >> 1
	return VerticalModeIndicator(bits)
}
