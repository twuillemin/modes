package fields

import (
	"fmt"
)

// TurbulenceFlag is the turbulence flag
//
// Specified in Doc 9871 / Table A-2-68
type TurbulenceFlag byte

const (
	// TFNoDataAvailable indicates turbulence data not available in Register 45.
	TFNoDataAvailable TurbulenceFlag = 0
	// TFDataAvailable indicates turbulence data available in Register 45.
	TFDataAvailable TurbulenceFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (tf TurbulenceFlag) ToString() string {

	switch tf {
	case TFNoDataAvailable:
		return "0 - turbulence data not available in Register 45"
	case TFDataAvailable:
		return "1 - turbulence data available in Register 45"
	default:
		return fmt.Sprintf("%v - Unknown code", tf)
	}
}

func ReadTurbulenceFlag(data []byte) TurbulenceFlag {
	return TurbulenceFlag(data[5] & 0x01)
}
