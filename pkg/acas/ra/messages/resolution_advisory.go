package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/acas/ra/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// ResolutionAdvisory is an ACAS message providing information about ResolutionAdvisory
//
// Defined at 3.1.2.8.3.1 and 4.3.8.4.2.4
type ResolutionAdvisory struct {
	ActiveRA                fields.ActiveResolutionAdvisory
	RAComplement            fields.RAComplement
	RATerminatedIndicator   fields.RATerminatedIndicator
	MultipleThreatEncounter fields.MultipleThreatEncounter
	ThreatIdentityAddress
}

// ToString returns a basic, but readable, representation of the field
func (message ResolutionAdvisory) ToString() string {
	return fmt.Sprintf("Active Resolution Advisory:\n%v\n"+
		"Active Resolution Advisory Complement:\n%v\n"+
		"Active Resolution Advisory Terminated Indicator: %v\n"+
		"Multiple Threat Encounter:                       %v",
		common.PrefixMultiLine(message.ActiveRA.ToString(), "    "),
		common.PrefixMultiLine(message.RAComplement.ToString(), "    "),
		message.RATerminatedIndicator.ToString(),
		message.MultipleThreatEncounter.ToString())
}

// ReadResolutionAdvisory reads a ResolutionAdvisory data message
//
// Params:
//    - data: The content of the message including the field VDS. This is for example the full content
//               content of the MV field from Mode S message
//
// Returns a properly formatted ResolutionAdvisory
func ReadResolutionAdvisory(data []byte) (*ResolutionAdvisory, error) {

	if len(data) != 6 {
		return nil, errors.New("the data for ACAS ResolutionAdvisory message must be 6 bytes long")
	}

	// Format of the message is as follow:
	//        0                 1                 2                 3                 4                 5
	//                 |             RAC |  R  R M  T  TID |       TID       |       TID       |       TID       |
	//       ARA       |   ARA       RAC |  A  A T  T  d d | d d d d d d d d | d d d d d d d d | d d d d d d _ _ |
	//                 |             RAC |  C  T E  I  TIDA|       TIDA      | TIDA    TIDR    |TIDR     TIDB    |
	// a a a a a a a a | a a a a a a c c | c c t m i i a a | a a a a a a a a | a a a r r r r r | r r b b b b b b |

	return &ResolutionAdvisory{
		ActiveRA:                fields.ReadActiveResolutionAdvisory(data),
		RAComplement:            fields.ReadRAComplement(data),
		RATerminatedIndicator:   fields.ReadRATerminatedIndicator(data),
		MultipleThreatEncounter: fields.ReadMultipleThreatEncounter(data),
	}, nil
}
