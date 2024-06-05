package bds65

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
)

type AircraftOperationalStatus interface {
	bds.Message
	GetADSBLevel() byte
}

// ReadAircraftOperationalStatus reads a message at the format AircraftOperationalStatus
func ReadAircraftOperationalStatus(data []byte) (AircraftOperationalStatus, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 31 {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadAircraftOperationalStatus", formatTypeCode)
	}

	adsbLevel := (data[5] & 0xE0) >> 5
	subType := data[0] & 0x07

	switch adsbLevel {
	case 0:
		return ReadAircraftOperationalStatusADSB0(data)
	case 1:
		switch subType {
		case 0:
			return ReadAircraftOperationalStatusAirborneADSB1(data)
		case 1:
			return ReadAircraftOperationalStatusSurfaceADSB1(data)
		default:
			return nil, fmt.Errorf("the field SubType must be comprised between 0 and 1 included, got %v", subType)
		}
	case 2:
		switch subType {
		case 0:
			return ReadAircraftOperationalStatusAirborneADSB2(data)
		case 1:
			return ReadAircraftOperationalStatusSurfaceADSB2(data)
		default:
			return nil, fmt.Errorf("the field SubType must be comprised between 0 and 1 included, got %v", subType)
		}
	default:
		return nil, fmt.Errorf("the field VersionNumber must be comprised between 0 and 2 included, got %v", subType)
	}
}
