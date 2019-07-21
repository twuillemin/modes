package messages

import (
	"fmt"
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

// GetName returns the name of the message
func (message *Format19GroundSupersonic) GetName() string {
	return bds09Name
}

// GetBDS returns the binary data format
func (message *Format19GroundSupersonic) GetBDS() string {
	return bds09Code
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
	return fmt.Sprintf("Message:                         %v - %v (%v)\n"+
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
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
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

// readFormat19GroundSupersonic reads a message at the format BDS 6,5
func readFormat19GroundSupersonic(data []byte) (*Format19GroundSupersonic, error) {

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
