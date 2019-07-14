package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Message Extended Squitter (ME)
//
// -----------------------------------------------------------------------------------------

// MessageExtendedSquitter field in DF = 17 shall be used to transmit broadcast messages. Extended squitter shall be
// supported by registers 05, 06, 07, 08, 09, 0A {HEX} and 61-6F {HEX} and shall conform to either version 0, version 1
// or version 2 message formats as described below:
//
// Defined at 3.1.2.8.6.2
type MessageExtendedSquitter struct {
	Data []byte
}

// ReadMessageExtendedSquitter reads the ME field from a message
func ReadMessageExtendedSquitter(message common.MessageData) MessageExtendedSquitter {

	return MessageExtendedSquitter{
		Data: message.Payload[3:10],
	}
}

// ToString returns a basic, but readable, representation of the field
func (messageExtendedSquitter MessageExtendedSquitter) ToString() string {
	return fmt.Sprintf("%X %X %X %X %X %X %X",
		messageExtendedSquitter.Data[0],
		messageExtendedSquitter.Data[1],
		messageExtendedSquitter.Data[2],
		messageExtendedSquitter.Data[3],
		messageExtendedSquitter.Data[4],
		messageExtendedSquitter.Data[5],
		messageExtendedSquitter.Data[6])
}

// ToExtendedString returns a complete representation of the field
func (messageExtendedSquitter MessageExtendedSquitter) ToExtendedString() string {
	return messageExtendedSquitter.ToString()
}
