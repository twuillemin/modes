package fields

import "fmt"

// AirborneVelocitySubtypeCode is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type AirborneVelocitySubtypeCode byte

const (
	// AVSCReserved0 indicates Airborne Status Message
	AVSCReserved0 AirborneVelocitySubtypeCode = 0
	// AVSCGroundSpeedNormal indicates Surface Status Message
	AVSCGroundSpeedNormal AirborneVelocitySubtypeCode = 1
	// AVSCGroundSpeedSupersonic is reserved
	AVSCGroundSpeedSupersonic AirborneVelocitySubtypeCode = 2
	// AVSCAirspeedNormal is reserved
	AVSCAirspeedNormal AirborneVelocitySubtypeCode = 3
	// AVSCAirspeedSupersonic is reserved
	AVSCAirspeedSupersonic AirborneVelocitySubtypeCode = 4
	// AVSCReserved5 is reserved
	AVSCReserved5 AirborneVelocitySubtypeCode = 5
	// AVSCReserved6 is reserved
	AVSCReserved6 AirborneVelocitySubtypeCode = 6
	// AVSCReserved7 is reserved
	AVSCReserved7 AirborneVelocitySubtypeCode = 7
)

// ToString returns a basic, but readable, representation of the field
func (code AirborneVelocitySubtypeCode) ToString() string {

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

// ReadAirborneVelocitySubtypeCode reads the AirborneVelocitySubtypeCode from a 56 bits data field
func ReadAirborneVelocitySubtypeCode(data []byte) AirborneVelocitySubtypeCode {
	bits := data[0] & 0x07
	return AirborneVelocitySubtypeCode(bits)
}
