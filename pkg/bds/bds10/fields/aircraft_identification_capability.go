package fields

import "fmt"

// AircraftIdentificationCapability is the Aircraft Identification Capability definition
//
// Specified in Doc 9871 / D.2.4.1
type AircraftIdentificationCapability byte

const (
	// IdentificationFromADLP indicates that is coming through ADLP.
	IdentificationFromADLP AircraftIdentificationCapability = 0
	// IdentificationFromSeparateInterface indicates that is coming through a separate interface.
	IdentificationFromSeparateInterface AircraftIdentificationCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (aic AircraftIdentificationCapability) ToString() string {

	switch aic {
	case IdentificationFromADLP:
		return "0 - Aircraft Identification is coming through ADLP"
	case IdentificationFromSeparateInterface:
		return "1 - Aircraft Identification is coming through a separate interface"
	default:
		return fmt.Sprintf("%v - Unknown code", aic)
	}
}

// ReadAircraftIdentificationCapability reads the AircraftIdentificationCapability from a 56 bits data field
func ReadAircraftIdentificationCapability(data []byte) AircraftIdentificationCapability {
	bits := (data[4] & 0x80) >> 7
	return AircraftIdentificationCapability(bits)
}
