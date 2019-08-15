package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19GroundSupersonic is a message at the format BDS 9,0
type Format19GroundSupersonic struct {
	IntentChange                  fields.IntentChange
	IFRCapability                 fields.IFRCapability
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	DirectionEastWest             fields.DirectionEastWest
	VelocityEWSupersonic          fields.VelocityEWSupersonic
	DirectionNorthSouth           fields.DirectionNorthSouth
	VelocityNSSupersonic          fields.VelocityNSSupersonic
	VerticalRateSource            fields.VerticalRateSource
	VerticalRateSign              fields.VerticalRateSign
	VerticalRate                  fields.VerticalRate
	DifferenceGNSSBaroSign        fields.DifferenceGNSSBaroSign
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format19GroundSupersonic) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format19V0OrMore
}

// GetRegister returns the register of the message
func (message *Format19GroundSupersonic) GetRegister() bds.Register {
	return adsb.Format19V0OrMore.GetRegister()
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format19GroundSupersonic) GetFormatTypeCode() byte {
	return 19
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message *Format19GroundSupersonic) GetSubtype() fields.Subtype {
	return fields.SubtypeGroundSpeedSupersonic
}

// GetIntentChange returns the IntentChange
func (message *Format19GroundSupersonic) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message *Format19GroundSupersonic) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message *Format19GroundSupersonic) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message *Format19GroundSupersonic) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message *Format19GroundSupersonic) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message *Format19GroundSupersonic) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message *Format19GroundSupersonic) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message *Format19GroundSupersonic) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19GroundSupersonic) ToString() string {
	return fmt.Sprintf("Message:                         %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change:                   %v\n"+
		"IFR Capability:                  %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Direction EW:                    %v\n"+
		"Velocity EW Supersonic:          %v\n"+
		"Direction NS:                    %v\n"+
		"Velocity NS Supersonic:          %v\n"+
		"Vertical Rate Source:            %v\n"+
		"Vertical Rate Sign:              %v\n"+
		"Vertical Rate:                   %v\n"+
		"Difference GNSS Baro Sign:       %v\n"+
		"Difference GNSS Baro:            %v",
		adsb.Format19V0OrMore.ToString(),
		message.GetSubtype().ToString(),
		message.IntentChange.ToString(),
		message.IFRCapability.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.DirectionEastWest.ToString(),
		message.VelocityEWSupersonic.ToString(),
		message.DirectionNorthSouth.ToString(),
		message.VelocityNSSupersonic.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadFormat19GroundSupersonic reads a message at the format Format19 / Subtype 2 (Ground speed supersonic)
func ReadFormat19GroundSupersonic(data []byte) (*Format19GroundSupersonic, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format19V0OrMore.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat19GroundSupersonic", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedSupersonic {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat19GroundSupersonic", subType.ToString())
	}

	return &Format19GroundSupersonic{
		IntentChange:                  fields.ReadIntentChange(data),
		IFRCapability:                 fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		DirectionEastWest:             fields.ReadDirectionEastWest(data),
		VelocityEWSupersonic:          fields.ReadVelocityEWSupersonic(data),
		DirectionNorthSouth:           fields.ReadDirectionNorthSouth(data),
		VelocityNSSupersonic:          fields.ReadVelocityNSSupersonic(data),
		VerticalRateSource:            fields.ReadVerticalRateSource(data),
		VerticalRateSign:              fields.ReadVerticalRateSign(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		DifferenceGNSSBaroSign:        fields.ReadDifferenceGNSSBaroSign(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
