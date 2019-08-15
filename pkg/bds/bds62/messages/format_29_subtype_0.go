package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
)

// Format29Subtype0 is a message at the format BDS 6,2
type Format29Subtype0 struct {
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

// GetMessageFormat returns the ADSB format of the message
func (message Format29Subtype0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format29
}

// GetRegister returns the register of the message
func (message Format29Subtype0) GetRegister() bds.Register {
	return adsb.Format29.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format29Subtype0) GetSubtype() adsb.Subtype {
	return fields.Subtype0
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format29Subtype0) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format29Subtype0) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// ToString returns a basic, but readable, representation of the message
func (message Format29Subtype0) ToString() string {
	return fmt.Sprintf("Message:                                      %v\n"+
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
		"Surveillance Integrity ReaderLevel:                 %v\n"+
		"Capability / Mode Codes:                      %v\n"+
		"Emergency / Priority Status:                  %v",
		adsb.GetMessageFormatInformation(&message),
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
		message.EmergencyPriorityStatus.ToString())
}

// ReadFormat29Subtype0 reads a message at the format Format 29 / Subtype 0
func ReadFormat29Subtype0(data []byte) (*Format29Subtype0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format29.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat29Subtype0", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.Subtype0 {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat29Subtype0", subType.ToString())
	}

	return &Format29Subtype0{
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
