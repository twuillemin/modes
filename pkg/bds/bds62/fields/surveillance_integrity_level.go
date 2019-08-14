package fields

import "fmt"

// SurveillanceIntegrityLevel is the Surveillance Integrity Level definition
//
// Specified in Doc 9871 / B.2.3.10.9
type SurveillanceIntegrityLevel byte

const (
	// SUILLevel0 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: unknown
	//     - Probability of exceeding the Vertical Integrity CR without notification: unknown
	SUILLevel0 SurveillanceIntegrityLevel = 0
	// SUILLevel1 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-3 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-3 per flight hour or per sample
	SUILLevel1 SurveillanceIntegrityLevel = 1
	// SUILLevel2 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-5 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-5 per flight hour or per sample
	SUILLevel2 SurveillanceIntegrityLevel = 2
	// SUILLevel3 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-7 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 2 * 10^-7 per 150 seconds or per sample
	SUILLevel3 SurveillanceIntegrityLevel = 3
)

var silLabelHorizontal = "Probability of exceeding the Horizontal Containment Radius reported in the NIC Subfield without an indication: "
var silLabelVertical = "Probability of exceeding the Vertical Integrity Containment Region (VPL) without an indication:                "

// ToString returns a basic, but readable, representation of the field
func (level SurveillanceIntegrityLevel) ToString() string {

	switch level {
	case SUILLevel0:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "Unknown",
			silLabelVertical, "Unknown")
	case SUILLevel1:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-3 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-3 per flight hour or per sample")
	case SUILLevel2:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-5 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-5 per flight hour or per sample")
	case SUILLevel3:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-7 per flight hour or per sample",
			silLabelVertical, "<= 2 * 10^-7 per 150 seconds or per sample")
	default:
		return fmt.Sprintf("%v - Unknown code", level)
	}
}

// ReadSurveillanceIntegrityLevel reads the SourceIntegrityLevel from a 56 bits data field
func ReadSurveillanceIntegrityLevel(data []byte) SurveillanceIntegrityLevel {
	bits := (data[5] & 0x0C) >> 2
	return SurveillanceIntegrityLevel(bits)
}
