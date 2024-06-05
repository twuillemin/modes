package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// EncodedLatitude is the encoded latitude
//
// Specified in Doc 9871 / C.2.6
type EncodedLatitude uint32

// ToString returns a basic, but readable, representation of the field
func (encodedLatitude EncodedLatitude) ToString() string {
	return fmt.Sprintf("%v", encodedLatitude)
}

// ReadEncodedLatitude reads the EncodedLatitude from a 56 bits data field
func ReadEncodedLatitude(data []byte) EncodedLatitude {

	b2 := (data[2] & 0x02) >> 1
	b1 := (data[2]&0x01)<<7 + (data[3]&0xFE)>>1
	b0 := (data[3]&0x01)<<7 + (data[4]&0xFE)>>1

	return EncodedLatitude(bitutils.Pack3Bytes(b2, b1, b0))
}
