package bds09

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirborneVelocityAirSpeedSupersonic is a message at the format BDS 9,0
type AirborneVelocityAirSpeedSupersonic struct {
	FormatTypeCode                 byte
	Subtype                        fields.Subtype
	IntentChange                   fields.IntentChange
	IFRCapability                  fields.IFRCapability
	NavigationUncertaintyCategory  fields.NavigationUncertaintyCategory
	MagneticHeadingStatus          fields.MagneticHeadingStatus
	MagneticHeading                float32
	AirspeedStatus                 fields.NumericValueStatus
	Airspeed                       uint16
	VerticalRateSource             fields.VerticalRateSource
	VerticalRateStatus             fields.NumericValueStatus
	VerticalRate                   int16
	HeightDifferenceFromBaroStatus fields.NumericValueStatus
	HeightDifferenceFromBaro       int16
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
	return fmt.Sprintf(""+
		"Message:                             %v\n"+
		"Subtype:                             %v\n"+
		"Intent Change:                       %v\n"+
		"IFR Capability:                      %v\n"+
		"Navigation Uncertainty Category:     %v\n"+
		"Magnetic Heading Status:             %v\n"+
		"Magnetic Heading (degrees):          %v\n"+
		"Air Speed Status:                    %v\n"+
		"Air Speed (knots):                   %v\n"+
		"Vertical Rate Source:                %v\n"+
		"Vertical Rate Status:                %v\n"+
		"Vertical Rate (ft/min):              %v\n"+
		"Geom. Height Diff. From Baro Status: %v\n"+
		"Geom. Height Diff. From Baro (ft):   %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.IntentChange.ToString(),
		message.IFRCapability.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.MagneticHeadingStatus.ToString(),
		message.MagneticHeading,
		message.AirspeedStatus.ToString(),
		message.Airspeed,
		message.VerticalRateSource.ToString(),
		message.VerticalRateStatus.ToString(),
		message.VerticalRate,
		message.HeightDifferenceFromBaroStatus.ToString(),
		message.HeightDifferenceFromBaro)
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

	magneticHeading, magneticHeadingStatus := fields.ReadMagneticHeading(data)
	airSpeed, airSpeedStatus := fields.ReadAirspeedSupersonic(data)
	verticalRate, verticalRateStatus := fields.ReadVerticalRate(data)
	diffBaro, diffBaroStatus := fields.ReadHeightDifference(data)

	return &AirborneVelocityAirSpeedSupersonic{
		FormatTypeCode:                 formatTypeCode,
		Subtype:                        subType,
		IntentChange:                   fields.ReadIntentChange(data),
		IFRCapability:                  fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory:  fields.ReadNavigationUncertaintyCategory(data),
		MagneticHeadingStatus:          magneticHeadingStatus,
		MagneticHeading:                magneticHeading,
		AirspeedStatus:                 airSpeedStatus,
		Airspeed:                       airSpeed,
		VerticalRateSource:             fields.ReadVerticalRateSource(data),
		VerticalRateStatus:             verticalRateStatus,
		VerticalRate:                   verticalRate,
		HeightDifferenceFromBaroStatus: diffBaroStatus,
		HeightDifferenceFromBaro:       diffBaro,
	}, nil
}
