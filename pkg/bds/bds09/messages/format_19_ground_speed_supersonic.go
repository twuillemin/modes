package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19GroundSpeedSupersonic is a message at the format BDS 9,0
type Format19GroundSpeedSupersonic struct {
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
func (message Format19GroundSpeedSupersonic) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format19
}

// GetRegister returns the register of the message
func (message Format19GroundSpeedSupersonic) GetRegister() bds.Register {
	return adsb.Format19.GetRegister()
}

// GetFormatTypeCode returns the Format Type Code
func (message Format19GroundSpeedSupersonic) GetFormatTypeCode() byte {
	return 19
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message Format19GroundSpeedSupersonic) GetSubtype() adsb.Subtype {
	return fields.SubtypeGroundSpeedSupersonic
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format19GroundSpeedSupersonic) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format19GroundSpeedSupersonic) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetIntentChange returns the IntentChange
func (message Format19GroundSpeedSupersonic) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message Format19GroundSpeedSupersonic) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message Format19GroundSpeedSupersonic) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message Format19GroundSpeedSupersonic) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message Format19GroundSpeedSupersonic) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message Format19GroundSpeedSupersonic) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message Format19GroundSpeedSupersonic) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message Format19GroundSpeedSupersonic) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19GroundSpeedSupersonic) ToString() string {
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
		adsb.GetMessageFormatInformation(&message),
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

// ReadFormat19GroundSpeedSupersonic reads a message at the format Format19 / Subtype 2 (Ground speed supersonic)
func ReadFormat19GroundSpeedSupersonic(data []byte) (*Format19GroundSpeedSupersonic, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format19.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat19GroundSpeedSupersonic", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedSupersonic {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat19GroundSpeedSupersonic", subType.ToString())
	}

	return &Format19GroundSpeedSupersonic{
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
