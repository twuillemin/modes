package aircraft_operational_status

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds65"
)

type AircraftOperationalStatus struct {
	bds65.AircraftOperationalStatus
	adsbVersion adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AircraftOperationalStatus) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

func (message AircraftOperationalStatus) ToString() string {
	return fmt.Sprintf(""+
		"%v",
		message.AircraftOperationalStatus.ToString())
}

// ReadAircraftOperationalStatus reads a message at the format AircraftOperationalStatus
func ReadAircraftOperationalStatus(_ adsb.ADSBVersion, data []byte) (*AircraftOperationalStatus, error) {

	bds, err := bds65.ReadAircraftOperationalStatus(data)
	if err != nil {
		return nil, err
	}

	return &AircraftOperationalStatus{
		AircraftOperationalStatus: bds,
		adsbVersion:               adsb.ADSBVersion(bds.GetADSBLevel()),
	}, nil
}
