package ra

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
	ThreatTypeIndicator     fields.ThreatTypeIndicator
	ThreatIdentityAddress   *fields.ThreatIdentityAddress
	ThreatIdentityAltitude  *fields.ThreatIdentityAltitude
	ThreatIdentityRange     *fields.ThreatIdentityRange
	ThreatIdentityBearing   *fields.ThreatIdentityBearing
}

// ToString returns a basic, but readable, representation of the field
func (message ResolutionAdvisory) ToString() string {

	threatIdentity := ""
	if message.ThreatTypeIndicator == fields.ThreatTypeNoIdentity {
		threatIdentity = ""
	} else if message.ThreatTypeIndicator == fields.ThreatTypeModeS {
		if message.ThreatIdentityAddress != nil {
			threatIdentity = fmt.Sprintf("\nThreat Mode-S Identity:                          %v", message.ThreatIdentityAddress.ToString())
		} else {
			threatIdentity = fmt.Sprintf("\nThreat Mode-S Identity:                          [Error: Missing Mod-S Address]")
		}
	} else if message.ThreatTypeIndicator == fields.ThreatTypeAltitudeRangeBearing {
		if message.ThreatIdentityAltitude != nil {
			threatIdentity = fmt.Sprintf("\nThreat Altitude:                                 %v", message.ThreatIdentityAltitude.ToString())
		} else {
			threatIdentity = fmt.Sprintf("\nThreat Altitude:                                 [Error: Missing Threat Altitude]")
		}
		if message.ThreatIdentityRange != nil {
			threatIdentity += fmt.Sprintf("\nThreat Range:                                    %v", message.ThreatIdentityRange.ToString())
		} else {
			threatIdentity += fmt.Sprintf("\nThreat Range:                                   [Error: Missing Threat Range]")
		}
		if message.ThreatIdentityBearing != nil {
			threatIdentity += fmt.Sprintf("\nThreat Bearing:                                  %v", message.ThreatIdentityBearing.ToString())
		} else {
			threatIdentity += fmt.Sprintf("\nThreat Bearing:                                 [Error: Missing Threat Bearing]")
		}
	}

	return fmt.Sprintf("Active Resolution Advisory:\n%v\n"+
		"Active Resolution Advisory Complement:\n%v\n"+
		"Active Resolution Advisory Terminated Indicator: %v\n"+
		"Multiple Threat Encounter:                       %v\n"+
		"Threat Type Indicator:                           %v"+
		"%v",
		common.PrefixMultiLine(message.ActiveRA.ToString(), "    "),
		common.PrefixMultiLine(message.RAComplement.ToString(), "    "),
		message.RATerminatedIndicator.ToString(),
		message.MultipleThreatEncounter.ToString(),
		message.ThreatTypeIndicator.ToString(),
		threatIdentity)
}

// ReadResolutionAdvisory reads a ResolutionAdvisory data message
//
// Params:
//   - data: The content of the message including the field VDS. This is for example the full content
//     of the MV field from Mode S message
//
// Returns a properly formatted ResolutionAdvisory
func ReadResolutionAdvisory(data []byte) (*ResolutionAdvisory, error) {

	if len(data) != 6 {
		return nil, errors.New("the data for ACAS ResolutionAdvisory message must be 6 bytes long")
	}

	// Format of the message is as follows:
	//        0                 1                 2                 3                 4                 5
	//                 |             RAC |  R  R M  T  TID |       TID       |       TID       |       TID       |
	//       ARA       |   ARA       RAC |  A  A T  T  d d | d d d d d d d d | d d d d d d d d | d d d d d d _ _ |
	//                 |             RAC |  C  T E  I  TIDA|       TIDA      | TIDA    TIDR    |TIDR     TIDB    |
	// a a a a a a a a | a a a a a a c c | c c t m i i a a | a a a a a a a a | a a a r r r r r | r r b b b b b b |

	var threatIdentityAddress *fields.ThreatIdentityAddress = nil
	var threatIdentityAltitude *fields.ThreatIdentityAltitude = nil
	var threatIdentityRange *fields.ThreatIdentityRange = nil
	var threatIdentityBearing *fields.ThreatIdentityBearing = nil

	threatTypeIndicator := fields.ReadThreatTypeIndicator(data)

	if threatTypeIndicator == fields.ThreatTypeModeS {
		var address = fields.ReadThreatIdentityAddress(data)
		threatIdentityAddress = &address
	} else if threatTypeIndicator == fields.ThreatTypeAltitudeRangeBearing {
		var threatAltitude = fields.ReadThreatIdentityAltitude(data)
		var threatRange = fields.ReadThreatIdentityRange(data)
		var threatBearing = fields.ReadThreatIdentityBearing(data)
		threatIdentityAltitude = &threatAltitude
		threatIdentityRange = &threatRange
		threatIdentityBearing = &threatBearing
	}

	return &ResolutionAdvisory{
		ActiveRA:                fields.ReadActiveResolutionAdvisory(data),
		RAComplement:            fields.ReadRAComplement(data),
		RATerminatedIndicator:   fields.ReadRATerminatedIndicator(data),
		MultipleThreatEncounter: fields.ReadMultipleThreatEncounter(data),
		ThreatTypeIndicator:     threatTypeIndicator,
		ThreatIdentityAddress:   threatIdentityAddress,
		ThreatIdentityAltitude:  threatIdentityAltitude,
		ThreatIdentityRange:     threatIdentityRange,
		ThreatIdentityBearing:   threatIdentityBearing,
	}, nil
}
