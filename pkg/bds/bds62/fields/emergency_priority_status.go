package fields

import "fmt"

// EmergencyPriorityStatus is the Emergency Priority Status definition
//
// Specified in Doc 9871 / B.2.3.9.16
type EmergencyPriorityStatus byte

const (
	// EPSNoEmergency indicates no emergency
	EPSNoEmergency EmergencyPriorityStatus = 0
	// EPSGeneralEmergency indicates general emergency
	EPSGeneralEmergency EmergencyPriorityStatus = 1
	// EPSLifeguardMedical indicates lifeguard/medical emergency
	EPSLifeguardMedical EmergencyPriorityStatus = 2
	// EPSMinimumFuel indicates minimum fuel
	EPSMinimumFuel EmergencyPriorityStatus = 3
	// EPSNoCommunication indicates no communications
	EPSNoCommunication EmergencyPriorityStatus = 4
	// EPSUnlawfulInterference indicates unlawful interference
	EPSUnlawfulInterference EmergencyPriorityStatus = 5
	// EPSDownedAircraft indicates downed aircraft
	EPSDownedAircraft EmergencyPriorityStatus = 6
	// EPSReserved7 is reserved
	EPSReserved7 EmergencyPriorityStatus = 7
)

// ToString returns a basic, but readable, representation of the field
func (capability EmergencyPriorityStatus) ToString() string {

	switch capability {
	case EPSNoEmergency:
		return "0 - no emergency"
	case EPSGeneralEmergency:
		return "1 - general emergency"
	case EPSLifeguardMedical:
		return "2 - lifeguard/medical emergency"
	case EPSMinimumFuel:
		return "3 - minimum fuel"
	case EPSNoCommunication:
		return "4 - no communications"
	case EPSUnlawfulInterference:
		return "5 - unlawful interference"
	case EPSDownedAircraft:
		return "6 - downed aircraft"
	case EPSReserved7:
		return "7 - reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", capability)
	}
}

// ReadEmergencyPriorityStatus reads the EmergencyPriorityStatus from a 56 bits data field
func ReadEmergencyPriorityStatus(data []byte) EmergencyPriorityStatus {
	bits := data[6] & 0x07
	return EmergencyPriorityStatus(bits)
}
