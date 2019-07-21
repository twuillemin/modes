package fields

import "fmt"

// GeometricVerticalAccuracy is the Geometric Vertical Accuracy definition
//
// Specified in Doc 9871 / C.2.3.10.8
type GeometricVerticalAccuracy byte

const (
	// GVAUnknownOrGreaterThan150m indicates Unknown or GVA >= 150 meters
	GVAUnknownOrGreaterThan150m GeometricVerticalAccuracy = 0
	// GVALowerThan150m indicates GVA <= 150 meters
	GVALowerThan150m GeometricVerticalAccuracy = 1
	// GVALowerThan45m indicates GVA <= 45 meters
	GVALowerThan45m GeometricVerticalAccuracy = 2
	// GVAReserved3 is reserved
	GVAReserved3 GeometricVerticalAccuracy = 3
)

// ToString returns a basic, but readable, representation of the field
func (baro GeometricVerticalAccuracy) ToString() string {

	switch baro {
	case GVAUnknownOrGreaterThan150m:
		return "0 - Unknown or GVA >= 150 meters"
	case GVALowerThan150m:
		return "0 - GVA <= 150 meters"
	case GVALowerThan45m:
		return "0 - GVA <= 45 meters"
	case GVAReserved3:
		return "0 - reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", baro)
	}
}

// ReadGeometricVerticalAccuracy reads the GeometricVerticalAccuracy from a 56 bits data field
func ReadGeometricVerticalAccuracy(data []byte) GeometricVerticalAccuracy {
	bits := (data[6] & 0xC0) >> 6
	return GeometricVerticalAccuracy(bits)
}
