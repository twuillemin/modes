package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadEncodedLongitude reads the EncodedLongitude from a 56 bits data field
func ReadEncodedLongitude(data []byte) uint32 {

	b2 := data[4] & 0x01
	b1 := data[5]
	b0 := data[6]

	return bitutils.Pack3Bytes(b2, b1, b0)
}
