package fields

import "fmt"

// SurveillanceIntegrityLevelSupplement is the Single Antenna Flag definition
//
// Specified in Doc 9871 / C.2.3.10.15
type SurveillanceIntegrityLevelSupplement byte

const (
	// SILSByHour indicates Probability of exceeding NIC radius of containment is based on "per sample"
	SILSByHour SurveillanceIntegrityLevelSupplement = 0
	// SILSBySample indicates Probability of exceeding NIC radius of containment is based on "per sample"
	SILSBySample SurveillanceIntegrityLevelSupplement = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SurveillanceIntegrityLevelSupplement) ToString() string {

	switch status {
	case SILSByHour:
		return "0 - Probability of exceeding NIC radius of containment is based on \"per sample\""
	case SILSBySample:
		return "1 - Probability of exceeding NIC radius of containment is based on \"per sample\""
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSurveillanceIntegrityLevelSupplement reads the SurveillanceIntegrityLevelSupplement from a 56 bits data field
func ReadSurveillanceIntegrityLevelSupplement(data []byte) SurveillanceIntegrityLevelSupplement {
	bits := (data[6] & 0x02) >> 1
	return SurveillanceIntegrityLevelSupplement(bits)
}
