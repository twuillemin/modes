package messages

import "github.com/twuillemin/modes/pkg/bds/common"

// ADSBMessage is the basic interface that ADSB messages are expected to implement
type ADSBMessage interface {
	common.BDSMessage

	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
}
