package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadEncodedLatitude reads the EncodedLatitude from a 56 bits data field
func ReadEncodedLatitude(data []byte) uint32 {

	b2 := data[2] & 0x03
	b1 := data[3]
	b0 := data[4] & 0xFE

	return bitutils.Pack3Bytes(b2, b1, b0) >> 1
}
