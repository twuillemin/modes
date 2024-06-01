package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// BarometricPressureSetting is the Selected Altitude definition
//
// Specified in Doc 9871 / C.2.3.9.5
type BarometricPressureSetting float64

// GetStatus returns the status of the altitude
func (barometricPressureSetting BarometricPressureSetting) GetStatus() BarometricPressureSettingStatus {
	if barometricPressureSetting == 0 {
		return BPSInvalid
	}
	return BPSValid
}

// ToString returns a basic, but readable, representation of the field
func (barometricPressureSetting BarometricPressureSetting) ToString() string {

	if barometricPressureSetting == 0 {
		return "0 - no data or invalid"
	}

	return fmt.Sprintf("%v feet", barometricPressureSetting.GetBarometricPressureSetting())

}

// GetBarometricPressureSetting returns the BarometricPressureSetting. Note that the returned value will be the 0 for SASInvalid
func (barometricPressureSetting BarometricPressureSetting) GetBarometricPressureSetting() float64 {

	if barometricPressureSetting == 0 {
		return 0
	}

	return (float64(barometricPressureSetting) - 1) * 0.8
}

// ReadBarometricPressureSetting reads the BarometricPressureSetting from a 56 bits data field
func ReadBarometricPressureSetting(data []byte) BarometricPressureSetting {
	bits1 := (data[2] & 0x0F) >> 3
	bits2 := (data[2]&0x07)<<5 + (data[3]&0xF8)>>3
	return BarometricPressureSetting(bitutils.Pack2Bytes(bits1, bits2))
}
