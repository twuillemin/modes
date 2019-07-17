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
type MessageCommB struct {
	Data []byte
}

// ReadMessageCommB reads the MB field from a message
func ReadMessageCommB(message common.MessageData) MessageCommB {

	return MessageCommB{
		Data: message.Payload[3:10],
	}
}

// ToString returns a basic, but readable, representation of the field
func (messageCommB MessageCommB) ToString() string {
	return fmt.Sprintf("%02X %02X %02X %02X %02X %02X %02X",
		messageCommB.Data[0],
		messageCommB.Data[1],
		messageCommB.Data[2],
		messageCommB.Data[3],
		messageCommB.Data[4],
		messageCommB.Data[5],
		messageCommB.Data[6])
}
