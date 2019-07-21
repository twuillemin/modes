package messages

import (
	"fmt"
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
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevelV2
	TrackAngleHeading                    fields.TrackAngleHeading
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
	SurveillanceIntegrityLevelSupplement fields.SurveillanceIntegrityLevelSupplement
}

// GetName returns the name of the message
func (message *Format31V2Surface) GetName() string {
	return bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V2Surface) GetBDS() string {
	return bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V2Surface) GetFormatTypeCode() byte {
	return 31
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V2Surface) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V2Surface) ToString() string {
	return fmt.Sprintf("Message:                                 %v - %v (%v)\n"+
		"Subtype:                                 %v\n"+
		"Airborne Capability Class:\n%v\n"+
		"Aircraft Length And Width                %v\n"+
		"Operational Mode:\n%v\n"+
		"ADSV Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity Level:            %v\n"+
		"Surveillance Integrity Level Supplement: %v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.Subtype.ToString(),
		common.PrefixMultiLine(message.SurfaceCapabilityClass.ToString(), "    "),
		message.LengthAndWidth.ToString(),
		common.PrefixMultiLine(message.SurfaceOperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.SurveillanceIntegrityLevel.ToString(),
		message.SurveillanceIntegrityLevelSupplement.ToString(),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V2Surface reads a message at the format Format31V2Surface
func ReadFormat31V2Surface(data []byte) (*Format31V2Surface, error) {

	return &Format31V2Surface{
		Subtype:                              fields.ReadSubtypeV2(data),
		SurfaceCapabilityClass:               fields.ReadSurfaceCapabilityClassV2(data),
		LengthAndWidth:                       fields.ReadAircraftLengthAndWidth(data),
		SurfaceOperationalMode:               fields.ReadSurfaceOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevelV2(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
		SurveillanceIntegrityLevelSupplement: fields.ReadSurveillanceIntegrityLevelSupplement(data),
	}, nil
}
