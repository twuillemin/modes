package fields

import "fmt"

// AircraftCategorySetC defines the type of the AircraftCategory
type AircraftCategorySetC byte

const (
	// ACSCNoCategory indicates No aircraft category information
	ACSCNoCategory AircraftCategorySetC = 0
	// ACSCSurfaceEmergency indicates Surface vehicle - emergency vehicle
	ACSCSurfaceEmergency AircraftCategorySetC = 1
	// ACSCSurfaceService indicates Surface vehicle - service vehicle
	ACSCSurfaceService AircraftCategorySetC = 2
	// ACSCFixedOrObstruction indicates Fixed ground or tethered obstruction
	ACSCFixedOrObstruction AircraftCategorySetC = 3
	// ACSCReserved1 indicates Reserved
	ACSCReserved1 AircraftCategorySetC = 4
	// ACSCReserved2 indicates Reserved
	ACSCReserved2 AircraftCategorySetC = 5
	// ACSCReserved3 indicates Reserved
	ACSCReserved3 AircraftCategorySetC = 6
	// ACSCReserved4 indicates Reserved
	ACSCReserved4 AircraftCategorySetC = 7
)

// GetCategorySetName returns the name of the category set
func (category AircraftCategorySetC) GetCategorySetName() string {
	return "Set C"
}

// ToString returns a basic, but readable, representation of the field
func (category AircraftCategorySetC) ToString() string {

	switch category {
	case ACSCNoCategory:
		return "0 - No aircraft category information"
	case ACSCSurfaceEmergency:
		return "1 - Surface vehicle - emergency vehicle"
	case ACSCSurfaceService:
		return "2 - Surface vehicle - service vehicle"
	case ACSCFixedOrObstruction:
		return "3 - Fixed ground or tethered obstruction"
	case ACSCReserved1:
		return "4 - Reserved"
	case ACSCReserved2:
		return "5 - Reserved"
	case ACSCReserved3:
		return "6 - Reserved"
	case ACSCReserved4:
		return "7 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadAircraftCategorySetC read the aircraft category from a 56 bits data field
func ReadAircraftCategorySetC(data []byte) AircraftCategorySetC {

	// The category are the the 3 lsb bits of the the fist byte of the message
	return AircraftCategorySetC(data[0] & 0x7)
}
