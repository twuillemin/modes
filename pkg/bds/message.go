package bds

import (
	"github.com/twuillemin/modes/pkg/bds/register"
	"github.com/twuillemin/modes/pkg/common"
)

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	common.Printable

	// GetRegister returns the Register the message
	GetRegister() register.Register

	// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
	CheckCoherency() error
}
