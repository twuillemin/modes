package fields

import "fmt"

// TerminalAreaOperationalCapabilities is the Terminal Area Operational Capabilities format definition
//
// Specified in Doc 9871 / A.2.3.11.4
type TerminalAreaOperationalCapabilities byte

const (
	// TAOCReserved0 is reserved
	TAOCReserved0 TerminalAreaOperationalCapabilities = 0
	// TAOCReserved1 is reserved
	TAOCReserved1 TerminalAreaOperationalCapabilities = 1
	// TAOCReserved2 is reserved
	TAOCReserved2 TerminalAreaOperationalCapabilities = 2
	// TAOCReserved3 is reserved
	TAOCReserved3 TerminalAreaOperationalCapabilities = 3
	// TAOCReserved4 is reserved
	TAOCReserved4 TerminalAreaOperationalCapabilities = 4
	// TAOCReserved5 is reserved
	TAOCReserved5 TerminalAreaOperationalCapabilities = 5
	// TAOCReserved6 is reserved
	TAOCReserved6 TerminalAreaOperationalCapabilities = 6
	// TAOCReserved7 is reserved
	TAOCReserved7 TerminalAreaOperationalCapabilities = 7
	// TAOCReserved8 is reserved
	TAOCReserved8 TerminalAreaOperationalCapabilities = 8
	// TAOCReserved9 is reserved
	TAOCReserved9 TerminalAreaOperationalCapabilities = 9
	// TAOCReserved10 is reserved
	TAOCReserved10 TerminalAreaOperationalCapabilities = 10
	// TAOCReserved11 is reserved
	TAOCReserved11 TerminalAreaOperationalCapabilities = 11
	// TAOCReserved12 is reserved
	TAOCReserved12 TerminalAreaOperationalCapabilities = 12
	// TAOCReserved13 is reserved
	TAOCReserved13 TerminalAreaOperationalCapabilities = 13
	// TAOCReserved14 is reserved
	TAOCReserved14 TerminalAreaOperationalCapabilities = 14
	// TAOCReserved15 is reserved
	TAOCReserved15 TerminalAreaOperationalCapabilities = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities TerminalAreaOperationalCapabilities) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadTerminalAreaOperationalCapabilities reads the TerminalAreaOperationalCapabilities from a 56 bits data field
func ReadTerminalAreaOperationalCapabilities(data []byte) TerminalAreaOperationalCapabilities {
	bits := data[1] & 0x0F
	return TerminalAreaOperationalCapabilities(bits)
}
