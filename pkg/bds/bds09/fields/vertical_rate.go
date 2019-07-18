package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// VerticalRateStatus is the status of the VerticalRate
type VerticalRateStatus int

const (
	// VerticalRateStatusNoInformation indicates no vertical rate
	VerticalRateStatusNoInformation VerticalRateStatus = 0
	// VerticalRateStatusRegular indicates that the VerticalRate is computed on the linear scale value of field * factor
	VerticalRateStatusRegular VerticalRateStatus = 1
	// VerticalRateStatusMaximum indicates that the VerticalRate field value indicates vertical rate greater the maximum of the scale
	VerticalRateStatusMaximum VerticalRateStatus = 2
)

// VerticalRate is the Velocity EW Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type VerticalRate uint16

// GetStatus returns the status of the velocity
func (verticalRate VerticalRate) GetStatus() VerticalRateStatus {
	if verticalRate == 0 {
		return VerticalRateStatusNoInformation
	} else if verticalRate >= 511 {
		return VerticalRateStatusMaximum
	} else {
		return VerticalRateStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (verticalRate VerticalRate) ToString() string {

	if verticalRate == 0 {
		return "no vertical rate information"
	} else if verticalRate >= 511 {
		return ">32608 ft/min"
	} else {
		return fmt.Sprintf("%v ft/min", verticalRate.GetVerticalRate())
	}
}

// GetVerticalRate returns the vertical rate. Note that the returned value will be 0 for VerticalRateStatusNoInformation and
// the maximum for VerticalRateStatusMaximum
func (verticalRate VerticalRate) GetVerticalRate() int {

	if verticalRate == 0 {
		return 0
	} else if verticalRate >= 511 {
		return 32640
	} else {
		return (int(verticalRate) - 1) * 64
	}
}

// ReadVerticalRate reads the VerticalRate from a 56 bits data field
func ReadVerticalRate(data []byte) VerticalRate {
	bit1 := (data[4] & 0x04) >> 2
	bit2 := (data[4]&0x03)<<6 + (data[5]&0xFC)>>2
	return VerticalRate(bitutils.Pack2Bytes(bit1, bit2))
}
