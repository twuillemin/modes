package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 ReplyInformationAirAir (RI)
//
// -----------------------------------------------------------------------------------------

// ReplyInformationAirAir (RI) downlink field shall report the aircraft’s maximum cruising true
// airspeed capability and type of reply to interrogating aircraft.
//   - Value of 0 signifies a reply to an air-air interrogation UF = 0 with AQ = 0, no operating ACAS.
//   - Values from 1 to 7 are reserved for ACAS
//   - Values from 8 to 715 signifies a reply to an air-air interrogation UF = 0 with AQ = 1 and that the
//   maximum airspeed is defined as detailed in the constant values.
//
// Defined at 3.1.2.8.2.2
type ReplyInformationAirAir int

const (
	// ReplyInformationNoOperatingACAS signifies no operating ACAS.
	ReplyInformationNoOperatingACAS ReplyInformationAirAir = 0
	// ReplyInformationNReservedACAS1 is reserved for ACAS (1).
	ReplyInformationReservedACAS1 ReplyInformationAirAir = 1
	// ReplyInformationNReservedACAS2 is reserved for ACAS (2).
	ReplyInformationReservedACAS2 ReplyInformationAirAir = 2
	// ReplyInformationNReservedACAS3 is reserved for ACAS (3).
	ReplyInformationReservedACAS3 ReplyInformationAirAir = 3
	// ReplyInformationNReservedACAS4 is reserved for ACAS (4).
	ReplyInformationReservedACAS4 ReplyInformationAirAir = 4
	// ReplyInformationNReservedACAS5 is reserved for ACAS (5).
	ReplyInformationReservedACAS5 ReplyInformationAirAir = 5
	// ReplyInformationNReservedACAS6 is reserved for ACAS (6).
	ReplyInformationReservedACAS6 ReplyInformationAirAir = 6
	// ReplyInformationNReservedACAS7 is reserved for ACAS (7).
	ReplyInformationReservedACAS7 ReplyInformationAirAir = 7
	// ReplyInformationMaximumAirSpeedNotAvailable signifies no maximum airspeed data available
	ReplyInformationMaximumAirSpeedNotAvailable ReplyInformationAirAir = 8
	// ReplyInformationMaximumAirSpeed0To140 signifies maximum airspeed is less than or equal to 140 km/h (75 kt)
	ReplyInformationMaximumAirSpeed0To140 ReplyInformationAirAir = 9
	// ReplyInformationMaximumAirSpeed140To280 signifies maximum airspeed is greater than 140 and less than or equal
	// to 280 km/h (75 and 150 kt)
	ReplyInformationMaximumAirSpeed140To280 ReplyInformationAirAir = 10
	// ReplyInformationMaximumAirSpeed280To560 signifies maximum airspeed is greater than 280 and less than or equal
	// to 560 km/h (150 and 300 kt)
	ReplyInformationMaximumAirSpeed280To560 ReplyInformationAirAir = 11
	// ReplyInformationMaximumAirSpeed560To1110 signifies maximum airspeed is greater than 560 and less than or equal
	// to 1 110 km/h (300 and 600 kt)
	ReplyInformationMaximumAirSpeed560To1110 ReplyInformationAirAir = 12
	// ReplyInformationMaximumAirSpeed1110To2220 signifies maximum airspeed is greater than 1 110 and less than or
	// equal to 2 220 km/h (600 and 1 200 kt)
	ReplyInformationMaximumAirSpeed1110To2220 ReplyInformationAirAir = 13
	// ReplyInformationMaximumAirSpeedOver2220 signifies maximum airspeed is more than 2 220 km/h (1 200 kt)
	ReplyInformationMaximumAirSpeedOver2220 ReplyInformationAirAir = 14
	// ReplyInformationNotAssigned is not assigned.
	ReplyInformationNotAssigned ReplyInformationAirAir = 15
)

func ReadReplyInformationAirAir(message common.MessageData) ReplyInformationAirAir {

	replyInformation := ((message.Payload[0] & 0x03) << 2) | ((message.Payload[1] & 0x080) >> 7)

	return ReplyInformationAirAir(replyInformation)
}

