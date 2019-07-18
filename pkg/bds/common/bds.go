package common

import "github.com/twuillemin/modes/pkg/common"

// BDSMessage is the basic interface that ADSB messages are expected to implement
type BDSMessage interface {
	common.Printable
	// GetName returns the name of the message
	GetName() string
	// GetBDS returns the binary data format
	GetBDS() string
}
