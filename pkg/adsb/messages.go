package adsb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/common"
	"strings"
)

// MessageLevel is used to represent a message minimum/maximum ADSB level.
type MessageLevel byte

const (
	// MessageLevel0 indicates that the message could be level ADSB level 0 or more
	MessageLevel0 MessageLevel = 0
	// MessageLevel1 indicates that the message must be read as ADSB 0 or has been determined as being level 0 only
	MessageLevel1 MessageLevel = 1
	// MessageLevel2 indicates that the message must be read as ADSB 2 or has been determined as being level 1 only
	MessageLevel2 MessageLevel = 2
)

// ToString returns a basic, but readable, representation of the message
func (level MessageLevel) ToString() string {
	switch level {
	case MessageLevel0:
		return "ADSB V0"
	case MessageLevel1:
		return "ADSB V1"
	case MessageLevel2:
		return "ADSB V2"
	default:
		return fmt.Sprintf("there is no ADSB ReaderLevel associated with value %v", level)
	}
}

// Subtype is the subtype of a Format of a message. Some message don't have subtype and so should use nil.
type Subtype interface {
	common.Printable

	ToSubtype() Subtype
}

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	bds.Message

	// GetMessageFormat returns the ADSB format of the message
	GetMessageFormat() MessageFormat

	// GetSubtype returns the subtype of the message if any. If the message does not have a subtype, should be null.
	GetSubtype() Subtype

	// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
	GetMinimumADSBLevel() MessageLevel

	// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
	GetMaximumADSBLevel() MessageLevel
}

// GetMessageFormatInformation generates a string presenting the common information to all ADSB messages, which is
// only the format, the register, the subtype, and the ADSB level supported.
func GetMessageFormatInformation(message Message) string {

	levels := make([]string, 0, 3)

	if message.GetMinimumADSBLevel() == MessageLevel0 {
		levels = append(levels, MessageLevel0.ToString())
	}

	if (message.GetMinimumADSBLevel() == MessageLevel1) ||
		(message.GetMinimumADSBLevel() == MessageLevel0 && message.GetMaximumADSBLevel() != MessageLevel0) {
		levels = append(levels, MessageLevel1.ToString())
	}

	if message.GetMaximumADSBLevel() == MessageLevel2 {
		levels = append(levels, MessageLevel2.ToString())
	}

	return fmt.Sprintf("%v [%v]",
		message.GetMessageFormat().ToString(),
		strings.Join(levels, ", "))
}
