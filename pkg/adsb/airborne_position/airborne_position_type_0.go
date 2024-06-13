package airborne_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/bds/bds05"
)

// AirbornePositionType0 is a special case of AirbornePosition, similar to type 20, 21 and 22, but
// without HPL and other enhanced information. See Table C-41
type AirbornePositionType0 struct {
	bds05.AirbornePosition
	NavigationIntegrityCategory byte
	adsbVersion                 adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AirbornePositionType0) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirbornePositionType0) CheckCoherency() error {
	return message.AirbornePosition.CheckCoherency()
}

func (message AirbornePositionType0) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Navigation Integrity Category:     %v",
		message.AirbornePosition.ToString(),
		message.NavigationIntegrityCategory)
}

// ReadAirbornePositionType0 reads a message at the format AirbornePositionType0
func ReadAirbornePositionType0(adsbVersion adsb.ADSBVersion, data []byte) (*AirbornePositionType0, error) {

	// Update the first byte to have a standard BDS code 0,5 (Airborne position with baro altitude) - format 22
	dataTmp := []byte{0xB0, data[1], data[2], data[3], data[4], data[5], data[6]}

	bds, err := bds05.ReadAirbornePosition(dataTmp)
	if err != nil {
		return nil, err
	}

	bds.FormatTypeCode = 0

	return &AirbornePositionType0{
		AirbornePosition:            *bds,
		NavigationIntegrityCategory: 0,
		adsbVersion:                 adsbVersion,
	}, nil
}
