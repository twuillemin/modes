package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// EncodedLongitude is the encoded longitude
//
// Specified in Doc 9871 / C.2.6
type EncodedLongitude uint32

// ToString returns a basic, but readable, representation of the field
func (encodedLongitude EncodedLongitude) ToString() string {
	return fmt.Sprintf("%v", encodedLongitude)
}

// ReadEncodedLongitude read the EncodedLongitude from a 56 bits data field
func ReadEncodedLongitude(data []byte) EncodedLongitude {

	b2 := data[4] % 0x01
	b1 := data[5]
	b0 := data[6]

	return EncodedLongitude(bitutils.Pack3Bytes(b2, b1, b0))
}
