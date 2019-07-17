package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/adsb/messages"
)

// ReadMessage reads and parse an ADSB message.
//
// params:
//    - message: The body of the message. The message must be 7 bytes long
//
// Return the parsed message or an error
func ReadMessage(data []byte) (messages.ADSBMessage, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for ADSB message must be 7 bytes long")
	}

	// Get the type
	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1, 2, 3, 4:
		return messages.ReadBDS08(data)
	case 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22:
		return messages.ReadBDS05(data)
	case 31:
		return messages.ReadBDS65(data)
	}

	return nil, nil
}
