package fields

import (
	"fmt"
)

// ThreatTypeIndicator indicates the type of information contained in the TID part of the message
type ThreatTypeIndicator int

const (
	// ThreatTypeNoIdentity signifies No identity data in TID
	ThreatTypeNoIdentity ThreatTypeIndicator = 0
	// ThreatTypeModeS signifies that the TID contains a Mode S transponder address
	ThreatTypeModeS ThreatTypeIndicator = 1
	// ThreatTypeAltitudeRangeBearing signifies that the TID contains altitude, range and bearing data
	ThreatTypeAltitudeRangeBearing ThreatTypeIndicator = 2
	// ThreatTypeReserved3 is not assigned
	ThreatTypeReserved3 ThreatTypeIndicator = 3
)

// ToString returns a basic, but readable, representation of the field
func (indicator ThreatTypeIndicator) ToString() string {
	switch indicator {
	case ThreatTypeNoIdentity:
		return "0 - No identity data in TID"
	case ThreatTypeModeS:
		return "1 - TID contains a Mode S transponder address"
	case ThreatTypeAltitudeRangeBearing:
		return "2 - TID contains altitude, range and bearing data"
	case ThreatTypeReserved3:
		return "3 - Not assigned"
	default:
		return fmt.Sprintf("%v - Unknown code", indicator)
	}
}

// ReadThreatTypeIndicator reads the ThreatTypeIndicator
func ReadThreatTypeIndicator(data []byte) ThreatTypeIndicator {

	bits := (data[2] & 0x0B) >> 2
	return ThreatTypeIndicator(bits)
}
