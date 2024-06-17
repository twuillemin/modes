package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// SelectedHeadingStatus is the Selected Heading Status definition
//
// Specified in Doc 9871 / C.2.3.9.6
type SelectedHeadingStatus byte

const (
	// SHSInvalid indicates that the Selected Heading is invalid or not available
	SHSInvalid = 0
	// SHSValid indicates that the Selected Heading is invalid or not available
	SHSValid SelectedHeadingStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SelectedHeadingStatus) ToString() string {

	switch status {
	case SHSInvalid:
		return "0 - selected heading data is either not available or is invalid"
	case SHSValid:
		return "1 - selected heading data is valid"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSelectedHeading reads the SelectedHeading from a 56 bits data field
func ReadSelectedHeading(data []byte) (float32, SelectedHeadingStatus) {
	status := SelectedHeadingStatus((data[3] & 0x04) >> 2)

	negative := data[3]&0x02 != 0
	byte1 := data[3] & 0x01
	byte2 := data[4] & 0xFE
	heading := float32(bitutils.Pack2Bytes(byte1, byte2) >> 1)

	if negative {
		heading = -heading
	}

	return heading * 180 / 256, status
}
