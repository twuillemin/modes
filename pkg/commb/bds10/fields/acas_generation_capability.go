package fields

import "fmt"

// ACASGenerationCapability is the ACAS generation capability definition
//
// Specified in Doc 9871 / D.2.4.1
type ACASGenerationCapability byte

const (
	// ACASGenerationNotCapable indicates that the ACAS is generating TAs only.
	ACASGenerationNotCapable ACASGenerationCapability = 0
	// ACASGenerationCapable indicates that the ACAS is generating both TAs and RAs.
	ACASGenerationCapable ACASGenerationCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (agc ACASGenerationCapability) ToString() string {

	switch agc {
	case ACASGenerationNotCapable:
		return "0 - The ACAS is generating TAs only"
	case ACASGenerationCapable:
		return "1 - The ACAS is generating both TAs and RAs"
	default:
		return fmt.Sprintf("%v - Unknown code", agc)
	}
}

// ReadACASGenerationCapability reads the ACASGenerationCapability from a 56 bits data field
func ReadACASGenerationCapability(data []byte) ACASGenerationCapability {
	bits := (data[4] & 0x04) >> 2
	return ACASGenerationCapability(bits)
}
