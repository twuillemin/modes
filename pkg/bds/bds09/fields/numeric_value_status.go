package fields

import "fmt"

// NumericValueStatus is the status of the numeric values, such as air speed, velocity, etc.
type NumericValueStatus int

const (
	// NVSNoInformation indicates no velocity information
	NVSNoInformation NumericValueStatus = 0
	// NVSRegular indicates that the Velocity is computed on the linear scale value of field * factor
	NVSRegular NumericValueStatus = 1
	// NVSMaximum indicates that the Velocity field value indicates velocity greater the maximum of the scale
	NVSMaximum NumericValueStatus = 2
)

// ToString returns a basic, but readable, representation of the field
func (vs NumericValueStatus) ToString() string {

	switch vs {
	case NVSNoInformation:
		return "No information"
	case NVSRegular:
		return "Regular value"
	case NVSMaximum:
		return "Maximum"
	default:
		return fmt.Sprintf("%v - Unknown code", vs)
	}
}
