package bds62

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"github.com/twuillemin/modes/pkg/common"
)

// TargetStateAndStatus1 is a message at the format BDS 6,2
type TargetStateAndStatus1 struct {
	FormatTypeCode                       byte
	Subtype                              fields.Subtype
	SourceIntegrityLevelSupplement       fields.SourceIntegrityLevelSupplement
	SelectedAltitudeType                 fields.SelectedAltitudeType
	SelectedAltitude                     fields.SelectedAltitude
	BarometricPressureSetting            fields.BarometricPressureSetting
	SelectedHeadingStatus                fields.SelectedHeadingStatus
	SelectedHeadingSign                  fields.SelectedHeadingSign
	SelectedHeading                      fields.SelectedHeading
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	NICBaro                              fields.NICBaro
	SourceIntegrityLevel                 fields.SourceIntegrityLevel
	StatusMCPFCUBits                     fields.StatusMCPFCUBits
	AutopilotEngaged                     fields.AutopilotEngaged
	VNAVModeEngaged                      fields.VNAVModeEngaged
	AltitudeHoldModeEngaged              fields.AltitudeHoldModeEngaged
	ApproachModeEngaged                  fields.ApproachModeEngaged
	ACASOperational                      fields.ACASOperational
	LNAVModeEngaged                      fields.LNAVModeEngaged
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
		"Selected Altitude :                            %v\n"+
		"Barometric Pressure Setting (minus 800 mbars): %v\n"+
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
		message.SelectedAltitude.ToString(),
		message.BarometricPressureSetting.ToString(),
		message.SelectedHeading.ToString(message.SelectedHeadingStatus, message.SelectedHeadingSign),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.NICBaro.ToString(),
		message.SourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.StatusMCPFCUBits.ToString(),
		printStatusBit(message.StatusMCPFCUBits, message.AutopilotEngaged),
		printStatusBit(message.StatusMCPFCUBits, message.VNAVModeEngaged),
		printStatusBit(message.StatusMCPFCUBits, message.AltitudeHoldModeEngaged),
		printStatusBit(message.StatusMCPFCUBits, message.ApproachModeEngaged),
		printStatusBit(message.StatusMCPFCUBits, message.LNAVModeEngaged),
		message.ACASOperational.ToString())
}

func printStatusBit(status fields.StatusMCPFCUBits, bit common.Printable) string {
	if status == fields.SMFBNoInformationProvided {
		return "No information provided from MCP/FCU"
	}
	return bit.ToString()
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

	return &TargetStateAndStatus1{
		FormatTypeCode:                       formatTypeCode,
		Subtype:                              subType,
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
		SelectedAltitudeType:                 fields.ReadSelectedAltitudeType(data),
		SelectedAltitude:                     fields.ReadSelectedAltitude(data),
		BarometricPressureSetting:            fields.ReadBarometricPressureSetting(data),
		SelectedHeadingStatus:                fields.ReadSelectedHeadingStatus(data),
		SelectedHeadingSign:                  fields.ReadSelectedHeadingSign(data),
		SelectedHeading:                      fields.ReadSelectedHeading(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		NICBaro:                              fields.ReadNICBaro(data),
		SourceIntegrityLevel:                 fields.ReadSourceIntegrityLevel(data),
		StatusMCPFCUBits:                     fields.ReadStatusMCPFCUBits(data),
		AutopilotEngaged:                     fields.ReadAutopilotEngaged(data),
		VNAVModeEngaged:                      fields.ReadVNAVModeEngaged(data),
		AltitudeHoldModeEngaged:              fields.ReadAltitudeHoldModeEngaged(data),
		ApproachModeEngaged:                  fields.ReadApproachModeEngaged(data),
		ACASOperational:                      fields.ReadACASOperational(data),
		LNAVModeEngaged:                      fields.ReadLNAVModeEngaged(data),
	}, nil
}
