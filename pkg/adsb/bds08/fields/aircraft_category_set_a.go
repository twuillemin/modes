package fields

import "fmt"

// AircraftCategorySetA defines the type of the AircraftCategory
//
// Specified in Doc 9871 / Table A-2-8
type AircraftCategorySetA byte

const (
	// ACSANoCategory indicates No aircraft category information
	ACSANoCategory AircraftCategorySetA = 0
	// ACSALight indicates Light (< 15 500 lbs or 7 031 kg)
	ACSALight AircraftCategorySetA = 1
	// ACSAMedium1 indicates (>15 500 to 75 000 lbs, or 7 031 to 34 019 kg)
	ACSAMedium1 AircraftCategorySetA = 2
	// ACSAMedium2 indicates (>75 000 to 300 000 lbs, or 34 019 to 136 078 kg)
	ACSAMedium2 AircraftCategorySetA = 3
	// ACSAHighVortex indicates High vortex aircraft
	ACSAHighVortex AircraftCategorySetA = 4
	// ACSAHeavy indicates (> 300 000 lbs or 136 078 kg)
	ACSAHeavy AircraftCategorySetA = 5
	// ACSAHighPerformance indicates High performance (> 5g acceleration) and high speed (> 400 kt)
	ACSAHighPerformance AircraftCategorySetA = 6
	// ACSARotorCraft indicates Rotorcraft
	ACSARotorCraft AircraftCategorySetA = 7
)

// GetCategorySetName returns the name of the category set
func (category AircraftCategorySetA) GetCategorySetName() string {
	return "Set A"
}

// ToString returns a basic, but readable, representation of the field
func (category AircraftCategorySetA) ToString() string {

	switch category {
	case ACSANoCategory:
		return "0 - No aircraft category information"
	case ACSALight:
		return "1 - Light (< 15_500 lbs or 7_031 kg)"
	case ACSAMedium1:
		return "2 - Medium 1 (>15_500 to 75_000 lbs, or 7_031 to 34_019 kg)"
	case ACSAMedium2:
		return "3 - Medium 2 (>75_000 to 300_000 lbs, or 34_019 to 136_078 kg)"
	case ACSAHighVortex:
		return "4 - High vortex aircraft"
	case ACSAHeavy:
		return "5 - Heavy (> 300_000 lbs or 136_078 kg)"
	case ACSAHighPerformance:
		return "6 - High performance (> 5g acceleration) and high speed (> 400 kt)"
	case ACSARotorCraft:
		return "7 - Rotorcraft"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadAircraftCategorySetA reads the aircraft category from a 56 bits data field
func ReadAircraftCategorySetA(data []byte) AircraftCategorySetA {

	// The category is the 3 lsb bits of the fist byte of the message
	return AircraftCategorySetA(data[0] & 0x7)
}
