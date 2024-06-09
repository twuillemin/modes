package fields

func ReadWakeVortexV0(data []byte) (bool, HazardLevel) {
	status := (data[1] & 0x08) != 0
	level := HazardLevel((data[1] & 0x06) >> 1)
	return status, level
}

func ReadWakeVortexV1(data []byte) (bool, HazardLevel) {
	status := (data[1] & 0x40) != 0
	level := HazardLevel((data[1] & 0x30) >> 4)
	return status, level
}
