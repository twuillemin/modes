package fields

import "fmt"

// Source is the source
//
// Specified in Doc 9871 / Table A-2-68
type Source byte

const (
	// SourceInvalid indicates turbulence data not available in Register 45.
	SourceInvalid Source = 0
	// SourceINS indicates INS source.
	SourceINS Source = 1
	// SourceGNSS indicates GNSS source.
	SourceGNSS Source = 2
	// SourceDMEDME indicates DME/DME source.
	SourceDMEDME Source = 3
	// SourceVORDME indicates VOR/DME source.
	SourceVORDME Source = 4
)

// ToString returns a basic, but readable, representation of the field
func (source Source) ToString() string {

	switch source {
	case SourceInvalid:
		return "0 - Invalid"
	case SourceINS:
		return "1 - INS"
	case SourceGNSS:
		return "2 - GNSS"
	case SourceDMEDME:
		return "3 - DME/DME"
	case SourceVORDME:
		return "4 - VOR/DMS"
	default:
		return fmt.Sprintf("%v - Unknown code", source)
	}
}

func ReadSource(data []byte) Source {
	return Source((data[0] & 0xF0) >> 4)
}
