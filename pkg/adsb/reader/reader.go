package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	messages05 "github.com/twuillemin/modes/pkg/bds/bds05/messages"
	messages06 "github.com/twuillemin/modes/pkg/bds/bds06/messages"
	messages08 "github.com/twuillemin/modes/pkg/bds/bds08/messages"
	messages09 "github.com/twuillemin/modes/pkg/bds/bds09/messages"
	messages61 "github.com/twuillemin/modes/pkg/bds/bds61/messages"
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
	// Code   BDS      V0      V1       V2
	//  0
	//  1     0,8      OK
	//  2     0,8      OK
	//  3     0,8      OK
	//  4     0,8      OK
	//  5     0,6      OK
	//  6     0,6      OK
	//  7     0,6      OK
	//  8     0,6      OK
	//  9     0,5      OK
	// 10     0,5      OK
	// 14     0,5      OK
	// 12     0,5      OK
	// 13     0,5      OK
	// 14     0,5      OK
	// 15     0,5      OK
	// 16     0,5      OK
	// 17     0,5      OK
	// 18     0,5      OK
	// 19     0,9      OK
	// 20     0,5      OK
	// 21     0,5      OK
	// 22     0,5      OK
	// 23   Reserved  Res.
	// 24   Reserved  Res.
	// 25   Reserved  Res.
	// 26   Reserved  Res.
	// 27   Reserved  Res.
	// 28     6,1      OK
	// 29   Reserved  Res.
	// 30   Reserved  Res.
	// 31     6,5      OK

	// Get the type
	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1, 2, 3, 4:
		return messages08.ReadBDS08(data)
	case 5, 6, 7, 8:
		return messages06.ReadBDS06(data)
	case 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22:
		return messages05.ReadBDS05(data)
	case 19:
		return messages09.ReadBDS09(data)
	case 28:
		return messages61.ReadBDS61(data)
	case 31:
		return messages65.ReadBDS65(data)
	}

	return nil, nil
}
