package messages

import (
	"fmt"
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

// GetName returns the name of the message
func (message *Format31V0) GetName() string {
	return bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V0) GetBDS() string {
	return bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V0) GetFormatTypeCode() byte {
	return 31
}

// GetSubtype returns the subtype of the Operational Status Sub Type
func (message *Format31V0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V0) ToString() string {
	return fmt.Sprintf("Message:                                        %v - %v (%v)\n"+
		"Subtype:                                        %v\n"+
		"En Route Operational Capabilities:              %v\n"+
		"En Route Operational Capability Status:         %v\n"+
		"Terminal Area Operational Capabilities:         %v\n"+
		"Terminal Area Operational Capability Status:    %v\n"+
		"Approach/Landing Operational Capabilities:      %v\n"+
		"Approach/Landing Operational Capability Status: %v\n"+
		"Surface Operational Capabilities:               %v\n"+
		"Surface Operational Capability Status:          %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
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
