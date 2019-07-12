package reader

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/messages"
)

// ReadMessage reads and parse a Mode S message
//
// params:
//    - message: The body of the message. The message must be 7 or 14 bytes long
//    - noCRC: indicates whether the CRC has been already subtracted from the parity field. In this case
//             the parity part of the the message is directly the ICAO 24 address
//
// Return the parsed message and its ICAO address (on 3 bytes) or an error
func ReadMessage(message []byte, noCRC bool) (messages.MessageModeS, []byte, error) {

	if len(message) != 7 && len(message) != 14 {
		return nil, nil, errors.New("a Mode S message must be 7 or 14 bytes long")
	}

	messageData := common.MessageData{
		DownLinkFormat: message[0] >> 3,
		FirstField:     message[0] & 0x07,
		Payload:        message[1 : len(message)-3],
		Parity:         message[len(message)-3:],
	}

	var modeS messages.MessageModeS
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
	case 24:
		modeS, err = messages.ParseDF24(messageData)
	default:
		err = fmt.Errorf("the downlink format %v is not supported", messageData.DownLinkFormat)
	}

	if err != nil {
		return nil, nil, err
	}

	var icao24 []byte

	// For DF11 and DF17, special case
	if messageData.DownLinkFormat == 11 || messageData.DownLinkFormat == 17 {
		icao24 = []byte{
			messageData.Payload[0],
			messageData.Payload[1],
			messageData.Payload[2],
		}
	} else {
		// If the CRC was already removed
		if noCRC {
			// Use directly the parity
			icao24 = []byte{
				messageData.Parity[0],
				messageData.Parity[1],
				messageData.Parity[2],
			}
		} else {

			// Compute parity on the whole message, except the 3 last bytes
			parity := computeParity(message[:len(message)-3])
			icao24 = xorArrays(parity, messageData.Parity)
		}
	}

	return modeS, icao24, nil
}
