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
	// ReplyInformationReservedACAS1 is reserved for ACAS (1).
	ReplyInformationReservedACAS1 ReplyInformationAirAir = 1
	// ReplyInformationACASResolutionProhibited signifies that ACAS with resolution capability inhibited
	ReplyInformationACASResolutionProhibited ReplyInformationAirAir = 2
	// ReplyInformationACASVerticalOnlyResolution signifies that ACAS with vertical-only resolution capability
	ReplyInformationACASVerticalOnlyResolution ReplyInformationAirAir = 3
	// ReplyInformationACASVerticalAndHorizontalResolution signifies that ACAS with vertical and horizontal resolution capability
	ReplyInformationACASVerticalAndHorizontalResolution ReplyInformationAirAir = 4
	// ReplyInformationReservedACAS5 is reserved for ACAS (5).
	ReplyInformationReservedACAS5 ReplyInformationAirAir = 5
	// ReplyInformationReservedACAS6 is reserved for ACAS (6).
	ReplyInformationReservedACAS6 ReplyInformationAirAir = 6
	// ReplyInformationReservedACAS7 is reserved for ACAS (7).
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

// ReadReplyInformationAirAir reads the RI field from a message
func ReadReplyInformationAirAir(message common.MessageData) ReplyInformationAirAir {

	replyInformation := ((message.Payload[0] & 0x07) << 1) | ((message.Payload[1] & 0x80) >> 7)

	return ReplyInformationAirAir(replyInformation)
}

// ToString returns a basic, but readable, representation of the field
func (replyInformationAirAir ReplyInformationAirAir) ToString() string {
	switch replyInformationAirAir {
	case ReplyInformationNoOperatingACAS:
		return "0 - No Operating ACAS"
	case ReplyInformationReservedACAS1:
		return "1 - Reserved ACAS"
	case ReplyInformationACASResolutionProhibited:
		return "2 - ACAS with resolution capability inhibited"
	case ReplyInformationACASVerticalOnlyResolution:
		return "3 - ACAS with vertical-only resolution capability"
	case ReplyInformationACASVerticalAndHorizontalResolution:
		return "4 - ACAS with vertical and horizontal resolution capability"
	case ReplyInformationReservedACAS5:
		return "5 - Reserved ACAS"
	case ReplyInformationReservedACAS6:
		return "6 - Reserved ACAS"
	case ReplyInformationReservedACAS7:
		return "7 - Reserved ACAS"
	case ReplyInformationMaximumAirSpeedNotAvailable:
		return "8 - Maximum Air Speed Not Available"
	case ReplyInformationMaximumAirSpeed0To140:
		return "9 - Maximum Air Speed 0 To 140 km/h (75 kt)"
	case ReplyInformationMaximumAirSpeed140To280:
		return "10 - Maximum AirSpeed 140 To 280 km/h (75 and 150 kt)"
	case ReplyInformationMaximumAirSpeed280To560:
		return "11 - Maximum AirSpeed 280 To 560 km/h (150 and 300 kt)"
	case ReplyInformationMaximumAirSpeed560To1110:
		return "12 - Maximum AirSpeed 560 To 1110 km/h (300 and 600 kt)"
	case ReplyInformationMaximumAirSpeed1110To2220:
		return "13 - Maximum AirSpeed 1110 To 2220 km/h (600 and 1 200 kt)"
	case ReplyInformationMaximumAirSpeedOver2220:
		return "14 - Maximum Air Speed Over 2220 km/h (1 200 kt)"
	case ReplyInformationNotAssigned:
		return "15 - Not Assigned"
	default:
		return fmt.Sprintf("%v - Unknown code", replyInformationAirAir)
	}
}
