package messages

import "github.com/twuillemin/modes/pkg/common"

// ModeSMessage is the basic interface that Mode S messages are expected to implement
type ModeSMessage interface {
	common.Printable

	// GetName returns the name of the message
	GetName() string
	// GetDownLinkFormat returns the downlink format of the message
	GetDownLinkFormat() int
}
