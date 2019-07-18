package fields

import "fmt"

// IFRCapabilityFlag is the IFR Capability Flag definition
//
// Specified in Doc 9871 / Table A-2-9
type IFRCapabilityFlag byte

const (
	// IFRFNoCapability indicates Transmitting aircraft has no capability for ADS-Bbased conflict detection or higher
	// level (class A1 or above) applications
	IFRFNoCapability IFRCapabilityFlag = 0
	// IFRFCapable indicates Transmitting aircraft has capability for ADS-B-based conflict detection and higher level
	// (class A1 or above) applications.
	IFRFCapable IFRCapabilityFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (flag IFRCapabilityFlag) ToString() string {

	switch flag {
	case IFRFNoCapability:
		return "0 - transmitting aircraft has no capability for ADS-Bbased conflict detection or higher level (class A1 or above) applications"
	case IFRFCapable:
		return "1 - transmitting aircraft has capability for ADS-B-based conflict detection and higher level (class A1 or above) applications"
	default:
		return fmt.Sprintf("%v - Unknown code", flag)
	}
}

// ReadIFRCapabilityFlag reads the IFRCapabilityFlag from a 56 bits data field
func ReadIFRCapabilityFlag(data []byte) IFRCapabilityFlag {
	bits := (data[1] & 0x40) >> 6
	return IFRCapabilityFlag(bits)
}
