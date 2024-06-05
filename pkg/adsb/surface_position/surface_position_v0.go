package surface_position

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"

	"github.com/twuillemin/modes/pkg/adsb/surface_position/fields"
	"github.com/twuillemin/modes/pkg/bds/bds06"
)

type SurfacePositionV0 struct {
	bds06.SurfacePosition
	HorizontalProtectionLimit fields.HorizontalProtectionLimit
	ContainmentRadius         fields.ContainmentRadius
}

func (message SurfacePositionV0) ToString() string {
	return fmt.Sprintf(""+
		"%v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v",
		message.SurfacePosition.ToString(),
		message.HorizontalProtectionLimit.ToString(),
		message.ContainmentRadius.ToString())
}

// GetADSBVersion returns the ADSB level used to read the data
func (message SurfacePositionV0) GetADSBVersion() adsb.ADSBVersion {
	return adsb.ADSBV0
}

// ReadSurfacePositionV0 reads a message at the format SurfacePositionV0
func ReadSurfacePositionV0(data []byte) (*SurfacePositionV0, error) {

	bds, err := bds06.ReadSurfacePosition(data)
	if err != nil {
		return nil, err
	}

	return &SurfacePositionV0{
		SurfacePosition:           *bds,
		HorizontalProtectionLimit: fields.HorizontalProtectionLimit(bds.FormatTypeCode),
		ContainmentRadius:         fields.ContainmentRadius(bds.FormatTypeCode),
	}, nil
}
