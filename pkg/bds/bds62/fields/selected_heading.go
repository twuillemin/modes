package fields

import (
	"fmt"
)

// SelectedHeading is the Selected Altitude definition
//
// Specified in Doc 9871 / C.2.3.9.8
type SelectedHeading byte

// ToString returns a basic, but readable, representation of the field
func (selectedHeading SelectedHeading) ToString(status SelectedHeadingStatus, sign SelectedHeadingSign) string {

	if status == SHSInvalid {
		return status.ToString()
	}

	return fmt.Sprintf("%v degrees", selectedHeading.GetSelectedHeading(status, sign))

}

// GetSelectedHeading returns the SelectedHeading. Note that the returned value will be the 0 for SASInvalid
func (selectedHeading SelectedHeading) GetSelectedHeading(status SelectedHeadingStatus, sign SelectedHeadingSign) float64 {

	if status == SHSInvalid {
		return 0
	}

	absAngle := float64(selectedHeading) * 0.703125

	if sign == SHSPositive {
		return absAngle
	}

	return -absAngle
}

// ReadSelectedHeading reads the SelectedHeading from a 56 bits data field
func ReadSelectedHeading(data []byte) SelectedHeading {
	bits := (data[3]&0x01)<<7 + (data[4]&0xFE)>>1
	return SelectedHeading(bits)
}
