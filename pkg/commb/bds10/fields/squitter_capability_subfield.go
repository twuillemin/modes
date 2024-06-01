package fields

import "fmt"

// SquitterCapabilitySubfield is the Squitter capability subfield definition
//
// Specified in Doc 9871 / D.2.4.1
type SquitterCapabilitySubfield byte

const (
	// Registers05And06NotUpdatedRecently indicates that the registers 05 and 06 have not been updated recently.
	Registers05And06NotUpdatedRecently SquitterCapabilitySubfield = 0
	// Registers05And06UpdatedRecently indicates that the registers 05 and 06 have been updated within the last ten,
	//plus or minus one, seconds.
	Registers05And06UpdatedRecently SquitterCapabilitySubfield = 1
)

// ToString returns a basic, but readable, representation of the field
func (scs SquitterCapabilitySubfield) ToString() string {

	switch scs {
	case Registers05And06NotUpdatedRecently:
		return "0 - Both registers 05 and 06 have not been updated recently"
	case Registers05And06UpdatedRecently:
		return "1 - Both registers 05 and 06 have been updated within the last ten, plus or minus one, seconds"
	default:
		return fmt.Sprintf("%v - Unknown code", scs)
	}
}

// ReadSquitterCapabilitySubfield reads the SquitterCapabilitySubfield from a 56 bits data field
func ReadSquitterCapabilitySubfield(data []byte) SquitterCapabilitySubfield {
	bits := (data[4] & 0x40) >> 6
	return SquitterCapabilitySubfield(bits)
}
