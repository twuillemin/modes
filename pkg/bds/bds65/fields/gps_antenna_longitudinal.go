package fields

import "fmt"

// GPSAntennaLongitudinal is the Lateral Axis GPS Antenna Offset Encoding definition
//
// Specified in Doc 9871 / Table C-33
type GPSAntennaLongitudinal byte

const (
	// GLONNoData indicates No Data
	GLONNoData GPSAntennaLongitudinal = 0
	// GLONCompensatedBySensor indicates the Antenna Offset is compensated by the Sensor to be the position of the ADS-B participant’s ADS-B Position Reference Point
	GLONCompensatedBySensor GPSAntennaLongitudinal = 1

	// GLON2Meters indicates 2 meters
	GLON2Meters GPSAntennaLongitudinal = 2
	// GLON4Meters indicates 4 meters
	GLON4Meters GPSAntennaLongitudinal = 3
	// GLON6Meters indicates 6 meters
	GLON6Meters GPSAntennaLongitudinal = 4
	// GLON8Meters indicates 8 meters
	GLON8Meters GPSAntennaLongitudinal = 5

	// GLON10Meters indicates 10 meters
	GLON10Meters GPSAntennaLongitudinal = 6
	// GLON12Meters indicates 12 meters
	GLON12Meters GPSAntennaLongitudinal = 7
	// GLON14Meters indicates 14 meters
	GLON14Meters GPSAntennaLongitudinal = 8
	// GLON16Meters indicates 16 meters
	GLON16Meters GPSAntennaLongitudinal = 9
	// GLON18Meters indicates 18 meters
	GLON18Meters GPSAntennaLongitudinal = 10

	// GLON20Meters indicates 20 meters
	GLON20Meters GPSAntennaLongitudinal = 11
	// GLON22Meters indicates 22 meters
	GLON22Meters GPSAntennaLongitudinal = 12
	// GLON24Meters indicates 24 meters
	GLON24Meters GPSAntennaLongitudinal = 13
	// GLON26Meters indicates 26 meters
	GLON26Meters GPSAntennaLongitudinal = 14
	// GLON28Meters indicates 28 meters
	GLON28Meters GPSAntennaLongitudinal = 15

	// GLON30Meters indicates 30 meters
	GLON30Meters GPSAntennaLongitudinal = 16
	// GLON32Meters indicates 32 meters
	GLON32Meters GPSAntennaLongitudinal = 17
	// GLON34Meters indicates 34 meters
	GLON34Meters GPSAntennaLongitudinal = 18
	// GLON36Meters indicates 36 meters
	GLON36Meters GPSAntennaLongitudinal = 19
	// GLON38Meters indicates 38 meters
	GLON38Meters GPSAntennaLongitudinal = 20

	// GLON40Meters indicates 40 meters
	GLON40Meters GPSAntennaLongitudinal = 21
	// GLON42Meters indicates 42 meters
	GLON42Meters GPSAntennaLongitudinal = 22
	// GLON44Meters indicates 44 meters
	GLON44Meters GPSAntennaLongitudinal = 23
	// GLON46Meters indicates 46 meters
	GLON46Meters GPSAntennaLongitudinal = 24
	// GLON48Meters indicates 48 meters
	GLON48Meters GPSAntennaLongitudinal = 25

	// GLON50Meters indicates 50 meters
	GLON50Meters GPSAntennaLongitudinal = 26
	// GLON52Meters indicates 52 meters
	GLON52Meters GPSAntennaLongitudinal = 27
	// GLON54Meters indicates 54 meters
	GLON54Meters GPSAntennaLongitudinal = 28
	// GLON56Meters indicates 56 meters
	GLON56Meters GPSAntennaLongitudinal = 29
	// GLON58Meters indicates 58 meters
	GLON58Meters GPSAntennaLongitudinal = 30

	// GLON60Meters indicates 60 meters
	GLON60Meters GPSAntennaLongitudinal = 31
)

// ToString returns a basic, but readable, representation of the field
func (data GPSAntennaLongitudinal) ToString() string {

	switch data {
	case GLONNoData:
		return "0 - no data"
	case GLONCompensatedBySensor:
		return "1 - Antenna Offset is compensated by the Sensor"
	case GLON2Meters, GLON4Meters, GLON6Meters, GLON8Meters,
		GLON10Meters, GLON12Meters, GLON14Meters, GLON16Meters, GLON18Meters,
		GLON20Meters, GLON22Meters, GLON24Meters, GLON26Meters, GLON28Meters,
		GLON30Meters, GLON32Meters, GLON34Meters, GLON36Meters, GLON38Meters,
		GLON40Meters, GLON42Meters, GLON44Meters, GLON46Meters, GLON48Meters,
		GLON50Meters, GLON52Meters, GLON54Meters, GLON56Meters, GLON58Meters,
		GLON60Meters:
		return fmt.Sprintf("%v - %v", data, data*2-2)
	default:
		return fmt.Sprintf("%v - Unknown code", data)
	}
}

// ReadGPSAntennaLongitudinal reads the GPSAntennaLongitudinal from a 56 bits data field
func ReadGPSAntennaLongitudinal(data []byte) GPSAntennaLongitudinal {
	bits := data[4] & 0x1F
	return GPSAntennaLongitudinal(bits)
}
