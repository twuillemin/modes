package fields

import "fmt"

// AirReferencedVelocityReportCapability is the Air Referenced Velocity Report Capability definition
//
// Specified in Doc 9871 / B.2.3.10.3
type AirReferencedVelocityReportCapability byte

const (
	// ARVNoCapability indicates No capability for sending messages to support Air-Referenced Velocity Reports
	ARVNoCapability AirReferencedVelocityReportCapability = 0
	// ARVCapable indicates Capability of sending messages to support Air-Referenced Velocity Reports
	ARVCapable AirReferencedVelocityReportCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (status AirReferencedVelocityReportCapability) ToString() string {

	switch status {
	case ARVNoCapability:
		return "0 - No capability for sending messages to support Air-Referenced Velocity Reports"
	case ARVCapable:
		return "1 - Capability of sending messages to support Air-Referenced Velocity Reports"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadAirReferencedVelocityReportCapability reads the AirReferencedVelocityReportCapability from a 56 bits data field
func ReadAirReferencedVelocityReportCapability(data []byte) AirReferencedVelocityReportCapability {
	bits := (data[1] & 0x02) >> 1
	return AirReferencedVelocityReportCapability(bits)
}
