package fields

import "fmt"

// TransponderEnhancedProtocolIndicator is the Transponder enhanced protocol indicator definition
//
// Specified in Doc 9871 / D.2.4.1
type TransponderEnhancedProtocolIndicator byte

const (
	// TransponderLevel2To4 indicates a transponder Level 2 to 4.
	TransponderLevel2To4 TransponderEnhancedProtocolIndicator = 0
	// TransponderLevel5 indicates a transponder Level 5.
	TransponderLevel5 TransponderEnhancedProtocolIndicator = 1
)

// ToString returns a basic, but readable, representation of the field
func (tepi TransponderEnhancedProtocolIndicator) ToString() string {

	switch tepi {
	case TransponderLevel2To4:
		return "0 - Transponder Level 2 to 4"
	case TransponderLevel5:
		return "1 - Transponder Level 5"
	default:
		return fmt.Sprintf("%v - Unknown code", tepi)
	}
}

// ReadTransponderEnhancedProtocolIndicator reads the TransponderEnhancedProtocolIndicator from a 56 bits data field
func ReadTransponderEnhancedProtocolIndicator(data []byte) TransponderEnhancedProtocolIndicator {
	bits := data[2] & 0x01
	return TransponderEnhancedProtocolIndicator(bits)
}
