package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Application Field (AF)
//
// -----------------------------------------------------------------------------------------

// ApplicationField (AF) field in DF = 19 shall be used to define the format of the 112-bit transmission as follows.
//
// Defined at 3.1.2.8.8.2
type ApplicationField int

const (
	// ApplicationFieldReserved0 is reserved
	ApplicationFieldReserved0 ApplicationField = 0
	// ApplicationFieldReserved1 is reserved
	ApplicationFieldReserved1 ApplicationField = 1
	// ApplicationFieldReserved2 is reserved
	ApplicationFieldReserved2 ApplicationField = 2
	// ApplicationFieldReserved3 is reserved
	ApplicationFieldReserved3 ApplicationField = 3
	// ApplicationFieldReserved4 is reserved
	ApplicationFieldReserved4 ApplicationField = 4
	// ApplicationFieldReserved5 is reserved
	ApplicationFieldReserved5 ApplicationField = 5
	// ApplicationFieldReserved6 is reserved
	ApplicationFieldReserved6 ApplicationField = 6
	// ApplicationFieldReserved7 is reserved
	ApplicationFieldReserved7 ApplicationField = 7
)

// readApplicationField reads the AF field from a message
func ReadApplicationField(message common.MessageData) ApplicationField {

	return ApplicationField(message.FirstField)
}

func (applicationField ApplicationField) PrettyPrint() string {
	switch applicationField {
	case ApplicationFieldReserved0:
		return "0 - Reserved"
	case ApplicationFieldReserved1:
		return "1 - Reserved"
	case ApplicationFieldReserved2:
		return "2 - Reserved"
	case ApplicationFieldReserved3:
		return "3 - Reserved"
	case ApplicationFieldReserved4:
		return "4 - Reserved"
	case ApplicationFieldReserved5:
		return "5 - Reserved"
	case ApplicationFieldReserved6:
		return "6 - Reserved"
	case ApplicationFieldReserved7:
		return "7 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", applicationField)
	}
}

func (applicationField ApplicationField) ExtendedPrettyPrint() string {
	return applicationField.PrettyPrint()
}
