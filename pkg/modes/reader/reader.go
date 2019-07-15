package reader

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/messages"
)

// ReadMessage reads and parse a Mode S message. The CRC is not verified at this point.
//
// params:
//    - message: The body of the message. The message must be 7 or 14 bytes long
//
// Return the parsed message or an error
func ReadMessage(message []byte) (messages.ModeSMessage, error) {

	if len(message) != 7 && len(message) != 14 {
		return nil, errors.New("a Mode S message must be 7 or 14 bytes long")
	}

	messageData := common.MessageData{
		DownLinkFormat: message[0] >> 3,
		FirstField:     message[0] & 0x07,
		Payload:        message[1 : len(message)-3],
		Parity:         message[len(message)-3:],
	}

	var modeS messages.ModeSMessage
	var err error

	// Extract the content of the message
	switch messageData.DownLinkFormat {
	case 0:
		modeS, err = messages.ParseDF0(messageData)
	case 4:
		modeS, err = messages.ParseDF4(messageData)
	case 5:
		modeS, err = messages.ParseDF5(messageData)
	case 11:
		modeS, err = messages.ParseDF11(messageData)
	case 16:
		modeS, err = messages.ParseDF16(messageData)
	case 17:
		modeS, err = messages.ParseDF17(messageData)
	case 18:
		modeS, err = messages.ParseDF18(messageData)
	case 19:
		modeS, err = messages.ParseDF19(messageData)
	case 20:
		modeS, err = messages.ParseDF20(messageData)
	case 21:
		modeS, err = messages.ParseDF21(messageData)
	case 24, 25, 26, 27, 28, 29, 30, 31:
		// For DF24, only the first two bits of the generally used 5 bits
		// are possible. So use all combinations 24->31
		modeS, err = messages.ParseDF24(messageData)
	default:
		err = fmt.Errorf("the downlink format %v is not supported", messageData.DownLinkFormat)
	}

	if err != nil {
		return nil, err
	}

	return modeS, nil
}
