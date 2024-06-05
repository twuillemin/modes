package aircraft_identification_and_category

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds08"
)

type AircraftIdentificationAndCategory struct {
	bds08.AircraftIdentificationAndCategory
	adsbVersion adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AircraftIdentificationAndCategory) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

func (message AircraftIdentificationAndCategory) ToString() string {
	return fmt.Sprintf(""+
		"%v",
		message.AircraftIdentificationAndCategory.ToString())
}

// ReadAircraftIdentificationAndCategory reads a message at the format AircraftIdentificationAndCategory
func ReadAircraftIdentificationAndCategory(adsbVersion adsb.ADSBVersion, data []byte) (*AircraftIdentificationAndCategory, error) {

	bds, err := bds08.ReadAircraftIdentificationAndCategory(data)
	if err != nil {
		return nil, err
	}

	return &AircraftIdentificationAndCategory{
		AircraftIdentificationAndCategory: *bds,
		adsbVersion:                       adsbVersion,
	}, nil
}
