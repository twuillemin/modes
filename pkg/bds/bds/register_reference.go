package bds

// -------------------------------------------------------------------------
//                         INTERFACE DEFINITION
// -------------------------------------------------------------------------

// Register is the reference information for Message
type Register interface {
	// GetId returns the id of the BDS message. Should always be in the "BDS x,y" format
	GetId() string

	// GetDescription returns the description of the BDS message
	GetDescription() string
}

// -------------------------------------------------------------------------
//                         INTERNAL STRUCTURE
// -------------------------------------------------------------------------

// The basic structure for keeping information about known messages
type registerDefinition struct {
	id          string
	description string
}

// GetId returns the id of the BDS message. Should always be in the "BDS x,y" format
func (register registerDefinition) GetId() string {
	return register.id
}

// GetDescription returns the description of the BDS message
func (register registerDefinition) GetDescription() string {
	return register.description
}

// -------------------------------------------------------------------------
//                              INSTANCES
// -------------------------------------------------------------------------

// BDS05 is the definition f the register 0,5
var BDS05 = registerDefinition{
	id:          "BDS 0,5",
	description: "Extended squitter airborne position",
}

// BDS06 is the definition of the register 0,6
var BDS06 = registerDefinition{
	id:          "BDS 0,6",
	description: "Extended squitter surface position",
}

// BDS08 is the definition of the register 0,8
var BDS08 = registerDefinition{
	id:          "BDS 0,8",
	description: "Extended squitter aircraft identification and category",
}

// BDS09 is the definition of the register 0,9
var BDS09 = registerDefinition{
	id:          "BDS 0,9",
	description: "Extended squitter airborne velocity",
}

// BDS61 is the definition of the register 6,1
var BDS61 = registerDefinition{
	id:          "BDS 6,1",
	description: "Extended squitter emergency/priority status",
}

// BDS62 is the definition of the register 6,2
var BDS62 = registerDefinition{
	id:          "BDS 6,2",
	description: "Target state and status information",
}

// BDS65 is the definition of the register 6,5
var BDS65 = registerDefinition{
	id:          "BDS 6,5",
	description: "Extended squitter aircraft operational status",
}
