package airborne_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
)

// ReadAirbornePosition reads a message at the format AirbornePosition
func ReadAirbornePosition(adsbVersion adsb.ADSBVersion, data []byte, nicSupplementA bool) (adsb.Message, error) {

	switch adsbVersion {
	case adsb.ADSBV0:
		return ReadAirbornePositionV0(data)
	case adsb.ADSBV1:
		return ReadAirbornePositionV1(data, nicSupplementA)
	case adsb.ADSBV2:
		return ReadAirbornePositionV2(data, nicSupplementA)
	default:
		return nil, fmt.Errorf("the ADS-B version %v does not allow to read AirbornePosition messages", adsbVersion)
	}
}
