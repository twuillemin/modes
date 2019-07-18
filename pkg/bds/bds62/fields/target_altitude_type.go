package fields

import "fmt"

// TargetAltitudeType is the Target Altitude Type definition
//
// Specified in Doc 9871 / B.2.3.9.4
type TargetAltitudeType byte

const (
	// TATReferencedToPressureAltitude indicates referenced to pressure-altitude (flight level)
	TATReferencedToPressureAltitude TargetAltitudeType = 0
	// TATReferencedToBarometricAltitude indicates referenced to barometric corrected altitude (mean sea level)
	TATReferencedToBarometricAltitude TargetAltitudeType = 1
)

// ToString returns a basic, but readable, representation of the field
func (targetType TargetAltitudeType) ToString() string {

	switch targetType {
	case TATReferencedToPressureAltitude:
		return "0 - referenced to pressure-altitude (flight level)"
	case TATReferencedToBarometricAltitude:
		return "1 - referenced to barometric corrected altitude (mean sea level)"
	default:
		return fmt.Sprintf("%v - Unknown code", targetType)
	}
}

// ReadTargetAltitudeType reads the TargetAltitudeType from a 56 bits data field
func ReadTargetAltitudeType(data []byte) TargetAltitudeType {
	bits := (data[1] & 0x40) >> 6
	return TargetAltitudeType(bits)
}
