package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/commb"
	bds07 "github.com/twuillemin/modes/pkg/commb/bds07/messages"
	bds10 "github.com/twuillemin/modes/pkg/commb/bds10/messages"
	bds20 "github.com/twuillemin/modes/pkg/commb/bds20/messages"
)

// ReadCommBMessage reads and parse a Comm-B message.
//
// params:
//   - data: The body of the message. The message must be 7 bytes long
//
// Return the parsed message, the detected ADSB ReaderLevel and an optional error. The detected ADSB ReaderLevel will generally be
// the same as the given one, except if the decoded message has information to change it.
func ReadCommBMessage(data []byte) (commb.Message, error) {
	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B message must be 7 bytes long")
	}

	messages := make([]commb.Message, 0, 10)

	if bds07.CheckIfDataReadable(data) == nil {
		message, err := bds07.ReadExtendedSquitterStatus(data)
		if err == nil {
			messages = append(messages, message)
		}
	}

	if bds10.CheckIfDataReadable(data) == nil {
		message, err := bds10.ReadDataLinkCapabilityReport(data)
		if err == nil {
			messages = append(messages, message)
		}
	}

	if bds20.CheckIfDataReadable(data) == nil {
		message, err := bds20.ReadAircraftIdentification(data)
		if err == nil {
			messages = append(messages, message)
		}
	}

	switch len(messages) {
	case 0:
		return nil, errors.New("message can not be read")
	case 1:
		return messages[0], nil
	default:
		return nil, errors.New("multiple format match the message")
	}
}