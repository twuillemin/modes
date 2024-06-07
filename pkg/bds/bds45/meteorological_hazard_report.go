package bds45

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds45/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// MeteorologicalHazardReport is a message at the format BDS 4,5
//
// Specified in Doc 9871 / Table A-2-69
type MeteorologicalHazardReport struct {
	TurbulenceStatus            bool
	Turbulence                  fields.HazardLevel
	WindShearStatus             bool
	WindShear                   fields.HazardLevel
	MicroBurstStatus            bool
	MicroBurst                  fields.HazardLevel
	IcingStatus                 bool
	Icing                       fields.HazardLevel
	WakeVortexStatus            bool
	WakeVortex                  fields.HazardLevel
	StaticAirTemperatureStatus  bool
	StaticAirTemperature        float32
	AverageStaticPressureStatus bool
	AverageStaticPressure       uint32
	RadioHeightStatus           bool
	RadioHeight                 uint32
}

// GetRegister returns the Register the message
func (message MeteorologicalHazardReport) GetRegister() register.Register {
	return register.BDS45
}

// ToString returns a basic, but readable, representation of the message
func (message MeteorologicalHazardReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                            %v\n"+
		"Turbulence Status:                  %v\n"+
		"Turbulence:                         %v\n"+
		"Wind Shear Status:                  %v\n"+
		"Wind Shear:                         %v\n"+
		"Micro Burst Status:                 %v\n"+
		"Micro Burst:                        %v\n"+
		"Icing Status:                       %v\n"+
		"Icing:                              %v\n"+
		"Wake Vortex Status:                 %v\n"+
		"Wake Vortex:                        %v\n"+
		"Static Air Temperature Status:      %v\n"+
		"Static Air Temperature (degrees C): %v\n"+
		"Average Static Pressure Status:     %v\n"+
		"Average Static Pressure (hPa):      %v\n"+
		"Radio Height Status:                %v\n"+
		"Radio Height (feet):                %v",
		message.GetRegister().ToString(),
		message.TurbulenceStatus,
		message.Turbulence.ToString(),
		message.WindShearStatus,
		message.WindShear.ToString(),
		message.MicroBurstStatus,
		message.MicroBurst.ToString(),
		message.IcingStatus,
		message.Icing.ToString(),
		message.WakeVortexStatus,
		message.WakeVortex.ToString(),
		message.StaticAirTemperatureStatus,
		message.StaticAirTemperature,
		message.AverageStaticPressureStatus,
		message.AverageStaticPressure,
		message.RadioHeightStatus,
		message.RadioHeight)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message MeteorologicalHazardReport) CheckCoherency() error {
	// If no data available, it is probably not coherent
	if !message.TurbulenceStatus && !message.WindShearStatus && !message.MicroBurstStatus && !message.IcingStatus && !message.WakeVortexStatus && !message.StaticAirTemperatureStatus && !message.AverageStaticPressureStatus && !message.RadioHeightStatus {
		return errors.New("the message does not convey any information")
	}

	if !message.TurbulenceStatus && message.Turbulence != fields.HLNil {
		return errors.New("the turbulence status is set to false, but a turbulence value is given")
	}

	if !message.WindShearStatus && message.WindShear != fields.HLNil {
		return errors.New("the wind shear status is set to false, but a wind shear value is given")
	}

	if !message.MicroBurstStatus && message.MicroBurst != fields.HLNil {
		return errors.New("the micro burst status is set to false, but a micro burst value is given")
	}

	if !message.IcingStatus && message.Icing != fields.HLNil {
		return errors.New("the icing status is set to false, but a icing value is given")
	}

	if !message.WakeVortexStatus && message.WakeVortex != fields.HLNil {
		return errors.New("the wake vortex status is set to false, but a wake vortex value is given")
	}

	if !message.StaticAirTemperatureStatus && message.StaticAirTemperature != 0 {
		return errors.New("the static air temperature status is set to false, but a static air temperature value is given")
	}

	if !message.AverageStaticPressureStatus && message.AverageStaticPressure != 0 {
		return errors.New("the average static pressure status is set to false, but a average static pressure value is given")
	}

	if !message.RadioHeightStatus && message.RadioHeight != 0 {
		return errors.New("the radioHeight status is set to false, but a radioHeight value is given")
	}

	if (message.StaticAirTemperature < -80) || (message.StaticAirTemperature > 60) {
		return errors.New("the static air temperature is to high or to low (-80 <= temp <= 60)")
	}

	return nil
}

// ReadMeteorologicalHazardReport reads a message as a MeteorologicalHazardReport
func ReadMeteorologicalHazardReport(data []byte) (*MeteorologicalHazardReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B MeteorologicalHazardReport message must be 7 bytes long")
	}

	if data[6]&0x1F != 0 {
		return nil, errors.New("the bits 52 to 56 must be zero")
	}

	turbulenceStatus, turbulence := fields.ReadTurbulence(data)
	windShearStatus, windShear := fields.ReadWindShear(data)
	microBurstStatus, microBurst := fields.ReadMicroBurst(data)
	icingStatus, icing := fields.ReadIcing(data)
	wakeVortexStatus, wakeVortex := fields.ReadWakeVortex(data)
	staticAirTemperatureStatus, staticAirTemperature := fields.ReadStaticAirTemperature(data)
	averageStaticPressureStatus, averageStaticPressure := fields.ReadAverageStaticPressure(data)
	radioHeightStatus, radioHeight := fields.ReadRadioHeight(data)

	return &MeteorologicalHazardReport{
		TurbulenceStatus:            turbulenceStatus,
		Turbulence:                  turbulence,
		WindShearStatus:             windShearStatus,
		WindShear:                   windShear,
		MicroBurstStatus:            microBurstStatus,
		MicroBurst:                  microBurst,
		IcingStatus:                 icingStatus,
		Icing:                       icing,
		WakeVortexStatus:            wakeVortexStatus,
		WakeVortex:                  wakeVortex,
		StaticAirTemperatureStatus:  staticAirTemperatureStatus,
		StaticAirTemperature:        staticAirTemperature,
		AverageStaticPressureStatus: averageStaticPressureStatus,
		AverageStaticPressure:       averageStaticPressure,
		RadioHeightStatus:           radioHeightStatus,
		RadioHeight:                 radioHeight,
	}, nil
}
