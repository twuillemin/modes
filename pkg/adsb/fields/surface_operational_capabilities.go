package fields

import "fmt"

// SurfaceOperationalCapabilities is the Surface Operational Capabilities format definition
//
// Specified in Doc 9871 / A.2.3.11.6
type SurfaceOperationalCapabilities byte

const (
	// SOCReserved0 is reserved
	SOCReserved0 SurfaceOperationalCapabilities = 0
	// SOCReserved1 is reserved
	SOCReserved1 SurfaceOperationalCapabilities = 1
	// SOCReserved2 is reserved
	SOCReserved2 SurfaceOperationalCapabilities = 2
	// SOCReserved3 is reserved
	SOCReserved3 SurfaceOperationalCapabilities = 3
	// SOCReserved4 is reserved
	SOCReserved4 SurfaceOperationalCapabilities = 4
	// SOCReserved5 is reserved
	SOCReserved5 SurfaceOperationalCapabilities = 5
	// SOCReserved6 is reserved
	SOCReserved6 SurfaceOperationalCapabilities = 6
	// SOCReserved7 is reserved
	SOCReserved7 SurfaceOperationalCapabilities = 7
	// SOCReserved8 is reserved
	SOCReserved8 SurfaceOperationalCapabilities = 8
	// SOCReserved9 is reserved
	SOCReserved9 SurfaceOperationalCapabilities = 9
	// SOCReserved10 is reserved
	SOCReserved10 SurfaceOperationalCapabilities = 10
	// SOCReserved11 is reserved
	SOCReserved11 SurfaceOperationalCapabilities = 11
	// SOCReserved12 is reserved
	SOCReserved12 SurfaceOperationalCapabilities = 12
	// SOCReserved13 is reserved
	SOCReserved13 SurfaceOperationalCapabilities = 13
	// SOCReserved14 is reserved
	SOCReserved14 SurfaceOperationalCapabilities = 14
	// SOCReserved15 is reserved
	SOCReserved15 SurfaceOperationalCapabilities = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities SurfaceOperationalCapabilities) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadSurfaceOperationalCapabilities reads the SurfaceOperationalCapabilities from a 56 bits data field
func ReadSurfaceOperationalCapabilities(data []byte) SurfaceOperationalCapabilities {
	bits := data[2] & 0x0F
	return SurfaceOperationalCapabilities(bits)
}
