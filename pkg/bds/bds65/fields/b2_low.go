package fields

import "fmt"

// B2Low is the B2 Low definition
//
// Specified in Doc 9871 / B.2.3.10.3
type B2Low byte

const (
	// B2LGreaterThan70W indicates Greater than or equal to 70 Watts transmit power
	B2LGreaterThan70W B2Low = 0
	// B2LLessThan70W indicates Less than 70 Watts transmit power
	B2LLessThan70W B2Low = 1
)

// ToString returns a basic, but readable, representation of the field
func (status B2Low) ToString() string {

	switch status {
	case B2LGreaterThan70W:
		return "0 - Greater than or equal to 70 Watts transmit power"
	case B2LLessThan70W:
		return "1 - Less than 70 Watts transmit power"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadB2Low reads the B2Low from a 56 bits data field
func ReadB2Low(data []byte) B2Low {
	bits := (data[1] & 0x02) >> 1
	return B2Low(bits)
}
