package messages

import "github.com/twuillemin/modes/pkg/common"

// ACASMessage is the basic interface that ACAS messages are expected to implement
type ACASMessage interface {
	common.Printable

	// GetName returns the name of the message
	GetName() string
	// GetVDS1 returns the VDS1
	GetVDS1() byte
	// GetVDS2 returns the VDS1
	GetVDS2() byte
}
