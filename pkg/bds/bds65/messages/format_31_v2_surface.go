package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V2Surface is a message at the format BDS 6,5 the ADSB V2 / Surface
type Format31V2Surface struct {
	Subtype                              fields.SubtypeV2
	SurfaceCapabilityClass               fields.SurfaceCapabilityClassV2
	LengthAndWidth                       fields.LengthWidth
	SurfaceOperationalMode               fields.SurfaceOperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	SSourceIntegrityLevel                fields.SourceIntegrityLevel
	TrackAngleHeading                    fields.TrackAngleHeading
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
	SourceIntegrityLevelSupplement       fields.SourceIntegrityLevelSupplement
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format31V2Surface) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31V2
}

// GetRegister returns the register of the message
func (message *Format31V2Surface) GetRegister() bds.Register {
	return adsb.Format31V2.GetRegister()
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V2Surface) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V2Surface) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Aircraft Length And Width                %v\n"+
		"Operational Mode:\n%v\n"+
		"ADSV Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Source Integrity Level:                  %v\n"+
		"Source Integrity Level Supplement:       %v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.Format31V2.ToString(),
		message.Subtype.ToString(),
		common.PrefixMultiLine(message.SurfaceCapabilityClass.ToString(), "    "),
		message.LengthAndWidth.ToString(),
		common.PrefixMultiLine(message.SurfaceOperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.SSourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V2Surface reads a message at the format Format31V2Surface
func ReadFormat31V2Surface(data []byte) (*Format31V2Surface, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format31V2", formatTypeCode)
	}

	return &Format31V2Surface{
		Subtype:                              fields.ReadSubtypeV2(data),
		SurfaceCapabilityClass:               fields.ReadSurfaceCapabilityClassV2(data),
		LengthAndWidth:                       fields.ReadAircraftLengthAndWidth(data),
		SurfaceOperationalMode:               fields.ReadSurfaceOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		SSourceIntegrityLevel:                fields.ReadSourceIntegrityLevel(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
	}, nil
}