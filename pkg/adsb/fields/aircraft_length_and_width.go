package fields

import "fmt"

// AircraftLengthAndWidth is the Aircraft Length And Width definition
//
// Specified in Doc 9871 / B.2.3.10.11
type AircraftLengthAndWidth byte

const (
	// ALWLength15Width11Point5 indicates Length lower than 15 meters and Width lower than 11.5 meters
	ALWLength15Width11Point5 AircraftLengthAndWidth = 0
	// ALWLength15Width23 indicates Length lower than 15 meters and Width lower than 23 meters
	ALWLength15Width23 AircraftLengthAndWidth = 1
	// ALWLength25Width28Point5 indicates Length lower than 25 meters and Width lower than 28.5 meters
	ALWLength25Width28Point5 AircraftLengthAndWidth = 2
	// ALWLength25Width34 indicates Length lower than 25 meters and Width lower than 34 meters
	ALWLength25Width34 AircraftLengthAndWidth = 3
	// ALWLength35Width33 indicates Length lower than 35 meters and Width lower than 33 meters
	ALWLength35Width33 AircraftLengthAndWidth = 4
	// ALWLength35Width38 indicates Length lower than 35 meters and Width lower than 38 meters
	ALWLength35Width38 AircraftLengthAndWidth = 5
	// ALWLength45Width39Point5 indicates Length lower than 45 meters and Width lower than 39.5 meters
	ALWLength45Width39Point5 AircraftLengthAndWidth = 6
	// ALWLength45Width45 indicates Length lower than 45 meters and Width lower than 45 meters
	ALWLength45Width45 AircraftLengthAndWidth = 7
	// ALWLength55Width45 indicates Length lower than 55 meters and Width lower than 45 meters
	ALWLength55Width45 AircraftLengthAndWidth = 8
	// ALWLength55Width52 indicates Length lower than 55 meters and Width lower than 52 meters
	ALWLength55Width52 AircraftLengthAndWidth = 9
	// ALWLength65Width59Point5 indicates Length lower than 65 meters and Width lower than 59 meters
	ALWLength65Width59Point5 AircraftLengthAndWidth = 10
	// ALWLength65Width67 indicates Length lower than 65 meters and Width lower than 67 meters
	ALWLength65Width67 AircraftLengthAndWidth = 11
	// ALWLength75Width72Point5 indicates Length lower than 75 meters and Width lower than 72.5 meters
	ALWLength75Width72Point5 AircraftLengthAndWidth = 12
	// ALWLength75Width80 indicates Length lower than 75 meters and Width lower than 80 meters
	ALWLength75Width80 AircraftLengthAndWidth = 13
	// ALWLength85Width80 indicates Length lower than 85 meters and Width lower than 80 meters
	ALWLength85Width80 AircraftLengthAndWidth = 14
	// ALWLength85Width90 indicates Length lower than 85 meters and Width lower than 90 meters
	ALWLength85Width90 AircraftLengthAndWidth = 15
)

// ToString returns a basic, but readable, representation of the field
func (status AircraftLengthAndWidth) ToString() string {

	switch status {
	case ALWLength15Width11Point5:
		return "0 - Length lower than 15 meters and Width lower than 11.5 meters"
	case ALWLength15Width23:
		return "1 - Length lower than 15 meters and Width lower than 23 meters"
	case ALWLength25Width28Point5:
		return "2  Length lower than 25 meters and Width lower than 28.5 meters"
	case ALWLength25Width34:
		return "3  Length lower than 25 meters and Width lower than 34 meters"
	case ALWLength35Width33:
		return "4  Length lower than 35 meters and Width lower than 33 meters"
	case ALWLength35Width38:
		return "5  Length lower than 35 meters and Width lower than 38 meters"
	case ALWLength45Width39Point5:
		return "6  Length lower than 45 meters and Width lower than 39.5 meters"
	case ALWLength45Width45:
		return "7  Length lower than 45 meters and Width lower than 45 meters"
	case ALWLength55Width45:
		return "8  Length lower than 55 meters and Width lower than 45 meters"
	case ALWLength55Width52:
		return "9  Length lower than 55 meters and Width lower than 52 meters"
	case ALWLength65Width59Point5:
		return "10 Length lower than 65 meters and Width lower than 59 meters"
	case ALWLength65Width67:
		return "11 Length lower than 65 meters and Width lower than 67 meters"
	case ALWLength75Width72Point5:
		return "12  Length lower than 75 meters and Width lower than 72.5 meters"
	case ALWLength75Width80:
		return "13  Length lower than 75 meters and Width lower than 80 meters"
	case ALWLength85Width80:
		return "14  Length lower than 85 meters and Width lower than 80 meters"
	case ALWLength85Width90:
		return "15  Length lower than 85 meters and Width lower than 90 meters"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadAircraftLengthAndWidth reads the AircraftLengthAndWidth from a 56 bits data field
func ReadAircraftLengthAndWidth(data []byte) AircraftLengthAndWidth {
	bits := data[2] & 0x0F
	return AircraftLengthAndWidth(bits)
}
