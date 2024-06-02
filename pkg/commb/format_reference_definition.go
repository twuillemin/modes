package commb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/common"
)

// -------------------------------------------------------------------------
//                         INTERFACE DEFINITION
// -------------------------------------------------------------------------

// MessageFormat is the definition of a format for a Comm-B message
type MessageFormat interface {
	common.Printable

	// GetRegister returns the BDS implementing the message
	GetRegister() bds.Register
}

// -------------------------------------------------------------------------
//                         INTERNAL STRUCTURE
// -------------------------------------------------------------------------

// The basic structure for keeping information about known commb messages
type commbFormatReferenceDefinition struct {
	register bds.Register
}

// GetRegister returns the BDS implementing the message
func (commb commbFormatReferenceDefinition) GetRegister() bds.Register {
	return commb.register
}

// ToString returns a basic, but readable, representation of the message
func (commb commbFormatReferenceDefinition) ToString() string {
	return fmt.Sprintf("%v (%v)",
		commb.register.GetDescription(),
		commb.register.GetId())
}

// -------------------------------------------------------------------------
//                             THE FORMATS
// -------------------------------------------------------------------------

// FormatDataLinkCapabilityReport is the definition of a message DataLinkCapabilityReport
var FormatDataLinkCapabilityReport = commbFormatReferenceDefinition{
	register: bds.BDS10,
}

// FormatAircraftIdentification is the definition of a message AircraftIdentification
var FormatAircraftIdentification = commbFormatReferenceDefinition{
	register: bds.BDS20,
}
