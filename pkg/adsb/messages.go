package adsb

import (
	"github.com/twuillemin/modes/pkg/common"
)

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	common.Printable

	// GetADSBVersion returns the ADSB level used to read the data
	GetADSBVersion() ADSBVersion

	// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
	CheckCoherency() error
}
