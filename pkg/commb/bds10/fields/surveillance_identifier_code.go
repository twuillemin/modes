package fields

import "fmt"

// SurveillanceIdentifierCode is the Surveillance identifier code definition
//
// Specified in Doc 9871 / D.2.4.1
type SurveillanceIdentifierCode byte

const (
	// SurveillanceIdentifierCodeAbsent indicates no surveillance identifier code capability.
	SurveillanceIdentifierCodeAbsent SurveillanceIdentifierCode = 0
	// SurveillanceIdentifierCodeCapable indicates surveillance identifier code capability.
	SurveillanceIdentifierCodeCapable SurveillanceIdentifierCode = 1
)

// ToString returns a basic, but readable, representation of the field
func (sic SurveillanceIdentifierCode) ToString() string {

	switch sic {
	case SurveillanceIdentifierCodeAbsent:
		return "0 - No surveillance identifier code capability"
	case SurveillanceIdentifierCodeCapable:
		return "1 - Surveillance identifier code capability"
	default:
		return fmt.Sprintf("%v - Unknown code", sic)
	}
}

// ReadSurveillanceIdentifierCode reads the SurveillanceIdentifierCode from a 56 bits data field
func ReadSurveillanceIdentifierCode(data []byte) SurveillanceIdentifierCode {
	bits := (data[2] & 0x40) >> 5
	return SurveillanceIdentifierCode(bits)
}
