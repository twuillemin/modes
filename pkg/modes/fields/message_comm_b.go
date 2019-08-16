package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Message Comm-B (MB)
//
// -----------------------------------------------------------------------------------------

// MessageCommB field shall be used to transmit data link messages to the ground
//
// Defined at 3.1.2.6.6.1
type MessageCommB []byte

// ReadMessageCommB reads the MB field from a message
func ReadMessageCommB(message common.MessageData) MessageCommB {

	data := make([]byte, 7)

	for i := 0; i < 7; i++ {
		data[i] = message.Payload[i+3]
	}

	return data
}

// ToString returns a basic, but readable, representation of the field
func (messageCommB MessageCommB) ToString() string {
	return fmt.Sprintf("%02X %02X %02X %02X %02X %02X %02X",
		messageCommB[0],
		messageCommB[1],
		messageCommB[2],
		messageCommB[3],
		messageCommB[4],
		messageCommB[5],
		messageCommB[6])
}
