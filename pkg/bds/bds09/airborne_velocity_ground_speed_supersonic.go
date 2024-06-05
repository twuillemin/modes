package bds09

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirborneVelocityGroundSpeedSupersonic is a message at the format BDS 9,0
type AirborneVelocityGroundSpeedSupersonic struct {
	FormatTypeCode                byte
	Subtype                       fields.Subtype
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

func (message AirborneVelocityGroundSpeedSupersonic) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AirborneVelocityGroundSpeedSupersonic) GetRegister() register.Register {
	return register.BDS09
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirborneVelocityGroundSpeedSupersonic) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AirborneVelocityGroundSpeedSupersonic) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Subtype:                           %v\n"+
		"Intent Change:                     %v\n"+
		"IFR Capability:                    %v\n"+
		"Navigation Uncertainty Category:   %v\n"+
		"Direction EW:                      %v\n"+
		"Velocity EW Supersonic:            %v\n"+
		"Direction NS:                      %v\n"+
		"Velocity NS Supersonic:            %v\n"+
		"Vertical Rate Source:              %v\n"+
		"Vertical Rate Sign:                %v\n"+
		"Vertical Rate:                     %v\n"+
		"Difference GNSS Baro Sign:         %v\n"+
		"Difference GNSS Baro:              %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
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

// ReadAirborneVelocityGroundSpeedSupersonic reads a message AirborneVelocity (Ground speed supersonic)
func ReadAirborneVelocityGroundSpeedSupersonic(data []byte) (*AirborneVelocityGroundSpeedSupersonic, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 19, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedSupersonic {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAirborneVelocityGroundSpeedSupersonic", subType.ToString())
	}

	return &AirborneVelocityGroundSpeedSupersonic{
		FormatTypeCode:                formatTypeCode,
		Subtype:                       subType,
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
