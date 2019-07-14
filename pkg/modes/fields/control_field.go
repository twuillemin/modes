package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Control Field (CF)
//
// -----------------------------------------------------------------------------------------

// ControlField (CF) field in DF = 18 shall be used to define the format of the 112-bit transmission as follows.
//
// Defined at 3.1.2.8.7.2
type ControlField int

const (
	// ControlFieldADSB signifies ADS-B ES/NT devices that report the ICAO 24-bit address in the AA field
	ControlFieldADSB ControlField = 0
	// ControlFieldADSBReserved signifies reserved for ADS-B for ES/NT devices that use other addressing techniques in the AA field
	ControlFieldADSBReserved ControlField = 1
	// ControlFieldTISBFineFormat signifies fine format TIS-B message
	ControlFieldTISBFineFormat ControlField = 2
	// ControlFieldTISBCoarseFormat signifies coarse format TIS-B message
	ControlFieldTISBCoarseFormat ControlField = 3
	// ControlFieldTISBReservedManagement signifies reserved for TIS-B management messages
	ControlFieldTISBReservedManagement ControlField = 4
	// ControlFieldTISBRelayADSB signifies TIS-B messages that relay ADS-B messages that use other addressing techniques in the AA field
	ControlFieldTISBRelayADSB ControlField = 5
	// ControlFieldADSBRebroadcast signifies  ADS-B rebroadcast using the same type codes and message formats as defined for DF = 17 ADS-B messages
	ControlFieldADSBRebroadcast ControlField = 6
	// ControlFieldReserved is reserved
	ControlFieldReserved ControlField = 7
)

// ReadControlField reads the CF field from a message
func ReadControlField(message common.MessageData) ControlField {

	return ControlField(message.FirstField)
}

// ToString returns a basic, but readable, representation of the field
func (controlField ControlField) ToString() string {
	switch controlField {
	case ControlFieldADSB:
		return "0 - ADSB (ADS-B ES/NT devices that report the ICAO 24-bit address in the AA field)"
	case ControlFieldADSBReserved:
		return "1 - ADSB Reserved (reserved for ADS-B for ES/NT devices that use other addressing techniques in the AA field)"
	case ControlFieldTISBFineFormat:
		return "2 - TISB Fine Format"
	case ControlFieldTISBCoarseFormat:
		return "3 - TISB Coarse Format"
	case ControlFieldTISBReservedManagement:
		return "4 - TISB Reserved Management"
	case ControlFieldTISBRelayADSB:
		return "5 - TISB Relay ADSB"
	case ControlFieldADSBRebroadcast:
		return "6 - ADSB Rebroadcast (ADS-B rebroadcast using the same type codes and message formats as defined for DF = 17 ADS-B messages)"
	case ControlFieldReserved:
		return "7 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", controlField)
	}
}
