package fields

import "fmt"

// SubTypeCode is the Sub type Code definition
//
// Specified in Doc 9871 / B-2-101
type SubTypeCode byte

const (
	// STCAirborne indicates Airborne Status Message
	STCAirborne SubTypeCode = 0
	// STCSurface indicates Surface Status Message
	STCSurface SubTypeCode = 1
	// STCReserved2 is reserved
	STCReserved2 SubTypeCode = 2
	// STCReserved3 is reserved
	STCReserved3 SubTypeCode = 3
	// STCReserved4 is reserved
	STCReserved4 SubTypeCode = 4
	// STCReserved5 is reserved
	STCReserved5 SubTypeCode = 5
	// STCReserved6 is reserved
	STCReserved6 SubTypeCode = 6
	// STCReserved7 is reserved
	STCReserved7 SubTypeCode = 7
)

// ToString returns a basic, but readable, representation of the field
func (code SubTypeCode) ToString() string {

	switch code {
	case STCAirborne:
		return "0 - Airborne Status Message"
	case STCSurface:
		return "1 - Surface Status Message"
	case STCReserved2, STCReserved3, STCReserved4, STCReserved5, STCReserved6, STCReserved7:
		return fmt.Sprintf("%v - Reserved", code)
	default:
		return fmt.Sprintf("%v - Unknown code", code)
	}
}

// ReadSubTypeCode reads the SubTypeCode from a 56 bits data field
func ReadSubTypeCode(data []byte) SubTypeCode {
	bits := data[0] & 0x07
	return SubTypeCode(bits)
}
