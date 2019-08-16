package bitutils

// GrayToBinary8 converts the given byte from gray code to "classical" binary
//
// Params:
//    - gray: The gray byte to convert
//
// Return the "standard" binary representation
func GrayToBinary8(gray byte) byte {

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

// GrayToBinary16 converts the given uint16 from gray code to "classical" binary
//
// Params:
//    - gray: The gray byte to convert
//
// Return the "standard" binary representation
func GrayToBinary16(gray uint16) uint16 {

	result := uint16(0)

	for i := 15; i >= 0; i-- {
		index := uint(i)
		previousComputedBit := (0x01 << (index + 1)) & result
		currentInputBit := (0x01 << index) & gray
		currentResultBit := previousComputedBit>>1 ^ currentInputBit
		result = result | currentResultBit
	}

	return result
}
