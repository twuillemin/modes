package commb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
)

// -------------------------------------------------------------------------
//                         DEFINITION
// -------------------------------------------------------------------------

// MessageFormat is the definition of a format for a Comm-B message
type MessageFormat struct {
	register bds.Register
}

// GetRegister returns the BDS implementing the message
func (commb MessageFormat) GetRegister() bds.Register {
	return commb.register
}

// ToString returns a basic, but readable, representation of the message
func (commb MessageFormat) ToString() string {
	return fmt.Sprintf("%v (%v)",
		commb.register.GetDescription(),
		commb.register.GetId())
}

// -------------------------------------------------------------------------
//                             THE FORMATS
// -------------------------------------------------------------------------

// FormatDataLinkCapabilityReport is the definition of a message DataLinkCapabilityReport
var FormatDataLinkCapabilityReport = MessageFormat{
	register: bds.BDS10,
}

// FormatAircraftIdentification is the definition of a message AircraftIdentification
var FormatAircraftIdentification = MessageFormat{
	register: bds.BDS20,
}
