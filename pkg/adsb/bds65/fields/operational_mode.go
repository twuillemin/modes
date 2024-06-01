package fields

import "fmt"

// OperationalMode is the Operational Mode definition
//
// Specified in Doc 9871 / B.2.3.10.4
type OperationalMode struct {
	ACASResolutionAdvisoryActive ACASResolutionAdvisoryActiveV1
	IdentSwitchActive            IdentSwitchActive
	ReceivingATCServices         ReceivingATCServices
}

// ToString returns a basic, but readable, representation of the field
func (mode OperationalMode) ToString() string {

	return fmt.Sprintf("ACAS Resolution Advisory Active: %v\n"+
		"Ident Switch Active:             %v\n"+
		"Receiving ATC Services:          %v",
		mode.ACASResolutionAdvisoryActive.ToString(),
		mode.IdentSwitchActive.ToString(),
		mode.ReceivingATCServices.ToString())
}

// ReadOperationalMode reads the OperationalMode from a 56 bits data field
func ReadOperationalMode(data []byte) OperationalMode {
	return OperationalMode{
		ACASResolutionAdvisoryActive: ReadACASResolutionAdvisoryActiveV1(data),
		IdentSwitchActive:            ReadIdentSwitchActive(data),
		ReceivingATCServices:         ReadReceivingATCServices(data),
	}
}
