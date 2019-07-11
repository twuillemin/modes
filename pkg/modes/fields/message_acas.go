package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Message ACAS (MV)
//
// -----------------------------------------------------------------------------------------

// MessageACAS (MV) field shall contain the aircraft address which provides unambiguous identification of
// the aircraft.
//
// Defined at 3.1.2.8.3.1
type MessageACAS struct {
	Data []byte
}

// readMessageACAS reads the MV field from a message
func ReadMessageACAS(message common.MessageData) MessageACAS {

	return MessageACAS{
		Data: message.Payload[3:10],
	}
}

func (messageACAS MessageACAS) PrettyPrint() string {
	return fmt.Sprintf("%X %X %X %X %X %X %X",
		messageACAS.Data[0],
		messageACAS.Data[1],
		messageACAS.Data[2],
		messageACAS.Data[3],
		messageACAS.Data[4],
		messageACAS.Data[5],
		messageACAS.Data[6])
}

func (messageACAS MessageACAS) ExtendedPrettyPrint() string {
	return messageACAS.PrettyPrint()
}
