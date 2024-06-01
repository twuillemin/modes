package fields

import (
	"fmt"
)

// DifferenceGNSSBaroStatus is the status of the DifferenceGNSSBaro
type DifferenceGNSSBaroStatus int

const (
	// DGBSNoInformation indicates no vertical rate
	DGBSNoInformation DifferenceGNSSBaroStatus = 0
	// DGBSRegular indicates that the DifferenceGNSSBaro is computed on the linear scale value of field * factor
	DGBSRegular DifferenceGNSSBaroStatus = 1
	// DGBSMaximum indicates that the DifferenceGNSSBaro field value indicates vertical rate greater the maximum of the scale
	DGBSMaximum DifferenceGNSSBaroStatus = 2
)

// DifferenceGNSSBaro is the Velocity EW Normal definition
//
// Specified in Doc 9871 / Table A-2-9
type DifferenceGNSSBaro byte

// GetStatus returns the status of the velocity
func (DifferenceGNSSBaro DifferenceGNSSBaro) GetStatus() DifferenceGNSSBaroStatus {
	if DifferenceGNSSBaro == 0 {
		return DGBSNoInformation
	} else if DifferenceGNSSBaro >= 127 {
		return DGBSMaximum
	} else {
		return DGBSRegular
	}
}

// ToString returns a basic, but readable, representation of the field
func (DifferenceGNSSBaro DifferenceGNSSBaro) ToString() string {

	if DifferenceGNSSBaro == 0 {
		return "no information"
	} else if DifferenceGNSSBaro >= 127 {
		return ">3137.5 ft"
	} else {
		return fmt.Sprintf("%v ft", DifferenceGNSSBaro.GetDifferenceGNSSBaro())
	}
}

// GetDifferenceGNSSBaro returns the vertical rate. Note that the returned value will be 0 for DGBSNoInformation and
// the maximum for DGBSMaximum
func (DifferenceGNSSBaro DifferenceGNSSBaro) GetDifferenceGNSSBaro() int {

	if DifferenceGNSSBaro == 0 {
		return 0
	} else if DifferenceGNSSBaro >= 127 {
		return 3150
	} else {
		return (int(DifferenceGNSSBaro) - 1) * 25
	}
}

// ReadDifferenceGNSSBaro reads the DifferenceGNSSBaro from a 56 bits data field
func ReadDifferenceGNSSBaro(data []byte) DifferenceGNSSBaro {
	return DifferenceGNSSBaro(data[6] & 0x7F)
}
