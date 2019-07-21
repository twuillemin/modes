package fields

import "fmt"

// UniversalAccessTransceiverCapability is the UAT IN (Universal Access Transceiver) definition
//
// Specified in Doc 9871 / C.2.3.10.3
type UniversalAccessTransceiverCapability byte

const (
	// UATNoCapability indicates Aircraft has No UAT Receive capability
	UATNoCapability UniversalAccessTransceiverCapability = 0
	// UATCapable indicates Aircraft has UAT Receive capability
	UATCapable UniversalAccessTransceiverCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (status UniversalAccessTransceiverCapability) ToString() string {

	switch status {
	case UATNoCapability:
		return "0 - Aircraft has No UAT Receive capability"
	case UATCapable:
		return "1 - Aircraft has UAT Receive capability"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadUniversalAccessTransceiverCapabilityAirborne reads the UAT IN (Universal Access Transceiver) from a 56 bits data field
func ReadUniversalAccessTransceiverCapabilityAirborne(data []byte) UniversalAccessTransceiverCapability {
	bits := (data[2] & 0x20) >> 5
	return UniversalAccessTransceiverCapability(bits)
}

// ReadUniversalAccessTransceiverCapabilitySurface reads the UAT IN (Universal Access Transceiver) from a 56 bits data field
func ReadUniversalAccessTransceiverCapabilitySurface(data []byte) UniversalAccessTransceiverCapability {
	bits := data[1] & 0x01
	return UniversalAccessTransceiverCapability(bits)
}
