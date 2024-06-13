package airborne_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/airborne_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds05"
)

type AirbornePositionV2 struct {
	bds05.AirbornePosition
	NavigationIntegrityCodeSupplementA bool
	NavigationIntegrityCodeSupplementB fields.NavigationIntegrityCodeSupplementB
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusV2
	NavigationIntegrityCategory        byte
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AirbornePositionV2) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV2
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirbornePositionV2) CheckCoherency() error {
	return message.AirbornePosition.CheckCoherency()
}

func (message AirbornePositionV2) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Navigation Integrity Code A:       %v\n"+
		"Navigation Integrity Code B:       %v\n"+
		"Horizontal Containment Radius:     %v\n"+
		"Navigation Integrity Category:     %v",
		message.AirbornePosition.ToString(),
		message.NavigationIntegrityCodeSupplementA,
		message.NavigationIntegrityCodeSupplementB.ToString(),
		message.HorizontalContainmentRadius.ToString(),
		message.NavigationIntegrityCategory)
}

// ReadAirbornePositionV2 reads a message at the format AirbornePositionV2
func ReadAirbornePositionV2(data []byte, nicSupplementA bool) (*AirbornePositionV2, error) {

	bds, err := bds05.ReadAirbornePosition(data)
	if err != nil {
		return nil, err
	}

	nicSupplementB := fields.ReadNavigationIntegritySupplementB(data)

	hcr, nic := getHCRAndNICForV2(bds.FormatTypeCode, nicSupplementA, nicSupplementB)

	return &AirbornePositionV2{
		AirbornePosition:                   *bds,
		NavigationIntegrityCodeSupplementA: nicSupplementA,
		NavigationIntegrityCodeSupplementB: nicSupplementB,
		HorizontalContainmentRadius:        hcr,
		NavigationIntegrityCategory:        nic,
	}, nil
}

func getHCRAndNICForV2(formatTypeCode byte,
	nicSupplementA bool,
	nicSupplementB fields.NavigationIntegrityCodeSupplementB) (fields.HorizontalContainmentRadiusV2, byte) {
	switch formatTypeCode {
	case 9:
		return fields.HCRBaroV2RcLowerThan7Point5M, 11
	case 10:
		return fields.HCRBaroV2RcLowerThan25M, 10
	case 11:
		if nicSupplementA {
			return fields.HCRBaroV2RcLowerThan75M, 9
		} else {
			return fields.HCRBaroV2RcLowerThan0Point1NM, 8
		}
	case 12:
		return fields.HCRBaroV2RcLowerThan0Point2NM, 7
	case 13:
		if !nicSupplementA {
			if nicSupplementB == fields.NICBOne {
				return fields.HCRBaroV2RcLowerThan0Point3NM, 6
			} else {
				return fields.HCRBaroV2RcLowerThan0Point5NM, 6
			}
		} else {
			return fields.HCRBaroV2RcLowerThan0Point6NM, 6
		}
	case 14:
		return fields.HCRBaroV2RcLowerThan1Point0NM, 5
	case 15:
		return fields.HCRBaroV2RcLowerThan2NM, 4
	case 16:
		if nicSupplementA {
			return fields.HCRBaroV2RcLowerThan4NM, 3
		} else {
			return fields.HCRBaroV2RcLowerThan8NM, 2
		}
	case 17:
		return fields.HCRBaroV2RcLowerThan20NM, 1
	case 18:
		return fields.HCRBaroV2RcGreaterThan20NM, 0
	case 20:
		return fields.HCRGNSSV2RcLowerThan7Point5M, 11
	case 21:
		return fields.HCRGNSSV2RcLowerThan25M, 10
	default:
		return fields.HCRGNSSV2RcGreaterThan25MOrUnknown, 0
	}
}
