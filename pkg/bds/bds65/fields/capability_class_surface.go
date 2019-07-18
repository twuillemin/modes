package fields

import "fmt"

// CapabilityClassSurface is the Capability Class Surface definition
//
// Specified in Doc 9871 / B.2.3.10.3
type CapabilityClassSurface struct {
	PositionOffsetApplied                    PositionOffsetApplied
	CockpitDisplayOfTrafficInformationStatus CockpitDisplayOfTrafficInformationStatus
	B2Low                                    B2Low
}

// ToString returns a basic, but readable, representation of the field
func (capability CapabilityClassSurface) ToString() string {

	return fmt.Sprintf("PositionOffsetApplied:                    %v\n"+
		"CockpitDisplayOfTrafficInformationStatus: %v\n"+
		"B2Low:                                    %v",
		capability.PositionOffsetApplied.ToString(),
		capability.CockpitDisplayOfTrafficInformationStatus.ToString(),
		capability.B2Low.ToString())
}

// ReadCapabilityClassSurface reads the CapabilityClassSurface from a 56 bits data field
func ReadCapabilityClassSurface(data []byte) CapabilityClassSurface {
	return CapabilityClassSurface{
		PositionOffsetApplied:                    ReadPositionOffsetApplied(data),
		CockpitDisplayOfTrafficInformationStatus: ReadCockpitDisplayOfTrafficInformationStatus(data),
		B2Low:                                    ReadB2Low(data),
	}
}
