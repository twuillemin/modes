package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadStaticAirTemperature(data []byte) float32 {

	negative := (data[3] & 0x01) != 0

	byte1 := data[3] & 0xFF
	byte2 := data[4] & 0xC0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 2
	temperature := float32(allBits) * 0.25

	if negative {
		temperature -= temperature
	}

	return temperature
}
