package airborne_velocity

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds09"
)

type AirborneVelocity struct {
	bds09.AirborneVelocity
	adsbVersion adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AirborneVelocity) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

func (message AirborneVelocity) ToString() string {
	return fmt.Sprintf(""+
		"%v",
		message.AirborneVelocity.ToString())
}

// ReadAirborneVelocity reads a message at the format AirborneVelocity
func ReadAirborneVelocity(adsbVersion adsb.ADSBVersion, data []byte) (*AirborneVelocity, error) {

	bds, err := bds09.ReadAirborneVelocity(data)
	if err != nil {
		return nil, err
	}

	return &AirborneVelocity{
		AirborneVelocity: bds,
		adsbVersion:      adsbVersion,
	}, nil
}
