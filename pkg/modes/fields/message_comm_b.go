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

// readMessageCommB reads the MB field from a message
func ReadMessageCommB(message common.MessageData) MessageCommB {

	return MessageCommB{
		Data: message.Payload[3:10],
	}
}

func (messageCommB MessageCommB) PrettyPrint() string {
	return fmt.Sprintf("%X %X %X %X %X %X %X",
		messageCommB.Data[0],
		messageCommB.Data[1],
		messageCommB.Data[2],
		messageCommB.Data[3],
		messageCommB.Data[4],
		messageCommB.Data[5],
		messageCommB.Data[6])
}

func (messageCommB MessageCommB) ExtendedPrettyPrint() string {
	return messageCommB.PrettyPrint()
}
