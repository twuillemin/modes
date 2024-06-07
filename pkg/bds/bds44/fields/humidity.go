package fields

func ReadHumidity(data []byte) (bool, uint32) {
	status := (data[6] & 0x80) != 0

	humidity := uint32(data[6] & 0x7F)

	return status, humidity
}
