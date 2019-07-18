package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// VelocityEWNormal is the Velocity EW Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type VelocityEWNormal uint16

// GetStatus returns the status of the velocity
func (velocity VelocityEWNormal) GetStatus() VelocityStatus {
	if velocity == 0 {
		return VelocityStatusNoInformation
	} else if velocity >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (velocity VelocityEWNormal) ToString() string {

	if velocity == 0 {
		return "no velocity information"
	} else if velocity >= 1023 {
		return ">1021.5 kt"
	} else {
		return fmt.Sprintf("%v kt", velocity.GetVelocity())
	}
}

// GetVelocityValue returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (velocity VelocityEWNormal) GetVelocity() int {

	if velocity == 0 {
		return 0
	} else if velocity >= 1023 {
		return 1022
	} else {
		return int(velocity) - 1
	}
}

// ReadVelocityEWNormal reads the VelocityEWNormal from a 56 bits data field
func ReadVelocityEWNormal(data []byte) VelocityEWNormal {
	bit1 := data[1] & 0x03
	bit2 := data[2]
	return VelocityEWNormal(bitutils.Pack2Bytes(bit1, bit2))
}
