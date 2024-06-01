package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// TargetHeadingTrackAngle is the Target Altitude definition
//
// Specified in Doc 9871 / B.2.3.9.9
type TargetHeadingTrackAngle uint16

// GetStatus returns the status of the target heading / track
func (targetAngle TargetHeadingTrackAngle) GetStatus() TargetHeadingTrackStatus {
	if targetAngle >= 1011 {
		return THTSInvalid
	}
	return THTSValid
}

// ToString returns a basic, but readable, representation of the field
func (targetAngle TargetHeadingTrackAngle) ToString() string {

	if targetAngle >= 360 {
		return "Invalid (out of range)"
	}

	return fmt.Sprintf("%v degrees", targetAngle.GetTargetHeadingTrackAngle())

}

// GetTargetHeadingTrackAngle returns the TargetHeadingTrackAngle. Note that the returned value will be the maximum
// for THTSValid
func (targetAngle TargetHeadingTrackAngle) GetTargetHeadingTrackAngle() int {

	if targetAngle >= 360 {
		return 360
	}

	return int(targetAngle)
}

// ReadTargetHeadingTrackAngle reads the TargetHeadingTrackAngle from a 56 bits data field
func ReadTargetHeadingTrackAngle(data []byte) TargetHeadingTrackAngle {
	bit1 := (data[3] & 0x10) >> 4
	bit2 := (data[3]&0x0F)<<4 + (data[4]&0xF0)>>4
	return TargetHeadingTrackAngle(bitutils.Pack2Bytes(bit1, bit2))
}
