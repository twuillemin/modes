package fields

import "fmt"

// SurveillanceStatus is the surveillance status
//
// Specified in Doc 9871 / A.2.3.2.6
type SurveillanceStatus byte

const (
	// SSNoCondition indicates No aircraft category information
	SSNoCondition SurveillanceStatus = 0
	// SSPermanentAlert indicates Permanent alert (emergency condition)
	SSPermanentAlert SurveillanceStatus = 1
	// SSTemporaryAlert indicates temporary  alert (change  in Mode  A identity code other than emergency condition)
	SSTemporaryAlert SurveillanceStatus = 2
	// SSSPICondition indicates SPI condition
	SSSPICondition SurveillanceStatus = 3
)

// ToString returns a basic, but readable, representation of the field
func (surveillanceStatus SurveillanceStatus) ToString() string {

	switch surveillanceStatus {
	case SSNoCondition:
		return "0 - no status information"
	case SSPermanentAlert:
		return "1 - permanent alert condition"
	case SSTemporaryAlert:
		return "2 - temporary alert condition"
	case SSSPICondition:
		return "3 - SPI condition"
	default:
		return fmt.Sprintf("%v - Unknown code", surveillanceStatus)
	}
}

// ReadSurveillanceStatus reads the SurveillanceStatus from a 56 bits data field
func ReadSurveillanceStatus(data []byte) SurveillanceStatus {
	bits := (data[0] & 0x06) >> 1
	return SurveillanceStatus(bits)
}
