package fields

import "fmt"

// NICBaro is the NIC Baro definition
//
// Specified in Doc 9871 / B.2.3.10.10
type NICBaro byte

const (
	// NICBGilhamNotCrossChecked indicates that the barometric altitude that is being reported in the Airborne
	// Position Message is based on a Gilham coded input that has not been cross-checked against another source of
	// pressure-altitude
	NICBGilhamNotCrossChecked NICBaro = 0
	// NICBGilhamCrossCheckedOrNonGilham indicates that the barometric altitude that is being reported in the Airborne
	// Position Message is either based on a Gilham code input that has been cross-checked against another source of
	// pressure-altitude and verified as being consistent, or is based on a non-Gilham coded source
	NICBGilhamCrossCheckedOrNonGilham NICBaro = 1
)

// ToString returns a basic, but readable, representation of the field
func (baro NICBaro) ToString() string {

	switch baro {
	case NICBGilhamNotCrossChecked:
		return "0 - barometric altitude based on a Gilham input that has not been cross-checked"
	case NICBGilhamCrossCheckedOrNonGilham:
		return "1 - barometric altitude based on a Gilham input that has been cross-checked or on a non Gilham input"
	default:
		return fmt.Sprintf("%v - Unknown code", baro)
	}
}

// ReadNICBaro reads the NICBaro from a 56 bits data field
func ReadNICBaro(data []byte) NICBaro {
	bits := (data[6] & 0x08) >> 3
	return NICBaro(bits)
}
