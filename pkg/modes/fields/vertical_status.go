package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 VerticalStatus (VS)
//
// -----------------------------------------------------------------------------------------

// VerticalStatus (VS) downlink field shall indicate the status of the aircraft
//
// Defined at 3.1.2.8.2.1
type VerticalStatus int

const (
	// VerticalStatusAirborne signifies that the aircraft is airborne
	VerticalStatusAirborne VerticalStatus = 0
	// VerticalStatusOnTheGround signifies that the aircraft is on the ground
	VerticalStatusOnTheGround VerticalStatus = 1
)

// readVerticalStatus reads the VS field from a message
func ReadVerticalStatus(message common.MessageData) VerticalStatus {
	if message.FirstField&0x04 != 0 {
		return VerticalStatusOnTheGround
	}
	return VerticalStatusAirborne
}

func (verticalStatus VerticalStatus) PrettyPrint() string {
	switch verticalStatus {
	case VerticalStatusAirborne:
		return "0 - Airborne"
	case VerticalStatusOnTheGround:
		return "1 - On The Ground"
	default:
		return fmt.Sprintf("%v - Unknown code", verticalStatus)
	}
}

func (verticalStatus VerticalStatus) ExtendedPrettyPrint() string {
	return verticalStatus.PrettyPrint()
}
