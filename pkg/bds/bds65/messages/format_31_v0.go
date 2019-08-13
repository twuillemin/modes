package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
)

// Format31V0 is a message at the format BDS 6,5 the ADSB V1
type Format31V0 struct {
	Subtype                                    fields.SubtypeV0
	EnRouteOperationalCapabilities             fields.EnRouteOperationalCapabilities
	EnRouteOperationalCapabilityStatus         fields.EnRouteOperationalCapabilityStatus
	TerminalAreaOperationalCapabilities        fields.TerminalAreaOperationalCapabilities
	TerminalAreaOperationalCapabilityStatus    fields.TerminalAreaOperationalCapabilityStatus
	ApproachLandingOperationalCapabilities     fields.ApproachLandingOperationalCapabilities
	ApproachLandingOperationalCapabilityStatus fields.ApproachLandingOperationalCapabilityStatus
	SurfaceOperationalCapabilities             fields.SurfaceOperationalCapabilities
	SurfaceOperationalCapabilityStatus         fields.SurfaceOperationalCapabilityStatus
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format31V0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31V0
}

// GetRegister returns the register of the message
func (message *Format31V0) GetRegister() bds.Register {
	return adsb.Format31V0.GetRegister()
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V0) ToString() string {
	return fmt.Sprintf("Message:                                        %v\n"+
		"Subtype:                                        %v\n"+
		"En Route Operational Capabilities:              %v\n"+
		"En Route Operational Capability Status:         %v\n"+
		"Terminal Area Operational Capabilities:         %v\n"+
		"Terminal Area Operational Capability Status:    %v\n"+
		"Approach/Landing Operational Capabilities:      %v\n"+
		"Approach/Landing Operational Capability Status: %v\n"+
		"Surface Operational Capabilities:               %v\n"+
		"Surface Operational Capability Status:          %v",
		adsb.Format31V0.ToString(),
		message.GetSubtype().ToString(),
		message.EnRouteOperationalCapabilities.ToString(),
		message.EnRouteOperationalCapabilityStatus.ToString(),
		message.TerminalAreaOperationalCapabilities.ToString(),
		message.TerminalAreaOperationalCapabilityStatus.ToString(),
		message.ApproachLandingOperationalCapabilities.ToString(),
		message.ApproachLandingOperationalCapabilityStatus.ToString(),
		message.SurfaceOperationalCapabilities.ToString(),
		message.SurfaceOperationalCapabilityStatus.ToString())
}

// ReadFormat31V0 reads a message at the format BDS 6,5
func ReadFormat31V0(data []byte) (*Format31V0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format31V0", formatTypeCode)
	}

	// Check the ADSB Level
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion0 {
		return nil, fmt.Errorf("the data are given at at %v format that can not be read by ReadFormat31V0", detectedADSBLevel)
	}

	return &Format31V0{
		Subtype:                                    fields.ReadSubtypeV0(data),
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
