package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V1Surface is a message at the format BDS 6,5 the ADSB V1 / Surface
//
// Specified in Doc 9871 / B.2.3.10
type Format31V1Surface struct {
	Subtype                              fields.SubtypeV1
	SurfaceCapabilityClass               fields.SurfaceCapabilityClassV1
	LengthAndWidth                       fields.LengthWidth
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV1
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	TrackAngleHeading                    fields.TrackAngleHeading
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format31V1Surface) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31V1
}

// GetRegister returns the register of the message
func (message *Format31V1Surface) GetRegister() bds.Register {
	return adsb.Format31V1.GetRegister()
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V1Surface) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Surface) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Aircraft Length And Width                %v\n"+
		"Operational Mode:\n%v\n"+
		"ADSV Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity Level:\n%v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.Format31V1.ToString(),
		message.Subtype.ToString(),
		common.PrefixMultiLine(message.SurfaceCapabilityClass.ToString(), "    "),
		message.LengthAndWidth.ToString(),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), "    "),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V1Surface reads a message at the format Format31V1Surface
func ReadFormat31V1Surface(data []byte) (*Format31V1Surface, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat31V1Surface", formatTypeCode)
	}

	subType := fields.ReadSubtypeV1(data)
	if subType != fields.SubtypeV1Surface {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat31V1Surface", subType.ToString())
	}

	// Check the ADSB Level
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion1 {
		return nil, fmt.Errorf("the data are given at %v format and can not be read by ReadFormat31V1Surface", detectedADSBLevel.ToString())
	}

	serviceLevel := (data[1]&0xC0)>>4 + (data[1]&0x0C)>>2
	if serviceLevel != 0 {
		return nil, fmt.Errorf("the ServiceLevel (field Capability Class) must be 0 (%v given)", serviceLevel)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	return &Format31V1Surface{
		Subtype:                              subType,
		SurfaceCapabilityClass:               fields.ReadSurfaceCapabilityClassV1(data),
		LengthAndWidth:                       fields.ReadAircraftLengthAndWidth(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
