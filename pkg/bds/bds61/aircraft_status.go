package bds61

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

type AircraftStatus interface {
	bds.Message
	GetSubtype() fields.Subtype
}

// ReadAircraftStatus reads a message at the format AircraftStatus
func ReadAircraftStatus(data []byte) (AircraftStatus, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadAircraftStatus", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	switch subType {
	case fields.SubtypeNoInformation:
		return ReadAircraftStatusNoInformation(data)
	case fields.SubtypeEmergencyPriorityStatus:
		return ReadAircraftStatusEmergency(data)
	case fields.SubtypeRABroadcast:
		return ReadAircraftStatusACAS(data)
	default:
		return nil, fmt.Errorf("the field SubType must be comprised between 0 and 2 included, got %v", subType)
	}
}
