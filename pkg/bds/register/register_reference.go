package register

import "fmt"

// -------------------------------------------------------------------------
//                         DEFINITION
// -------------------------------------------------------------------------

// Register is the reference information for Message
type Register struct {
	id          string
	description string
}

// GetId returns the id of the BDS message. Should always be in the "BDS x,y" format
func (register Register) GetId() string {
	return register.id
}

// GetDescription returns the description of the BDS message
func (register Register) GetDescription() string {
	return register.description
}

// ToString returns a printable screen
func (register Register) ToString() string {

	return fmt.Sprintf("%v - %v", register.id, register.description)
}

// -------------------------------------------------------------------------
//                              INSTANCES
// -------------------------------------------------------------------------

// BDS00 is the definition of the register 0,0 (Does not exist in reality)
// See https://www.icao.int/APAC/Meetings/2023%20Mode%20S%20and%20DAPs%20WG6/IP03_SGP%20AI.6%20-%20IC%20codes%20for%20ADS-B%20and%20M-LAT.pdf
// pge 12
var BDS00 = Register{
	id:          "BDS 0,0",
	description: "No message waiting for transmission",
}

// BDS05 is the definition of the register 0,5
var BDS05 = Register{
	id:          "BDS 0,5",
	description: "Extended squitter airborne position",
}

// BDS06 is the definition of the register 0,6
var BDS06 = Register{
	id:          "BDS 0,6",
	description: "Extended squitter surface position",
}

// BDS07 is the definition of the register 0,7
var BDS07 = Register{
	id:          "BDS 0,7",
	description: " Extended squitter status",
}

// BDS08 is the definition of the register 0,8
var BDS08 = Register{
	id:          "BDS 0,8",
	description: "Extended squitter aircraft identification and category",
}

// BDS09 is the definition of the register 0,9
var BDS09 = Register{
	id:          "BDS 0,9",
	description: "Extended squitter airborne velocity",
}

// BDS10 is the definition of the register 1,0
var BDS10 = Register{
	id:          "BDS 1,0",
	description: "Data link capability report",
}

// BDS17 is the definition of the register 1,7
var BDS17 = Register{
	id:          "BDS 1,7",
	description: "Common usage GICB capability report",
}

// BDS20 is the definition of the register 2,0
var BDS20 = Register{
	id:          "BDS 2,0",
	description: "Aircraft identification",
}

// BDS30 is the definition of the register 3,0
var BDS30 = Register{
	id:          "BDS 3,0",
	description: "ACAS active resolution advisory",
}

// BDS40 is the definition of the register 4,0
var BDS40 = Register{
	id:          "BDS 4,0",
	description: "Selected vertical intention",
}

// BDS50 is the definition of the register 5,0
var BDS50 = Register{
	id:          "BDS 5,0",
	description: "Track and turn report",
}

// BDS60 is the definition of the register 6,0
var BDS60 = Register{
	id:          "BDS 6,0",
	description: "Heading and speed report",
}

// BDS61 is the definition of the register 6,1
var BDS61 = Register{
	id:          "BDS 6,1",
	description: "Extended squitter emergency/priority status",
}

// BDS62 is the definition of the register 6,2
var BDS62 = Register{
	id:          "BDS 6,2",
	description: "Target state and status information",
}

// BDS65 is the definition of the register 6,5
var BDS65 = Register{
	id:          "BDS 6,5",
	description: "Extended squitter aircraft operational status",
}
