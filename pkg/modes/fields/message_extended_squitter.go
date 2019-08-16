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
type MessageExtendedSquitter []byte

// ReadMessageExtendedSquitter reads the ME field from a message
func ReadMessageExtendedSquitter(message common.MessageData) MessageExtendedSquitter {

	data := make([]byte, 7)

	for i := 0; i < 7; i++ {
		data[i] = message.Payload[i+3]
	}

	return data
}

// ToString returns a basic, but readable, representation of the field
func (messageExtendedSquitter MessageExtendedSquitter) ToString() string {
	return fmt.Sprintf("%02X %02X %02X %02X %02X %02X %02X",
		messageExtendedSquitter[0],
		messageExtendedSquitter[1],
		messageExtendedSquitter[2],
		messageExtendedSquitter[3],
		messageExtendedSquitter[4],
		messageExtendedSquitter[5],
		messageExtendedSquitter[6])
}
