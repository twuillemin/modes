package fields

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"strconv"
)

// -----------------------------------------------------------------------------------------
//
//                                 Number Of D-Segment (ND)
//
// -----------------------------------------------------------------------------------------

// NumberOfDSegment (ND) downlink field shall designate the number of the message segment contained in MD
//
// Defined at 3.1.2.7.3.2
type NumberOfDSegment uint8

// readNumberOfDSegment reads the ND field from a message
func ReadNumberOfDSegment(message common.MessageData) NumberOfDSegment {

	return NumberOfDSegment((message.DownLinkFormat&0x01)<<3 | message.FirstField)
}

func (numberOfDSegment NumberOfDSegment) PrettyPrint() string {
	return strconv.Itoa(int(numberOfDSegment))
}

func (numberOfDSegment NumberOfDSegment) ExtendedPrettyPrint() string {
	return numberOfDSegment.PrettyPrint()
}
