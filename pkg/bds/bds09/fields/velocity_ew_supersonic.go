package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// VelocityEWSupersonic is the Velocity EW Supersonic definition
//
// Specified in Doc 9871 / Table A-2-9
type VelocityEWSupersonic uint16

// GetStatus returns the status of the velocity
func (velocity VelocityEWSupersonic) GetStatus() VelocityStatus {
	if velocity == 0 {
		return VelocityStatusNoInformation
	} else if velocity >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (velocity VelocityEWSupersonic) ToString() string {

	if velocity == 0 {
		return "no velocity information"
	} else if velocity >= 1023 {
		return ">4086 kt"
	} else {
		return fmt.Sprintf("%v kt", velocity.GetVelocity())
	}
}

// GetVelocity returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (velocity VelocityEWSupersonic) GetVelocity() int {

	if velocity == 0 {
		return 0
	} else if velocity >= 1023 {
		return 4088
	} else {
		return (int(velocity) - 1) * 4
	}
}

// ReadVelocityEWSupersonic reads the VelocityEWSupersonic from a 56 bits data field
func ReadVelocityEWSupersonic(data []byte) VelocityEWSupersonic {
	bit1 := data[1] & 0x03
	bit2 := data[2]
	return VelocityEWSupersonic(bitutils.Pack2Bytes(bit1, bit2))
}
