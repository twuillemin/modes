package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// MessageBDS09 is the basic interface that ADSB messages at the format BDS 0,9 are expected to implement
type MessageBDS09 interface {
	adsb.Message

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

// ReadBDS09 reads a message at the format BDS 0,9. As this format does not have changes from ADSB V0 to
// ADSB V2, the returned ADSBLevel is always the given one.
//
// Params:
//    - adsbLevel: The ADSB level request (not used, but present for coherency)
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS09(adsbLevel adsb.ReaderLevel, data []byte) (MessageBDS09, adsb.ReaderLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != 19 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,9 format", formatTypeCode)
	}

	// Read the subtype
	subType := fields.ReadSubtype(data)

	switch subType {
	case fields.SubtypeGroundSpeedNormal:
		message, err := ReadFormat19GroundSpeedNormal(data)
		return message, adsbLevel, err
	case fields.SubtypeGroundSpeedSupersonic:
		message, err := ReadFormat19GroundSpeedSupersonic(data)
		return message, adsbLevel, err
	case fields.SubtypeAirspeedNormal:
		message, err := ReadFormat19AirSpeedNormal(data)
		return message, adsbLevel, err
	case fields.SubtypeAirspeedSupersonic:
		message, err := ReadFormat19AirSpeedSupersonic(data)
		return message, adsbLevel, err

	default:
		return nil, adsbLevel, fmt.Errorf("the subtype %v of SubtypeAirborne Velocity is not supported", formatTypeCode)
	}
}
