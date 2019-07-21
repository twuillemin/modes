package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19AirspeedNormal is a message at the format BDS 9,0
type Format19AirspeedNormal struct {
	IntentChange                  fields.IntentChange
	IFRCapability                 fields.IFRCapability
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	MagneticHeadingStatus         fields.MagneticHeadingStatus
	MagneticHeading               fields.MagneticHeading
	AirspeedType                  fields.AirspeedType
	AirspeedNormal                fields.AirspeedNormal
	VerticalRateSource            fields.VerticalRateSource
	VerticalRateSign              fields.VerticalRateSign
	VerticalRate                  fields.VerticalRate
	DifferenceGNSSBaroSign        fields.DifferenceGNSSBaroSign
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

// GetName returns the name of the message
func (message *Format19AirspeedNormal) GetName() string {
	return bds09Name
}

// GetBDS returns the binary data format
func (message *Format19AirspeedNormal) GetBDS() string {
	return bds09Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format19AirspeedNormal) GetFormatTypeCode() byte {
	return 19
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message *Format19AirspeedNormal) GetSubtype() fields.Subtype {
	return fields.SubtypeAirspeedNormal
}

// GetIntentChange returns the IntentChange
func (message *Format19AirspeedNormal) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message *Format19AirspeedNormal) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message *Format19AirspeedNormal) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message *Format19AirspeedNormal) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message *Format19AirspeedNormal) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message *Format19AirspeedNormal) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message *Format19AirspeedNormal) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message *Format19AirspeedNormal) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19AirspeedNormal) ToString() string {
	return fmt.Sprintf("Message:                         %v - %v (%v)\n"+
		"Subtype:                         %v\n"+
		"Intent Change:                   %v\n"+
		"IFR Capability:                  %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Magnetic Heading Status:         %v\n"+
		"Magnetic Heading:                %v\n"+
		"Airspeed Type:                   %v\n"+
		"Airspeed:                        %v\n"+
		"Vertical Rate Source :           %v\n"+
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
		message.MagneticHeadingStatus.ToString(),
		message.MagneticHeading.ToString(),
		message.AirspeedType.ToString(),
		message.AirspeedNormal.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// readFormat19AirspeedNormal reads a message at the format BDS 6,5
func readFormat19AirspeedNormal(data []byte) (*Format19AirspeedNormal, error) {

	return &Format19AirspeedNormal{
		IntentChange:                  fields.ReadIntentChange(data),
		IFRCapability:                 fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		MagneticHeadingStatus:         fields.ReadMagneticHeadingStatus(data),
		MagneticHeading:               fields.ReadMagneticHeading(data),
		AirspeedType:                  fields.ReadAirspeedType(data),
		AirspeedNormal:                fields.ReadAirspeedNormal(data),
		VerticalRateSource:            fields.ReadVerticalRateSource(data),
		VerticalRateSign:              fields.ReadVerticalRateSign(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		DifferenceGNSSBaroSign:        fields.ReadDifferenceGNSSBaroSign(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
