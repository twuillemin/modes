package fields

import "fmt"

// AirborneOperationalMode is the Operational Mode definition
//
// Specified in Doc 9871 / C.2.3.10.4
type AirborneOperationalMode struct {
	ACASResolutionAdvisoryActive ACASResolutionAdvisoryActiveV2
	IdentSwitchActive            IdentSwitchActive
	SingleAntennaFlag            SingleAntennaFlag
	SystemDesignAssurance        SystemDesignAssurance
}

// ToString returns a basic, but readable, representation of the field
func (mode AirborneOperationalMode) ToString() string {

	return fmt.Sprintf(""+
		"ACAS Resolution Advisory Active: %v\n"+
		"Ident Switch Active:             %v\n"+
		"Single Antenna Flag:             %v\n"+
		"System Design Assurance:         %v",
		mode.ACASResolutionAdvisoryActive.ToString(),
		mode.IdentSwitchActive.ToString(),
		mode.SingleAntennaFlag.ToString(),
		mode.SystemDesignAssurance.ToString())
}

// ReadAirborneOperationalMode reads the AirborneOperationalMode from a 56 bits data field
func ReadAirborneOperationalMode(data []byte) AirborneOperationalMode {
	return AirborneOperationalMode{
		ACASResolutionAdvisoryActive: ReadACASResolutionAdvisoryActiveV2(data),
		IdentSwitchActive:            ReadIdentSwitchActive(data),
		SingleAntennaFlag:            ReadSingleAntennaFlag(data),
		SystemDesignAssurance:        ReadSystemDesignAssurance(data),
	}
}
