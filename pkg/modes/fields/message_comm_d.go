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
type MessageCommD []byte

// ReadMessageCommD reads the MB field from a message
func ReadMessageCommD(message common.MessageData) MessageCommD {

	data := make([]byte, 10)

	for i := 0; i < 10; i++ {
		data[i] = message.Payload[i]
	}

	return data
}

// ToString returns a basic, but readable, representation of the field
func (messageCommD MessageCommD) ToString() string {
	return fmt.Sprintf("%02X %02X %02X %02X %02X %02X %02X %02X %02X %02X",
		messageCommD[0],
		messageCommD[1],
		messageCommD[2],
		messageCommD[3],
		messageCommD[4],
		messageCommD[5],
		messageCommD[6],
		messageCommD[7],
		messageCommD[8],
		messageCommD[9])
}
