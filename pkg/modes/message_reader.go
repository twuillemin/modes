package modes

import (
	"errors"
	"fmt"
)

// MessageData is the basic message structure that is applicable to all messages, except DF24
type MessageData struct {
	DownLinkFormat uint8
	FirstField     uint8
	Payload        []uint8
	Parity         []uint8
}

func ReadMessage(message []uint8) (Message, error) {

	if len(message) < 4 {
		return nil, errors.New("a message can not be shorted than 4 bytes")
	}

	messageData := MessageData{
		DownLinkFormat: message[0] >> 3,
		FirstField:     message[0] & 0x07,
		Payload:        message[1 : len(message)-3],
		Parity:         message[len(message)-3:],
	}

	switch messageData.DownLinkFormat {
	case 0:
		return ParseDF0(messageData)
	case 4:
		return ParseDF4(messageData)
	case 5:
		return ParseDF5(messageData)
	case 11:
		return ParseDF11(messageData)
	case 16:
		return ParseDF16(messageData)
	case 17:
		return ParseDF17(messageData)
	case 18:
		return ParseDF18(messageData)
	case 19:
		return ParseDF19(messageData)
	case 20:
		return ParseDF20(messageData)
	case 21:
		return ParseDF21(messageData)
	case 24:
		return ParseDF24(messageData)
	default:
		return nil, fmt.Errorf("the downlink format %v is not supported", messageData.DownLinkFormat)
	}
}
