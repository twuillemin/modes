package fields

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 AltitudeCode (AC)
//
// -----------------------------------------------------------------------------------------

// AltitudeReportMethod details how the altitude of the AltitudeCode (AC) field is reported. It corresponds to AC
// fields bit 26 (M-bit) and bit 28 (Q-bit):
//   - M = 0 and Q = 0 => 100 foot increments
//   - M = 0 and Q = 1 => 25 foot increments. For this method, the coding method is only able to provide values between
//     minus 1 000 ft and plus 50 175 ft.
//   - M = 1 => metric units
//
// Defined at 3.1.2.6.5.4
type AltitudeReportMethod int

const (
	// AltitudeCodeReportNotAvailable signifies that altitude information is not available or that the altitude
	// has been determined invalid.
	AltitudeCodeReportNotAvailable AltitudeReportMethod = 1
	// AltitudeCodeReportMetricUnits signifies that altitude is reported in metric units.
	AltitudeCodeReportMetricUnits AltitudeReportMethod = 2
	// AltitudeCodeReport100FootIncrements signifies that altitude is reported in 100-foot increments.
	AltitudeCodeReport100FootIncrements AltitudeReportMethod = 3
	// AltitudeCodeReport25FootIncrements signifies that altitude is reported in 25-foot increments.
	AltitudeCodeReport25FootIncrements AltitudeReportMethod = 4
)

// ReadAltitude reads the altitude code from a message
func ReadAltitude(message common.MessageData) (int32, AltitudeReportMethod, error) {

	// Altitude code is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// message     |_  _  _  x  x  x  x  x |x  M  x  Q  x  x  x  x
	// 100 foot    |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 0  B2 D2 B4 D4

	// Get the raw altitude code
	altitudeCode := uint16(message.Payload[1]&0x1f)<<8 + uint16(message.Payload[2])

	if altitudeCode == 0 {
		return 0, AltitudeCodeReportNotAvailable, nil
	}

	// Get the M bit
	mBit := (altitudeCode & 0x0040) != 0

	// If altitude reported in metric units
	if mBit {
		// Not specified for now
		return 0, AltitudeCodeReportMetricUnits, nil
	}

	// Get the Q bit
	qBit := (altitudeCode & 0x0010) != 0

	// If altitude reported 25-foot increments
	if qBit {

		// If the M bit equals 0 and the Q bit equals 1, the 11-bit field represented by bits 20 to 25, 27 and 29 to 32
		// shall represent a binary coded field with a "least significant bit" (LSB) of 25 ft. The binary value of the
		// positive decimal integer “N” shall be encoded to report pressure-altitude in the range [(25 N – 1 000)
		// plus or minus 12.5 ft]. The coding of 3.1.2.6.5.4 c) shall be used to report pressure-altitude
		// above 50 187.5 ft.
		n := uint16(0)
		n |= (altitudeCode & 0x1F80) >> 2
		n |= (altitudeCode & 0x0020) >> 1
		n |= altitudeCode & 0x000F

		altitude := 25*int32(n) - 1000

		return altitude, AltitudeCodeReport25FootIncrements, nil
	}

	// Otherwise, altitude is reported in 100 foot increment
	// The altitude shall be coded according to the pattern for Mode C replies of 3.1.1.7.12.2.3.
	// Starting with bit 20 the sequence shall be C1, A1, C2, A2, C4, A4, ZERO, B1, ZERO, B2, D2, B4, D4.
	c1 := (altitudeCode & 0x1000) != 0
	a1 := (altitudeCode & 0x0800) != 0
	c2 := (altitudeCode & 0x0400) != 0
	a2 := (altitudeCode & 0x0200) != 0
	c4 := (altitudeCode & 0x0100) != 0
	a4 := (altitudeCode & 0x0080) != 0
	b1 := (altitudeCode & 0x0020) != 0
	b2 := (altitudeCode & 0x0008) != 0
	d2 := (altitudeCode & 0x0004) != 0
	b4 := (altitudeCode & 0x0002) != 0
	d4 := (altitudeCode & 0x0001) != 0

	altitude, err := bitutils.GillhamToAltitude(false, d2, d4, a1, a2, a4, b1, b2, b4, c1, c2, c4)
	if err != nil {
		return 0, AltitudeCodeReportNotAvailable, errors.New("the Altitude field is malformed")
	}

	return altitude, AltitudeCodeReport100FootIncrements, nil
}

// ToString returns a basic, but readable, representation of the field
func (altitudeReportMethod AltitudeReportMethod) ToString() string {
	switch altitudeReportMethod {
	case AltitudeCodeReportNotAvailable:
		return "1 - Not Available"
	case AltitudeCodeReportMetricUnits:
		return "2 - Metric Units"
	case AltitudeCodeReport100FootIncrements:
		return "3 - 100 Foot Increments"
	case AltitudeCodeReport25FootIncrements:
		return "4 - 25 Foot Increments"
	default:
		return fmt.Sprintf("%v - Unknown code", altitudeReportMethod)
	}
}
