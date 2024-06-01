package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/common"
)

// SurfaceOperationalMode is the Operational Mode definition
//
// Specified in Doc 9871 / C.2.3.10.4
type SurfaceOperationalMode struct {
	ACASResolutionAdvisoryActive ACASResolutionAdvisoryActiveV2
	IdentSwitchActive            IdentSwitchActive
	SingleAntennaFlag            SingleAntennaFlag
	SystemDesignAssurance        SystemDesignAssurance
	GPSAntenna                   GPSAntenna
}

// ToString returns a basic, but readable, representation of the field
func (mode SurfaceOperationalMode) ToString() string {

	return fmt.Sprintf(""+
		"ACAS Resolution Advisory Active: %v\n"+
		"Ident Switch Active:             %v\n"+
		"Single Antenna Flag:             %v\n"+
		"System Design Assurance:         %v\n"+
		"GPS Antenna:                     %v",
		mode.ACASResolutionAdvisoryActive.ToString(),
		mode.IdentSwitchActive.ToString(),
		mode.SingleAntennaFlag.ToString(),
		mode.SystemDesignAssurance.ToString(),
		common.PrefixMultiLine(mode.GPSAntenna.ToString(), "    "))
}

// ReadSurfaceOperationalMode reads the SurfaceOperationalMode from a 56 bits data field
func ReadSurfaceOperationalMode(data []byte) SurfaceOperationalMode {
	return SurfaceOperationalMode{
		ACASResolutionAdvisoryActive: ReadACASResolutionAdvisoryActiveV2(data),
		IdentSwitchActive:            ReadIdentSwitchActive(data),
		SingleAntennaFlag:            ReadSingleAntennaFlag(data),
		SystemDesignAssurance:        ReadSystemDesignAssurance(data),
		GPSAntenna:                   ReadGPSAntenna(data),
	}
}
