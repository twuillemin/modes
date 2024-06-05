package bds09

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirborneVelocityGroundSpeedNormal is a message at the format BDS 9,0
type AirborneVelocityGroundSpeedNormal struct {
	FormatTypeCode                byte
	Subtype                       fields.Subtype
	IntentChange                  fields.IntentChange
	IFRCapability                 fields.IFRCapability
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	DirectionEastWest             fields.DirectionEastWest
	VelocityEWNormal              fields.VelocityEWNormal
	DirectionNorthSouth           fields.DirectionNorthSouth
	VelocityNSNormal              fields.VelocityNSNormal
	VerticalRateSource            fields.VerticalRateSource
	VerticalRateSign              fields.VerticalRateSign
	VerticalRate                  fields.VerticalRate
	DifferenceGNSSBaroSign        fields.DifferenceGNSSBaroSign
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

func (message AirborneVelocityGroundSpeedNormal) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AirborneVelocityGroundSpeedNormal) GetRegister() register.Register {
	return register.BDS09
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirborneVelocityGroundSpeedNormal) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AirborneVelocityGroundSpeedNormal) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Subtype:                           %v\n"+
		"Intent Change:                     %v\n"+
		"IFR Capability:                    %v\n"+
		"Navigation Uncertainty Category:   %v\n"+
		"Direction EW:                      %v\n"+
		"Velocity EW Normal:                %v\n"+
		"Direction NS:                      %v\n"+
		"Velocity NS Normal:                %v\n"+
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
		message.VelocityEWNormal.ToString(),
		message.DirectionNorthSouth.ToString(),
		message.VelocityNSNormal.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadAirborneVelocityGroundSpeedNormal reads a message AirborneVelocity (Ground speed normal)
func ReadAirborneVelocityGroundSpeedNormal(data []byte) (*AirborneVelocityGroundSpeedNormal, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 19, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedNormal {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat19GroundSpeedNormal", subType.ToString())
	}

	return &AirborneVelocityGroundSpeedNormal{
		FormatTypeCode:                formatTypeCode,
		Subtype:                       subType,
		IntentChange:                  fields.ReadIntentChange(data),
		IFRCapability:                 fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		DirectionEastWest:             fields.ReadDirectionEastWest(data),
		VelocityEWNormal:              fields.ReadVelocityEWNormal(data),
		DirectionNorthSouth:           fields.ReadDirectionNorthSouth(data),
		VelocityNSNormal:              fields.ReadVelocityNSNormal(data),
		VerticalRateSource:            fields.ReadVerticalRateSource(data),
		VerticalRateSign:              fields.ReadVerticalRateSign(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		DifferenceGNSSBaroSign:        fields.ReadDifferenceGNSSBaroSign(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
