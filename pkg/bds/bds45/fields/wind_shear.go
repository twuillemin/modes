package fields

func ReadWindShear(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x10) != 0
	level := HazardLevel((data[0] & 0x0C) >> 2)
	return status, level
}
