package fields

import "fmt"

// NavigationIntegrityCodeSupplementB is the Navigation Integrity Code Supplement B
//
// Specified in Doc 9871 / C.2.3.2.5
type NavigationIntegrityCodeSupplementB byte

const (
	// NICBZero indicates a value of 0
	NICBZero NavigationIntegrityCodeSupplementB = 0
	// NICBOne indicates a value of one
	NICBOne NavigationIntegrityCodeSupplementB = 1
)

// ToString returns a basic, but readable, representation of the field
func (nicb NavigationIntegrityCodeSupplementB) ToString() string {

	switch nicb {
	case NICBZero:
		return "0"
	case NICBOne:
		return "1"
	default:
		return fmt.Sprintf("%v - Unknown code", nicb)
	}
}

// ReadNavigationIntegritySupplementB reads the NavigationIntegrityCodeSupplementB from a 56 bits data field
func ReadNavigationIntegritySupplementB(data []byte) NavigationIntegrityCodeSupplementB {
	bits := data[0] & 0x01
	return NavigationIntegrityCodeSupplementB(bits)
}
