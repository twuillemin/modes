package airborne_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/airborne_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds05"
)

type AirbornePositionV0 struct {
	bds05.AirbornePosition
	HorizontalProtectionLimit fields.HorizontalProtectionLimit
	ContainmentRadius         fields.ContainmentRadius
}

// GetADSBVersion returns the ADSB level used to read the data
func (message AirbornePositionV0) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV0
}

func (message AirbornePositionV0) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v",
		message.AirbornePosition.ToString(),
		message.HorizontalProtectionLimit.ToString(),
		message.ContainmentRadius.ToString())
}

// ReadAirbornePositionV0 reads a message at the format AirbornePositionV0
func ReadAirbornePositionV0(data []byte) (*AirbornePositionV0, error) {

	bds, err := bds05.ReadAirbornePosition(data)
	if err != nil {
		return nil, err
	}

	return &AirbornePositionV0{
		AirbornePosition:          *bds,
		HorizontalProtectionLimit: fields.HorizontalProtectionLimit(bds.FormatTypeCode),
		ContainmentRadius:         fields.ContainmentRadius(bds.FormatTypeCode),
	}, nil
}
