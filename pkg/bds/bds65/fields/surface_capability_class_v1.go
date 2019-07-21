package fields

import "fmt"

// SurfaceCapabilityClassV1 is the Capability Class Surface definition
//
// Specified in Doc 9871 / B.2.3.10.3
type SurfaceCapabilityClassV1 struct {
	PositionOffsetApplied                    PositionOffsetApplied
	CockpitDisplayOfTrafficInformationStatus CockpitDisplayOfTrafficInformationStatus
	B2Low                                    B2Low
}

// ToString returns a basic, but readable, representation of the field
func (capability SurfaceCapabilityClassV1) ToString() string {

	return fmt.Sprintf("Position Offset Applied:                       %v\n"+
		"Cockpit Display Of Traffic Information Status: %v\n"+
		"B2Low:                                         %v",
		capability.PositionOffsetApplied.ToString(),
		capability.CockpitDisplayOfTrafficInformationStatus.ToString(),
		capability.B2Low.ToString())
}

// ReadSurfaceCapabilityClassV1 reads the SurfaceCapabilityClassV1 from a 56 bits data field
func ReadSurfaceCapabilityClassV1(data []byte) SurfaceCapabilityClassV1 {
	return SurfaceCapabilityClassV1{
		PositionOffsetApplied:                    ReadPositionOffsetApplied(data),
		CockpitDisplayOfTrafficInformationStatus: ReadCockpitDisplayOfTrafficInformationStatus(data),
		B2Low:                                    ReadB2Low(data),
	}
}
