package fields

import "fmt"

// PositionOffsetApplied is the Position Offset Applied definition
//
// Specified in Doc 9871 / B.2.3.10.3
type PositionOffsetApplied byte

const (
	// POANotADSB indicates Position transmitted is not the ADS-B position reference point
	POANotADSB PositionOffsetApplied = 0
	// POAIsADSB indicates Position transmitted is the ADS-B position reference point
	POAIsADSB PositionOffsetApplied = 1
)

// ToString returns a basic, but readable, representation of the field
func (status PositionOffsetApplied) ToString() string {

	switch status {
	case POANotADSB:
		return "0 - Position transmitted is not the ADS-B position reference point"
	case POAIsADSB:
		return "1 - Position transmitted is the ADS-B position reference point"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadPositionOffsetApplied reads the PositionOffsetApplied from a 56 bits data field
func ReadPositionOffsetApplied(data []byte) PositionOffsetApplied {
	bits := (data[1] & 0x20) >> 5
	return PositionOffsetApplied(bits)
}
