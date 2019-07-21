package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
)

// Format29Subtype0 is a message at the format BDS 6,2
type Format29Subtype0 struct {
	Subtype                                fields.Subtype
	VerticalDataAvailableSourceIndicator   fields.VerticalDataAvailableSourceIndicator
	TargetAltitudeType                     fields.TargetAltitudeType
	TargetAltitudeCapability               fields.TargetAltitudeCapability
	VerticalModeIndicator                  fields.VerticalModeIndicator
	TargetAltitude                         fields.TargetAltitude
	HorizontalDataAvailableSourceIndicator fields.HorizontalDataAvailableSourceIndicator
	TargetHeadingTrackAngle                fields.TargetHeadingTrackAngle
	TargetHeadingTrackIndicator            fields.TargetHeadingTrackIndicator
	HorizontalModeIndicator                fields.HorizontalModeIndicator
	NavigationalAccuracyCategoryPosition   fields.NavigationalAccuracyCategoryPositionV1
	NICBaro                                fields.NICBaro
	SurveillanceIntegrityLevel             fields.SurveillanceIntegrityLevel
	CapabilityModeCode                     fields.CapabilityModeCode
	EmergencyPriorityStatus                fields.EmergencyPriorityStatus
}

// GetName returns the name of the message
func (message *Format29Subtype0) GetName() string {
	return bds62Name
}

// GetBDS returns the binary data format
func (message *Format29Subtype0) GetBDS() string {
	return bds62Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format29Subtype0) GetFormatTypeCode() byte {
	return 29
}

// GetSubtype returns the Subtype
func (message *Format29Subtype0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format29Subtype0) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format29Subtype0) ToString() string {
	return fmt.Sprintf("Message:                                      %v - %v (%v)\n"+
		"Subtype:                                      %v\n"+
		"Vertical Data Available / Source Indicator:   %v\n"+
		"Target Altitude Type:                         %v\n"+
		"Target Altitude Capability:                   %v\n"+
		"Vertical Mode Indicator:                      %v\n"+
		"Target Altitude:                              %v\n"+
		"Horizontal Data Available / Source Indicator: %v\n"+
		"Target Heading / Track Angle:                 %v\n"+
		"Target Heading / Track Indicator:             %v\n"+
		"Horizontal Mode Indicator:                    %v\n"+
		"Navigation Accuracy Category - Position:      %v\n"+
		"Navigation Integrity Category - Baro:         %v\n"+
		"Surveillance Integrity Level:                 %v\n"+
		"Capability / Mode Codes:                      %v\n"+
		"Emergency / Priority Status:                  %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSubtype().ToString(),
		message.VerticalDataAvailableSourceIndicator.ToString(),
		message.TargetAltitudeType.ToString(),
		message.TargetAltitudeCapability.ToString(),
		message.VerticalModeIndicator.ToString(),
		message.TargetAltitude.ToString(),
		message.HorizontalDataAvailableSourceIndicator.ToString(),
		message.TargetHeadingTrackAngle.ToString(),
		message.TargetHeadingTrackIndicator.ToString(),
		message.HorizontalModeIndicator.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.NICBaro.ToString(),
		message.SurveillanceIntegrityLevel.ToString(),
		message.CapabilityModeCode.ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// readFormat29Subtype0 reads a message at the format BDS 6,2
func readFormat29Subtype0(data []byte) (*Format29Subtype0, error) {

	return &Format29Subtype0{
		Subtype:                                fields.ReadSubtypeV1(data),
		VerticalDataAvailableSourceIndicator:   fields.ReadVerticalDataAvailableSourceIndicator(data),
		TargetAltitudeType:                     fields.ReadTargetAltitudeType(data),
		TargetAltitudeCapability:               fields.ReadTargetAltitudeCapability(data),
		VerticalModeIndicator:                  fields.ReadVerticalModeIndicator(data),
		TargetAltitude:                         fields.ReadTargetAltitude(data),
		HorizontalDataAvailableSourceIndicator: fields.ReadHorizontalDataAvailableSourceIndicator(data),
		TargetHeadingTrackAngle:                fields.ReadTargetHeadingTrackAngle(data),
		TargetHeadingTrackIndicator:            fields.ReadTargetHeadingTrackIndicator(data),
		HorizontalModeIndicator:                fields.ReadHorizontalModeIndicator(data),
		NavigationalAccuracyCategoryPosition:   fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		NICBaro:                                fields.ReadNICBaro(data),
		SurveillanceIntegrityLevel:             fields.ReadSurveillanceIntegrityLevel(data),
		CapabilityModeCode:                     fields.ReadCapabilityModeCode(data),
		EmergencyPriorityStatus:                fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
