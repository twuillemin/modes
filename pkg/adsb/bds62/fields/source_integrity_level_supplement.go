package fields

import "fmt"

// SourceIntegrityLevelSupplement is the Source Integrity Level Supplement definition
//
// Specified in Doc 9871 / C.2.3.10.15
type SourceIntegrityLevelSupplement byte

const (
	// SILSByHour indicates Probability of exceeding NIC radius of containment is based on "per hour"
	SILSByHour SourceIntegrityLevelSupplement = 0
	// SILSBySample indicates Probability of exceeding NIC radius of containment is based on "per sample"
	SILSBySample SourceIntegrityLevelSupplement = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SourceIntegrityLevelSupplement) ToString() string {

	switch status {
	case SILSByHour:
		return "0 - Probability of exceeding NIC radius of containment is based on \"per hour\""
	case SILSBySample:
		return "1 - Probability of exceeding NIC radius of containment is based on \"per sample\""
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSourceIntegrityLevelSupplement reads the SourceIntegrityLevelSupplement from a 56 bits data field
func ReadSourceIntegrityLevelSupplement(data []byte) SourceIntegrityLevelSupplement {
	bits := data[0] & 0x01
	return SourceIntegrityLevelSupplement(bits)
}
