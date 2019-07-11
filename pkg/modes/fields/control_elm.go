package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 ControlELM (KE)
//
// -----------------------------------------------------------------------------------------

// ControlELM (KE) downlink field shall define the content of the ND and MD fields.
//
// Defined at 3.1.2.7.3.1
type ControlELM int

const (
	// DownlinkELMTransmission signifies downlink ELM transmission
	DownlinkELMTransmission ControlELM = 0
	// UplinkELMAcknowledgement signifies uplink ELM acknowledgement
	UplinkELMAcknowledgement ControlELM = 1
)

// readControlELM reads the KE field from a message
func ReadControlELM(message common.MessageData) ControlELM {
	if message.DownLinkFormat&0x02 != 0 {
		return UplinkELMAcknowledgement
	}
	return DownlinkELMTransmission
}

func (controlELM ControlELM) PrettyPrint() string {
	switch controlELM {
	case DownlinkELMTransmission:
		return "0 - Downlink ELM Transmission"
	case UplinkELMAcknowledgement:
		return "1 - Uplink ELM Acknowledgement"
	default:
		return fmt.Sprintf("%v - Unknown code", controlELM)
	}
}

func (controlELM ControlELM) ExtendedPrettyPrint() string {
	return controlELM.PrettyPrint()
}
