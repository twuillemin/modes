package fields

import (
	"fmt"
)

// ThreatIdentityRange is TIDR (threat identity data range subfield). This 7-bit subfield (76-82) shall contain the
// most recent threat range estimated by ACAS.
type ThreatIdentityRange byte

// ToString returns a string representation of an address
func (threatRange ThreatIdentityRange) ToString() string {

	if threatRange == 0 {
		return " 0 - No range estimate available"
	} else if threatRange == 1 {
		return " 1 - Less than 0.05"
	} else if 2 <= threatRange && threatRange <= 126 {
		return fmt.Sprintf("%v - %v +/- 0.05 NM", threatRange, threatRange.GetRange())
	} else if threatRange == 127 {
		return "Greater than 12.55"
	}
	return fmt.Sprintf("%v - Unknown code", threatRange)
}

// GetRange returns the numeric value of the range. Note that for extreme values, approximations are returned.
// A message's range with a value of 0 or 1 is returned as 0.0. A message's range with a value of 127 or more is
// returned as 12.6.
func (threatRange ThreatIdentityRange) GetRange() float64 {
	if threatRange == 0 || threatRange == 1 {
		return 0.0
	} else if threatRange >= 127 {
		return 12.6
	}

	return (float64(threatRange) - 1.0) / 10.0
}

// ReadThreatIdentityRange reads the ThreatIdentityRange
func ReadThreatIdentityRange(data []byte) ThreatIdentityRange {

	bits := (data[4]&0x1F)<<2 + (data[5]&0xA0)>>6
	return ThreatIdentityRange(bits)
}
