package fields

import "fmt"

// LengthWidth is the Aircraft Length And Width definition
//
// Specified in Doc 9871 / B.2.3.10.11
type LengthWidth byte

const (
	// LWLength15Width11Point5 indicates Length lower than 15 meters and Width lower than 11.5 meters
	LWLength15Width11Point5 LengthWidth = 0
	// LWLength15Width23 indicates Length lower than 15 meters and Width lower than 23 meters
	LWLength15Width23 LengthWidth = 1
	// LWLength25Width28Point5 indicates Length lower than 25 meters and Width lower than 28.5 meters
	LWLength25Width28Point5 LengthWidth = 2
	// LWLength25Width34 indicates Length lower than 25 meters and Width lower than 34 meters
	LWLength25Width34 LengthWidth = 3
	// LWLength35Width33 indicates Length lower than 35 meters and Width lower than 33 meters
	LWLength35Width33 LengthWidth = 4
	// LWLength35Width38 indicates Length lower than 35 meters and Width lower than 38 meters
	LWLength35Width38 LengthWidth = 5
	// LWLength45Width39Point5 indicates Length lower than 45 meters and Width lower than 39.5 meters
	LWLength45Width39Point5 LengthWidth = 6
	// LWLength45Width45 indicates Length lower than 45 meters and Width lower than 45 meters
	LWLength45Width45 LengthWidth = 7
	// LWLength55Width45 indicates Length lower than 55 meters and Width lower than 45 meters
	LWLength55Width45 LengthWidth = 8
	// LWLength55Width52 indicates Length lower than 55 meters and Width lower than 52 meters
	LWLength55Width52 LengthWidth = 9
	// LWLength65Width59Point5 indicates Length lower than 65 meters and Width lower than 59 meters
	LWLength65Width59Point5 LengthWidth = 10
	// LWLength65Width67 indicates Length lower than 65 meters and Width lower than 67 meters
	LWLength65Width67 LengthWidth = 11
	// LWLength75Width72Point5 indicates Length lower than 75 meters and Width lower than 72.5 meters
	LWLength75Width72Point5 LengthWidth = 12
	// LWLength75Width80 indicates Length lower than 75 meters and Width lower than 80 meters
	LWLength75Width80 LengthWidth = 13
	// LWLength85Width80 indicates Length lower than 85 meters and Width lower than 80 meters
	LWLength85Width80 LengthWidth = 14
	// LWLength85Width90 indicates Length lower than 85 meters and Width lower than 90 meters
	LWLength85Width90 LengthWidth = 15
)

// ToString returns a basic, but readable, representation of the field
func (status LengthWidth) ToString() string {

	switch status {
	case LWLength15Width11Point5:
		return "0 - Length lower than 15 meters and Width lower than 11.5 meters"
	case LWLength15Width23:
		return "1 - Length lower than 15 meters and Width lower than 23 meters"
	case LWLength25Width28Point5:
		return "2 - Length lower than 25 meters and Width lower than 28.5 meters"
	case LWLength25Width34:
		return "3 - Length lower than 25 meters and Width lower than 34 meters"
	case LWLength35Width33:
		return "4 - Length lower than 35 meters and Width lower than 33 meters"
	case LWLength35Width38:
		return "5 - Length lower than 35 meters and Width lower than 38 meters"
	case LWLength45Width39Point5:
		return "6 - Length lower than 45 meters and Width lower than 39.5 meters"
	case LWLength45Width45:
		return "7 - Length lower than 45 meters and Width lower than 45 meters"
	case LWLength55Width45:
		return "8 - Length lower than 55 meters and Width lower than 45 meters"
	case LWLength55Width52:
		return "9 - Length lower than 55 meters and Width lower than 52 meters"
	case LWLength65Width59Point5:
		return "10 - Length lower than 65 meters and Width lower than 59 meters"
	case LWLength65Width67:
		return "11 - Length lower than 65 meters and Width lower than 67 meters"
	case LWLength75Width72Point5:
		return "12 - Length lower than 75 meters and Width lower than 72.5 meters"
	case LWLength75Width80:
		return "13 - Length lower than 75 meters and Width lower than 80 meters"
	case LWLength85Width80:
		return "14 - Length lower than 85 meters and Width lower than 80 meters"
	case LWLength85Width90:
		return "15 - Length lower than 85 meters and Width lower than 90 meters"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadAircraftLengthAndWidth reads the LengthWidth from a 56 bits data field
func ReadAircraftLengthAndWidth(data []byte) LengthWidth {
	bits := data[2] & 0x0F
	return LengthWidth(bits)
}
