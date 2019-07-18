package fields

import "fmt"

// HorizontalDataAvailableSourceIndicator is the Horizontal Data Available / Source Indicator definition
//
// Specified in Doc 9871 / B.2.3.9.8
type HorizontalDataAvailableSourceIndicator byte

const (
	// HDANoValidDataAvailable indicates no valid horizontal target state data is available
	HDANoValidDataAvailable HorizontalDataAvailableSourceIndicator = 0
	// HDAAutopilot indicates autopilot control panel selected value, such as Mode Control Panel (MCP) or Flight Control Unit (FCU)
	HDAAutopilot HorizontalDataAvailableSourceIndicator = 1
	// HDAHoldingAltitude indicates maintaining current heading or track angle (e.g. autopilot mode select)
	HDAHoldingAltitude HorizontalDataAvailableSourceIndicator = 2
	// HDAFMS indicates FMS/RNAV system (indicates track angle specified by leg type)
	HDAFMS HorizontalDataAvailableSourceIndicator = 3
)

// ToString returns a basic, but readable, representation of the field
func (data HorizontalDataAvailableSourceIndicator) ToString() string {

	switch data {
	case HDANoValidDataAvailable:
		return "0 - no valid horizontal target state data is available"
	case HDAAutopilot:
		return "1 - autopilot control panel selected value, such as Mode Control Panel (MCP) or Flight Control Unit (FCU)"
	case HDAHoldingAltitude:
		return "2 - maintaining current heading or track angle (e.g. autopilot mode select)"
	case HDAFMS:
		return "3 - FMS/RNAV system (indicates track angle specified by leg type)"
	default:
		return fmt.Sprintf("%v - Unknown code", data)
	}
}

// ReadHorizontalDataAvailableSourceIndicator reads the HorizontalDataAvailableSourceIndicator from a 56 bits data field
func ReadHorizontalDataAvailableSourceIndicator(data []byte) HorizontalDataAvailableSourceIndicator {
	bits := (data[3] & 0x60) >> 5
	return HorizontalDataAvailableSourceIndicator(bits)
}
