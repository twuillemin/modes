package reader

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/messages"
)

func ReadMessage(message []uint8) (messages.Message, error) {

	if len(message) < 4 {
		return nil, errors.New("a message can not be shorted than 4 bytes")
	}

	messageData := common.MessageData{
		DownLinkFormat: message[0] >> 3,
		FirstField:     message[0] & 0x07,
		Payload:        message[1 : len(message)-3],
		Parity:         message[len(message)-3:],
	}

	switch messageData.DownLinkFormat {
	case 0:
		return messages.ParseDF0(messageData)
	case 4:
		return messages.ParseDF4(messageData)
	case 5:
		return messages.ParseDF5(messageData)
	case 11:
		return messages.ParseDF11(messageData)
	case 16:
		return messages.ParseDF16(messageData)
	case 17:
		return messages.ParseDF17(messageData)
	case 18:
		return messages.ParseDF18(messageData)
	case 19:
		return messages.ParseDF19(messageData)
	case 20:
		return messages.ParseDF20(messageData)
	case 21:
		return messages.ParseDF21(messageData)
	case 24:
		return messages.ParseDF24(messageData)
	default:
		return nil, fmt.Errorf("the downlink format %v is not supported", messageData.DownLinkFormat)
	}
}
