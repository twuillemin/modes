package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05V0 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB V0 are expected to implement
type MessageBDS05V0 interface {
	MessageBDS05
	GetSingleAntennaFlag() fields.SingleAntennaFlag
	// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
	GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit
	// GetContainmentRadius returns the ContainmentRadius
	GetContainmentRadius() fields.ContainmentRadius
}

func bds05v0ToString(message MessageBDS05V0) string {
	return fmt.Sprintf("Message:                           %v - %v (%v)\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v\n"+
		"Single Antenna:                    %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Altitude:                          %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalProtectionLimit().ToString(),
		message.GetContainmentRadius().ToString(),
		message.GetSingleAntennaFlag().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}
