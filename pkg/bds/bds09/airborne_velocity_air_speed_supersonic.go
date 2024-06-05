package bds09

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirborneVelocityAirSpeedSupersonic is a message at the format BDS 9,0
type AirborneVelocityAirSpeedSupersonic struct {
	FormatTypeCode                byte
	Subtype                       fields.Subtype
	IntentChange                  fields.IntentChange
	IFRCapability                 fields.IFRCapability
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	MagneticHeadingStatus         fields.MagneticHeadingStatus
	MagneticHeading               fields.MagneticHeading
	AirspeedType                  fields.AirspeedType
	AirspeedSupersonic            fields.AirspeedSupersonic
	VerticalRateSource            fields.VerticalRateSource
	VerticalRateSign              fields.VerticalRateSign
	VerticalRate                  fields.VerticalRate
	DifferenceGNSSBaroSign        fields.DifferenceGNSSBaroSign
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

func (message AirborneVelocityAirSpeedSupersonic) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AirborneVelocityAirSpeedSupersonic) GetRegister() register.Register {
	return register.BDS09
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirborneVelocityAirSpeedSupersonic) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AirborneVelocityAirSpeedSupersonic) ToString() string {
	return fmt.Sprintf("Message:                         %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change:                   %v\n"+
		"IFR Capability:                  %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Magnetic Heading Status:         %v\n"+
		"Magnetic Heading:                %v\n"+
		"Airspeed Type:                   %v\n"+
		"Airspeed:                        %v\n"+
		"Vertical Rate Source:            %v\n"+
		"Vertical Rate Sign:              %v\n"+
		"Vertical Rate:                   %v\n"+
		"Difference GNSS Baro Sign:       %v\n"+
		"Difference GNSS Baro:            %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.IntentChange.ToString(),
		message.IFRCapability.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.MagneticHeadingStatus.ToString(),
		message.MagneticHeading.ToString(),
		message.AirspeedType.ToString(),
		message.AirspeedSupersonic.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadAirborneVelocityAirSpeedSupersonic reads a message at the format AirborneVelocity / Subtype 4 (Airspeed supersonic)
func ReadAirborneVelocityAirSpeedSupersonic(data []byte) (*AirborneVelocityAirSpeedSupersonic, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 19, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeAirspeedSupersonic {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAirborneVelocityAirSpeedSupersonic", subType.ToString())
	}

	return &AirborneVelocityAirSpeedSupersonic{
		IntentChange:                  fields.ReadIntentChange(data),
		IFRCapability:                 fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		MagneticHeadingStatus:         fields.ReadMagneticHeadingStatus(data),
		MagneticHeading:               fields.ReadMagneticHeading(data),
		AirspeedType:                  fields.ReadAirspeedType(data),
		AirspeedSupersonic:            fields.ReadAirspeedSupersonic(data),
		VerticalRateSource:            fields.ReadVerticalRateSource(data),
		VerticalRateSign:              fields.ReadVerticalRateSign(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		DifferenceGNSSBaroSign:        fields.ReadDifferenceGNSSBaroSign(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
