package fields

import "fmt"

// SurveillanceIntegrityLevel is the Surveillance Integrity Level definition
//
// Specified in Doc 9871 / B.2.3.9.14
type SurveillanceIntegrityLevel byte

const (
	// SIL0 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: unknown
	//     - Probability of exceeding the Vertical Integrity CR without notification: unknown
	SIL0 SurveillanceIntegrityLevel = 0
	// SIL1 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-3 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-3 per flight hour or per sample
	SIL1 SurveillanceIntegrityLevel = 1
	// SIL2 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-5 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-5 per flight hour or per sample
	SIL2 SurveillanceIntegrityLevel = 2
	// SIL3 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-7 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 2 * 10^-7 per 150 seconds or per sample
	SIL3 SurveillanceIntegrityLevel = 3
)

var silLabelHorizontal = "Probability of exceeding the Horizontal Containment Radius reported in the NIC Subfield without an indication: "
var silLabelVertical = "Probability of exceeding the Vertical Integrity Containment Region (VPL) without an indication:                "

// ToString returns a basic, but readable, representation of the field
func (level SurveillanceIntegrityLevel) ToString() string {

	switch level {
	case SIL0:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "Unknown",
			silLabelVertical, "Unknown")
	case SIL1:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-3 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-3 per flight hour or per sample")
	case SIL2:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-5 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-5 per flight hour or per sample")
	case SIL3:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-7 per flight hour or per sample",
			silLabelVertical, "<= 2 * 10^-7 per 150 seconds or per sample")
	default:
		return fmt.Sprintf("%v - Unknown code", level)
	}
}

// ReadSurveillanceIntegrityLevel reads the SurveillanceIntegrityLevel from a 56 bits data field
func ReadSurveillanceIntegrityLevel(data []byte) SurveillanceIntegrityLevel {
	bits := (data[5] & 0x0C) >> 2
	return SurveillanceIntegrityLevel(bits)
}
