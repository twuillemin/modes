package fields

import "fmt"

// AirborneCapabilityClassV2 is the Airborne Capability Class definition
//
// Specified in Doc 9871 / C.2.3.10.3
type AirborneCapabilityClassV2 struct {
	ACASOperational                       ACASOperational
	ExtendedSquitterIn                    ExtendedSquitterIn
	AirReferencedVelocityReportCapability AirReferencedVelocityReportCapability
	TargetStateReportCapability           TargetStateReportCapability
	TargetChangeReportCapability          TargetChangeReportCapability
	UniversalAccessTransceiverCapability  UniversalAccessTransceiverCapability
}

// ToString returns a basic, but readable, representation of the field
func (capability AirborneCapabilityClassV2) ToString() string {

	return fmt.Sprintf("ACAS Operational:                                %v\n"+
		"1090ES IN (1090 MHz Extended Squitter):          %v\n"+
		"ARV (Air-Referenced Velocity Report Capability): %v\n"+
		"Target State Report Capability:                  %v\n"+
		"Target Change Report Capability:                 %v\n"+
		"UAT IN (Universal Access Transceiver):           %v",
		capability.ACASOperational.ToString(),
		capability.ExtendedSquitterIn.ToString(),
		capability.AirReferencedVelocityReportCapability.ToString(),
		capability.TargetStateReportCapability.ToString(),
		capability.TargetChangeReportCapability.ToString(),
		capability.UniversalAccessTransceiverCapability.ToString())
}

// ReadAirborneCapabilityClassV2 reads the AirborneCapabilityClassV2 from a 56 bits data field
func ReadAirborneCapabilityClassV2(data []byte) AirborneCapabilityClassV2 {
	return AirborneCapabilityClassV2{
		ACASOperational:                       ReadACASOperational(data),
		ExtendedSquitterIn:                    ReadExtendedSquitterIn(data),
		AirReferencedVelocityReportCapability: ReadAirReferencedVelocityReportCapability(data),
		TargetStateReportCapability:           ReadTargetStateReportCapability(data),
		TargetChangeReportCapability:          ReadTargetChangeReportCapability(data),
		UniversalAccessTransceiverCapability:  ReadUniversalAccessTransceiverCapabilityAirborne(data),
	}
}
