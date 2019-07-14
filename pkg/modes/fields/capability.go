package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Capability (CA)
//
// -----------------------------------------------------------------------------------------

// Capability (CA) field shall convey information on the transponder level, the additional information below,
// and shall be used in formats DF = 11 and DF = 17.
//
// Defined at 3.1.2.5.2.2.1
type Capability int

const (
	// CapabilityLevel1Transponder signifies Level 1 transponder (surveillance only), and no ability to set CA
	// code 7 and either airborne or on the ground
	CapabilityLevel1Transponder Capability = 0
	// CapabilityReserved1 is reserved
	CapabilityReserved1 Capability = 1
	// CapabilityReserved2 is reserved
	CapabilityReserved2 Capability = 2
	// CapabilityReserved3 is reserved
	CapabilityReserved3 Capability = 3
	// CapabilityLevel2OnTheGround signifies Level 2 or above transponder and ability to set CA code 7 and on the ground
	CapabilityLevel2OnTheGround Capability = 4
	// CapabilityLevel2Airborne signifies Level 2 or above transponder and ability to set CA code 7 and airborne
	CapabilityLevel2Airborne Capability = 5
	// CapabilityLevel2OnTheGroundOrAirborne signifies Level 2 or above transponder and ability to set CA code 7 and
	// either airborne or on the ground
	CapabilityLevel2OnTheGroundOrAirborne Capability = 6
	// CapabilityFSOrDR signifies the DR field is not equal to 0 or the FS field equals 2, 3, 4 or 5, and either
	// airborne or on the ground
	CapabilityFSOrDR Capability = 7
)

// ReadCapability reads the CA field from a message
func ReadCapability(message common.MessageData) Capability {

	return Capability(message.FirstField)
}

// ToString returns a basic, but readable, representation of the field
func (capability Capability) ToString() string {
	switch capability {
	case CapabilityLevel1Transponder:
		return "0 - Level 1 Transponder (surveillance only)"
	case CapabilityReserved1:
		return "1 - Reserved"
	case CapabilityReserved2:
		return "2 - Reserved"
	case CapabilityReserved3:
		return "3 - Reserved"
	case CapabilityLevel2OnTheGround:
		return "4 - Level 2 On The Ground"
	case CapabilityLevel2Airborne:
		return "5 - Level 2 Airborne"
	case CapabilityLevel2OnTheGroundOrAirborne:
		return "6 - Level 2 On The Ground or Airborne"
	case CapabilityFSOrDR:
		return "7 - Flight Status (FS) with alert rr downlink request (DR)"
	default:
		return fmt.Sprintf("%v - Unknown code", capability)
	}
}
