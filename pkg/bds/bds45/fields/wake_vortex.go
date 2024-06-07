package fields

func ReadWakeVortex(data []byte) (bool, HazardLevel) {
	status := (data[1] & 0x08) != 0
	level := HazardLevel((data[1] & 0x06) >> 1)
	return status, level
}
