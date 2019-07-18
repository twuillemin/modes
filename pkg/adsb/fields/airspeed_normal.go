package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// AirspeedNormal is the Airspeed Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type AirspeedNormal uint16

// GetStatus returns the status of the velocity
func (airspeed AirspeedNormal) GetStatus() VelocityStatus {
	if airspeed == 0 {
		return VelocityStatusNoInformation
	} else if airspeed >= 1023 {
		return VelocityStatusMaximum
	} else {
		return VelocityStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (airspeed AirspeedNormal) ToString() string {

	if airspeed == 0 {
		return "no airspeed information"
	} else if airspeed >= 1023 {
		return ">1021.5 kt"
	} else {
		return fmt.Sprintf("%v kt", airspeed.GetAirspeed())
	}
}

// GetAirspeed returns the airspeed. Note that the returned value will be 0 for VelocityStatusNoInformation and
// the maximum for VelocityMaximum
func (airspeed AirspeedNormal) GetAirspeed() int {

	if airspeed == 0 {
		return 0
	} else if airspeed >= 1023 {
		return 1022
	} else {
		return int(airspeed) - 1
	}
}

// ReadAirspeedNormal reads the AirspeedNormal from a 56 bits data field
func ReadAirspeedNormal(data []byte) AirspeedNormal {
	bit1 := (data[3] & 0x60) >> 5
	bit2 := (data[3]&0x1F)<<3 + (data[4]&0xE0)>>5
	return AirspeedNormal(bitutils.Pack2Bytes(bit1, bit2))
}
