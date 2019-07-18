package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// VelocityNSNormal is the Velocity NS Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type VelocityNSNormal uint16

// GetStatus returns the status of the velocity
func (velocity VelocityNSNormal) GetStatus() VelocityStatus {
	if velocity == 0 {
		return VelocityStatusNoInformation
	} else if velocity >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (velocity VelocityNSNormal) ToString() string {

	if velocity == 0 {
		return "no velocity information"
	} else if velocity >= 1023 {
		return ">1021.5 kt"
	} else {
		return fmt.Sprintf("%v kt", velocity.GetVelocity())
	}
}

// GetVelocity returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (velocity VelocityNSNormal) GetVelocity() int {

	if velocity == 0 {
		return 0
	} else if velocity >= 1023 {
		return 1022
	} else {
		return int(velocity) - 1
	}
}

// ReadVelocityNSNormal reads the VelocityNSNormal from a 56 bits data field
func ReadVelocityNSNormal(data []byte) VelocityNSNormal {
	bit1 := (data[3] & 0x60) >> 5
	bit2 := (data[3]&0x1F)<<3 + (data[4]&0xE0)>>5
	return VelocityNSNormal(bitutils.Pack2Bytes(bit1, bit2))
}
