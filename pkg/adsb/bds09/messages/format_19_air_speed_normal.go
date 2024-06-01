package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds"
)

// Format19AirSpeedNormal is a message at the format BDS 9,0
type Format19AirSpeedNormal struct {
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

// GetMessageFormat returns the ADSB format of the message
func (message Format19AirSpeedNormal) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format19
}

// GetRegister returns the register of the message
func (message Format19AirSpeedNormal) GetRegister() bds.Register {
	return adsb.Format19.GetRegister()
}

// GetSubtype returns the code of the Operational Status Sub Type
func (message Format19AirSpeedNormal) GetSubtype() adsb.Subtype {
	return fields.SubtypeAirspeedNormal
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format19AirSpeedNormal) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format19AirSpeedNormal) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetIntentChange returns the IntentChange
func (message Format19AirSpeedNormal) GetIntentChange() fields.IntentChange {
	return message.IntentChange
}

// GetIFRCapability returns the IFRCapability
func (message Format19AirSpeedNormal) GetIFRCapability() fields.IFRCapability {
	return message.IFRCapability
}

// GetNavigationUncertaintyCategory returns the NavigationUncertaintyCategory
func (message Format19AirSpeedNormal) GetNavigationUncertaintyCategory() fields.NavigationUncertaintyCategory {
	return message.NavigationUncertaintyCategory
}

// GetVerticalRateSource returns the VerticalRateSource
func (message Format19AirSpeedNormal) GetVerticalRateSource() fields.VerticalRateSource {
	return message.VerticalRateSource
}

// GetVerticalRateSign returns the VerticalRateSign
func (message Format19AirSpeedNormal) GetVerticalRateSign() fields.VerticalRateSign {
	return message.VerticalRateSign
}

// GetVerticalRate returns the VerticalRate
func (message Format19AirSpeedNormal) GetVerticalRate() fields.VerticalRate {
	return message.VerticalRate
}

// GetDifferenceGNSSBaroSign returns the DifferenceGNSSBaroSign
func (message Format19AirSpeedNormal) GetDifferenceGNSSBaroSign() fields.DifferenceGNSSBaroSign {
	return message.DifferenceGNSSBaroSign
}

// GetDifferenceGNSSBaro returns the DifferenceGNSSBaro
func (message Format19AirSpeedNormal) GetDifferenceGNSSBaro() fields.DifferenceGNSSBaro {
	return message.DifferenceGNSSBaro
}

// ToString returns a basic, but readable, representation of the message
func (message Format19AirSpeedNormal) ToString() string {
	return fmt.Sprintf("Message:                         %v\n"+
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
		adsb.GetMessageFormatInformation(&message),
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

// ReadFormat19AirSpeedNormal reads a message at the format Format19 / Subtype 3 (Airspeed normal)
func ReadFormat19AirSpeedNormal(data []byte) (*Format19AirSpeedNormal, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format19.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat19AirSpeedNormal", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeAirspeedNormal {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat19AirSpeedNormal", subType.ToString())
	}

	return &Format19AirSpeedNormal{
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
