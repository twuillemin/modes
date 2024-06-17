package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
	"strings"
)

// DTESubAddressStatuses is the support status of DTE Sub-addresses 0 to 15  definition
//
// Specified in Doc 9871 / D.2.4.1
type DTESubAddressStatuses uint16

// ToString returns a basic, but readable, representation of the field
func (dss DTESubAddressStatuses) ToString() string {
	statuses := make([]string, 16)

	for i := uint(0); i < 16; i++ {
		bitShift := 15 - i
		dteStatus := (dss >> bitShift) & 0x01
		statuses[i] = fmt.Sprintf("%d", dteStatus)
	}

	return strings.Join(statuses, ", ")
}

// ReadDTESubAddressStatuses reads the DTESubAddressStatuses from a 56 bits data field
func ReadDTESubAddressStatuses(data []byte) DTESubAddressStatuses {
	return DTESubAddressStatuses(bitutils.Pack2Bytes(data[5], data[6]))
}
