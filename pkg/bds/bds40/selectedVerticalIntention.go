package bds40

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds40/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// SelectedVerticalIntention is a message at the format BDS 4,0
//
// Specified in Doc 9871 / Table A-2-48
type SelectedVerticalIntention struct {
	MCPFCUSelectedAltitudeAvailable    bool
	MCPFCUSelectedAltitude             uint32
	FMSSelectedAltitudeAvailable       bool
	FMSSelectedAltitude                uint32
	BarometricPressureSettingAvailable bool
	BarometricPressureSetting          float32
	StatusOfModeBits                   fields.StatusOfModeBits
	VNAVMode                           bool
	AltitudeHoldMode                   bool
	ApproachMode                       bool
	StatusOfTargetSource               fields.StatusOfTargetSource
	TargetSource                       fields.TargetSource
}

// GetRegister returns the Register the message
func (message SelectedVerticalIntention) GetRegister() register.Register {
	return register.BDS40
}

// ToString returns a basic, but readable, representation of the message
func (message SelectedVerticalIntention) ToString() string {
	return fmt.Sprintf(""+
		"Message:                               %v\n"+
		"MCP/FCU Selected Altitude Available:   %v\n"+
		"MCP/FCU Selected Altitude:             %v\n"+
		"FMS Selected Altitude Available:       %v\n"+
		"FMS Selected Altitude:                 %v\n"+
		"Barometric Pressure Setting Available: %v\n"+
		"Barometric Pressure Setting:           %v\n"+
		"Status Of Mode Bits:                   %v\n"+
		"VNAV Mode:                             %v\n"+
		"Altitude Hold Mode:                    %v\n"+
		"Approach Mode:                         %v\n"+
		"Status Of Target Source:               %v\n"+
		"Target Source:                         %v",
		message.GetRegister().ToString(),
		message.MCPFCUSelectedAltitudeAvailable,
		message.MCPFCUSelectedAltitude,
		message.FMSSelectedAltitudeAvailable,
		message.FMSSelectedAltitude,
		message.BarometricPressureSettingAvailable,
		message.BarometricPressureSetting,
		message.StatusOfModeBits.ToString(),
		message.VNAVMode,
		message.AltitudeHoldMode,
		message.ApproachMode,
		message.StatusOfTargetSource.ToString(),
		message.TargetSource.ToString())
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message SelectedVerticalIntention) CheckCoherency() error {
	return nil
}

// ReadSelectedVerticalIntention reads a message as a SelectedVerticalIntention
func ReadSelectedVerticalIntention(data []byte) (*SelectedVerticalIntention, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B SelectedVerticalIntention message must be 7 bytes long")
	}

	if data[4]&0x01 != 0 {
		return nil, errors.New("the bit 40 must be zero")
	}

	if data[5]&0xFE != 0 {
		return nil, errors.New("the bits 41 to 47 must be zero")
	}

	if data[6]&0x18 != 0 {
		return nil, errors.New("the bits 52 to 53 must be zero")
	}

	MCPFCUSelectedAltitudeAvailable, MCPFCUSelectedAltitude := fields.ReadMCPSelectedAltitude(data)
	FMSSelectedAltitudeAvailable, FMSSelectedAltitude := fields.ReadFMSSelectedAltitude(data)
	barometricPressureSettingAvailable, barometricPressureSetting := fields.ReadBarometricPressure(data)

	return &SelectedVerticalIntention{
		MCPFCUSelectedAltitudeAvailable:    MCPFCUSelectedAltitudeAvailable,
		MCPFCUSelectedAltitude:             MCPFCUSelectedAltitude,
		FMSSelectedAltitudeAvailable:       FMSSelectedAltitudeAvailable,
		FMSSelectedAltitude:                FMSSelectedAltitude,
		BarometricPressureSettingAvailable: barometricPressureSettingAvailable,
		BarometricPressureSetting:          barometricPressureSetting,
		StatusOfModeBits:                   fields.ReadStatusOfModeBits(data),
		VNAVMode:                           (data[6]&0x80)>>7 == 0x01,
		AltitudeHoldMode:                   (data[6]&0x40)>>6 == 0x01,
		ApproachMode:                       (data[6]&0x20)>>5 == 0x01,
		StatusOfTargetSource:               fields.ReadStatusOfTargetSource(data),
		TargetSource:                       fields.ReadTargetSource(data),
	}, nil
}
