package fields

func ReadTurbulence(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x80) != 0
	level := HazardLevel((data[0] & 0x60) >> 5)
	return status, level
}
