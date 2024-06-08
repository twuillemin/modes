package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadAirspeedNormal reads the AirSpeed Normal from a 56 bits data field
func ReadAirspeedNormal(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	velocity := bitutils.Pack2Bytes(byte1, byte2) >> 5

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 1023, NVSMaximum
	}

	return velocity - 1, NVSRegular
}

// ReadAirspeedSupersonic reads the AirSpeed Supersonic from a 56 bits data field
func ReadAirspeedSupersonic(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	velocity := bitutils.Pack2Bytes(byte1, byte2) >> 5

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 4088, NVSMaximum
	}

	return (velocity - 1) * 4, NVSRegular
}
