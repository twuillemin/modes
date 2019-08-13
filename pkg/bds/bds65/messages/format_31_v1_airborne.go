package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V1Airborne is a message at the format BDS 6,5 the ADSB V1 / Airborne
//
// Specified in Doc 9871 / B.2.3.10
type Format31V1Airborne struct {
	Subtype                              fields.SubtypeV1
	AirborneCapabilityClass              fields.AirborneCapabilityClassV1
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV1
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	NICBaro                              fields.NICBaro
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format31V1Airborne) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31V1
}

// GetRegister returns the register of the message
func (message *Format31V1Airborne) GetRegister() bds.Register {
	return adsb.Format31V1.GetRegister()
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V1Airborne) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Airborne) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Operational Mode:\n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity Level:\n%v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.Format31V1.ToString(),
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

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format31V1", formatTypeCode)
	}

	return &Format31V1Airborne{
		Subtype:                              fields.ReadSubtypeV1(data),
		AirborneCapabilityClass:              fields.ReadAirborneCapabilityClassV1(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
