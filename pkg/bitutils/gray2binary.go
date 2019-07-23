package bitutils

// GrayToBinary converts the given bits (b0 being the MSB and b7 the LSB) from gray code to "classical" binary
func GrayToBinary(b0, b1, b2, b3, b4, b5, b6, b7 bool) byte {

	gray := byte(0)
	if b0 {
		gray |= 0x80
	}
	if b1 {
		gray |= 0x40
	}
	if b2 {
		gray |= 0x20
	}
	if b3 {
		gray |= 0x10
	}
	if b4 {
		gray |= 0x08
	}
	if b5 {
		gray |= 0x04
	}
	if b6 {
		gray |= 0x02
	}
	if b7 {
		gray |= 0x01
	}

	result := byte(0)

	for i := 7; i >= 0; i-- {
		index := uint(i)
		previousComputedBit := (0x01 << (index + 1)) & result
		currentInputBit := (0x01 << index) & gray
		currentResultBit := previousComputedBit>>1 ^ currentInputBit
		result = result | currentResultBit
	}

	return result
}
