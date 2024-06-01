package fields

import "fmt"

// ACASHybridSurveillanceCapability is the ACAS hybrid surveillance capability definition
//
// Specified in Doc 9871 / D.2.4.1
type ACASHybridSurveillanceCapability byte

const (
	// ACASHybridSurveillanceNotCapable indicates that there is no hybrid surveillance capability.
	ACASHybridSurveillanceNotCapable ACASHybridSurveillanceCapability = 0
	// ACASHybridSurveillanceCapable indicates the capability of hybrid surveillance.
	ACASHybridSurveillanceCapable ACASHybridSurveillanceCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (ahsc ACASHybridSurveillanceCapability) ToString() string {

	switch ahsc {
	case ACASHybridSurveillanceNotCapable:
		return "0 - No capability of hybrid surveillance"
	case ACASHybridSurveillanceCapable:
		return "1 - Capability of hybrid surveillance"
	default:
		return fmt.Sprintf("%v - Unknown code", ahsc)
	}
}

// ReadACASHybridSurveillanceCapability reads the ACASHybridSurveillanceCapability from a 56 bits data field
func ReadACASHybridSurveillanceCapability(data []byte) ACASHybridSurveillanceCapability {
	bits := (data[4] & 0x08) >> 3
	return ACASHybridSurveillanceCapability(bits)
}
