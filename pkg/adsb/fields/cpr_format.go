package fields

import "fmt"

// CPRFormat is the CPR (Compact Position Reporting) format definition
//
// Specified in Doc 9871 / A.2.3.2.1
type CPRFormat byte

const (
	// CPRFormatEven denotes an even format coding
	CPRFormatEven CPRFormat = 0
	// SSPermanentAlert indicates an odd format coding
	CPRFormatOdd CPRFormat = 1
)

// ToString returns a basic, but readable, representation of the field
func (format CPRFormat) ToString() string {

	switch format {
	case CPRFormatEven:
		return "0 - even format coding"
	case CPRFormatOdd:
		return "1 - odd format coding"
	default:
		return fmt.Sprintf("%v - Unknown code", format)
	}
}

// ReadCPRFormat read the CPRFormat from a 56 bits data field
func ReadCPRFormat(data []byte) CPRFormat {
	bits := (data[2] & 0x04) >> 2
	return CPRFormat(bits)
}
