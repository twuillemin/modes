package fields

import (
	"fmt"
)

// SurfaceCapabilityClassV2 is the Capability Class Surface definition
//
// Specified in Doc 9871 / C.2.3.10.3
type SurfaceCapabilityClassV2 struct {
	ExtendedSquitterIn                   ExtendedSquitterIn
	B2Low                                B2Low
	UniversalAccessTransceiverCapability UniversalAccessTransceiverCapability
	NavigationAccuracyCategory           NavigationAccuracyCategory
	NICSupplementC                       NICSupplementC
}

// ToString returns a basic, but readable, representation of the field
func (capability SurfaceCapabilityClassV2) ToString() string {

	return fmt.Sprintf("1090ES IN (1090 MHz Extended Squitter):                   %v\n"+
		"B2 Low (Class B2 Transmit Power Less Than 70 Watts):      %v\n"+
		"UAT IN (Universal Access Transceiver):                    %v\n"+
		"NACV (Navigation Accuracy Category for Velocity):         %v\n"+
		"NIC Supplement-C (NIC Supplement for use on the Surface): %v",
		capability.ExtendedSquitterIn.ToString(),
		capability.B2Low.ToString(),
		capability.UniversalAccessTransceiverCapability.ToString(),
		capability.NavigationAccuracyCategory.ToString(),
		capability.NICSupplementC.ToString())
}

// ReadSurfaceCapabilityClassV2 reads the SurfaceCapabilityClassV2 from a 56 bits data field
func ReadSurfaceCapabilityClassV2(data []byte) SurfaceCapabilityClassV2 {
	return SurfaceCapabilityClassV2{
		ExtendedSquitterIn:                   ReadExtendedSquitterIn(data),
		B2Low:                                ReadB2Low(data),
		UniversalAccessTransceiverCapability: ReadUniversalAccessTransceiverCapabilitySurface(data),
		NavigationAccuracyCategory:           ReadNavigationAccuracyCategory(data),
		NICSupplementC:                       ReadNICSupplementC(data),
	}
}
