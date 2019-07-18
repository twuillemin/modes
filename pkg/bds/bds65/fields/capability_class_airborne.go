package fields

import "fmt"

// CapabilityClassAirborne is the Capability Class Airborne definition
//
// Specified in Doc 9871 / B.2.3.10.3
type CapabilityClassAirborne struct {
	NotACASStatus                            NotACASStatus
	CockpitDisplayOfTrafficInformationStatus CockpitDisplayOfTrafficInformationStatus
	AirReferencedVelocityReportCapability    AirReferencedVelocityReportCapability
	TargetStateReportCapability              TargetStateReportCapability
	TargetChangeReportCapability             TargetChangeReportCapability
}

// ToString returns a basic, but readable, representation of the field
func (capability CapabilityClassAirborne) ToString() string {

	return fmt.Sprintf("NotACASStatus:                            %v\n"+
		"CockpitDisplayOfTrafficInformationStatus: %v\n"+
		"AirReferencedVelocityReportCapability:    %v\n"+
		"TargetStateReportCapability:              %v\n"+
		"TargetChangeReportCapability:             %v",
		capability.NotACASStatus.ToString(),
		capability.CockpitDisplayOfTrafficInformationStatus.ToString(),
		capability.AirReferencedVelocityReportCapability.ToString(),
		capability.TargetStateReportCapability.ToString(),
		capability.TargetChangeReportCapability.ToString())
}

// ReadCapabilityClassAirborne reads the CapabilityClassAirborne from a 56 bits data field
func ReadCapabilityClassAirborne(data []byte) CapabilityClassAirborne {
	return CapabilityClassAirborne{
		NotACASStatus:                            ReadNotACASStatus(data),
		CockpitDisplayOfTrafficInformationStatus: ReadCockpitDisplayOfTrafficInformationStatus(data),
		AirReferencedVelocityReportCapability:    ReadAirReferencedVelocityReportCapability(data),
		TargetStateReportCapability:              ReadTargetStateReportCapability(data),
		TargetChangeReportCapability:             ReadTargetChangeReportCapability(data),
	}
}
