package bitutils

// Pack2Bytes packs 2 bytes in a single uint16.
//
// Params:
//    - b1: the first bit
//    - b0: the second bit
//
// The two bytes are packed as b1 b1 ... b1 b1 b0 b0 ... b0 b0
//
// Return the packed byte
func Pack2Bytes(b1 byte, b0 byte) uint16 {
	return uint16(b1)<<8 + uint16(b0)
}

// Pack3Bytes packs 3 bytes in a single uint32.
//
// Params:
//    - b2: the first bit
//    - b1: the second bit
//    - b0: the third bit
//
// The two bytes are packed as b2 b2 ... b2 b2 b1 b1 ... b1 b1 b0 b0 ... b0 b0
//
// Return the packed byte
func Pack3Bytes(b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}

// Pack4Bytes packs 4 bytes in a single uint32.
//
// Params:
//    - b3: the first bit
//    - b2: the second bit
//    - b1: the third bit
//    - b0: the fourth bit
//
// The two bytes are packed as b3 ... b3 b2 ... b2 b1 ... b1 b0 ... b0
//
// Return the packed byte
func Pack4Bytes(b3 byte, b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b3)<<24 + uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}
