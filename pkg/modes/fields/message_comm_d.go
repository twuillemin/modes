package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Message Comm-D (MD)
//
// -----------------------------------------------------------------------------------------

// MessageCommD field shall contain:
//   a) one of the segments of a sequence used to transmit a downlink ELM to the interrogator; or
//   b) control codes for an uplink ELM.
//
// Defined at 3.1.2.7.3.3
type MessageCommD struct {
	Data []byte
}

// ReadMessageCommD reads the MB field from a message
func ReadMessageCommD(message common.MessageData) MessageCommD {

	return MessageCommD{
		Data: message.Payload,
	}
}

// ToString returns a basic, but readable, representation of the field
func (messageCommD MessageCommD) ToString() string {
	return fmt.Sprintf("%X %X %X %X %X %X %X %X %X %X",
		messageCommD.Data[0],
		messageCommD.Data[1],
		messageCommD.Data[2],
		messageCommD.Data[3],
		messageCommD.Data[4],
		messageCommD.Data[5],
		messageCommD.Data[6],
		messageCommD.Data[7],
		messageCommD.Data[8],
		messageCommD.Data[9])
}
