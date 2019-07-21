package fields

import "fmt"

// ExtendedSquitterIn is the 1090ES IN (1090 MHz Extended Squitter) definition
//
// Specified in Doc 9871 / C.2.3.10.3
type ExtendedSquitterIn byte

const (
	// ESINoCapability indicates Aircraft has NO 1090ES Receive capability
	ESINoCapability ExtendedSquitterIn = 0
	// ESICapable indicates Aircraft has 1090ES Receive capability
	ESICapable ExtendedSquitterIn = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ExtendedSquitterIn) ToString() string {

	switch status {
	case ESINoCapability:
		return "0 - Aircraft has NO 1090ES Receive capability"
	case ESICapable:
		return "1 - Aircraft has 1090ES Receive capability"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadExtendedSquitterIn reads the ExtendedSquitterIn from a 56 bits data field
func ReadExtendedSquitterIn(data []byte) ExtendedSquitterIn {
	bits := (data[1] & 0x10) >> 4
	return ExtendedSquitterIn(bits)
}
