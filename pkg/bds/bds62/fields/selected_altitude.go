package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// SelectedAltitude is the Selected Altitude definition
//
// Specified in Doc 9871 / C.2.3.9.4
type SelectedAltitude uint16

// GetStatus returns the status of the altitude
func (selectedAltitude SelectedAltitude) GetStatus() SelectedAltitudeStatus {
	if selectedAltitude == 0 {
		return SASInvalid
	}
	return SASValid
}

// ToString returns a basic, but readable, representation of the field
func (selectedAltitude SelectedAltitude) ToString() string {

	if selectedAltitude == 1011 {
		return "0 - no data or invalid"
	}

	return fmt.Sprintf("%v feet", selectedAltitude.GetSelectedAltitude())

}

// GetSelectedAltitude returns the SelectedAltitude. Note that the returned value will be the 0 for SASInvalid
func (selectedAltitude SelectedAltitude) GetSelectedAltitude() int {

	if selectedAltitude == 0 {
		return 0
	}

	return (int(selectedAltitude) - 1) * 32
}

// ReadSelectedAltitude reads the SelectedAltitude from a 56 bits data field
func ReadSelectedAltitude(data []byte) SelectedAltitude {
	bits1 := (data[1] & 0x70) >> 4
	bits2 := (data[1]&0x0F)<<4 + (data[2]&0xF0)>>4
	return SelectedAltitude(bitutils.Pack2Bytes(bits1, bits2))
}
