package fields

import "fmt"

// GPSAntennaLongitudinal is the Lateral Axis GPS Antenna Offset Encoding definition
//
// Specified in Doc 9871 / Table C-33
type GPSAntennaLongitudinal byte

// ToString returns a basic, but readable, representation of the field
func (data GPSAntennaLongitudinal) ToString() string {

	switch data {
	case 0:
		return "0 - no data"
	case 1:
		return "1 - Position Offset Applied by Sensor"
	default:
		return fmt.Sprintf("%v - %v", data, data*2-2)
	}
}

// ReadGPSAntennaLongitudinal reads the GPSAntennaLongitudinal from a 56 bits data field
func ReadGPSAntennaLongitudinal(data []byte) GPSAntennaLongitudinal {
	bits := data[4] & 0x1F
	return GPSAntennaLongitudinal(bits)
}
