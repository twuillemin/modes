package reader

// 	crcPolynomial is the polynomial for the CRC redundancy check
//  Note: we assume that the degree of the polynomial is divisible by 8 (holds for Mode S) and the msb is left out
//
// Values defined according to Annex 10 V4
var crcPolynomial = []uint8{0xFF, 0xF4, 0x09}

// computeParity computes the parity of a slice of byte as 3-byte array. We used the implementation from:
// http://www.eurocontrol.int/eec/gallery/content/public/document/eec/report/1994/022_CRC_calculations_for_Mode_S.pdf
//
// params:
//    - data: The data for which to compute parity
//
// Returns the CRC (3 bytes)
func computeParity(data []byte) []byte {

	crcLength := len(crcPolynomial)

	// Initialize with the beginning of the message
	crc := make([]byte, crcLength)
	for i := 0; i < crcLength; i++ {
		crc[i] = data[i]
	}

	// For all bit
	for i := 0; i < len(data)*8; i++ {

		// Keep msb
		invert := (crc[0] & 0x80) != 0

		// Shift left
		crc[0] <<= 1
		for b := 1; b < crcLength; b++ {
			crc[b-1] |= (crc[b] >> 7) & 0x1
			crc[b] <<= 1
		}

		// Get next bit from message
		byteIdx := (crcLength*8 + i) / 8
		bitShift := uint(7 - (i % 8))
		if byteIdx < len(data) {
			crc[len(crc)-1] |= (data[byteIdx] >> bitShift) & 0x1
		}

		// xor
		if invert {
			for b := 0; b < crcLength; b++ {
				crc[b] ^= crcPolynomial[b]
			}
		}
	}

	return crc

}

// xor applies a xor operation between two arrays. If the size of the two array is different, the operation is only
// using the smallest size
//
// params:
//    - array1: The first array
//    - array2: The second array
//
// Return the result of the xor
func xorArrays(array1 []uint8, array2 []uint8) []uint8 {

	// Get the smallest size
	size := len(array1)
	if sizeArray2 := len(array2); sizeArray2 < size {
		size = sizeArray2
	}

	result := make([]uint8, size)

	for i := 0; i < size; i++ {
		result[i] = array1[i] ^ array2[i]
	}

	return result
}
