package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Downlink Request (DR)
//
// -----------------------------------------------------------------------------------------

// DownlinkRequest (DR) field shall contain requests to downlink information.
//
// Defined at 3.1.2.6.5.2 for codes 0 to 15 and 3.1.2.7.7.1 for codes 16 to 31
type DownlinkRequest int

const (
	// DownlinkRequestNoDownlinkRequest signifies no downlink request
	DownlinkRequestNoDownlinkRequest DownlinkRequest = 0
	// DownlinkRequestToSendCommBMessage signifies request to send Comm-B message
	DownlinkRequestToSendCommBMessage DownlinkRequest = 1
	// DownlinkRequestReservedACAS2 is reserved for ACAS
	DownlinkRequestReservedACAS2 DownlinkRequest = 2
	// DownlinkRequestReservedACAS3 is reserved for ACAS
	DownlinkRequestReservedACAS3 DownlinkRequest = 3
	// DownlinkRequestCommBMessage1Available signifies Comm-B broadcast message 1 available
	DownlinkRequestCommBMessage1Available DownlinkRequest = 4
	// DownlinkRequestCommBMessage2Available signifies Comm-B broadcast message 2 available
	DownlinkRequestCommBMessage2Available DownlinkRequest = 5
	// DownlinkRequestReservedACAS6 is reserved for ACAS
	DownlinkRequestReservedACAS6 DownlinkRequest = 6
	// DownlinkRequestReservedACAS7 is reserved for ACAS
	DownlinkRequestReservedACAS7 DownlinkRequest = 7
	// DownlinkRequestNotAssigned8 is not assigned
	DownlinkRequestNotAssigned8 DownlinkRequest = 8
	// DownlinkRequestNotAssigned9 is not assigned
	DownlinkRequestNotAssigned9 DownlinkRequest = 9
	// DownlinkRequestNotAssigned10 is not assigned
	DownlinkRequestNotAssigned10 DownlinkRequest = 10
	// DownlinkRequestNotAssigned11 is not assigned
	DownlinkRequestNotAssigned11 DownlinkRequest = 11
	// DownlinkRequestNotAssigned12 is not assigned
	DownlinkRequestNotAssigned12 DownlinkRequest = 12
	// DownlinkRequestNotAssigned13 is not assigned
	DownlinkRequestNotAssigned13 DownlinkRequest = 13
	// DownlinkRequestNotAssigned14 is not assigned
	DownlinkRequestNotAssigned14 DownlinkRequest = 14
	// DownlinkRequestNotAssigned15 is not assigned
	DownlinkRequestNotAssigned15 DownlinkRequest = 15
	// DownlinkRequestELMAvailable1Segments announces the presence of a downlink ELM of 1 segments
	DownlinkRequestELMAvailable1Segments DownlinkRequest = 16
	// DownlinkRequestELMAvailable2Segments announces the presence of a downlink ELM of 2 segments
	DownlinkRequestELMAvailable2Segments DownlinkRequest = 17
	// DownlinkRequestELMAvailable3Segments announces the presence of a downlink ELM of 3 segments
	DownlinkRequestELMAvailable3Segments DownlinkRequest = 18
	// DownlinkRequestELMAvailable4Segments announces the presence of a downlink ELM of 4 segments
	DownlinkRequestELMAvailable4Segments DownlinkRequest = 19
	// DownlinkRequestELMAvailable5Segments announces the presence of a downlink ELM of 5 segments
	DownlinkRequestELMAvailable5Segments DownlinkRequest = 20
	// DownlinkRequestELMAvailable6Segments announces the presence of a downlink ELM of 6 segments
	DownlinkRequestELMAvailable6Segments DownlinkRequest = 21
	// DownlinkRequestELMAvailable7Segments announces the presence of a downlink ELM of 7 segments
	DownlinkRequestELMAvailable7Segments DownlinkRequest = 22
	// DownlinkRequestELMAvailable8Segments announces the presence of a downlink ELM of 8 segments
	DownlinkRequestELMAvailable8Segments DownlinkRequest = 23
	// DownlinkRequestELMAvailable9Segments announces the presence of a downlink ELM of 9 segments
	DownlinkRequestELMAvailable9Segments DownlinkRequest = 24
	// DownlinkRequestELMAvailable10Segments announces the presence of a downlink ELM of 10 segments
	DownlinkRequestELMAvailable10Segments DownlinkRequest = 25
	// DownlinkRequestELMAvailable11Segments announces the presence of a downlink ELM of 11 segments
	DownlinkRequestELMAvailable11Segments DownlinkRequest = 26
	// DownlinkRequestELMAvailable12Segments announces the presence of a downlink ELM of 12 segments
	DownlinkRequestELMAvailable12Segments DownlinkRequest = 27
	// DownlinkRequestELMAvailable13Segments announces the presence of a downlink ELM of 13 segments
	DownlinkRequestELMAvailable13Segments DownlinkRequest = 28
	// DownlinkRequestELMAvailable14Segments announces the presence of a downlink ELM of 14 segments
	DownlinkRequestELMAvailable14Segments DownlinkRequest = 29
	// DownlinkRequestELMAvailable15Segments announces the presence of a downlink ELM of 15 segments
	DownlinkRequestELMAvailable15Segments DownlinkRequest = 30
	// DownlinkRequestELMAvailable16Segments announces the presence of a downlink ELM of 16 segments
	DownlinkRequestELMAvailable16Segments DownlinkRequest = 31
)

