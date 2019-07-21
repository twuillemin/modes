package fields

import "fmt"

// SystemDesignAssurance is the System Design Assurance definition
//
// Specified in Doc 9871 / Table C-32
type SystemDesignAssurance byte

const (
	// SDANotApplicable indicates Not Applicable
	SDANotApplicable SystemDesignAssurance = 0
	// SDALevelD indicates Level D
	SDALevelD SystemDesignAssurance = 1
	// SDALevelC indicates Level C
	SDALevelC SystemDesignAssurance = 2
	// SDALevelB indicates Level B
	SDALevelB SystemDesignAssurance = 3
)

// ToString returns a basic, but readable, representation of the field
func (status SystemDesignAssurance) ToString() string {

	switch status {
	case SDANotApplicable:
		return "0 - Not Applicable"
	case SDALevelD:
		return "1 - Level D"
	case SDALevelC:
		return "2 - Level C"
	case SDALevelB:
		return "3 - Level B"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSystemDesignAssurance reads the SystemDesignAssurance from a 56 bits data field
func ReadSystemDesignAssurance(data []byte) SystemDesignAssurance {
	bits := data[3] & 0x03
	return SystemDesignAssurance(bits)
}
