package fields

import "fmt"

// VerticalDataAvailableSourceIndicator is the Vertical Data Available / Source Indicator definition
//
// Specified in Doc 9871 / B.2.3.9.3
type VerticalDataAvailableSourceIndicator byte

const (
	// VDANoValidDataAvailable indicates no valid Vertical Target State data is available
	VDANoValidDataAvailable VerticalDataAvailableSourceIndicator = 0
	// VDAAutopilot indicates autopilot control panel selected value, such as Mode Control Panel (MCP) or Flight Control Unit (FCU)
	VDAAutopilot VerticalDataAvailableSourceIndicator = 1
	// VDAHoldingAltitude indicates holding altitude
	VDAHoldingAltitude VerticalDataAvailableSourceIndicator = 2
	// VDAFMS indicates FMS/RNAV system
	VDAFMS VerticalDataAvailableSourceIndicator = 3
)

// ToString returns a basic, but readable, representation of the field
func (data VerticalDataAvailableSourceIndicator) ToString() string {

	switch data {
	case VDANoValidDataAvailable:
		return "0 - no valid Vertical Target State data is available"
	case VDAAutopilot:
		return "1 - autopilot control panel selected value, such as Mode Control Panel (MCP) or Flight Control Unit (FCU)"
	case VDAHoldingAltitude:
		return "2 - holding altitude"
	case VDAFMS:
		return "3 - FMS/RNAV system"
	default:
		return fmt.Sprintf("%v - Unknown code", data)
	}
}

// ReadVerticalDataAvailableSourceIndicator reads the VerticalDataAvailableSourceIndicator from a 56 bits data field
func ReadVerticalDataAvailableSourceIndicator(data []byte) VerticalDataAvailableSourceIndicator {
	bits := (data[0]&0x01)<<1 + (data[1]&0x80)>>7
	return VerticalDataAvailableSourceIndicator(bits)
}
