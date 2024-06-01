package fields

import (
	"fmt"
)

// ThreatIdentityBearing is TIDR (threat identity data bearing subfield). This 6-bit subfield (83-88) shall contain the
// most recent estimated bearing of the threat aircraft, relative to the ACAS aircraft heading.
type ThreatIdentityBearing byte

// ToString returns a string representation of an address
func (threatBearing ThreatIdentityBearing) ToString() string {

	if threatBearing == 0 {
		return " 0 - No bearing estimate available"
	} else if 1 <= threatBearing && threatBearing <= 60 {
		return fmt.Sprintf("%v - between %v and %v", threatBearing, (threatBearing-1)*6, threatBearing*6)
	} else if 61 <= threatBearing && threatBearing <= 63 {
		return fmt.Sprintf("%v - Not assigned", threatBearing)
	}

	return fmt.Sprintf("%v - Unknown code", threatBearing)
}

// GetBearing returns the approximate  numeric value of the bearing. Note that for extreme values, nonsense are
// returned.
// A message's bearing with a value of 0 or 1 is returned as 0.0. A message's bearing with a value of 61 or more is
// returned as 363.
// The bearing is normally a range between 6*(bearing value-1) and 6*bearing vale. To return just a single value, this
// function returns the approximate, which is  6*bearing vale - 3
func (threatBearing ThreatIdentityBearing) GetBearing() int {
	if threatBearing == 0 {
		return 0
	} else if threatBearing >= 61 {
		return 363
	}

	return int(threatBearing)*6 - 3
}

// ReadThreatIdentityBearing reads the ThreatIdentityBearing
func ReadThreatIdentityBearing(data []byte) ThreatIdentityBearing {

	bits := data[5] & 0x3F
	return ThreatIdentityBearing(bits)
}
