package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Flight Status (FS)
//
// -----------------------------------------------------------------------------------------

// FlightStatus (FS) field shall contain information about the flight status
//
// Defined at 3.1.2.6.5.1
type FlightStatus int

const (
	// FlightStatusNoAlertNoSPIAirborne signifies no alert and no SPI, aircraft is airborne
	FlightStatusNoAlertNoSPIAirborne FlightStatus = 0
	// FlightStatusNoAlertNoSPIOnTheGround signifies no alert and no SPI, aircraft is on the ground
	FlightStatusNoAlertNoSPIOnTheGround FlightStatus = 1
	// FlightStatusAlertNoSPIAirborne signifies alert and no SPI, aircraft is airborne
	FlightStatusAlertNoSPIAirborne FlightStatus = 2
	// FlightStatusAlertNoSPIOnTheGround signifies alert and no SPI, aircraft is on the ground
	FlightStatusAlertNoSPIOnTheGround FlightStatus = 3
	// FlightStatusAlertSPIAirborne signifies alert and SPI, aircraft is airborne
	FlightStatusAlertSPIAirborne FlightStatus = 4
	// FlightStatusAlertSPIOnTheGround signifies alert and SPI, aircraft is on the ground
	FlightStatusAlertSPIOnTheGround FlightStatus = 5
	// FlightStatusReserved is reserved
	FlightStatusReserved FlightStatus = 6
	// FlightStatusNotAssigned is not assigned
	FlightStatusNotAssigned FlightStatus = 7
)

// readFlightStatus reads the FS field from a message
func ReadFlightStatus(message common.MessageData) FlightStatus {

	return FlightStatus(message.FirstField)
}

func (flightStatus FlightStatus) PrettyPrint() string {
	switch flightStatus {
	case FlightStatusNoAlertNoSPIAirborne:
		return "0 - No Alert, No SPI, Airborne"
	case FlightStatusNoAlertNoSPIOnTheGround:
		return "1 - No Alert, No SPI, On The Ground"
	case FlightStatusAlertNoSPIAirborne:
		return "2 - Alert, No SPI, Airborne"
	case FlightStatusAlertNoSPIOnTheGround:
		return "3 - Alert, No SPI, On The Ground"
	case FlightStatusAlertSPIAirborne:
		return "4 - Alert, SPI, Airborne"
	case FlightStatusAlertSPIOnTheGround:
		return "5 - Alert, SPI, On The Ground"
	case FlightStatusReserved:
		return "6 - Reserved"
	case FlightStatusNotAssigned:
		return "7 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", flightStatus)
	}
}

func (flightStatus FlightStatus) ExtendedPrettyPrint() string {
	return flightStatus.PrettyPrint()
}
