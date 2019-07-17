package fields

import "fmt"

// CockpitDisplayOfTrafficInformationStatus is the CDTI Status definition
//
// Specified in Doc 9871 / B.2.3.10.3
type CockpitDisplayOfTrafficInformationStatus byte

const (
	// CDTINotOperational indicates Traffic display not operational
	CDTINotOperational CockpitDisplayOfTrafficInformationStatus = 0
	// CDTIOperational indicates Traffic display operational
	CDTIOperational CockpitDisplayOfTrafficInformationStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (status CockpitDisplayOfTrafficInformationStatus) ToString() string {

	switch status {
	case CDTINotOperational:
		return "0 - Traffic display not operational"
	case CDTIOperational:
		return "1 - Traffic display operational"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadCockpitDisplayOfTrafficInformationStatus reads the CockpitDisplayOfTrafficInformationStatus from a 56 bits data field
func ReadCockpitDisplayOfTrafficInformationStatus(data []byte) CockpitDisplayOfTrafficInformationStatus {
	bits := (data[1] & 0x10) >> 4
	return CockpitDisplayOfTrafficInformationStatus(bits)
}
