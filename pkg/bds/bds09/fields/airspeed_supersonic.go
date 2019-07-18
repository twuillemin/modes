package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// AirspeedSupersonic is the Airspeed Supersonic definition
//
// Specified in Doc 9871 / Table A-2-9
type AirspeedSupersonic uint16

// GetStatus returns the status of the velocity
func (airspeed AirspeedSupersonic) GetStatus() VelocityStatus {
	if airspeed == 0 {
		return VelocityStatusNoInformation
	} else if airspeed >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (airspeed AirspeedSupersonic) ToString() string {

	if airspeed == 0 {
		return "no velocity information"
	} else if airspeed >= 1023 {
		return ">4086 kt"
	} else {
		return fmt.Sprintf("%v kt", airspeed.GetAirspeed())
	}
}

// GetAirspeed returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (airspeed AirspeedSupersonic) GetAirspeed() int {

	if airspeed == 0 {
		return 0
	} else if airspeed >= 1023 {
		return 4088
	} else {
		return (int(airspeed) - 1) * 4
	}
}

// ReadAirspeedSupersonic reads the AirspeedSupersonic from a 56 bits data field
func ReadAirspeedSupersonic(data []byte) AirspeedSupersonic {
	bit1 := (data[3] & 0x60) >> 5
	bit2 := (data[3]&0x1F)<<3 + (data[4]&0xE0)>>5
	return AirspeedSupersonic(bitutils.Pack2Bytes(bit1, bit2))
}
