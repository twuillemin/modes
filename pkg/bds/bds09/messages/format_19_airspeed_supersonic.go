package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19AirspeedSupersonic is a message at the format BDS 9,0
type Format19AirspeedSupersonic struct {
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

// GetName returns the name of the message
func (message *Format19AirspeedSupersonic) GetName() string {
	return bds09Name
}

// GetBDS returns the binary data format
func (message *Format19AirspeedSupersonic) GetBDS() string {
	return bds09Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format19AirspeedSupersonic) GetFormatTypeCode() byte {
	return 19
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message *Format19AirspeedSupersonic) GetSubtype() fields.Subtype {
	return fields.SubtypeAirspeedSupersonic
}

// GetIntentChange returns the IntentChange
func (message *Format19AirspeedSupersonic) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message *Format19AirspeedSupersonic) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message *Format19AirspeedSupersonic) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message *Format19AirspeedSupersonic) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message *Format19AirspeedSupersonic) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message *Format19AirspeedSupersonic) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message *Format19AirspeedSupersonic) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message *Format19AirspeedSupersonic) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19AirspeedSupersonic) ToString() string {
	return fmt.Sprintf("Message:                         %v - %v (%v)\n"+
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
		message.AirspeedSupersonic.ToString(),
		message.VerticalRateSource.ToString(),
		message.VerticalRateSign.ToString(),
		message.VerticalRate.ToString(),
		message.DifferenceGNSSBaroSign.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// readFormat19AirspeedSupersonic reads a message at the format BDS 6,5
func readFormat19AirspeedSupersonic(data []byte) (*Format19AirspeedSupersonic, error) {

	return &Format19AirspeedSupersonic{
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
