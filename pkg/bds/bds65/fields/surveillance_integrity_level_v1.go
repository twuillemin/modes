package fields

import "fmt"

// SurveillanceIntegrityLevelV1 is the Surveillance Integrity Level definition
//
// Specified in Doc 9871 / B.2.3.10.9
type SurveillanceIntegrityLevelV1 byte

const (
	// SILV1Level0 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: unknown
	//     - Probability of exceeding the Vertical Integrity CR without notification: unknown
	SILV1Level0 SurveillanceIntegrityLevelV1 = 0
	// SILV1Level1 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-3 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-3 per flight hour or per sample
	SILV1Level1 SurveillanceIntegrityLevelV1 = 1
	// SILV1Level2 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-5 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 1 * 10^-5 per flight hour or per sample
	SILV1Level2 SurveillanceIntegrityLevelV1 = 2
	// SILV1Level3 indicates:
	//     - Probability of exceeding the Horizontal CR without notification: <=1 * 10^-7 per flight hour or per sample
	//     - Probability of exceeding the Vertical Integrity CR without notification: <= 2 * 10^-7 per 150 seconds or per sample
	SILV1Level3 SurveillanceIntegrityLevelV1 = 3
)

var silLabelHorizontal = "Probability of exceeding the Horizontal Containment Radius reported in the NIC Subfield without an indication: "
var silLabelVertical = "Probability of exceeding the Vertical Integrity Containment Region (VPL) without an indication:                "

// ToString returns a basic, but readable, representation of the field
func (level SurveillanceIntegrityLevelV1) ToString() string {

	switch level {
	case SILV1Level0:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "Unknown",
			silLabelVertical, "Unknown")
	case SILV1Level1:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-3 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-3 per flight hour or per sample")
	case SILV1Level2:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-5 per flight hour or per sample",
			silLabelVertical, "<= 1 * 10^-5 per flight hour or per sample")
	case SILV1Level3:
		return fmt.Sprintf("%v%v\n%v%v",
			silLabelHorizontal, "<= 1 * 10^-7 per flight hour or per sample",
			silLabelVertical, "<= 2 * 10^-7 per 150 seconds or per sample")
	default:
		return fmt.Sprintf("%v - Unknown code", level)
	}
}

// ReadSurveillanceIntegrityLevelV1 reads the SurveillanceIntegrityLevelV1 from a 56 bits data field
func ReadSurveillanceIntegrityLevelV1(data []byte) SurveillanceIntegrityLevelV1 {
	bits := (data[6] & 0x30) >> 4
	return SurveillanceIntegrityLevelV1(bits)
}
