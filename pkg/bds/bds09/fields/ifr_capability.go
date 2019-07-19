package fields

import "fmt"

// IFRCapability is the IFR Capability Flag definition
//
// Specified in Doc 9871 / Table A-2-9
type IFRCapability byte

const (
	// IFRNoCapability indicates Transmitting aircraft has no capability for ADS-B based conflict detection or higher
	// level (class A1 or above) applications
	IFRNoCapability IFRCapability = 0
	// IFRCapable indicates Transmitting aircraft has capability for ADS-B-based conflict detection and higher level
	// (class A1 or above) applications.
	IFRCapable IFRCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (flag IFRCapability) ToString() string {

	switch flag {
	case IFRNoCapability:
		return "0 - transmitting aircraft has no capability for ADS-B based conflict detection or higher level (class A1 or above) applications"
	case IFRCapable:
		return "1 - transmitting aircraft has capability for ADS-B-based conflict detection and higher level (class A1 or above) applications"
	default:
		return fmt.Sprintf("%v - Unknown code", flag)
	}
}

// ReadIFRCapability reads the IFRCapability from a 56 bits data field
func ReadIFRCapability(data []byte) IFRCapability {
	bits := (data[1] & 0x40) >> 6
	return IFRCapability(bits)
}
