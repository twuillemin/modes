package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/adsb/messages"
)

// ReadMessage reads and parse a Mode S message. The CRC is not verified at this point.
//
// params:
//    - message: The body of the message. The message must be 7 or 14 bytes long
//
// Return the parsed message or an error
func ReadMessage(data []byte) (messages.ADSBMessage, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for ADSB message must be 7 bytes long")
	}

	// Get the type
	formatTypeCode := data[0] >> 3

	switch formatTypeCode {
	case 1, 2, 3, 4:
		return messages.ReadBDS08(data)
	}

	return nil, nil
}
