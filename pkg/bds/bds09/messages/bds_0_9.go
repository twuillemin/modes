package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS09 is the basic interface that ADSB messages at the format BDS 6,5 are expected to implement
type MessageBDS09 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetAirborneVelocitySubtype returns the Airborne Velocity Subtype
	GetAirborneVelocitySubtype() fields.AirborneVelocitySubtype
}

var bds09Code = "BDS 0,9"
var bds09Name = "Extended squitter airborne velocity"

// ReadBDS09 reads a message at the format BDS 6,5
func ReadBDS09(data []byte) (MessageBDS09, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,9 format", formatTypeCode)
	}

	// Read the version of ADSB and the subtype
	subType := fields.ReadAirborneVelocitySubtype(data)

	switch subType {
	case fields.AVSCGroundSpeedNormal:
		return ReadFormat19GroundNormal(data)
	case fields.AVSCGroundSpeedSupersonic:
		return ReadFormat19GroundSupersonic(data)
	case fields.AVSCAirspeedNormal:
		return ReadFormat19AirspeedNormal(data)
	case fields.AVSCAirspeedSupersonic:
		return ReadFormat19AirspeedSupersonic(data)

	default:
		return nil, fmt.Errorf("the subtype %v of Airborne Velocity is not supported", formatTypeCode)
	}
}
