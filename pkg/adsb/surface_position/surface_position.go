package surface_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
)

// ReadSurfacePosition reads a message at the format SurfacePosition
func ReadSurfacePosition(adsbVersion adsb.ADSBVersion, data []byte, nicSupplementA bool, nicSupplementC bool) (adsb.Message, error) {

	switch adsbVersion {
	case adsb.ADSBV0:
		return ReadSurfacePositionV0(data)
	case adsb.ADSBV1:
		return ReadSurfacePositionV1(data, nicSupplementA)
	case adsb.ADSBV2:
		return ReadSurfacePositionV2(data, nicSupplementA, nicSupplementC)
	default:
		return nil, fmt.Errorf("the ADS-B version %v does not allow to read AirbornePosition messages", adsbVersion)
	}
}
