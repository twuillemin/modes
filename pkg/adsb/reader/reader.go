package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	messages05 "github.com/twuillemin/modes/pkg/bds/bds05/messages"
	messages08 "github.com/twuillemin/modes/pkg/bds/bds08/messages"
	messages09 "github.com/twuillemin/modes/pkg/bds/bds09/messages"
	messages65 "github.com/twuillemin/modes/pkg/bds/bds65/messages"
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

	// -----------------------------------
	// Code   BDS     V0      V1       V2
	//  0
	//  1     0,8
	//  2     0,8
	//  3     0,8
	//  4     0,8
	//  5
	//  6
	//  7
	//  8
	//  9     0,5
	// 10     0,5
	// 14     0,5
	// 12     0,5
	// 13     0,5
	// 14     0,5
	// 15     0,5
	// 16     0,5
	// 17     0,5
	// 18     0,5
	// 19     0,9
	// 20     0,5
	// 21     0,5
	// 22     0,5
	// 23   Reserved
	// 24   Reserved
	// 25   Reserved
	// 26   Reserved
	// 27   Reserved
	// 28
	// 29   Reserved
	// 30   Reserved
	// 31     6,5      OK

	// Get the type
	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1, 2, 3, 4:
		return messages08.ReadBDS08(data)
	case 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22:
		return messages05.ReadBDS05(data)
	case 19:
		return messages09.ReadBDS09(data)
	case 31:
		return messages65.ReadBDS65(data)
	}

	return nil, nil
}