func (replyInformationAirAir ReplyInformationAirAir) PrettyPrint() string {
	switch replyInformationAirAir {
	case ReplyInformationNoOperatingACAS:
		return "0 - No Operating ACAS"
	case ReplyInformationReservedACAS1:
		return "1 - Reserved ACAS"
	case ReplyInformationReservedACAS2:
		return "2 - Reserved ACAS"
	case ReplyInformationReservedACAS3:
		return "3 - Reserved ACAS"
	case ReplyInformationReservedACAS4:
		return "4 - Reserved ACAS"
	case ReplyInformationReservedACAS5:
		return "5 - Reserved ACAS"
	case ReplyInformationReservedACAS6:
		return "6 - Reserved ACAS"
	case ReplyInformationReservedACAS7:
		return "7 - Reserved ACAS"
	case ReplyInformationMaximumAirSpeedNotAvailable:
		return "8 - Maximum Air Speed Not Available"
	case ReplyInformationMaximumAirSpeed0To140:
		return "9 - Maximum Air Speed 0 To 140"
	case ReplyInformationMaximumAirSpeed140To280:
		return "10 - Maximum AirSpeed 140 To 280"
	case ReplyInformationMaximumAirSpeed280To560:
		return "11 - Maximum AirSpeed 280 To 560"
	case ReplyInformationMaximumAirSpeed560To1110:
		return "12 - Maximum AirSpeed 560 To 1110"
	case ReplyInformationMaximumAirSpeed1110To2220:
		return "13 - Maximum AirSpeed 1110 To 2220"
	case ReplyInformationMaximumAirSpeedOver2220:
		return "14 - Maximum Air Speed Over 2220"
	case ReplyInformationNotAssigned:
		return "15 - Not Assigned"
	default:
		return fmt.Sprintf("%v - Unknown code", replyInformationAirAir)
	}
}

func (replyInformationAirAir ReplyInformationAirAir) ExtendedPrettyPrint() string {
	switch replyInformationAirAir {
	case ReplyInformationNoOperatingACAS:
		return "0 - No Operating ACAS"
	case ReplyInformationReservedACAS1:
		return "1 - Reserved ACAS"
	case ReplyInformationReservedACAS2:
		return "2 - Reserved ACAS"
	case ReplyInformationReservedACAS3:
		return "3 - Reserved ACAS"
	case ReplyInformationReservedACAS4:
		return "4 - Reserved ACAS"
	case ReplyInformationReservedACAS5:
		return "5 - Reserved ACAS"
	case ReplyInformationReservedACAS6:
		return "6 - Reserved ACAS"
	case ReplyInformationReservedACAS7:
		return "7 - Reserved ACAS"
	case ReplyInformationMaximumAirSpeedNotAvailable:
		return "8 - Maximum Air Speed Not Available"
	case ReplyInformationMaximumAirSpeed0To140:
		return "9 - Maximum Air Speed 0 To 140: maximum airspeed is less than or equal to 140 km/h (75 kt)"
	case ReplyInformationMaximumAirSpeed140To280:
		return "10 - Maximum AirSpeed 140 To 280: maximum airspeed is greater than 140 and less than or equal to 280 km/h (75 and 150 kt)"
	case ReplyInformationMaximumAirSpeed280To560:
		return "11 - Maximum AirSpeed 280 To 560: maximum airspeed is greater than 280 and less than or equal to 560 km/h (150 and 300 kt)"
	case ReplyInformationMaximumAirSpeed560To1110:
		return "12 - Maximum AirSpeed 560 To 1110: airspeed is greater than 560 and less than or equal to 1 110 km/h (300 and 600 kt)"
	case ReplyInformationMaximumAirSpeed1110To2220:
		return "13 - Maximum AirSpeed 1110 To 2220: airspeed is greater than 1 110 and less than or equal to 2 220 km/h (600 and 1 200 kt)"
	case ReplyInformationMaximumAirSpeedOver2220:
		return "14 - Maximum Air Speed Over 2220: maximum airspeed is more than 2 220 km/h (1 200 kt)"
	case ReplyInformationNotAssigned:
		return "15 - Not Assigned"
	default:
		return fmt.Sprintf("%v - Unknown code", replyInformationAirAir)
	}
}
