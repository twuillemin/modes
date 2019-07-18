package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V1Airborne is a message at the format BDS 6,5 the ADSB V1 / Airborne
type Format31V1Airborne struct {
	AirborneCapabilityClass              fields.CapabilityClassAirborne
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplement
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPosition
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	NICBaro                              fields.NICBaro
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

// GetName returns the name of the message
func (message *Format31V1Airborne) GetName() string {
	return bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V1Airborne) GetBDS() string {
	return bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V1Airborne) GetFormatTypeCode() byte {
	return 31
}

// GetOperationalStatusSubTypeCode returns the code of the Operational Status Sub Type
func (message *Format31V1Airborne) GetOperationalStatusSubTypeCode() byte {
	return 0
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Airborne) ToString() string {
	return fmt.Sprintf("Message:          %v (%v)\n"+
		"SubType:                              0 - Airborne\n"+
		"AirborneCapabilityClass:\n%v\n"+
		"OperationalMode:\n%v\n"+
		"VersionNumber:                        %v\n"+
		"NICSupplement:                        %v\n"+
		"NavigationalAccuracyCategoryPosition: %v\n"+
		"SurveillanceIntegrityLevel:\n%v\n"+
		"NICBaro:                              %v\n"+
		"HorizontalReferenceDirection:         %v",
		message.GetBDS(),
		message.GetName(),
		common.PrefixMultiLine(message.AirborneCapabilityClass.ToString(), " - "),
		common.PrefixMultiLine(message.OperationalMode.ToString(), " - "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), " - "),
		message.NICBaro.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V1Airborne reads a message at the format Format31V1Airborne
func ReadFormat31V1Airborne(data []byte) (*Format31V1Airborne, error) {

	return &Format31V1Airborne{
		AirborneCapabilityClass:              fields.ReadCapabilityClassAirborne(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplement(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPosition(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
