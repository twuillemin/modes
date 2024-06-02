package fields

import "fmt"

// TransmissionRateSubfield is the Transmission rate subfield definition
//
// Specified in Doc 9871 / A.2.7
type TransmissionRateSubfield byte

const (
	// TransmissionRateNoCapability indicates No capability to determine surface squitter rate.
	TransmissionRateNoCapability TransmissionRateSubfield = 0
	// TransmissionRateHigh indicates High surface squitter rate selected.
	TransmissionRateHigh TransmissionRateSubfield = 1
	// TransmissionRateLow indicates Low surface squitter rate selected.
	TransmissionRateLow TransmissionRateSubfield = 2
)

// ToString returns a basic, but readable, representation of the field
func (trs TransmissionRateSubfield) ToString() string {

	switch trs {
	case TransmissionRateNoCapability:
		return "0 - Barometric altitude"
	case TransmissionRateHigh:
		return "1 - High surface squitter rate selected"
	case TransmissionRateLow:
		return "2 - Low surface squitter rate selected"
	default:
		return fmt.Sprintf("%v - Unknown code", trs)
	}
}

// ReadTransmissionRateSubfield reads the TransmissionRateSubfield from a 56 bits data field
func ReadTransmissionRateSubfield(data []byte) TransmissionRateSubfield {
	bits := (data[0] & 0xC0) >> 6
	return TransmissionRateSubfield(bits)
}
