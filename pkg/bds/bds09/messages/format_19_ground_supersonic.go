package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
)

// Format19GroundSupersonic is a message at the format BDS 9,0
type Format19GroundSupersonic struct {
	IntentChangeFlag              fields.IntentChangeFlag
	IFRCapabilityFlag             fields.IFRCapabilityFlag
	NavigationUncertaintyCategory fields.NavigationUncertaintyCategory
	DirectionBitEWVelocity        fields.DirectionBitEWVelocity
	VelocityEWSupersonic          fields.VelocityEWSupersonic
	DirectionBitNSVelocity        fields.DirectionBitNSVelocity
	VelocityNSSupersonic          fields.VelocityNSSupersonic
	SourceBitVerticalRate         fields.SourceBitVerticalRate
	SignBitVerticalRate           fields.SignBitVerticalRate
	VerticalRate                  fields.VerticalRate
	GNSSAltitudeSignBit           fields.GNSSAltitudeSignBit
	DifferenceGNSSBaro            fields.DifferenceGNSSBaro
}

// GetName returns the name of the message
func (message *Format19GroundSupersonic) GetName() string {
	return bds09Name
}

// GetBDS returns the binary data format
func (message *Format19GroundSupersonic) GetBDS() string {
	return bds09Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format19GroundSupersonic) GetFormatTypeCode() byte {
	return 19
}

// GetAirborneVelocitySubtype returns the code of the Operational Status Sub Type
func (message *Format19GroundSupersonic) GetAirborneVelocitySubtype() fields.AirborneVelocitySubtype {
	return fields.AVSCGroundSpeedSupersonic
}

// ToString returns a basic, but readable, representation of the message
func (message Format19GroundSupersonic) ToString() string {
	return fmt.Sprintf("Message:                         %v (%v)\n"+
		"Format Type Code:                %v\n"+
		"Subtype:                         %v\n"+
		"Intent Change Flag:              %v\n"+
		"IFR Capability Flag:             %v\n"+
		"Navigation Uncertainty Category: %v\n"+
		"Direction Bit EW Velocity:       %v\n"+
		"Velocity EW Supersonic:          %v\n"+
		"Direction Bit NS Velocity:       %v\n"+
		"Velocity NS Supersonic:          %v\n"+
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
		message.VelocityEWSupersonic.ToString(),
		message.DirectionBitNSVelocity.ToString(),
		message.VelocityNSSupersonic.ToString(),
		message.SourceBitVerticalRate.ToString(),
		message.SignBitVerticalRate.ToString(),
		message.VerticalRate.ToString(),
		message.GNSSAltitudeSignBit.ToString(),
		message.DifferenceGNSSBaro.ToString())
}

// ReadFormat19GroundSupersonic reads a message at the format BDS 6,5
func ReadFormat19GroundSupersonic(data []byte) (*Format19GroundSupersonic, error) {

	return &Format19GroundSupersonic{
		IntentChangeFlag:              fields.ReadIntentChangeFlag(data),
		IFRCapabilityFlag:             fields.ReadIFRCapabilityFlag(data),
		NavigationUncertaintyCategory: fields.ReadNavigationUncertaintyCategory(data),
		DirectionBitEWVelocity:        fields.ReadDirectionBitEWVelocity(data),
		VelocityEWSupersonic:          fields.ReadVelocityEWSupersonic(data),
		DirectionBitNSVelocity:        fields.ReadDirectionBitNSVelocity(data),
		VelocityNSSupersonic:          fields.ReadVelocityNSSupersonic(data),
		SourceBitVerticalRate:         fields.ReadSourceBitVerticalRate(data),
		SignBitVerticalRate:           fields.ReadSignBitVerticalRate(data),
		VerticalRate:                  fields.ReadVerticalRate(data),
		GNSSAltitudeSignBit:           fields.ReadGNSSAltitudeSignBit(data),
		DifferenceGNSSBaro:            fields.ReadDifferenceGNSSBaro(data),
	}, nil
}
