package fields

import "fmt"

// OperationalMode is the Operational Mode definition
//
// Specified in Doc 9871 / B.2.3.10.4
type OperationalMode struct {
	ACASResolutionAdvisoryActive ACASResolutionAdvisoryActive
	IdentSwitchActive            IdentSwitchActive
	ReceivingATCServices         ReceivingATCServices
}

// ToString returns a basic, but readable, representation of the field
func (mode OperationalMode) ToString() string {

	return fmt.Sprintf("ACASResolutionAdvisoryActive:  %v\n"+
		"IdentSwitchActive:             %v\n"+
		"ReceivingATCServices:          %v",
		mode.ACASResolutionAdvisoryActive.ToString(),
		mode.IdentSwitchActive.ToString(),
		mode.ReceivingATCServices.ToString())
}

// ReadOperationalMode reads the OperationalMode from a 56 bits data field
func ReadOperationalMode(data []byte) OperationalMode {
	return OperationalMode{
		ACASResolutionAdvisoryActive: ReadACASResolutionAdvisoryActive(data),
		IdentSwitchActive:            ReadIdentSwitchActive(data),
		ReceivingATCServices:         ReadReceivingATCServices(data),
	}
}
