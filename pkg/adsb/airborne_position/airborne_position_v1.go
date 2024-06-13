package airborne_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/airborne_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds05"
)

type AirbornePositionV1 struct {
	bds05.AirbornePosition
	NavigationIntegrityCodeSupplementA bool
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusV1
	NavigationIntegrityCategory        byte
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AirbornePositionV1) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV1
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirbornePositionV1) CheckCoherency() error {
	return message.AirbornePosition.CheckCoherency()
}

func (message AirbornePositionV1) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Navigation Integrity Code A:       %v\n"+
		"Horizontal Containment Radius:     %v\n"+
		"Navigation Integrity Category:     %v",
		message.AirbornePosition.ToString(),
		message.NavigationIntegrityCodeSupplementA,
		message.HorizontalContainmentRadius.ToString(),
		message.NavigationIntegrityCategory)
}

// ReadAirbornePositionV1 reads a message at the format AirbornePositionV1
func ReadAirbornePositionV1(data []byte, nicSupplementA bool) (*AirbornePositionV1, error) {

	bds, err := bds05.ReadAirbornePosition(data)
	if err != nil {
		return nil, err
	}

	hcr, nic := getHCRAndNICForV1(bds.FormatTypeCode, nicSupplementA)

	return &AirbornePositionV1{
		AirbornePosition:                   *bds,
		NavigationIntegrityCodeSupplementA: nicSupplementA,
		HorizontalContainmentRadius:        hcr,
		NavigationIntegrityCategory:        nic,
	}, nil
}

func getHCRAndNICForV1(formatTypeCode byte, nicSupplementA bool) (fields.HorizontalContainmentRadiusV1, byte) {
	switch formatTypeCode {
	case 9:
		return fields.HCRBaroV1RcLowerThan7Point5MAndVPLLowerThan11M, 11
	case 10:
		return fields.HCRBaroV1RcLowerThan25MAndVPLLowerThan37Point5M, 10
	case 11:
		if nicSupplementA {
			return fields.HCRBaroV1RcLowerThan75MAndVPLLowerThan112M, 9
		} else {
			return fields.HCRBaroV1RcLowerThan0Point1NM, 8
		}

	case 12:
		return fields.HCRBaroV1RcLowerThan0Point2NM, 7
	case 13:
		if nicSupplementA {
			return fields.HCRBaroV1RcLowerThan0Point6NM, 6
		} else {
			return fields.HCRBaroV1RcLowerThan0Point5NM, 6
		}
	case 14:
		return fields.HCRBaroV1RcLowerThan1Point0NM, 5
	case 15:
		return fields.HCRBaroV1RcLowerThan2NM, 4
	case 16:
		if nicSupplementA {
			return fields.HCRBaroV1RcLowerThan4NM, 3
		} else {
			return fields.HCRBaroV1RcLowerThan8NM, 2
		}
	case 17:
		return fields.HCRBaroV1RcLowerThan20NM, 1
	case 18:
		return fields.HCRBaroV1RcGreaterThan20NM, 0
	case 20:
		return fields.HCRGNSSV1RcLowerThan7Point5MAndVPLLowerThan11M, 11
	case 21:
		return fields.HCRGNSSV1RcLowerThan25MAndVPLLowerThan37Point5M, 10
	default:
		return fields.HCRGNSSV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown, 0
	}
}
