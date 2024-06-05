package bds65

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

type AircraftOperationalStatusADSB0 struct {
	FormatTypeCode                             byte
	Subtype                                    byte
	EnRouteOperationalCapabilities             fields.EnRouteOperationalCapabilities
	EnRouteOperationalCapabilityStatus         fields.EnRouteOperationalCapabilityStatus
	TerminalAreaOperationalCapabilities        fields.TerminalAreaOperationalCapabilities
	TerminalAreaOperationalCapabilityStatus    fields.TerminalAreaOperationalCapabilityStatus
	ApproachLandingOperationalCapabilities     fields.ApproachLandingOperationalCapabilities
	ApproachLandingOperationalCapabilityStatus fields.ApproachLandingOperationalCapabilityStatus
	SurfaceOperationalCapabilities             fields.SurfaceOperationalCapabilities
	SurfaceOperationalCapabilityStatus         fields.SurfaceOperationalCapabilityStatus
}

func (message AircraftOperationalStatusADSB0) GetADSBLevel() byte {
	return 0
}

// GetRegister returns the Register the message
func (message AircraftOperationalStatusADSB0) GetRegister() register.Register {
	return register.BDS65
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftOperationalStatusADSB0) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftOperationalStatusADSB0) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                        %v\n"+
		"En Route Operational Capabilities:              %v\n"+
		"En Route Operational Capability Status:         %v\n"+
		"Terminal Area Operational Capabilities:         %v\n"+
		"Terminal Area Operational Capability Status:    %v\n"+
		"Approach/Landing Operational Capabilities:      %v\n"+
		"Approach/Landing Operational Capability Status: %v\n"+
		"SubtypeSurface Operational Capabilities:               %v\n"+
		"SubtypeSurface Operational Capability Status:          %v",
		message.GetRegister().ToString(),
		message.EnRouteOperationalCapabilities.ToString(),
		message.EnRouteOperationalCapabilityStatus.ToString(),
		message.TerminalAreaOperationalCapabilities.ToString(),
		message.TerminalAreaOperationalCapabilityStatus.ToString(),
		message.ApproachLandingOperationalCapabilities.ToString(),
		message.ApproachLandingOperationalCapabilityStatus.ToString(),
		message.SurfaceOperationalCapabilities.ToString(),
		message.SurfaceOperationalCapabilityStatus.ToString())
}

// ReadAircraftOperationalStatusADSB0 reads a message at the format Format31 for ADSB V0
func ReadAircraftOperationalStatusADSB0(data []byte) (*AircraftOperationalStatusADSB0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 31 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 29, got %v", formatTypeCode)
	}

	subType := data[0] & 0x07
	if subType != 0 {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftOperationalStatusADSB0", subType)
	}

	if data[5] != 0 || data[6] != 0 {
		return nil, errors.New("the bits 41 to 56 are expected to be 0")
	}

	return &AircraftOperationalStatusADSB0{
		FormatTypeCode:                             formatTypeCode,
		Subtype:                                    subType,
		EnRouteOperationalCapabilities:             fields.ReadEnRouteOperationalCapabilities(data),
		EnRouteOperationalCapabilityStatus:         fields.ReadEnRouteOperationalCapabilityStatus(data),
		TerminalAreaOperationalCapabilities:        fields.ReadTerminalAreaOperationalCapabilities(data),
		TerminalAreaOperationalCapabilityStatus:    fields.ReadTerminalAreaOperationalCapabilityStatus(data),
		ApproachLandingOperationalCapabilities:     fields.ReadApproachLandingOperationalCapabilities(data),
		ApproachLandingOperationalCapabilityStatus: fields.ReadApproachLandingOperationalCapabilityStatus(data),
		SurfaceOperationalCapabilities:             fields.ReadSurfaceOperationalCapabilities(data),
		SurfaceOperationalCapabilityStatus:         fields.ReadSurfaceOperationalCapabilityStatus(data),
	}, nil
}
