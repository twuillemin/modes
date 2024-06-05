package aircraft_status

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds61"
)

type AircraftStatus struct {
	bds61.AircraftStatus
	adsbVersion adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AircraftStatus) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

func (message AircraftStatus) ToString() string {
	return fmt.Sprintf(""+
		"%v",
		message.AircraftStatus.ToString())
}

// ReadAircraftStatus reads a message at the format AircraftStatus
func ReadAircraftStatus(adsbVersion adsb.ADSBVersion, data []byte) (*AircraftStatus, error) {

	bds, err := bds61.ReadAircraftStatus(data)
	if err != nil {
		return nil, err
	}

	return &AircraftStatus{
		AircraftStatus: bds,
		adsbVersion:    adsbVersion,
	}, nil
}
