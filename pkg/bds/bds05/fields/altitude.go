package fields

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// AltitudeSource is the type of source of the Altitude: Barometric or GNSS
type AltitudeSource int

const (
	// AltitudeBarometric signifies that altitude is barometric altitude
	AltitudeBarometric AltitudeSource = 0
	// AltitudeGNSS signifies that altitude is GNSS height (HAE)
	AltitudeGNSS AltitudeSource = 1
)

// ToString returns a basic, but readable, representation of the field
func (altitudeSource AltitudeSource) ToString() string {
	switch altitudeSource {
	case AltitudeBarometric:
		return "0 - Barometric"
	case AltitudeGNSS:
		return "1 - GNSS"
	default:
		return fmt.Sprintf("%v - Unknown code", altitudeSource)
	}
}

// AltitudeReportMethod defines how the altitude is reported
type AltitudeReportMethod int

const (
	// AltitudeReport100FootIncrements signifies that altitude is reported in 100-foot increments
	AltitudeReport100FootIncrements AltitudeReportMethod = 0
	// AltitudeReport25FootIncrements signifies that altitude is reported in 25-foot increments
	AltitudeReport25FootIncrements AltitudeReportMethod = 1
)

// ToString returns a basic, but readable, representation of the field
func (altitudeReportMethod AltitudeReportMethod) ToString() string {
	switch altitudeReportMethod {
	case AltitudeReport100FootIncrements:
		return "0 - 100 Foot Increments"
	case AltitudeReport25FootIncrements:
		return "1 - 25 Foot Increments"
	default:
		return fmt.Sprintf("%v - Unknown code", altitudeReportMethod)
	}
}

// ReadAltitude reads the altitude code from a message.
func ReadAltitude(data []byte) (int32, AltitudeSource, AltitudeReportMethod, error) {

	// Determines the source of the altitude. Only format 20 to 22 are using GNSS altitude
	source := AltitudeBarometric
	format := (data[0] & 0xF8) >> 3
	if 20 <= format && format <= 22 {
		source = AltitudeGNSS
	}

	// Altitude code is a 12 bits fields, so read a uint16
	// bit         | 8  9 10 11 12 13 14 15| 16 17 18 19 20 21 22 23 |
	// message     | x  x  x  x  x  x  x  x|  x  x  x  x  _  _  _  _ |
	// 100 foot    |C1 A1 C2 A2 C4 A4 B1  Q| B2 D2 B4 D4  _  _  _  _ |

	// Get the Q bit
	qBit := (data[1] & 0x01) != 0

	// If altitude reported 25-foot increments
	if qBit {

		// If the Q bit equals 1, the 11-bit field represented by bits 8 to 14 and 16 to 18
		// shall represent a binary coded field with a "least significant bit" (LSB) of 25 ft. The binary value of the
		// positive decimal integer “N” shall be encoded to report pressure-altitude in the range [(25 N – 1 000)
		// plus or minus 12.5 ft]. The coding of 3.1.2.6.5.4 c) shall be used to report pressure-altitude
		// above 50 187.5 ft.
		// Annexe 10, volume IV, 3.1.2.6.5.4
		n := uint16(0)
		n |= uint16(data[1]&0xFE) << 3
		n |= uint16(data[2]&0xF0) >> 4

		return 25*int32(n) - 1000, source, AltitudeReport25FootIncrements, nil
	}

	// Otherwise, altitude is reported in 100 foot increment
	// The altitude shall be coded according to the pattern for Mode C replies of 3.1.1.7.12.2.3.
	// Starting with bit 20 the sequence shall be C1, A1, C2, A2, C4, A4, <<ZERO>>, B1, ZERO, B2, D2, B4, D4.
	// The M bit <<ZERO>> is ignored for ADSB
	c1 := (data[1] & 0x80) != 0
	a1 := (data[1] & 0x40) != 0
	c2 := (data[1] & 0x20) != 0
	a2 := (data[1] & 0x10) != 0
	c4 := (data[1] & 0x08) != 0
	a4 := (data[1] & 0x04) != 0
	b1 := (data[1] & 0x02) != 0
	b2 := (data[2] & 0x80) != 0
	d2 := (data[2] & 0x40) != 0
	b4 := (data[2] & 0x20) != 0
	d4 := (data[2] & 0x10) != 0

	altitudeFeet, err := bitutils.GillhamToAltitude(false, d2, d4, a1, a2, a4, b1, b2, b4, c1, c2, c4)
	if err != nil {
		return 0, 0, 0, errors.New("the Altitude field is malformed")
	}

	return altitudeFeet, source, AltitudeReport100FootIncrements, nil
}
