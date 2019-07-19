package fields

import "fmt"

// RAComplement indicates all the currently active RACs, if any, received from other ACAS aircraft.
type RAComplement struct {
	// DoNotPassBelow signifies that the other aircraft ask to not pass below
	DoNotPassBelow bool
	// DoNotPassAbove signifies that the other aircraft ask to not pass above
	DoNotPassAbove bool
	// DoNotTurnLeft signifies that the other aircraft ask to not turn left
	DoNotTurnLeft bool
	// DoNotTurnRight signifies that the other aircraft ask to not turn right
	DoNotTurnRight bool
}

// ToString returns a basic, but readable, representation of the field
func (complement RAComplement) ToString() string {
	return fmt.Sprintf("Do not pass below: %v\n"+
		"Do not pass above: %v\n"+
		"Do not turn left:  %v\n"+
		"Do not turn right: %v",
		complement.DoNotPassBelow,
		complement.DoNotPassAbove,
		complement.DoNotTurnLeft,
		complement.DoNotTurnRight)
}

// ReadRAComplement reads the bit data that constitutes the Resolution Advisory Complement (RAC)
func ReadRAComplement(data []byte) RAComplement {

	return RAComplement{
		DoNotPassBelow: (data[1] & 0x02) != 0,
		DoNotPassAbove: (data[1] & 0x01) != 0,
		DoNotTurnLeft:  (data[2] & 0x80) != 0,
		DoNotTurnRight: (data[2] & 0x40) != 0,
	}
}
