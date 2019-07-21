package fields

import "fmt"

// AirborneCapabilityClassV1 is the Airborne Capability Class definition
//
// Specified in Doc 9871 / B.2.3.10.3
type AirborneCapabilityClassV1 struct {
	NotACASStatus                            NotACASStatus
	CockpitDisplayOfTrafficInformationStatus CockpitDisplayOfTrafficInformationStatus
	AirReferencedVelocityReportCapability    AirReferencedVelocityReportCapability
	TargetStateReportCapability              TargetStateReportCapability
	TargetChangeReportCapability             TargetChangeReportCapability
}

// ToString returns a basic, but readable, representation of the field
func (capability AirborneCapabilityClassV1) ToString() string {

	return fmt.Sprintf("Not ACAS Status:                               %v\n"+
		"Cockpit Display Of Traffic Information Status: %v\n"+
		"Air Referenced Velocity Report Capability:     %v\n"+
		"Target State Report Capability:                %v\n"+
		"Target Change Report Capability:               %v",
		capability.NotACASStatus.ToString(),
		capability.CockpitDisplayOfTrafficInformationStatus.ToString(),
		capability.AirReferencedVelocityReportCapability.ToString(),
		capability.TargetStateReportCapability.ToString(),
		capability.TargetChangeReportCapability.ToString())
}

// ReadAirborneCapabilityClassV1 reads the AirborneCapabilityClassV1 from a 56 bits data field
func ReadAirborneCapabilityClassV1(data []byte) AirborneCapabilityClassV1 {
	return AirborneCapabilityClassV1{
		NotACASStatus:                            ReadNotACASStatus(data),
		CockpitDisplayOfTrafficInformationStatus: ReadCockpitDisplayOfTrafficInformationStatus(data),
		AirReferencedVelocityReportCapability:    ReadAirReferencedVelocityReportCapability(data),
		TargetStateReportCapability:              ReadTargetStateReportCapability(data),
		TargetChangeReportCapability:             ReadTargetChangeReportCapability(data),
	}
}
