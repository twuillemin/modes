package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds65/fields"
)

// Format31Reserved is a message at the format BDS 6,5 the ADSB V1
type Format31Reserved struct {
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
func (message Format31Reserved) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31
}

// GetSubtype returns the Subtype
func (message Format31Reserved) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format31Reserved) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format31Reserved) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// ToString returns a basic, but readable, representation of the message
func (message Format31Reserved) ToString() string {
	return fmt.Sprintf("Message:                                        %v\n"+
		"En Route Operational Capabilities:              %v\n"+
		"En Route Operational Capability Status:         %v\n"+
		"Terminal Area Operational Capabilities:         %v\n"+
		"Terminal Area Operational Capability Status:    %v\n"+
		"Approach/Landing Operational Capabilities:      %v\n"+
		"Approach/Landing Operational Capability Status: %v\n"+
		"SubtypeSurface Operational Capabilities:               %v\n"+
		"SubtypeSurface Operational Capability Status:          %v",
		adsb.GetMessageFormatInformation(&message),
		message.EnRouteOperationalCapabilities.ToString(),
		message.EnRouteOperationalCapabilityStatus.ToString(),
		message.TerminalAreaOperationalCapabilities.ToString(),
		message.TerminalAreaOperationalCapabilityStatus.ToString(),
		message.ApproachLandingOperationalCapabilities.ToString(),
		message.ApproachLandingOperationalCapabilityStatus.ToString(),
		message.SurfaceOperationalCapabilities.ToString(),
		message.SurfaceOperationalCapabilityStatus.ToString())
}

// ReadFormat31Reserved reads a message at the format Format31 for ADSB V0
func ReadFormat31Reserved(data []byte) (*Format31Reserved, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat31Reserved", formatTypeCode)
	}

	// Check the ADSB ReaderLevel
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion0 {
		return nil, fmt.Errorf("the data are given at ADSB ReaderLevel %v format that can not be read by ReadFormat31Reserved", detectedADSBLevel)
	}

	return &Format31Reserved{
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
