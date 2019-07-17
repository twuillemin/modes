package fields

import "fmt"

// TargetStateReportCapability is the Target State Report Capability definition
//
// Specified in Doc 9871 / B.2.3.10.3
type TargetStateReportCapability byte

const (
	// TSRCNoCapability indicates No capability for sending messages to support Target State Reports
	TSRCNoCapability TargetStateReportCapability = 0
	// TSRCCapable indicates Capability of sending messages to support Target State Reports
	TSRCCapable TargetStateReportCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (status TargetStateReportCapability) ToString() string {

	switch status {
	case TSRCNoCapability:
		return "0 - No capability for sending messages to support Target State Reports"
	case TSRCCapable:
		return "1 - Capability of sending messages to support Target State Reports"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadTargetStateReportCapability reads the TargetStateReportCapability from a 56 bits data field
func ReadTargetStateReportCapability(data []byte) TargetStateReportCapability {
	bits := data[1] & 0x01
	return TargetStateReportCapability(bits)
}
