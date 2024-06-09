package fields

func ReadWindShearV0(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x10) != 0
	level := HazardLevel((data[0] & 0x0C) >> 2)
	return status, level
}

func ReadWindShearV1(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x80) != 0
	level := HazardLevel((data[0] & 0x60) >> 5)
	return status, level
}
