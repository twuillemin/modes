package fields

func ReadIcing(data []byte) (bool, HazardLevel) {
	status := (data[1] & 0x40) != 0
	level := HazardLevel((data[1] & 0x30) >> 4)
	return status, level
}
