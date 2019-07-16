package bitutils

func Pack2Bytes(b1 byte, b0 byte) uint16 {
	return uint16(b1)<<8 + uint16(b0)
}

func Pack3Bytes(b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}

func Pack4Bytes(b3 byte, b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b3)<<24 + uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}
