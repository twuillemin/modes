package commb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/common"
)

// Message is the basic interface that ADSB messages are expected to implement
type Message interface {
	common.Printable

	// GetMessageFormat returns the Comm-B format of the message
	GetMessageFormat() MessageFormat
}

// GetMessageFormatInformation generates a string presenting the common information to all ADSB messages, which is
// only the format, the register, the subtype, and the ADSB level supported.
func GetMessageFormatInformation(message Message) string {

	return fmt.Sprintf("%v", message.GetMessageFormat().ToString())
}
