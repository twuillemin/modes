package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadTurbulenceV0(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x80) != 0
	level := HazardLevel((data[0] & 0x60) >> 5)
	return status, level
}

func ReadTurbulenceV1(data []byte) (bool, float32, float32, byte) {
	status := (data[4] & 0x04) != 0

	byte1 := data[4] & 0x03
	byte2 := data[5] & 0xF0
	average := float32(bitutils.Pack2Bytes(byte1, byte2)>>4) * 0.2

	if average > 1.26 {
		average = 1.26
	}

	byte1 = data[5] & 0x0F
	byte2 = data[6] & 0xC0
	peak := float32(bitutils.Pack2Bytes(byte1, byte2)>>4) * 0.2

	if peak > 1.26 {
		peak = 1.26
	}

	byte1 = data[5] & 0x0F
	byte2 = data[6] & 0xC0
	interval := (data[6] & 0x3C) >> 2

	return status, average, peak, interval
}
