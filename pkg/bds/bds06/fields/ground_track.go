package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadGroundTrack reads the GroundTrack from a 56 bits data field
func ReadGroundTrack(data []byte) (float32, bool) {

	status := (data[1] & 0x08) != 0

	byte1 := data[1] & 0x07
	byte2 := data[2] & 0xF0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 4

	groundTrack := float32(allBits) * 360.0 / 128.0

	return groundTrack, status
}
