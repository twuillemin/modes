package fields

import "fmt"

// SourceIntegrityLevel is the Source Integrity Level definition
//
// Specified in Doc 9871 / C.5.3.2.8
type SourceIntegrityLevel byte

const (
	// SILLevel0 indicates Unknown or SIL > 1 * 10^-3 per flight hour or per sample
	SILLevel0 SourceIntegrityLevel = 0
	// SILLevel1 indicates SIL <= 1 * 10^-3 per flight hour or per sample
	SILLevel1 SourceIntegrityLevel = 1
	// SILLevel2 indicates SIL <= 1 * 10^-5 per flight hour or per sample
	SILLevel2 SourceIntegrityLevel = 2
	// SILLevel3 indicates SIL <= 1 * 10^-7 per flight hour or per sample
	SILLevel3 SourceIntegrityLevel = 3
)

// ToString returns a basic, but readable, representation of the field
func (level SourceIntegrityLevel) ToString() string {

	switch level {
	case SILLevel0:
		return "0 - Unknown or SIL > 1 * 10^-3 per flight hour or per sample"
	case SILLevel1:
		return "1 - SIL <= 1 * 10^-3 per flight hour or per sample"
	case SILLevel2:
		return "1 - SIL <= 1 * 10^-5 per flight hour or per sample"
	case SILLevel3:
		return "1 - SIL <= 1 * 10^-7 per flight hour or per sample"
	default:
		return fmt.Sprintf("%v - Unknown code", level)
	}
}

// ReadSourceIntegrityLevel reads the SourceIntegrityLevel from a 56 bits data field
func ReadSourceIntegrityLevel(data []byte) SourceIntegrityLevel {
	bits := (data[6] & 0x30) >> 4
	return SourceIntegrityLevel(bits)
}
