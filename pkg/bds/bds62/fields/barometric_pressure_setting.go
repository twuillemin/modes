package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadBarometricPressureSetting reads the Barometric Pressure Setting from a 56 bits data field
func ReadBarometricPressureSetting(data []byte) (float32, NumericValueStatus) {
	byte1 := data[2] & 0x0F
	byte2 := data[3] & 0xF8
	allByte := bitutils.Pack2Bytes(byte1, byte2) >> 3

	if allByte == 0 {
		return 0.0, NVSNoInformation
	}

	return (float32(allByte-1) * 0.8) + 800, NVSRegular
}
