package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS61 is the basic interface that ADSB messages at the format BDS 6,1 are expected to implement
type MessageBDS61 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetSubtype returns the Subtype
	GetSubtype() fields.Subtype
	// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
	GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus
}

var bds61Code = "BDS 6,1"
var bds61Name = "Extended squitter emergency/priority status"

// ReadBDS61 reads a message at the format BDS 6,1
func ReadBDS61(data []byte) (MessageBDS61, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the format type code %v can not be read as a BDS 6,1 format", formatTypeCode)
	}

	// Read the version the subtype
	subType := fields.ReadSubtypeV0(data)

	switch subType {
	case fields.SubtypeV0EmergencyPriorityStatus:
		return ReadFormat28V0(data)

	default:
		return nil, fmt.Errorf("the subtype %v of Emergency/Priority Status is not supported", formatTypeCode)
	}
}
