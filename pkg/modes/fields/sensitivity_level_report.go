package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 SensitivityLevelReport (SL)
//
// -----------------------------------------------------------------------------------------

// SensitivityLevelReport (SL) downlink field shall be included in both short and long air-air reply formats
// (DF = 0 and 16). This field shall denote the sensitivity level at which ACAS is currently operating.
//
// Defined at 4.3.8.4.2.5
type SensitivityLevelReport int

const (
	// SensitivityLevelACASInoperative signifies that ACAS is inoperative
	SensitivityLevelACASInoperative SensitivityLevelReport = 0
	// SensitivityLevel1 signifies that ACAS is operating at sensitivity level 1
	SensitivityLevel1 SensitivityLevelReport = 1
	// SensitivityLevel2 signifies that ACAS is operating at sensitivity level 2
	SensitivityLevel2 SensitivityLevelReport = 2
	// SensitivityLevel3 signifies that ACAS is operating at sensitivity level 3
	SensitivityLevel3 SensitivityLevelReport = 3
	// SensitivityLevel4 signifies that ACAS is operating at sensitivity level 4
	SensitivityLevel4 SensitivityLevelReport = 4
	// SensitivityLevel5 signifies that ACAS is operating at sensitivity level 5
	SensitivityLevel5 SensitivityLevelReport = 5
	// SensitivityLevel6 signifies that ACAS is operating at sensitivity level 6
	SensitivityLevel6 SensitivityLevelReport = 6
	// SensitivityLevel7 signifies that ACAS is operating at sensitivity level 7
	SensitivityLevel7 SensitivityLevelReport = 7
)

// readSensitivityLevelReport reads the SL field from a message
func ReadSensitivityLevelReport(message common.MessageData) SensitivityLevelReport {

	// The 3 first bits of the message
	sensitivity := message.Payload[0] >> 5

	return SensitivityLevelReport(sensitivity)
}

func (sensitivityLevelReport SensitivityLevelReport) PrettyPrint() string {
	switch sensitivityLevelReport {
	case SensitivityLevelACASInoperative:
		return "0 - ACAS Inoperative"
	case SensitivityLevel1:
		return "1 - Level 1"
	case SensitivityLevel2:
		return "2 - Level 2"
	case SensitivityLevel3:
		return "3 - Level 3"
	case SensitivityLevel4:
		return "4 - Level 4"
	case SensitivityLevel5:
		return "5 - Level 5"
	case SensitivityLevel6:
		return "6 - Level 6"
	case SensitivityLevel7:
		return "7 - Level 7"
	default:
		return fmt.Sprintf("%v - Unknown code", sensitivityLevelReport)
	}
}

func (sensitivityLevelReport SensitivityLevelReport) ExtendedPrettyPrint() string {
	switch sensitivityLevelReport {
	case SensitivityLevelACASInoperative:
		return "0 - ACAS Inoperative"
	case SensitivityLevel1:
		return "1 - Level 1: ACAS is operating at sensitivity level 1"
	case SensitivityLevel2:
		return "2 - Level 2: ACAS is operating at sensitivity level 2"
	case SensitivityLevel3:
		return "3 - Level 3: ACAS is operating at sensitivity level 3"
	case SensitivityLevel4:
		return "4 - Level 4: ACAS is operating at sensitivity level 4"
	case SensitivityLevel5:
		return "5 - Level 5: ACAS is operating at sensitivity level 5"
	case SensitivityLevel6:
		return "6 - Level 6: ACAS is operating at sensitivity level 6"
	case SensitivityLevel7:
		return "7 - Level 7: ACAS is operating at sensitivity level 7"
	default:
		return fmt.Sprintf("%v - Unknown code", sensitivityLevelReport)
	}
}
