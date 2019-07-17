package fields

import "fmt"

// EnRouteOperationalCapabilities is the En Route Operational Capabilities
//
// Specified in Doc 9871 / A.2.3.11.3
type EnRouteOperationalCapabilities byte

const (
	// EROCReserved0 is reserved
	EROCReserved0 EnRouteOperationalCapabilities = 0
	// EROCReserved1 is reserved
	EROCReserved1 EnRouteOperationalCapabilities = 1
	// EROCReserved2 is reserved
	EROCReserved2 EnRouteOperationalCapabilities = 2
	// EROCReserved3 is reserved
	EROCReserved3 EnRouteOperationalCapabilities = 3
	// EROCReserved4 is reserved
	EROCReserved4 EnRouteOperationalCapabilities = 4
	// EROCReserved5 is reserved
	EROCReserved5 EnRouteOperationalCapabilities = 5
	// EROCReserved6 is reserved
	EROCReserved6 EnRouteOperationalCapabilities = 6
	// EROCReserved7 is reserved
	EROCReserved7 EnRouteOperationalCapabilities = 7
	// EROCReserved8 is reserved
	EROCReserved8 EnRouteOperationalCapabilities = 8
	// EROCReserved9 is reserved
	EROCReserved9 EnRouteOperationalCapabilities = 9
	// EROCReserved10 is reserved
	EROCReserved10 EnRouteOperationalCapabilities = 10
	// EROCReserved11 is reserved
	EROCReserved11 EnRouteOperationalCapabilities = 11
	// EROCReserved12 is reserved
	EROCReserved12 EnRouteOperationalCapabilities = 12
	// EROCReserved13 is reserved
	EROCReserved13 EnRouteOperationalCapabilities = 13
	// EROCReserved14 is reserved
	EROCReserved14 EnRouteOperationalCapabilities = 14
	// EROCReserved15 is reserved
	EROCReserved15 EnRouteOperationalCapabilities = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities EnRouteOperationalCapabilities) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadEnRouteOperationalCapabilities reads the EnRouteOperationalCapabilities from a 56 bits data field
func ReadEnRouteOperationalCapabilities(data []byte) EnRouteOperationalCapabilities {
	bits := (data[1] & 0xF0) >> 4
	return EnRouteOperationalCapabilities(bits)
}
