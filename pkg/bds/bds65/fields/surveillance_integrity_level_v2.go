package fields

import "fmt"

// SurveillanceIntegrityLevelV2 is the Surveillance Integrity Level definition
//
// Specified in Doc 9871 / B.2.3.10.9
type SurveillanceIntegrityLevelV2 byte

const (
	// SILV2Level0 indicates Unknown or SIL > 1 * 10^-3 per flight hour or per sample
	SILV2Level0 SurveillanceIntegrityLevelV2 = 0
	// SILV2Level1 indicates SIL <= 1 * 10^-3 per flight hour or per sample
	SILV2Level1 SurveillanceIntegrityLevelV2 = 1
	// SILV2Level2 indicates SIL <= 1 * 10^-5 per flight hour or per sample
	SILV2Level2 SurveillanceIntegrityLevelV2 = 2
	// SILV2Level3 indicates SIL <= 1 * 10^-7 per flight hour or per sample
	SILV2Level3 SurveillanceIntegrityLevelV2 = 3
)

// ToString returns a basic, but readable, representation of the field
func (level SurveillanceIntegrityLevelV2) ToString() string {

	switch level {
	case SILV2Level0:
		return "0 - Unknown or SIL > 1 * 10^-3 per flight hour or per sample"
	case SILV2Level1:
		return "1 - SIL <= 1 * 10^-3 per flight hour or per sample"
	case SILV2Level2:
		return "1 - SIL <= 1 * 10^-5 per flight hour or per sample"
	case SILV2Level3:
		return "1 - SIL <= 1 * 10^-7 per flight hour or per sample"
	default:
		return fmt.Sprintf("%v - Unknown code", level)
	}
}

// ReadSurveillanceIntegrityLevelV2 reads the SurveillanceIntegrityLevelV2 from a 56 bits data field
func ReadSurveillanceIntegrityLevelV2(data []byte) SurveillanceIntegrityLevelV2 {
	bits := (data[6] & 0x30) >> 4
	return SurveillanceIntegrityLevelV2(bits)
}
