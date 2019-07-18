package fields

import "fmt"

// PositionSubtypeCode is the Sub type Code definition
//
// Specified in Doc 9871 / B-2-101
type PositionSubtypeCode byte

const (
	// PSCAirborne indicates Airborne Status Message
	PSCAirborne PositionSubtypeCode = 0
	// PSCSurface indicates Surface Status Message
	PSCSurface PositionSubtypeCode = 1
	// PSCReserved2 is reserved
	PSCReserved2 PositionSubtypeCode = 2
	// PSCReserved3 is reserved
	PSCReserved3 PositionSubtypeCode = 3
	// PSCReserved4 is reserved
	PSCReserved4 PositionSubtypeCode = 4
	// PSCReserved5 is reserved
	PSCReserved5 PositionSubtypeCode = 5
	// PSCReserved6 is reserved
	PSCReserved6 PositionSubtypeCode = 6
	// PSCReserved7 is reserved
	PSCReserved7 PositionSubtypeCode = 7
)

// ToString returns a basic, but readable, representation of the field
func (code PositionSubtypeCode) ToString() string {

	switch code {
	case PSCAirborne:
		return "0 - Airborne Status Message"
	case PSCSurface:
		return "1 - Surface Status Message"
	case PSCReserved2, PSCReserved3, PSCReserved4, PSCReserved5, PSCReserved6, PSCReserved7:
		return fmt.Sprintf("%v - Reserved", code)
	default:
		return fmt.Sprintf("%v - Unknown code", code)
	}
}

// ReadPositionSubtypeCode reads the PositionSubtypeCode from a 56 bits data field
func ReadPositionSubtypeCode(data []byte) PositionSubtypeCode {
	bits := data[0] & 0x07
	return PositionSubtypeCode(bits)
}
