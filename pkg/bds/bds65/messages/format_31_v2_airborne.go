package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V2Airborne is a message at the format BDS 6,5 the ADSB V2 / Airborne
type Format31V2Airborne struct {
	Subtype                              fields.SubtypeV2
	AirborneCapabilityClass              fields.AirborneCapabilityClassV2
	OperationalMode                      fields.AirborneOperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplementA                       fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	GeometricVerticalAccuracy            fields.GeometricVerticalAccuracy
	SourceIntegrityLevel                 fields.SourceIntegrityLevel
	NICBaro                              fields.NICBaro
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
	SourceIntegrityLevelSupplement       fields.SourceIntegrityLevelSupplement
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format31V2Airborne) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31V2
}

// GetRegister returns the register of the message
func (message *Format31V2Airborne) GetRegister() bds.Register {
	return adsb.Format31V2.GetRegister()
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V2Airborne) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V2Airborne) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Operational Mode:\n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Geometric Vertical Accuracy:             %v\n"+
		"Source Integrity Level:                  %v\n"+
		"Source Integrity Level Supplement:       %v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.Format31V2.ToString(),
		message.GetSubtype().ToString(),
		common.PrefixMultiLine(message.AirborneCapabilityClass.ToString(), "    "),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplementA.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.GeometricVerticalAccuracy.ToString(),
		message.SourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.NICBaro.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V2Airborne reads a message at the format Format31V2Airborne
func ReadFormat31V2Airborne(data []byte) (*Format31V2Airborne, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format31V2", formatTypeCode)
	}

	return &Format31V2Airborne{
		Subtype:                              fields.ReadSubtypeV2(data),
		AirborneCapabilityClass:              fields.ReadAirborneCapabilityClassV2(data),
		OperationalMode:                      fields.ReadAirborneOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplementA:                       fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		GeometricVerticalAccuracy:            fields.ReadGeometricVerticalAccuracy(data),
		SourceIntegrityLevel:                 fields.ReadSourceIntegrityLevel(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
	}, nil
}
