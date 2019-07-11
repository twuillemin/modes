package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 CrossLinkCompatibility (CC)
//
// -----------------------------------------------------------------------------------------

// CrossLinkCompatibility (CC) downlink field shall indicate the ability of the transponder to support the cross-link
// capability, i.e. decode the contents of the DS field in an interrogation with UF equals 0 and respond with the
// contents of the specified GICB register in the corresponding reply with DF equals 16.
//
// Defined at 3.1.2.8.2.3
type CrossLinkCompatibility int

const (
	// CrossLinkCompatibilityNotSupported signifies that the transponder cannot support the cross-link capability.
	CrossLinkCompatibilityNotSupported CrossLinkCompatibility = 0
	// CrossLinkCompatibilitySupported signifies that the transponder supports the cross-link capability.
	CrossLinkCompatibilitySupported CrossLinkCompatibility = 1
)

// readCrossLinkCompatibility reads the CC field from a message
func ReadCrossLinkCompatibility(message common.MessageData) CrossLinkCompatibility {
	if message.FirstField&0x02 != 0 {
		return CrossLinkCompatibilitySupported
	}
	return CrossLinkCompatibilityNotSupported
}

func (crossLinkCompatibility CrossLinkCompatibility) PrettyPrint() string {
	switch crossLinkCompatibility {
	case CrossLinkCompatibilityNotSupported:
		return "0 - Not Supported"
	case CrossLinkCompatibilitySupported:
		return "1 - Supported"
	default:
		return fmt.Sprintf("%v - Unknown code", crossLinkCompatibility)
	}
}

func (crossLinkCompatibility CrossLinkCompatibility) ExtendedPrettyPrint() string {
	switch crossLinkCompatibility {
	case CrossLinkCompatibilityNotSupported:
		return "0 - Not Supported: the transponder cannot support the cross-link capability"
	case CrossLinkCompatibilitySupported:
		return "1 - Supported: the transponder supports the cross-link capability"
	default:
		return fmt.Sprintf("%v - Unknown code", crossLinkCompatibility)
	}
}
