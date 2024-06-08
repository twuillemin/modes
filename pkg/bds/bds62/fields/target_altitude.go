package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// TargetAltitude is the Target Altitude definition
//
// Specified in Doc 9871 / B.2.3.9.7

// ReadTargetAltitude reads the TargetAltitude from a 56 bits data field
func ReadTargetAltitude(data []byte) (int32, NumericValueStatus) {
	byte1 := data[1] & 0x01
	byte2 := data[2] & 0xFF
	byte3 := data[3] & 0x80
	targetAltitude := int32(bitutils.Pack3Bytes(byte1, byte2, byte3) >> 7)

	if targetAltitude > 1010 {
		return 0, NVSMaximum
	}

	return (targetAltitude * 100) - 1000, NVSRegular
}
