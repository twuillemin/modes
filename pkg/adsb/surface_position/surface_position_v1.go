package surface_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/surface_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds06"
)

type SurfacePositionV1 struct {
	bds06.SurfacePosition
	nicSupplementA              bool
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusV1
	NavigationIntegrityCategory byte
}

func (message SurfacePositionV1) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Navigation Integrity Code A:       %v\n"+
		"Horizontal Containment Limit:      %v\n"+
		"Navigation Integrity Category:     %v",
		message.SurfacePosition.ToString(),
		message.nicSupplementA,
		message.HorizontalContainmentRadius.ToString(),
		message.NavigationIntegrityCategory)
}

// GetADSBVersion returns the ADSB level used to read the data
func (message SurfacePositionV1) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV1
}

// ReadSurfacePositionV1 reads a message at the format SurfacePositionV1
func ReadSurfacePositionV1(data []byte, nicSupplementA bool) (*SurfacePositionV1, error) {

	bds, err := bds06.ReadSurfacePosition(data)
	if err != nil {
		return nil, err
	}

	hcr, nic := getHCRAndNICForV1(bds.FormatTypeCode, nicSupplementA)

	return &SurfacePositionV1{
		SurfacePosition:             *bds,
		nicSupplementA:              nicSupplementA,
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}

func getHCRAndNICForV1(formatTypeCode byte, nicSupplementA bool) (fields.HorizontalContainmentRadiusV1, byte) {
	switch formatTypeCode {
	case 5:
		return fields.HCRV1RcLowerThan7Point5M, 11
	case 6:
		return fields.HCRV1RcLowerThan25M, 10
	case 7:
		if nicSupplementA {
			return fields.HCRV1RcLowerThan75M, 9
		} else {
			return fields.HCRV1RcGreaterThan0Point1NM, 8
		}
	default:
		return fields.HCRV1RcGreaterThan0Point1NM, 0
	}
}
