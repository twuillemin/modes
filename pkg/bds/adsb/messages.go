package adsb

import (
	"github.com/twuillemin/modes/pkg/bds/bds"
)

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	bds.Message

	// GetMessageFormat returns the ADSB format of the message
	GetMessageFormat() MessageFormat
}
