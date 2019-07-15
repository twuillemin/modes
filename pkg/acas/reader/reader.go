package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/acas/messages"
)

// ReadMessage reads and parse an ACAS message.
//
// params:
//    - message: The body of the message.
//
// Return the parsed message or an error.
func ReadMessage(message []byte) (messages.ACASMessage, error) {

	if len(message) < 1 {
		return nil, errors.New("unable to parse an empty message")
	}

	vds1 := message[0] >> 4
	vds2 := message[0] & 0x0F

	if vds1 == 3 && vds2 == 0 {
		return messages.ReadMessageACAS30(message), nil
	} else {
		return messages.ReadMessageACASUnknown(message), nil
	}
}
