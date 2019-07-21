package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS09 is the basic interface that ADSB messages at the format BDS 0,9 are expected to implement
type MessageBDS09 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetSubtype returns the Airborne Velocity Subtype
	GetSubtype() fields.Subtype
	// GetIntentChange returns the IntentChange
	GetIntentChange() fields.IntentChange
	// GetIFRCapability returns the IFRCapability
	GetIFRCapability() fields.IFRCapability
	// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
	GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory
	// GetVerticalRateSource returns the VerticalRateSource
	GetVerticalRateSource() fields.VerticalRateSource
	// GetVerticalRateSign returns the VerticalRateSign
	GetVerticalRateSign() fields.VerticalRateSign
	// GetVerticalRate returns the VerticalRate
	GetVerticalRate() fields.VerticalRate
	// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
	GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign
	// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
	GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro
}

var bds09Code = "BDS 0,9"
var bds09Name = "Extended squitter airborne velocity"

// ReadBDS09 reads a message at the format BDS 0,9. As this format does not have changes from ADSB V0 to
// ADSB V2, the returned ADSBLevel is always the given one.
//
// Params:
//    - adsbLevel: The ADSB level request (not used, but present for coherency)
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS09(adsbLevel common.ADSBLevel, data []byte) (MessageBDS09, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != 19 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,9 format", formatTypeCode)
	}

	// Read the subtype
	subType := fields.ReadAirborneVelocitySubtype(data)

	switch subType {
	case fields.SubtypeGroundSpeedNormal:
		message, err := readFormat19GroundNormal(data)
		return message, adsbLevel, err
	case fields.SubtypeGroundSpeedSupersonic:
		message, err := readFormat19GroundSupersonic(data)
		return message, adsbLevel, err
	case fields.SubtypeAirspeedNormal:
		message, err := readFormat19AirspeedNormal(data)
		return message, adsbLevel, err
	case fields.SubtypeAirspeedSupersonic:
		message, err := readFormat19AirspeedSupersonic(data)
		return message, adsbLevel, err

	default:
		return nil, adsbLevel, fmt.Errorf("the subtype %v of Airborne Velocity is not supported", formatTypeCode)
	}
}
