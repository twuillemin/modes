package fields

import "fmt"

// ModeSSpecificServicesCapability is the Mode S specific services capability definition
//
// Specified in Doc 9871 / D.2.4.1
type ModeSSpecificServicesCapability byte

const (
	// ModeSSpecificServicesNotPresent indicates that no Mode S specific service is present.
	ModeSSpecificServicesNotPresent ModeSSpecificServicesCapability = 0
	// ModeSSpecificServicesAvailable indicates that at least one Mode S specific service is present.
	ModeSSpecificServicesAvailable ModeSSpecificServicesCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (mssc ModeSSpecificServicesCapability) ToString() string {

	switch mssc {
	case ModeSSpecificServicesNotPresent:
		return "0 - no Mode S specific service is present"
	case ModeSSpecificServicesAvailable:
		return "1 - at least one Mode S specific service is present"
	default:
		return fmt.Sprintf("%v - Unknown code", mssc)
	}
}

// ReadModeSSpecificServicesCapability reads the ModeSSpecificServicesCapability from a 56 bits data field
func ReadModeSSpecificServicesCapability(data []byte) ModeSSpecificServicesCapability {
	bits := (data[3] & 0x80) >> 7
	return ModeSSpecificServicesCapability(bits)
}
