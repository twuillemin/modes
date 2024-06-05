package fields

import "fmt"

// HorizontalModeIndicator is the Horizontal Mode Indicator definition
//
// Specified in Doc 9871 / B.2.3.9.11
type HorizontalModeIndicator byte

const (
	// HMIUnknown indicates unknown mode or information unavailable
	HMIUnknown HorizontalModeIndicator = 0
	// HMIAcquiringMode indicates "Acquiring" Mode
	HMIAcquiringMode HorizontalModeIndicator = 1
	// HMICapturingMode indicates "Capturing" or "Maintaining" Mode
	HMICapturingMode HorizontalModeIndicator = 2
	// HMIReserved3 is reserved
	HMIReserved3 HorizontalModeIndicator = 3
)

// ToString returns a basic, but readable, representation of the field
func (indicator HorizontalModeIndicator) ToString() string {

	switch indicator {
	case HMIUnknown:
		return "0 - unknown mode or information unavailable"
	case HMIAcquiringMode:
		return "1 - \"Acquiring\" Mode"
	case HMICapturingMode:
		return "2 - \"Capturing\" or \"Maintaining\" Mode"
	case HMIReserved3:
		return "3 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", indicator)
	}
}

// ReadHorizontalModeIndicator reads the HorizontalModeIndicator from a 56 bits data field
func ReadHorizontalModeIndicator(data []byte) HorizontalModeIndicator {
	bits := (data[4] & 0x06) >> 1
	return HorizontalModeIndicator(bits)
}
