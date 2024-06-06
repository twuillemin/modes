package bds09

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

type AirborneVelocity interface {
	bds.Message
	GetSubtype() fields.Subtype
}

// ReadAirborneVelocity reads a message at the format AirborneVelocity
func ReadAirborneVelocity(data []byte) (AirborneVelocity, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadAirborneVelocityAirSpeedNormal", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	switch subType {
	case fields.SubtypeGroundSpeedNormal:
		return ReadAirborneVelocityGroundSpeedNormal(data)
	case fields.SubtypeGroundSpeedSupersonic:
		return ReadAirborneVelocityGroundSpeedSupersonic(data)
	case fields.SubtypeAirspeedNormal:
		return ReadAirborneVelocityAirSpeedNormal(data)
	case fields.SubtypeAirspeedSupersonic:
		return ReadAirborneVelocityAirSpeedSupersonic(data)
	default:
		return nil, fmt.Errorf("the field SubType must be comprised between 1 and 4 included, got %v", subType)
	}
}
