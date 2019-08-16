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

// MessageACAS (MV) field shall contain the aircraft address which provides the long ACAS message.
//
// Defined at 3.1.2.8.3.1
type MessageACAS []byte

// ReadMessageACAS reads the MV field from a message. Note that the
// content is not parsed
func ReadMessageACAS(message common.MessageData) MessageACAS {

	// Format of the message is as follow:
	//        0               1                   2                 3                 4                  5               6           7, 8, 9
	//   SL  _ _   RI  |RI _ _     AC    |        AC       | <---------------------------------------- MV ----------------------------------> |
	//                 |                 |                 |  VDS1     VDS2  |       ARA       |   ARA       RAC | RAC RAT MTE Res | Reserved |
	// x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | v v v v v v v v | a a a a a a a a | a a a a a a c c | c c t m _ _ _ _ | 18 bits  |

	data := make([]byte, 7)

	for i := 0; i < 7; i++ {
		data[i] = message.Payload[i+3]
	}

	return data
}

// ToString returns a basic, but readable, representation of the field
func (messageACAS MessageACAS) ToString() string {
	return fmt.Sprintf("%v", messageACAS)
}
