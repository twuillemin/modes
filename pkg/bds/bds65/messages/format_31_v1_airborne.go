package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V1Airborne is a message at the format BDS 6,5 the ADSB V1 / Airborne
type Format31V1Airborne struct {
	Subtype                              fields.SubtypeV1
	AirborneCapabilityClass              fields.AirborneCapabilityClassV1
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV1
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevelV1
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

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V1Airborne) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Airborne) ToString() string {
	return fmt.Sprintf("Message:                                 %v - %v (%v)\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Operational Mode:\n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity Level:\n%v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetBDS(),
		message.GetSubtype().ToString(),
		common.PrefixMultiLine(message.AirborneCapabilityClass.ToString(), "    "),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), "    "),
		message.NICBaro.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V1Airborne reads a message at the format Format31V1Airborne
func ReadFormat31V1Airborne(data []byte) (*Format31V1Airborne, error) {

	return &Format31V1Airborne{
		Subtype:                              fields.ReadSubtypeV1(data),
		AirborneCapabilityClass:              fields.ReadAirborneCapabilityClassV1(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevelV1(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
