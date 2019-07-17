package fields

import "fmt"

// AircraftCategorySetB defines the type of the AircraftCategory
type AircraftCategorySetB byte

const (
	// ACSBNoCategory indicates No aircraft category information
	ACSBNoCategory AircraftCategorySetB = 0
	// ACSBGliderSailplane indicates  Glider/sailplane
	ACSBGliderSailplane AircraftCategorySetB = 1
	// ACSBLighterThanAir indicates  Lighter-than-air
	ACSBLighterThanAir AircraftCategorySetB = 2
	// ACSBParachutistSkydiver indicates Parachutist/skydive
	ACSBParachutistSkydiver AircraftCategorySetB = 3
	// ACSBUltralightParaglider indicates Ultralight/hang-glider/paraglider
	ACSBUltralightParaglider AircraftCategorySetB = 4
	// ACSBReserved indicates Reserved
	ACSBReserved AircraftCategorySetB = 5
	// ACSBUnmannedAerialVehicle indicates Unmanned aerial vehicle
	ACSBUnmannedAerialVehicle AircraftCategorySetB = 6
	// ACSBSpace indicates Space/transatmospheric vehicle
	ACSBSpace AircraftCategorySetB = 7
)

// GetCategorySetName returns the name of the category set
func (category AircraftCategorySetB) GetCategorySetName() string {
	return "Set B"
}

// ToString returns a basic, but readable, representation of the field
func (category AircraftCategorySetB) ToString() string {

	switch category {
	case ACSBNoCategory:
		return "0 - No aircraft category information"
	case ACSBGliderSailplane:
		return "1 - Glider / sailplane"
	case ACSBLighterThanAir:
		return "2 - Lighter-than-air"
	case ACSBParachutistSkydiver:
		return "3 - Parachutist / skydiver"
	case ACSBUltralightParaglider:
		return "4 - Ultralight / hang-glider / paraglider"
	case ACSBReserved:
		return "5 - Reserved"
	case ACSBUnmannedAerialVehicle:
		return "6 - Unmanned aerial vehicle"
	case ACSBSpace:
		return "7 - Space / transatmospheric vehicle"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadAircraftCategorySetB reads the aircraft category from a 56 bits data field
func ReadAircraftCategorySetB(data []byte) AircraftCategorySetB {

	// The category are the the 3 lsb bits of the the fist byte of the message
	return AircraftCategorySetB(data[0] & 0x7)
}
