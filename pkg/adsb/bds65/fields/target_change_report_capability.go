package fields

import "fmt"

// TargetChangeReportCapability is the Target Change Report Capability definition
//
// Specified in Doc 9871 / B.2.3.10.3
type TargetChangeReportCapability byte

const (
	// TCRCNoCapability indicates No capability for sending messages to support Trajectory Change Reports
	TCRCNoCapability TargetChangeReportCapability = 0
	// TCRCCapableTC0Only indicates Capability of sending messages to support TC+0 Report only
	TCRCCapableTC0Only TargetChangeReportCapability = 1
	// TCRCCapableMultipleTC indicates Capability of sending information for multiple TC Reports
	TCRCCapableMultipleTC TargetChangeReportCapability = 2
	// TCRCCapableReserved is Reserved
	TCRCCapableReserved TargetChangeReportCapability = 3
)

// ToString returns a basic, but readable, representation of the field
func (status TargetChangeReportCapability) ToString() string {

	switch status {
	case TCRCNoCapability:
		return "0 - No capability for sending messages to support Trajectory Change Reports"
	case TCRCCapableTC0Only:
		return "1 - Capability of sending messages to support TC+0 Report only"
	case TCRCCapableMultipleTC:
		return "2 - Capability of sending information for multiple TC Reports"
	case TCRCCapableReserved:
		return "3 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadTargetChangeReportCapability reads the TargetStateReportCapability from a 56 bits data field
func ReadTargetChangeReportCapability(data []byte) TargetChangeReportCapability {
	bits := (data[2] & 0xC0) >> 6
	return TargetChangeReportCapability(bits)
}
