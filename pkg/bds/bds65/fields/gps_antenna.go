package fields

import "fmt"

// GPSAntenna is the Lateral Axis GPS Antenna Offset Encoding definition
//
// Specified in Doc 9871 / Table C-26, C33 and C-34
type GPSAntenna struct {
	GPSAntennaLateral      GPSAntennaLateral
	GPSAntennaLongitudinal GPSAntennaLongitudinal
}

// ToString returns a basic, but readable, representation of the field
func (antenna GPSAntenna) ToString() string {
	return fmt.Sprintf(""+
		"Lateral Axis GPS Antenna:      %v\n"+
		"Longitudinal Axis GPS Antenna: %v",
		antenna.GPSAntennaLateral.ToString(),
		antenna.GPSAntennaLongitudinal.ToString())
}

// ReadGPSAntenna reads the GPS Antenna from a 56 bits data field
func ReadGPSAntenna(data []byte) GPSAntenna {
	return GPSAntenna{
		GPSAntennaLateral:      ReadGPSAntennaLateral(data),
		GPSAntennaLongitudinal: ReadGPSAntennaLongitudinal(data),
	}
}
