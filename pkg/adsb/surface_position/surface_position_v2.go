package surface_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/surface_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds06"
)

type SurfacePositionV2 struct {
	bds06.SurfacePosition
	nicSupplementA              bool
	nicSupplementC              bool
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusV2
	NavigationIntegrityCategory byte
}

func (message SurfacePositionV2) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Navigation Integrity Code A:       %v\n"+
		"Navigation Integrity Code C:       %v\n"+
		"Horizontal Containment Limit:      %v\n"+
		"Navigation Integrity Category:     %v",
		message.SurfacePosition.ToString(),
		message.nicSupplementA,
		message.nicSupplementC,
		message.HorizontalContainmentRadius.ToString(),
		message.NavigationIntegrityCategory)
}

// GetADSBVersion returns the ADSB level used to read the data
func (message SurfacePositionV2) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV2
}

// ReadSurfacePositionV2 reads a message at the format SurfacePositionV2
func ReadSurfacePositionV2(data []byte, nicSupplementA bool, nicSupplementC bool) (*SurfacePositionV2, error) {

	bds, err := bds06.ReadSurfacePosition(data)
	if err != nil {
		return nil, err
	}

	hcr, nic := getHCRAndNICForV2(bds.FormatTypeCode, nicSupplementA, nicSupplementC)

	return &SurfacePositionV2{
		SurfacePosition:             *bds,
		nicSupplementA:              nicSupplementA,
		nicSupplementC:              nicSupplementC,
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}

func getHCRAndNICForV2(formatTypeCode byte, nicSupplementA bool, nicSupplementC bool) (fields.HorizontalContainmentRadiusV2, byte) {
	switch formatTypeCode {
	case 5:
		return fields.HCRV2RcLowerThan7Point5M, 11
	case 6:
		return fields.HCRV2RcLowerThan25M, 10
	case 7:
		if nicSupplementA {
			return fields.HCRV2RcLowerThan75M, 9
		} else {
			return fields.HCRV2RcLowerThan0Point1NM, 8
		}
	default:
		if nicSupplementA {
			if nicSupplementC {
				return fields.HCRV2RcLowerThan0Point2NM, 7
			} else {
				return fields.HCRV2RcLowerThan0Point3NM, 6
			}
		} else {
			if nicSupplementC {
				return fields.HCRV2RcLowerThan0Point6NM, 6
			} else {
				return fields.HCRV2RcGreaterThan0Point6NM, 0
			}
		}
	}
}
