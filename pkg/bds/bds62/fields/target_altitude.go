package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// TargetAltitude is the Target Altitude definition
//
// Specified in Doc 9871 / B.2.3.9.7
type TargetAltitude uint16

// GetStatus returns the status of the altitude
func (targetAltitude TargetAltitude) GetStatus() TargetAltitudeStatus {
	if targetAltitude >= 1011 {
		return TASInvalid
	}
	return TASValid
}

// ToString returns a basic, but readable, representation of the field
func (targetAltitude TargetAltitude) ToString() string {

	if targetAltitude >= 1011 {
		return "Invalid (out of range)"
	}

	return fmt.Sprintf("%v feet", targetAltitude.GetTargetAltitude())

}

// GetTargetAltitude returns the TargetAltitude. Note that the returned value will be the maximum for TASInvalid
func (targetAltitude TargetAltitude) GetTargetAltitude() int {

	if targetAltitude >= 1011 {
		return 100100
	}

	return 100*int(targetAltitude) - 1000
}

// ReadTargetAltitude reads the TargetAltitude from a 56 bits data field
func ReadTargetAltitude(data []byte) TargetAltitude {
	bit1 := (data[1]&0x01)<<1 + (data[2]&0x80)>>7
	bit2 := (data[2]&0x7F)<<1 + (data[3]&0x80)>>7
	return TargetAltitude(bitutils.Pack2Bytes(bit1, bit2))
}
