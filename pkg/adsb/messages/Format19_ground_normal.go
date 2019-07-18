package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format19GroundNormal is a message at the format BDS 9,0
type Format19GroundNormal struct {
	IntentChangeFlag              fields.IntentChangeFlag
	IFRCapabilityFlag             fields.IFRCapabilityFlag
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	DirectionBitEWVelocity        fields.DirectionBitEWVelocity
	VelocityEWNormal              fields.VelocityEWNormal
	DirectionBitNSVelocity        fields.DirectionBitNSVelocity
	VelocityNSNormal              fields.VelocityNSNormal
	SourceBitVerticalRate         fields.SourceBitVerticalRate
	SignBitVerticalRate           fields.SignBitVerticalRate
	VerticalRate                  fields.VerticalRate
	GNSSAltitudeSignBit           fields.GNSSAltitudeSignBit
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

// GetName returns the name of the message
func (message *Format19GroundNormal) GetName() string {
	return bds09Name
}

// GetBDS returns the binary data format
func (message *Format19GroundNormal) GetBDS() string {
	return bds09Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format19GroundNormal) GetFormatTypeCode() byte {
	return 19
}

// GetAirborneVelocitySubtype returns the code of the Operational Status Sub Type
func (message *Format19GroundNormal) GetAirborneVelocitySubtype() fields.AirborneVelocitySubtype {
	return fields.AVSCGroundSpeedNormal
}

// ToString returns a basic, but readable, representation of the message
func (message Format19GroundNormal) ToString() string {
	return fmt.Sprintf("Message:                         %v (%v)\n"+
		"Format Type Code:                %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change Flag:              %v\n"+
		"IFR Capability Flag:             %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Direction Bit EW Velocity:       %v\n"+
		"Velocity EW Normal:              %v\n"+
		"Direction Bit NS Velocity:       %v\n"+
		"Velocity NS Normal:              %v\n"+
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
		message.DirectionBitEWVelocity.ToString(),
		message.VelocityEWNormal.ToString(),
		message.DirectionBitNSVelocity.ToString(),
		message.VelocityNSNormal.ToString(),
		message.SourceBitVerticalRate.ToString(),
		message.SignBitVerticalRate.ToString(),
		message.VerticalRate.ToString(),
		message.GNSSAltitudeSignBit.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadFormat19GroundNormal reads a message at the format BDS 6,5
func ReadFormat19GroundNormal(data []byte) (*Format19GroundNormal, error) {

	return &Format19GroundNormal{
		IntentChangeFlag:              fields.ReadIntentChangeFlag(data),
		IFRCapabilityFlag:             fields.ReadIFRCapabilityFlag(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		DirectionBitEWVelocity:        fields.ReadDirectionBitEWVelocity(data),
		VelocityEWNormal:              fields.ReadVelocityEWNormal(data),
		DirectionBitNSVelocity:        fields.ReadDirectionBitNSVelocity(data),
		VelocityNSNormal:              fields.ReadVelocityNSNormal(data),
		SourceBitVerticalRate:         fields.ReadSourceBitVerticalRate(data),
		SignBitVerticalRate:           fields.ReadSignBitVerticalRate(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		GNSSAltitudeSignBit:           fields.ReadGNSSAltitudeSignBit(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
