package fields

// RAType defines the type of the Resolution Advisory
type RAType int

const (
	// OneThreatOrSameSeparation indicates that the Resolution Advisory is for single threat or multiple threats with
	// the same separation
	OneThreatOrSameSeparation RAType = 0
	// MultipleThreatDifferentSeparation indicates that the Resolution Advisory is for multiple threats with
	// different separation
	MultipleThreatDifferentSeparation RAType = 1
	// NoVerticalRAGenerated indicates that no Resolution Advisory was generated
	NoVerticalRAGenerated RAType = 2
)

// ActiveResolutionAdvisory is the base type that all Resolution Advisory should implement
type ActiveResolutionAdvisory interface {
	// GetType returns the type of Resolution Advisory
	GetType() RAType
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}

// ReadActiveResolutionAdvisory reads a ActiveResolutionAdvisory field
func ReadActiveResolutionAdvisory(data []byte) ActiveResolutionAdvisory {

	// The type of Active RA depends on the first bit of the RA and of the presence
	// of multiple threats
	mte := (data[2] & 0x10) >> 4
	araFirstBit := (data[0] & 0x80) >> 7

	if araFirstBit == 1 {
		return ReadARAOneThreatOrSameSeparation(data)
	} else if mte != 0 {
		return ReadARAMultipleThreatsDifferentSeparation(data)
	} else {
		return ActiveRANoVerticalRAGenerated{}
	}
}
