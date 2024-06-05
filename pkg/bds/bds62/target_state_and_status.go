package bds62

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
)

type TargetStateAndStatus interface {
	bds.Message
	GetSubtype() fields.Subtype
}

// ReadTargetStateAndStatus reads a message at the format TargetStateAndStatus
func ReadTargetStateAndStatus(data []byte) (TargetStateAndStatus, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 29 {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadTargetStateAndStatus", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	switch subType {
	case fields.Subtype0:
		return ReadTargetStateAndStatus0(data)
	case fields.Subtype1:
		return ReadTargetStateAndStatus1(data)
	default:
		return nil, fmt.Errorf("the field SubType must be comprised between 0 and 1 included, got %v", subType)
	}
}
