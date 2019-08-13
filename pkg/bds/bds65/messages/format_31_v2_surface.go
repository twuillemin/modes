package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V2Surface is a message at the format BDS 6,5 the ADSB V2 / Surface
//
// Specified in Doc 9871 / C.2.3.10
type Format31V2Surface struct {
	Subtype                              fields.SubtypeV2
	SurfaceCapabilityClass               fields.SurfaceCapabilityClassV2
	LengthAndWidth                       fields.LengthWidth
	OperationalMode                      fields.SurfaceOperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplementA                       fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	SourceIntegrityLevel                 fields.SourceIntegrityLevel
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
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplementA.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.SourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V2Surface reads a message at the format Format31V2Surface
func ReadFormat31V2Surface(data []byte) (*Format31V2Surface, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat31V2Surface", formatTypeCode)
	}

	subType := fields.ReadSubtypeV2(data)
	if subType != fields.SubtypeV2Surface {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat31V2Surface", subType.ToString())
	}

	// Check the ADSB Level
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion2 {
		return nil, fmt.Errorf("the data are given at at %v format and can not be read by ReadFormat31V2Surface", detectedADSBLevel.ToString())
	}

	capabilityContent := (data[1] & 0xC0) >> 6
	if capabilityContent != 0 {
		return nil, fmt.Errorf("the Capability Content (field Capability Class) must be 0 (%v given)", capabilityContent)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	return &Format31V2Surface{
		Subtype:                              subType,
		SurfaceCapabilityClass:               fields.ReadSurfaceCapabilityClassV2(data),
		LengthAndWidth:                       fields.ReadAircraftLengthAndWidth(data),
		OperationalMode:                      fields.ReadSurfaceOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplementA:                       fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		SourceIntegrityLevel:                 fields.ReadSourceIntegrityLevel(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
	}, nil
}
