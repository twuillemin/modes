package fields

import "fmt"

// TargetAltitudeCapability is the Target Altitude Capability definition
//
// Specified in Doc 9871 / B.2.3.9.5
type TargetAltitudeCapability byte

const (
	// TACAltitudeOnly indicates capability for reporting holding altitude only
	TACAltitudeOnly TargetAltitudeCapability = 0
	// TACAltitudeOrAutopilot indicates capability for reporting either holding altitude or autopilot control panel
	// selected altitude
	TACAltitudeOrAutopilot TargetAltitudeCapability = 1
	// TACAltitudeOrAutopilotOrFMS indicates capability for reporting either holding altitude, autopilot control panel
	// selected altitude, or any FMS/RNAV level-off altitude
	TACAltitudeOrAutopilotOrFMS TargetAltitudeCapability = 2
	// TACReserved3 is reserved
	TACReserved3 TargetAltitudeCapability = 3
)

// ToString returns a basic, but readable, representation of the field
func (capability TargetAltitudeCapability) ToString() string {

	switch capability {
	case TACAltitudeOnly:
		return "0 - capability for reporting holding altitude only"
	case TACAltitudeOrAutopilot:
		return "1 - capability for reporting either holding altitude or autopilot control panel	selected altitude"
	case TACAltitudeOrAutopilotOrFMS:
		return "2 - capability for reporting either holding altitude, autopilot control panel selected altitude, or any FMS/RNAV level-off altitude"
	case TACReserved3:
		return "3 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", capability)
	}
}

// ReadTargetAltitudeCapability reads the TargetAltitudeCapability from a 56 bits data field
func ReadTargetAltitudeCapability(data []byte) TargetAltitudeCapability {
	bits := (data[1] & 0x18) >> 3
	return TargetAltitudeCapability(bits)
}
