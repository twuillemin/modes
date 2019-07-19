package fields

import "fmt"

// EmergencyState is the Emergency State definition
//
// Specified in Doc 9871 / Table B-2-97a
type EmergencyState byte

const (
	// ESNoEmergency indicates no emergency
	ESNoEmergency EmergencyState = 0
	// ESGeneralEmergency indicates general emergency
	ESGeneralEmergency EmergencyState = 1
	// ESLifeguardMedical indicates lifeguard/medical emergency
	ESLifeguardMedical EmergencyState = 2
	// ESMinimumFuel indicates minimum fuel
	ESMinimumFuel EmergencyState = 3
	// ESNoCommunication indicates no communications
	ESNoCommunication EmergencyState = 4
	// ESUnlawfulInterference indicates unlawful interference
	ESUnlawfulInterference EmergencyState = 5
	// ESDownedAircraft indicates downed aircraft
	ESDownedAircraft EmergencyState = 6
	// ESReserved7 is reserved
	ESReserved7 EmergencyState = 7
)

// ToString returns a basic, but readable, representation of the field
func (capability EmergencyState) ToString() string {

	switch capability {
	case ESNoEmergency:
		return "0 - no emergency"
	case ESGeneralEmergency:
		return "1 - general emergency"
	case ESLifeguardMedical:
		return "2 - lifeguard/medical emergency"
	case ESMinimumFuel:
		return "3 - minimum fuel"
	case ESNoCommunication:
		return "4 - no communications"
	case ESUnlawfulInterference:
		return "5 - unlawful interference"
	case ESDownedAircraft:
		return "6 - downed aircraft"
	case ESReserved7:
		return "7 - reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", capability)
	}
}

// ReadEmergencyState reads the EmergencyState from a 56 bits data field
func ReadEmergencyState(data []byte) EmergencyState {
	bits := data[6] & 0x0E
	return EmergencyState(bits)
}
