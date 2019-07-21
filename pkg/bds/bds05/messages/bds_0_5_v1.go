package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05V1 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB v1 are expected to implement
type MessageBDS05V1 interface {
	MessageBDS05
	GetSingleAntennaFlag() fields.SingleAntennaFlag
	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
}

func bds05v1ToString(message MessageBDS05V1) string {
	return fmt.Sprintf("Message:                           %v - %v (%v)\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Containment Radius      %v\n"+
		"Navigation Integrity Category      %v\n"+
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
		message.GetHorizontalContainmentRadius().ToString(),
		message.GetNavigationIntegrityCategory(),
		message.GetSingleAntennaFlag().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}
