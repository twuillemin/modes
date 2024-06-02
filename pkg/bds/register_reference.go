package bds

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

// -------------------------------------------------------------------------
//                              INSTANCES
// -------------------------------------------------------------------------

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

// BDS20 is the definition of the register 2,0
var BDS20 = Register{
	id:          "BDS 2,0",
	description: "Aircraft identification",
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
