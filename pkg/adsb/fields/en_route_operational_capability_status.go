package fields

import "fmt"

// EnRouteOperationalCapabilityStatus is the En Route Operational Capability Status
//
// Specified in Doc 9871 / A.2.3.11.7
type EnRouteOperationalCapabilityStatus byte

const (
	// EROCSReserved0 is reserved
	EROCSReserved0 EnRouteOperationalCapabilityStatus = 0
	// EROCSReserved1 is reserved
	EROCSReserved1 EnRouteOperationalCapabilityStatus = 1
	// EROCSReserved2 is reserved
	EROCSReserved2 EnRouteOperationalCapabilityStatus = 2
	// EROCSReserved3 is reserved
	EROCSReserved3 EnRouteOperationalCapabilityStatus = 3
	// EROCSReserved4 is reserved
	EROCSReserved4 EnRouteOperationalCapabilityStatus = 4
	// EROCSReserved5 is reserved
	EROCSReserved5 EnRouteOperationalCapabilityStatus = 5
	// EROCSReserved6 is reserved
	EROCSReserved6 EnRouteOperationalCapabilityStatus = 6
	// EROCSReserved7 is reserved
	EROCSReserved7 EnRouteOperationalCapabilityStatus = 7
	// EROCSReserved8 is reserved
	EROCSReserved8 EnRouteOperationalCapabilityStatus = 8
	// EROCSReserved9 is reserved
	EROCSReserved9 EnRouteOperationalCapabilityStatus = 9
	// EROCSReserved10 is reserved
	EROCSReserved10 EnRouteOperationalCapabilityStatus = 10
	// EROCSReserved11 is reserved
	EROCSReserved11 EnRouteOperationalCapabilityStatus = 11
	// EROCSReserved12 is reserved
	EROCSReserved12 EnRouteOperationalCapabilityStatus = 12
	// EROCSReserved13 is reserved
	EROCSReserved13 EnRouteOperationalCapabilityStatus = 13
	// EROCSReserved14 is reserved
	EROCSReserved14 EnRouteOperationalCapabilityStatus = 14
	// EROCSReserved15 is reserved
	EROCSReserved15 EnRouteOperationalCapabilityStatus = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities EnRouteOperationalCapabilityStatus) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadEnRouteOperationalCapabilityStatus reads the EnRouteOperationalCapabilityStatus from a 56 bits data field
func ReadEnRouteOperationalCapabilityStatus(data []byte) EnRouteOperationalCapabilityStatus {
	bits := (data[3] & 0xF0) >> 4
	return EnRouteOperationalCapabilityStatus(bits)
}
