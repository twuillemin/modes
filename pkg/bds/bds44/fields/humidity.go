package fields

func ReadHumidityV0(data []byte) (bool, float32) {
	status := (data[6] & 0x40) != 0
	humidity := float32(data[6]&0x3F) * 100 / 64

	return status, humidity
}

func ReadHumidityV1(data []byte) (bool, float32) {
	status := (data[6] & 0x80) != 0
	humidity := float32(data[6] & 0x7F)

	return status, humidity
}
