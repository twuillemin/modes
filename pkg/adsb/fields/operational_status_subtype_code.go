package fields

import "fmt"

// OperationalStatusSubtypeCode is the Operational Status Subtype Code definition
//
// Specified in Doc 9871 / B-2-101
type OperationalStatusSubtypeCode byte

const (
	// OSSCAirborne indicates Airborne Status Message
	OSSCAirborne OperationalStatusSubtypeCode = 0
	// OSSCSurface indicates Surface Status Message
	OSSCSurface OperationalStatusSubtypeCode = 1
	// OSSCReserved2 is reserved
	OSSCReserved2 OperationalStatusSubtypeCode = 2
	// OSSCReserved3 is reserved
	OSSCReserved3 OperationalStatusSubtypeCode = 3
	// OSSCReserved4 is reserved
	OSSCReserved4 OperationalStatusSubtypeCode = 4
	// OSSCReserved5 is reserved
	OSSCReserved5 OperationalStatusSubtypeCode = 5
	// OSSCReserved6 is reserved
	OSSCReserved6 OperationalStatusSubtypeCode = 6
	// OSSCReserved7 is reserved
	OSSCReserved7 OperationalStatusSubtypeCode = 7
)

// ToString returns a basic, but readable, representation of the field
func (code OperationalStatusSubtypeCode) ToString() string {

	switch code {
	case OSSCAirborne:
		return "0 - Airborne Status Message"
	case OSSCSurface:
		return "1 - Surface Status Message"
	case OSSCReserved2, OSSCReserved3, OSSCReserved4, OSSCReserved5, OSSCReserved6, OSSCReserved7:
		return fmt.Sprintf("%v - Reserved", code)
	default:
		return fmt.Sprintf("%v - Unknown code", code)
	}
}

// ReadOperationalStatusSubtypeCode reads the OperationalStatusSubtypeCode from a 56 bits data field
func ReadOperationalStatusSubtypeCode(data []byte) OperationalStatusSubtypeCode {
	bits := data[0] & 0x07
	return OperationalStatusSubtypeCode(bits)
}
