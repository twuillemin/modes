package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Utility Message (UM)
//
// -----------------------------------------------------------------------------------------

// UtilityMessage (UM) field shall contain transponder communications status information
//
// Defined at 3.1.2.6.5.3
type UtilityMessage struct {
	// InterrogatorIdentifier subfield reports the identifier of the interrogator that is reserved for
	// multi site communications.
	InterrogatorIdentifier uint8
	// IdentifierDesignator subfield reports the type of reservation made by the interrogator identified in IIS
	IdentifierDesignator UtilityMessageIdentifierDesignator
}

// UtilityMessageIdentifierDesignator subfield reports the type of reservation made by the interrogator
// identified in IIS.
//
// Defined at 3.1.2.6.5.3.1
type UtilityMessageIdentifierDesignator int

const (
	// UtilityMessageIdentifierDesignatorNoInformation signifies no information
	UtilityMessageIdentifierDesignatorNoInformation UtilityMessageIdentifierDesignator = 0
	// UtilityMessageIdentifierDesignatorCommB signifies IIS contains Comm-B II code
	UtilityMessageIdentifierDesignatorCommB UtilityMessageIdentifierDesignator = 1
	// UtilityMessageIdentifierDesignatorCommC signifies IIS contains Comm-C II code
	UtilityMessageIdentifierDesignatorCommC UtilityMessageIdentifierDesignator = 2
	// UtilityMessageIdentifierDesignatorCommD signifies IIS contains Comm-D II code
	UtilityMessageIdentifierDesignatorCommD UtilityMessageIdentifierDesignator = 3
)

// ReadUtilityMessage reads the UM field from a message
func ReadUtilityMessage(message common.MessageData) UtilityMessage {

	ii := ((message.Payload[0] & 0x07) << 1) | ((message.Payload[1] & 0x80) >> 7)
	id := (message.Payload[1] & 0x60) >> 5

	return UtilityMessage{
		InterrogatorIdentifier: ii,
		IdentifierDesignator:   UtilityMessageIdentifierDesignator(id),
	}
}

// ToString returns a basic, but readable, representation of the field
func (utilityMessageIdentifierDesignator UtilityMessageIdentifierDesignator) ToString() string {
	switch utilityMessageIdentifierDesignator {
	case UtilityMessageIdentifierDesignatorNoInformation:
		return "0 - No Information"
	case UtilityMessageIdentifierDesignatorCommB:
		return "1 - Designator Comm-B"
	case UtilityMessageIdentifierDesignatorCommC:
		return "2 - Designator Comm-C"
	case UtilityMessageIdentifierDesignatorCommD:
		return "3 - Designator Comm-D"
	default:
		return fmt.Sprintf("%v - Unknown code", utilityMessageIdentifierDesignator)
	}
}

// ToString returns a basic, but readable, representation of the field
func (utilityMessage UtilityMessage) ToString() string {
	return fmt.Sprintf("InterrogatorIdentifier: %v, Reservation: %v", utilityMessage.InterrogatorIdentifier, utilityMessage.IdentifierDesignator.ToString())
}
