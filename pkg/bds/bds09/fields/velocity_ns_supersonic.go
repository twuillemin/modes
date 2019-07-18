package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// VelocityNSSupersonic is the Velocity NS Supersonic definition
//
// Specified in Doc 9871 / Table A-2-9
type VelocityNSSupersonic uint16

// GetStatus returns the status of the velocity
func (velocity VelocityNSSupersonic) GetStatus() VelocityStatus {
	if velocity == 0 {
		return VelocityStatusNoInformation
	} else if velocity >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (velocity VelocityNSSupersonic) ToString() string {

	if velocity == 0 {
		return "no velocity information"
	} else if velocity >= 1023 {
		return ">4086 kt"
	} else {
		return fmt.Sprintf("%v kt", velocity.GetVelocity())
	}
}

// GetVelocityValue returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (velocity VelocityNSSupersonic) GetVelocity() int {

	if velocity == 0 {
		return 0
	} else if velocity >= 1023 {
		return 4088
	} else {
		return (int(velocity) - 1) * 4
	}
}

// ReadVelocityNSSupersonic reads the VelocityNSSupersonic from a 56 bits data field
func ReadVelocityNSSupersonic(data []byte) VelocityNSSupersonic {
	bit1 := (data[3] & 0x60) >> 5
	bit2 := (data[3]&0x1F)<<3 + (data[4]&0xE0)>>5
	return VelocityNSSupersonic(bitutils.Pack2Bytes(bit1, bit2))
}