// readSensitivityLevelReport reads the DR field from a message
func ReadDownlinkRequest(message common.MessageData) DownlinkRequest {

	return DownlinkRequest(message.Payload[0] >> 3)
}

func (downlinkRequest DownlinkRequest) PrettyPrint() string {
	switch downlinkRequest {
	case DownlinkRequestNoDownlinkRequest:
		return "0 - No Downlink Request"
	case DownlinkRequestToSendCommBMessage:
		return "1 - Request To Send Comm-B Message"
	case DownlinkRequestReservedACAS2:
		return "2 - Reserved ACAS"
	case DownlinkRequestReservedACAS3:
		return "3 - Reserved ACAS"
	case DownlinkRequestCommBMessage1Available:
		return "4 - Comm-B Message 1 Available"
	case DownlinkRequestCommBMessage2Available:
		return "5 - Comm-B Message 1 Available"
	case DownlinkRequestReservedACAS6:
		return "6 - Reserved ACAS"
	case DownlinkRequestReservedACAS7:
		return "7 - Reserved ACAS"
	case DownlinkRequestNotAssigned8:
		return "8 - Not Assigned"
	case DownlinkRequestNotAssigned9:
		return "9 - Not Assigned"
	case DownlinkRequestNotAssigned10:
		return "10 - Not Assigned"
	case DownlinkRequestNotAssigned11:
		return "11 - Not Assigned"
	case DownlinkRequestNotAssigned12:
		return "12 - Not Assigned"
	case DownlinkRequestNotAssigned13:
		return "13 - Not Assigned"
	case DownlinkRequestNotAssigned14:
		return "14 - Not Assigned"
	case DownlinkRequestNotAssigned15:
		return "15 - Not Assigned"
	case DownlinkRequestELMAvailable1Segments:
		return "16 - ELM Available 1 Segments"
	case DownlinkRequestELMAvailable2Segments:
		return "17 - ELM Available 2 Segments"
	case DownlinkRequestELMAvailable3Segments:
		return "18 - ELM Available 3 Segments"
	case DownlinkRequestELMAvailable4Segments:
		return "19 - ELM Available 4 Segments"
	case DownlinkRequestELMAvailable5Segments:
		return "20 - ELM Available 5 Segments"
	case DownlinkRequestELMAvailable6Segments:
		return "21 - ELM Available 6 Segments"
	case DownlinkRequestELMAvailable7Segments:
		return "22 - ELM Available 7 Segments"
	case DownlinkRequestELMAvailable8Segments:
		return "23 - ELM Available 8 Segments"
	case DownlinkRequestELMAvailable9Segments:
		return "24 - ELM Available 9 Segments"
	case DownlinkRequestELMAvailable10Segments:
		return "25 - ELM Available 10 Segment"
	case DownlinkRequestELMAvailable11Segments:
		return "26 - ELM Available 11 Segments"
	case DownlinkRequestELMAvailable12Segments:
		return "27 - ELM Available 12 Segments"
	case DownlinkRequestELMAvailable13Segments:
		return "28 - ELM Available 13 Segments"
	case DownlinkRequestELMAvailable14Segments:
		return "29 - ELM Available 14 Segments"
	case DownlinkRequestELMAvailable15Segments:
		return "30 - ELM Available 15 Segments"
	case DownlinkRequestELMAvailable16Segments:
		return "31 - ELM Available 16 Segments"
	default:
		return fmt.Sprintf("%v - Unknown code", downlinkRequest)
	}
}

func (downlinkRequest DownlinkRequest) ExtendedPrettyPrint() string {
	return downlinkRequest.PrettyPrint()
}
