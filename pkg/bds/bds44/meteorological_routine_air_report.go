package bds44

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds44/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// MeteorologicalRoutineAirReport is a message at the format BDS 5,0
//
// Specified in Doc 9871 / Table A-2-48
type MeteorologicalRoutineAirReport struct {
	Source                      fields.Source
	WindSpeedStatus             bool
	WindSpeed                   uint32
	WindDirectionStatus         bool
	WindDirection               float32
	StaticAirTemperatureStatus  bool
	StaticAirTemperature        float32
	AverageStaticPressureStatus bool
	AverageStaticPressure       uint32
	TurbulenceFlag              fields.TurbulenceFlag
	HumidityStatus              bool
	Humidity                    uint32
}

// GetRegister returns the Register the message
func (message MeteorologicalRoutineAirReport) GetRegister() register.Register {
	return register.BDS60
}

// ToString returns a basic, but readable, representation of the message
func (message MeteorologicalRoutineAirReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                            %v\n"+
		"Source:                             %v\n"+
		"Wind Speed Status:                  %v\n"+
		"Wind Speed (knot):                  %v\n"+
		"Wind Direction Status:              %v\n"+
		"Wind Direction (degrees):           %v\n"+
		"Static Air Temperature Status:      %v\n"+
		"Static Air Temperature (degrees C): %v\n"+
		"Average Static Pressure Status:     %v\n"+
		"Average Static Pressure (hPa):      %v\n"+
		"Magnetic Heading Orientation:       %v\n"+
		"Humidity Status:                    %v\n"+
		"Humidity (%%):                       %v",
		message.GetRegister().ToString(),
		message.Source.ToString(),
		message.WindSpeedStatus,
		message.WindSpeed,
		message.WindDirectionStatus,
		message.WindDirection,
		message.StaticAirTemperatureStatus,
		message.StaticAirTemperature,
		message.AverageStaticPressureStatus,
		message.AverageStaticPressure,
		message.TurbulenceFlag.ToString(),
		message.HumidityStatus,
		message.Humidity)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message MeteorologicalRoutineAirReport) CheckCoherency() error {
	// If no data available, it is probably not coherent
	if !message.WindSpeedStatus && !message.WindDirectionStatus && !message.StaticAirTemperatureStatus && !message.AverageStaticPressureStatus && !message.HumidityStatus {
		return errors.New("the message does not convey any information")
	}

	if message.Source >= 5 {
		return errors.New("field Source is a Reserved value")
	}

	if !message.WindSpeedStatus && message.WindSpeed != 0 {
		return errors.New("the wind speed status is set to false, but a wind speed value is given")
	}

	if !message.WindDirectionStatus && message.WindDirection != 0 {
		return errors.New("the wind direction status is set to false, but a wind direction value is given")
	}

	if !message.StaticAirTemperatureStatus && message.StaticAirTemperature != 0 {
		return errors.New("the static air temperature status is set to false, but a static air temperature value is given")
	}

	if !message.AverageStaticPressureStatus && message.AverageStaticPressure != 0 {
		return errors.New("the average static pressure status is set to false, but a average static pressure value is given")
	}

	if !message.HumidityStatus && message.Humidity != 0 {
		return errors.New("the humidity status is set to false, but a humidity value is given")
	}

	if message.WindSpeed > 250 {
		return errors.New("the wind speed is too high (above 250 knots)")
	}

	if (message.StaticAirTemperature < -80) || (message.StaticAirTemperature > 60) {
		return errors.New("the static air temperature is to high or to low (-80 <= temp <= 60)")
	}

	return nil
}

// ReadMeteorologicalRoutineAirReport reads a message as a MeteorologicalRoutineAirReport
func ReadMeteorologicalRoutineAirReport(data []byte) (*MeteorologicalRoutineAirReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B MeteorologicalRoutineAirReport message must be 7 bytes long")
	}

	windSpeedStatus, windSpeed := fields.ReadWindSpeed(data)
	windDirectionStatus, windDirection := fields.ReadWindDirection(data)
	staticAirTemperatureStatus, staticAirTemperature := fields.ReadStaticAirTemperature(data)
	averageStaticPressureStatus, averageStaticPressure := fields.ReadAverageStaticPressure(data)
	humidityStatus, humidity := fields.ReadHumidity(data)

	return &MeteorologicalRoutineAirReport{
		Source:                      fields.ReadSource(data),
		WindSpeedStatus:             windSpeedStatus,
		WindSpeed:                   windSpeed,
		WindDirectionStatus:         windDirectionStatus,
		WindDirection:               windDirection,
		StaticAirTemperatureStatus:  staticAirTemperatureStatus,
		StaticAirTemperature:        staticAirTemperature,
		AverageStaticPressureStatus: averageStaticPressureStatus,
		AverageStaticPressure:       averageStaticPressure,
		TurbulenceFlag:              fields.ReadTurbulenceFlag(data),
		HumidityStatus:              humidityStatus,
		Humidity:                    humidity,
	}, nil
}
