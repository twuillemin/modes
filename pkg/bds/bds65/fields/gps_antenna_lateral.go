package fields

import "fmt"

// GPSAntennaLateral is the Lateral Axis GPS Antenna Offset Encoding definition
//
// Specified in Doc 9871 / Table C-33
type GPSAntennaLateral byte

const (
	// GLATNoData indicates No Data
	GLATNoData GPSAntennaLateral = 0
	// GLATLeft2Meters indicates Direction Left at 2 Meters
	GLATLeft2Meters GPSAntennaLateral = 1
	// GLATLeft4Meters indicates Direction Left at 4 Meters
	GLATLeft4Meters GPSAntennaLateral = 2
	// GLATLeft6Meters indicates Direction Left at 6 Meters
	GLATLeft6Meters GPSAntennaLateral = 3
	// GLATRight0Meter indicates Direction Right at 0 Meters
	GLATRight0Meter GPSAntennaLateral = 4
	// GLATRight02Meters indicates Direction Right 2 Meters
	GLATRight02Meters GPSAntennaLateral = 5
	// GLATRight4Meters indicates Direction Right 4 Meters
	GLATRight4Meters GPSAntennaLateral = 6
	// GLATRight6Meters indicates Direction Right 6 Meters
	GLATRight6Meters GPSAntennaLateral = 7
)

// ToString returns a basic, but readable, representation of the field
func (supplement GPSAntennaLateral) ToString() string {

	switch supplement {
	case GLATNoData:
		return "0 - no data"
	case GLATLeft2Meters:
		return "1 - direction left, 2 meters"
	case GLATLeft4Meters:
		return "2 - direction left, 4 meters"
	case GLATLeft6Meters:
		return "3 - direction left, 6 meters"
	case GLATRight0Meter:
		return "4 - direction right, 0 meter"
	case GLATRight02Meters:
		return "5 - direction right, 2 meters"
	case GLATRight4Meters:
		return "6 - direction right, 4 meters"
	case GLATRight6Meters:
		return "7 - direction right, 6 meters"
	default:
		return fmt.Sprintf("%v - Unknown code", supplement)
	}
}

// ReadGPSAntennaLateral reads the GPSAntennaLateral from a 56 bits data field
func ReadGPSAntennaLateral(data []byte) GPSAntennaLateral {
	bits := (data[4] & 0xE0) >> 5
	return GPSAntennaLateral(bits)
}
