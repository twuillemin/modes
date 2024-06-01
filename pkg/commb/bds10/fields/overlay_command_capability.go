package fields

import "fmt"

// OverlayCommandCapability is the overlay command capability definition
//
// Specified in Doc 9871 / D.2.4.1
type OverlayCommandCapability byte

const (
	// NoOverlayCommandCapability indicates that this is no overlay command capability.
	NoOverlayCommandCapability OverlayCommandCapability = 0
	// OverlayCommandCapabilityPresent indicates that there is an overlay command capability.
	OverlayCommandCapabilityPresent OverlayCommandCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (occ OverlayCommandCapability) ToString() string {

	switch occ {
	case NoOverlayCommandCapability:
		return "0 - No Overlay Command Capability"
	case OverlayCommandCapabilityPresent:
		return "1 - Overlay Command Capability"
	default:
		return fmt.Sprintf("%v - Unknown code", occ)
	}
}

// ReadOverlayCommandCapability reads the OverlayCommandCapability from a 56 bits data field
func ReadOverlayCommandCapability(data []byte) OverlayCommandCapability {
	bits := (data[1] & 0x02) >> 1
	return OverlayCommandCapability(bits)
}
