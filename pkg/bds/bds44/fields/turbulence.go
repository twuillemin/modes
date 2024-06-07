package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// TurbulenceLevel is the turbulence flag
//
// Specified in Doc 9871 / Table A-2-68
type TurbulenceLevel byte

const (
	// TLNil indicates Nil.
	TLNil TurbulenceLevel = 0
	// TLLight indicates Light.
	TLLight TurbulenceLevel = 1
	// TLModerate indicates Moderate.
	TLModerate TurbulenceLevel = 2
	// TLSevere indicates Severe.
	TLSevere TurbulenceLevel = 3
)

// ToString returns a basic, but readable, representation of the field
func (tl TurbulenceLevel) ToString() string {

	switch tl {
	case TLNil:
		return "0 - Nil"
	case TLLight:
		return "1 - Light"
	case TLModerate:
		return "2 - Moderate"
	case TLSevere:
		return "3 - Severe"
	default:
		return fmt.Sprintf("%v - Unknown code", tl)
	}
}

func ReadTurbulence(data []byte) (bool, TurbulenceLevel) {
	status := (data[5] & 0x02) != 0

	byte1 := data[5] & 0x01
	byte2 := data[6] & 0x80
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 7
	level := TurbulenceLevel(allBits)

	return status, level
}
