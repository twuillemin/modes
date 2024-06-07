package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// WingDown is the roll angle wing
//
// Specified in Doc 9871 / Table A-2-80
type WingDown byte

const (
	// RARightWingDown indicates the right wing is down.
	RARightWingDown WingDown = 0
	// RALeftWingDown indicates the left wing is down.
	RALeftWingDown WingDown = 1
)

// ToString returns a basic, but readable, representation of the field
func (wd WingDown) ToString() string {

	switch wd {
	case RARightWingDown:
		return "0 - Right Wing Down"
	case RALeftWingDown:
		return "1 - Left Wing Down"
	default:
		return fmt.Sprintf("%v - Unknown code", wd)
	}
}

func ReadRollAngle(data []byte) (bool, WingDown, float32) {
	status := (data[0] & 0x80) != 0

	wingDown := WingDown((data[0] & 0x40) >> 6)

	byte1 := data[0] & 0x3F
	byte2 := data[1] & 0xE0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 5
	rollAngle := float32(allBits) * 45.0 / 256.0

	return status, wingDown, rollAngle
}
