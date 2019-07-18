package fields

import (
	"fmt"
)

// DifferenceGNSSBaroStatus is the status of the DifferenceGNSSBaro
type DifferenceGNSSBaroStatus int

const (
	// DifferenceGNSSBaroStatusNoInformation indicates no vertical rate
	DifferenceGNSSBaroStatusNoInformation DifferenceGNSSBaroStatus = 0
	// DifferenceGNSSBaroStatusRegular indicates that the DifferenceGNSSBaro is computed on the linear scale value of field * factor
	DifferenceGNSSBaroStatusRegular DifferenceGNSSBaroStatus = 1
	// DifferenceGNSSBaroStatusMaximum indicates that the DifferenceGNSSBaro field value indicates vertical rate greater the maximum of the scale
	DifferenceGNSSBaroStatusMaximum DifferenceGNSSBaroStatus = 2
)

// DifferenceGNSSBaro is the Velocity EW Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type DifferenceGNSSBaro uint16

// GetStatus returns the status of the velocity
func (DifferenceGNSSBaro DifferenceGNSSBaro) GetStatus() DifferenceGNSSBaroStatus {
	if DifferenceGNSSBaro == 0 {
		return DifferenceGNSSBaroStatusNoInformation
	} else if DifferenceGNSSBaro >= 127 {
		return DifferenceGNSSBaroStatusMaximum
	} else {
		return DifferenceGNSSBaroStatusRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (DifferenceGNSSBaro DifferenceGNSSBaro) ToString() string {

	if DifferenceGNSSBaro == 0 {
		return "no information"
	} else if DifferenceGNSSBaro >= 511 {
		return ">3137.5 ft"
	} else {
		return fmt.Sprintf("%v ft", DifferenceGNSSBaro.GetDifferenceGNSSBaro())
	}
}

// GetDifferenceGNSSBaro returns the vertical rate. Note that the returned value will be 0 for DifferenceGNSSBaroStatusNoInformation and
// the maximum for DifferenceGNSSBaroStatusMaximum
func (DifferenceGNSSBaro DifferenceGNSSBaro) GetDifferenceGNSSBaro() int {

	if DifferenceGNSSBaro == 0 {
		return 0
	} else if DifferenceGNSSBaro >= 511 {
		return 3150
	} else {
		return (int(DifferenceGNSSBaro) - 1) * 25
	}
}

// ReadDifferenceGNSSBaro reads the DifferenceGNSSBaro from a 56 bits data field
func ReadDifferenceGNSSBaro(data []byte) DifferenceGNSSBaro {
	return DifferenceGNSSBaro(data[6] & 0x7F)
}
