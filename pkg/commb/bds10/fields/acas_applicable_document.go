package fields

import "fmt"

// ACASApplicableDocument is the ACAS applicable MOPS documents definition
//
// Specified in Doc 9871 / D.2.4.1
type ACASApplicableDocument byte

// Note: As per definition (see Table A-2-16), the number is actually represented at reverse for the MSB and the LSB

const (
	// ApplicableDocument185 indicates RTCA DO-185.
	ApplicableDocument185 ACASApplicableDocument = 0
	// ApplicableDocument185A indicates RTCA DO-185A.
	ApplicableDocument185A ACASApplicableDocument = 2
	// ApplicableDocument185B indicates RTCA DO-185B.
	ApplicableDocument185B ACASApplicableDocument = 1
)

// ToString returns a basic, but readable, representation of the field
func (aad ACASApplicableDocument) ToString() string {

	switch aad {
	case ApplicableDocument185:
		return "RTCA DO-185"
	case ApplicableDocument185A:
		return "RTCA DO-185A"
	case ApplicableDocument185B:
		return "RTCA DO-185B"
	default:
		return fmt.Sprintf("%v - Unknown code", aad)
	}
}

// ReadACASApplicableDocument reads the ACASApplicableDocument from a 56 bits data field
func ReadACASApplicableDocument(data []byte) ACASApplicableDocument {
	bits := data[4] & 0x03
	return ACASApplicableDocument(bits)
}
