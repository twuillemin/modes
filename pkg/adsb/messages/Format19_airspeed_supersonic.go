package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format19AirspeedSupersonic is a message at the format BDS 9,0
type Format19AirspeedSupersonic struct {
	IntentChangeFlag              fields.IntentChangeFlag
	IFRCapabilityFlag             fields.IFRCapabilityFlag
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	MagneticHeadingStatus         fields.MagneticHeadingStatus
	MagneticHeading               fields.MagneticHeading
	AirspeedType                  fields.AirspeedType
	AirspeedSupersonic            fields.AirspeedSupersonic
	SourceBitVerticalRate         fields.SourceBitVerticalRate
	SignBitVerticalRate           fields.SignBitVerticalRate
	VerticalRate                  fields.VerticalRate
	GNSSAltitudeSignBit           fields.GNSSAltitudeSignBit
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

// GetAirborneVelocitySubtype returns the code of the Operational Status Sub Type
func (message *Format19AirspeedSupersonic) GetAirborneVelocitySubtype() fields.AirborneVelocitySubtype {
	return fields.AVSCAirspeedSupersonic
}

// ToString returns a basic, but readable, representation of the message
func (message Format19AirspeedSupersonic) ToString() string {
	return fmt.Sprintf("Message:                         %v (%v)\n"+
		"Format Type Code:                %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change Flag:              %v\n"+
		"IFR Capability Flag:             %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Magnetic Heading Status:         %v\n"+
		"Magnetic Heading:                %v\n"+
		"Airspeed Type:                   %v\n"+
		"Airspeed:                        %v\n"+
		"Source Bit Vertical Rate:        %v\n"+
		"Sign Bit Vertical Rate:          %v\n"+
		"Vertical Rate:                   %v\n"+
		"GNSS Altitude Sign Bit:          %v\n"+
		"Difference GNSS Baro:            %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetAirborneVelocitySubtype().ToString(),
		message.IntentChangeFlag.ToString(),
		message.IFRCapabilityFlag.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.MagneticHeadingStatus.ToString(),
		message.MagneticHeading.ToString(),
		message.AirspeedType.ToString(),
		message.AirspeedSupersonic.ToString(),
		message.SourceBitVerticalRate.ToString(),
		message.SignBitVerticalRate.ToString(),
		message.VerticalRate.ToString(),
		message.GNSSAltitudeSignBit.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadFormat19AirspeedSupersonic reads a message at the format BDS 6,5
func ReadFormat19AirspeedSupersonic(data []byte) (*Format19AirspeedSupersonic, error) {

	return &Format19AirspeedSupersonic{
		IntentChangeFlag:              fields.ReadIntentChangeFlag(data),
		IFRCapabilityFlag:             fields.ReadIFRCapabilityFlag(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		MagneticHeadingStatus:         fields.ReadMagneticHeadingStatus(data),
		MagneticHeading:               fields.ReadMagneticHeading(data),
		AirspeedType:                  fields.ReadAirspeedType(data),
		AirspeedSupersonic:            fields.ReadAirspeedSupersonic(data),
		SourceBitVerticalRate:         fields.ReadSourceBitVerticalRate(data),
		SignBitVerticalRate:           fields.ReadSignBitVerticalRate(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		GNSSAltitudeSignBit:           fields.ReadGNSSAltitudeSignBit(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
