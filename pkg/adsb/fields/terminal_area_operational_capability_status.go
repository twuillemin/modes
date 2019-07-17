package fields

import "fmt"

// TerminalAreaOperationalCapabilityStatus is the Terminal Area Operational Capabilities format definition
//
// Specified in Doc 9871 / A.2.3.11.8
type TerminalAreaOperationalCapabilityStatus byte

const (
	// TAOCSReserved0 is reserved
	TAOCSReserved0 TerminalAreaOperationalCapabilityStatus = 0
	// TAOCSReserved1 is reserved
	TAOCSReserved1 TerminalAreaOperationalCapabilityStatus = 1
	// TAOCSReserved2 is reserved
	TAOCSReserved2 TerminalAreaOperationalCapabilityStatus = 2
	// TAOCSReserved3 is reserved
	TAOCSReserved3 TerminalAreaOperationalCapabilityStatus = 3
	// TAOCSReserved4 is reserved
	TAOCSReserved4 TerminalAreaOperationalCapabilityStatus = 4
	// TAOCSReserved5 is reserved
	TAOCSReserved5 TerminalAreaOperationalCapabilityStatus = 5
	// TAOCSReserved6 is reserved
	TAOCSReserved6 TerminalAreaOperationalCapabilityStatus = 6
	// TAOCSReserved7 is reserved
	TAOCSReserved7 TerminalAreaOperationalCapabilityStatus = 7
	// TAOCSReserved8 is reserved
	TAOCSReserved8 TerminalAreaOperationalCapabilityStatus = 8
	// TAOCSReserved9 is reserved
	TAOCSReserved9 TerminalAreaOperationalCapabilityStatus = 9
	// TAOCSReserved10 is reserved
	TAOCSReserved10 TerminalAreaOperationalCapabilityStatus = 10
	// TAOCSReserved11 is reserved
	TAOCSReserved11 TerminalAreaOperationalCapabilityStatus = 11
	// TAOCSReserved12 is reserved
	TAOCSReserved12 TerminalAreaOperationalCapabilityStatus = 12
	// TAOCSReserved13 is reserved
	TAOCSReserved13 TerminalAreaOperationalCapabilityStatus = 13
	// TAOCSReserved14 is reserved
	TAOCSReserved14 TerminalAreaOperationalCapabilityStatus = 14
	// TAOCSReserved15 is reserved
	TAOCSReserved15 TerminalAreaOperationalCapabilityStatus = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities TerminalAreaOperationalCapabilityStatus) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadTerminalAreaOperationalCapabilityStatus reads the TerminalAreaOperationalCapabilityStatus from a 56 bits data field
func ReadTerminalAreaOperationalCapabilityStatus(data []byte) TerminalAreaOperationalCapabilityStatus {
	bits := data[3] & 0x0F
	return TerminalAreaOperationalCapabilityStatus(bits)
}
