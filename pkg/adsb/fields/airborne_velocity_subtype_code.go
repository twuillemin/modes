package fields

import "fmt"

// AirborneVelocitySubtype is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type AirborneVelocitySubtype byte

const (
	// AVSCReserved0 indicates Airborne Status Message
	AVSCReserved0 AirborneVelocitySubtype = 0
	// AVSCGroundSpeedNormal indicates Surface Status Message
	AVSCGroundSpeedNormal AirborneVelocitySubtype = 1
	// AVSCGroundSpeedSupersonic is reserved
	AVSCGroundSpeedSupersonic AirborneVelocitySubtype = 2
	// AVSCAirspeedNormal is reserved
	AVSCAirspeedNormal AirborneVelocitySubtype = 3
	// AVSCAirspeedSupersonic is reserved
	AVSCAirspeedSupersonic AirborneVelocitySubtype = 4
	// AVSCReserved5 is reserved
	AVSCReserved5 AirborneVelocitySubtype = 5
	// AVSCReserved6 is reserved
	AVSCReserved6 AirborneVelocitySubtype = 6
	// AVSCReserved7 is reserved
	AVSCReserved7 AirborneVelocitySubtype = 7
)

// ToString returns a basic, but readable, representation of the field
func (code AirborneVelocitySubtype) ToString() string {

	switch code {
	case AVSCGroundSpeedNormal:
		return "1 - GroundSpeed / Normal"
	case AVSCGroundSpeedSupersonic:
		return "2 - GroundSpeed / Supersonic"
	case AVSCAirspeedNormal:
		return "1 - Airspeed,Heading / Normal"
	case AVSCAirspeedSupersonic:
		return "2 - Airspeed,Heading / Supersonic"
	case AVSCReserved0, AVSCReserved5, AVSCReserved6, AVSCReserved7:
		return fmt.Sprintf("%v - Reserved", code)
	default:
		return fmt.Sprintf("%v - Unknown code", code)
	}
}

// ReadAirborneVelocitySubtype reads the AirborneVelocitySubtype from a 56 bits data field
func ReadAirborneVelocitySubtype(data []byte) AirborneVelocitySubtype {
	bits := data[0] & 0x07
	return AirborneVelocitySubtype(bits)
}
