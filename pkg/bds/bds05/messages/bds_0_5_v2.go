package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05V2 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB v2 are expected to implement
type MessageBDS05V2 interface {
	MessageBDS05
	// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
	GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB
	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
}

func bds05v2ToString(message MessageBDS05V2) string {
	return fmt.Sprintf("Message:                           %v (%v)\n"+
		"Format Type Code:                  %v\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Containment Radius      %v\n"+
		"Navigation Integrity Category      %v\n"+
		"NIC Supplement B:                  %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Altitude:                          %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalContainmentRadius().ToString(),
		message.GetNavigationIntegrityCategory(),
		message.GetNavigationIntegrityCodeSupplementB().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}
