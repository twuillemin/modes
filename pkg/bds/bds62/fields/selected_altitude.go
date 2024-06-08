package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadSelectedAltitude reads the SelectedAltitude from a 56 bits data field
func ReadSelectedAltitude(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[1] & 0x7F
	byte2 := data[2] & 0xF0
	allByte := bitutils.Pack2Bytes(byte1, byte2) >> 4

	if allByte == 0 {
		return 0, NVSNoInformation
	}

	return (allByte - 1) * 32, NVSRegular
}
