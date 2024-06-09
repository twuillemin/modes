package bds62

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// TargetStateAndStatus0 is a message at the format BDS 6,2
//
// Specified in Doc 9871 / Table B-2-98
type TargetStateAndStatus0 struct {
	FormatTypeCode                         byte
	Subtype                                fields.Subtype
	VerticalDataAvailableSourceIndicator   fields.VerticalDataAvailableSourceIndicator
	TargetAltitudeType                     fields.TargetAltitudeType
	TargetAltitudeCapability               fields.TargetAltitudeCapability
	VerticalModeIndicator                  fields.VerticalModeIndicator
	TargetAltitudeStatus                   fields.NumericValueStatus
	TargetAltitude                         int32
	HorizontalDataAvailableSourceIndicator fields.HorizontalDataAvailableSourceIndicator
	TargetHeadingTrackAngleStatus          fields.NumericValueStatus
	TargetHeadingTrackAngle                uint16
	TargetHeadingTrackIndicator            fields.TargetHeadingTrackIndicator
	HorizontalModeIndicator                fields.HorizontalModeIndicator
	NavigationalAccuracyCategoryPosition   fields.NavigationalAccuracyCategoryPositionV1
	NICBaro                                fields.NICBaro
	SurveillanceIntegrityLevel             fields.SurveillanceIntegrityLevel
	CapabilityModeCode                     fields.CapabilityModeCode
	EmergencyPriorityStatus                fields.EmergencyPriorityStatus
}

func (message TargetStateAndStatus0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message TargetStateAndStatus0) GetRegister() register.Register {
	return register.BDS62
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message TargetStateAndStatus0) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message TargetStateAndStatus0) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                      %v\n"+
		"Subtype:                                      %v\n"+
		"Vertical Data Available / Source Indicator:   %v\n"+
		"Target Altitude Type:                         %v\n"+
		"Target Altitude Capability:                   %v\n"+
		"Vertical Mode Indicator:                      %v\n"+
		"Target Altitude Status:                       %v\n"+
		"Target Altitude:                              %v\n"+
		"Horizontal Data Available / Source Indicator: %v\n"+
		"Target Heading / Track Angle Status:          %v\n"+
		"Target Heading / Track Angle:                 %v\n"+
		"Target Heading / Track Indicator:             %v\n"+
		"Horizontal Mode Indicator:                    %v\n"+
		"Navigation Accuracy Category - Position:      %v\n"+
		"Navigation Integrity Category - Baro:         %v\n"+
		"Surveillance Integrity ReaderLevel:           %v\n"+
		"Capability / Mode Codes:                      %v\n"+
		"Emergency / Priority Status:                  %v",
		message.GetRegister().ToString(),
		message.GetSubtype().ToString(),
		message.VerticalDataAvailableSourceIndicator.ToString(),
		message.TargetAltitudeType.ToString(),
		message.TargetAltitudeCapability.ToString(),
		message.VerticalModeIndicator.ToString(),
		message.TargetAltitudeStatus.ToString(),
		message.TargetAltitude,
		message.HorizontalDataAvailableSourceIndicator.ToString(),
		message.TargetHeadingTrackAngleStatus.ToString(),
		message.TargetHeadingTrackAngle,
		message.TargetHeadingTrackIndicator.ToString(),
		message.HorizontalModeIndicator.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.NICBaro.ToString(),
		message.SurveillanceIntegrityLevel.ToString(),
		message.CapabilityModeCode.ToString(),
		message.EmergencyPriorityStatus.ToString())
}

// ReadTargetStateAndStatus0 reads a TargetStateAndStatus / subtype 0
func ReadTargetStateAndStatus0(data []byte) (*TargetStateAndStatus0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 29 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 29, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.Subtype0 {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadTargetStateAndStatus0", subType.ToString())
	}

	if (data[5] & 0x03) != 0 {
		return nil, errors.New("the bits 47 to 48 are expected to be 0")
	}

	if (data[6] & 0xE0) != 0 {
		return nil, errors.New("the bits 49 to 51 are expected to be 0")
	}

	targetAltitude, targetAltitudeStatus := fields.ReadTargetAltitude(data)
	targetHeading, targetHeadingStatus := fields.ReadTargetHeadingTrackAngle(data)

	return &TargetStateAndStatus0{
		FormatTypeCode:                         formatTypeCode,
		Subtype:                                subType,
		VerticalDataAvailableSourceIndicator:   fields.ReadVerticalDataAvailableSourceIndicator(data),
		TargetAltitudeType:                     fields.ReadTargetAltitudeType(data),
		TargetAltitudeCapability:               fields.ReadTargetAltitudeCapability(data),
		VerticalModeIndicator:                  fields.ReadVerticalModeIndicator(data),
		TargetAltitudeStatus:                   targetAltitudeStatus,
		TargetAltitude:                         targetAltitude,
		HorizontalDataAvailableSourceIndicator: fields.ReadHorizontalDataAvailableSourceIndicator(data),
		TargetHeadingTrackAngleStatus:          targetHeadingStatus,
		TargetHeadingTrackAngle:                targetHeading,
		TargetHeadingTrackIndicator:            fields.ReadTargetHeadingTrackIndicator(data),
		HorizontalModeIndicator:                fields.ReadHorizontalModeIndicator(data),
		NavigationalAccuracyCategoryPosition:   fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		NICBaro:                                fields.ReadNICBaro(data),
		SurveillanceIntegrityLevel:             fields.ReadSurveillanceIntegrityLevel(data),
		CapabilityModeCode:                     fields.ReadCapabilityModeCode(data),
		EmergencyPriorityStatus:                fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
