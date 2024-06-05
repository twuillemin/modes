package fields

import "fmt"

// SelectedHeadingSign is the Selected Heading Status definition
//
// Specified in Doc 9871 / C.2.3.9.7
type SelectedHeadingSign byte

const (
	// SHSPositive indicates that the Selected Heading is positive
	SHSPositive = 0
	// SHSNegative indicates that the Selected Heading is negative
	SHSNegative SelectedHeadingSign = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SelectedHeadingSign) ToString() string {

	switch status {
	case SHSPositive:
		return "0 - selected heading data is positive in an angular system having a range between +180 and –180 degrees"
	case SHSNegative:
		return "1 - selected heading data is negative in an angular system having a range between +180 and –180 degrees"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSelectedHeadingSign reads the SelectedHeadingSign from a 56 bits data field
func ReadSelectedHeadingSign(data []byte) SelectedHeadingSign {
	bits := (data[3] & 0x02) >> 1
	return SelectedHeadingSign(bits)
}
