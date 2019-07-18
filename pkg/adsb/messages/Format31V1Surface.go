package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31V1Surface is a message at the format BDS 6,5 the ADSB V1 / Surface
type Format31V1Surface struct {
	SurfaceCapabilityClass               fields.CapabilityClassSurface
	OperationalMode                      fields.OperationalMode
	AircraftLengthAndWidth               fields.AircraftLengthAndWidth
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplement
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPosition
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	TrackAngleHeading                    fields.TrackAngleHeading
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

// GetName returns the name of the message
func (message *Format31V1Surface) GetName() string {
	return bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V1Surface) GetBDS() string {
	return bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V1Surface) GetFormatTypeCode() byte {
	return 31
}

// GetOperationalStatusSubTypeCode returns the code of the Operational Status Sub Type
func (message *Format31V1Surface) GetOperationalStatusSubTypeCode() byte {
	return 1
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Surface) ToString() string {
	return fmt.Sprintf("Message:          %v (%v)\n"+
		"SubType:                              0 - Airborne\n"+
		"AirborneCapabilityClass:\n%v\n"+
		"AircraftLengthAndWidth                %v\n"+
		"OperationalMode:\n%v\n"+
		"VersionNumber:                        %v\n"+
		"NICSupplement:                        %v\n"+
		"NavigationalAccuracyCategoryPosition: %v\n"+
		"SurveillanceIntegrityLevel:\n%v\n"+
		"NICBaro:                              %v\n"+
		"HorizontalReferenceDirection:         %v",
		message.GetBDS(),
		message.GetName(),
		common.PrefixMultiLine(message.SurfaceCapabilityClass.ToString(), " - "),
		message.AircraftLengthAndWidth.ToString(),
		common.PrefixMultiLine(message.OperationalMode.ToString(), " - "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), " - "),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31V1Surface reads a message at the format Format31V1Surface
func ReadFormat31V1Surface(data []byte) (*Format31V1Surface, error) {

	return &Format31V1Surface{
		SurfaceCapabilityClass:               fields.ReadCapabilityClassSurface(data),
		AircraftLengthAndWidth:               fields.ReadAircraftLengthAndWidth(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplement(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPosition(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
