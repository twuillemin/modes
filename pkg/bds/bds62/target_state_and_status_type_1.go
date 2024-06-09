package bds62

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// TargetStateAndStatus1 is a message at the format BDS 6,2
//
// Specified in Doc 9871 / Figure C-9
type TargetStateAndStatus1 struct {
	FormatTypeCode                       byte
	Subtype                              fields.Subtype
	SourceIntegrityLevelSupplement       fields.SourceIntegrityLevelSupplement
	SelectedAltitudeType                 fields.SelectedAltitudeType
	SelectedAltitudeStatus               fields.NumericValueStatus
	SelectedAltitude                     uint16
	BarometricPressureSettingStatus      fields.NumericValueStatus
	BarometricPressureSetting            float32
	SelectedHeadingStatus                fields.SelectedHeadingStatus
	SelectedHeading                      float32
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	NICBaro                              fields.NICBaro
	SourceIntegrityLevel                 fields.SourceIntegrityLevel
	StatusMCPFCUBits                     fields.StatusMCPFCUBits
	AutopilotEngaged                     bool
	VNAVModeEngaged                      bool
	AltitudeHoldModeEngaged              bool
	ApproachModeEngaged                  bool
	TCASACASOperational                  bool
	LNAVModeEngaged                      bool
}

func (message TargetStateAndStatus1) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message TargetStateAndStatus1) GetRegister() register.Register {
	return register.BDS62
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message TargetStateAndStatus1) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message TargetStateAndStatus1) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                       %v\n"+
		"Subtype:                                       %v\n"+
		"Selected Altitude Type:                        %v\n"+
		"Selected Altitude Status:                      %v\n"+
		"Selected Altitude:                             %v\n"+
		"Barometric Pressure Setting Status:            %v\n"+
		"Barometric Pressure Setting:                   %v\n"+
		"Selected Heading Status:                       %v\n"+
		"Selected Heading :                             %v\n"+
		"Navigation Accuracy Category - Position:       %v\n"+
		"Navigation Integrity Category - Baro:          %v\n"+
		"Source Integrity ReaderLevel:                  %v\n"+
		"Source Integrity ReaderLevel Supplement:       %v\n"+
		"Status of MCP/FPU Mode Bits:                   %v\n"+
		"Autopilot Engaged:                             %v\n"+
		"VNAV Mode Engaged:                             %v\n"+
		"Altitude Hold Mode Engaged:                    %v\n"+
		"Approach Mode Engaged:                         %v\n"+
		"LNAV Mode Engaged:                             %v\n"+
		"TCAS / ACAS Operational :                      %v",
		message.GetRegister().ToString(),
		message.GetSubtype().ToString(),
		message.SelectedAltitudeType.ToString(),
		message.SelectedAltitudeStatus.ToString(),
		message.SelectedAltitude,
		message.BarometricPressureSettingStatus.ToString(),
		message.BarometricPressureSetting,
		message.SelectedHeadingStatus.ToString(),
		message.SelectedHeading,
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.NICBaro.ToString(),
		message.SourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.StatusMCPFCUBits.ToString(),
		message.AutopilotEngaged,
		message.VNAVModeEngaged,
		message.AltitudeHoldModeEngaged,
		message.ApproachModeEngaged,
		message.LNAVModeEngaged,
		message.TCASACASOperational)
}

// ReadTargetStateAndStatus1 reads a TargetStateAndStatus / Subtype 1
func ReadTargetStateAndStatus1(data []byte) (*TargetStateAndStatus1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 29 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 29, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.Subtype1 {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadTargetStateAndStatus1", subType.ToString())
	}

	selectedAltitude, selectedAltitudeStatus := fields.ReadSelectedAltitude(data)
	barometricPressureSetting, barometricPressureSettingStatus := fields.ReadBarometricPressureSetting(data)
	selectedHeading, selectedHeadingStatus := fields.ReadSelectedHeading(data)

	return &TargetStateAndStatus1{
		FormatTypeCode:                       formatTypeCode,
		Subtype:                              subType,
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
		SelectedAltitudeType:                 fields.ReadSelectedAltitudeType(data),
		SelectedAltitudeStatus:               selectedAltitudeStatus,
		SelectedAltitude:                     selectedAltitude,
		BarometricPressureSettingStatus:      barometricPressureSettingStatus,
		BarometricPressureSetting:            barometricPressureSetting,
		SelectedHeadingStatus:                selectedHeadingStatus,
		SelectedHeading:                      selectedHeading,
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		NICBaro:                              fields.ReadNICBaro(data),
		SourceIntegrityLevel:                 fields.ReadSourceIntegrityLevel(data),
		StatusMCPFCUBits:                     fields.ReadStatusMCPFCUBits(data),
		AutopilotEngaged:                     (data[5] & 0x01) != 0,
		VNAVModeEngaged:                      (data[6] & 0x80) != 0,
		AltitudeHoldModeEngaged:              (data[6] & 0x40) != 0,
		ApproachModeEngaged:                  (data[6] & 0x10) != 0,
		TCASACASOperational:                  (data[6] & 0x08) != 0,
		LNAVModeEngaged:                      (data[6] & 0x04) != 0,
	}, nil
}
