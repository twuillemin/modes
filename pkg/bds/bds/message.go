package bds

import "github.com/twuillemin/modes/pkg/common"

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	common.Printable

	// GetRegister returns the register (Binary Data Store) of the message
	GetRegister() Register
}
