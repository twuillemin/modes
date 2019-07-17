package messages

import "github.com/twuillemin/modes/pkg/common"

// ADSBMessage is the basic interface that ADSB messages are expected to implement
type ADSBMessage interface {
	common.Printable

	// GetName returns the name of the message
	GetName() string
	// GetBDS returns the binary data format
	GetBDS() string
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
}
