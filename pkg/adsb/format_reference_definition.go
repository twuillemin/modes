package adsb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
)

// -------------------------------------------------------------------------
//                         DEFINITION
// -------------------------------------------------------------------------

// MessageFormat is the definition of a format for an ADSB message
type MessageFormat struct {
	typeCode byte
	register bds.Register
}

func (message MessageFormat) GetTypeCode() byte {
	return message.typeCode
}

func (message MessageFormat) GetRegister() bds.Register {
	return message.register
}

// ToString returns a basic, but readable, representation of the message
func (message MessageFormat) ToString() string {
	return fmt.Sprintf("%v - %v (%v)",
		message.typeCode,
		message.register.GetDescription(),
		message.register.GetId())
}

// -------------------------------------------------------------------------
//                             THE FORMATS
// -------------------------------------------------------------------------

// Format01 is the definition of a message at Format01
var Format01 = MessageFormat{
	typeCode: 1,
	register: bds.BDS08,
}

// Format02 is the definition of a message at Format02
var Format02 = MessageFormat{
	typeCode: 2,
	register: bds.BDS08,
}

// Format03 is the definition of a message at Format03
var Format03 = MessageFormat{
	typeCode: 3,
	register: bds.BDS08,
}

// Format04 is the definition of a message at Format04
var Format04 = MessageFormat{
	typeCode: 4,
	register: bds.BDS08,
}

// Format05 is the definition of a message at Format05
var Format05 = MessageFormat{
	typeCode: 5,
	register: bds.BDS06,
}

// Format06 is the definition of a message at Format06
var Format06 = MessageFormat{
	typeCode: 6,
	register: bds.BDS06,
}

// Format07 is the definition of a message at Format07
var Format07 = MessageFormat{
	typeCode: 7,
	register: bds.BDS06,
}

// Format08 is the definition of a message at Format08
var Format08 = MessageFormat{
	typeCode: 8,
	register: bds.BDS06,
}

// Format09 is the definition of a message at Format09
var Format09 = MessageFormat{
	typeCode: 9,
	register: bds.BDS05,
}

// Format10 is the definition of a message at Format10
var Format10 = MessageFormat{
	typeCode: 10,
	register: bds.BDS05,
}

// Format11 is the definition of a message at Format11
var Format11 = MessageFormat{
	typeCode: 11,
	register: bds.BDS05,
}

// Format12 is the definition of a message at Format12
var Format12 = MessageFormat{
	typeCode: 12,
	register: bds.BDS05,
}

// Format13 is the definition of a message at Format13
var Format13 = MessageFormat{
	typeCode: 13,
	register: bds.BDS05,
}

// Format14 is the definition of a message at Format14
var Format14 = MessageFormat{
	typeCode: 14,
	register: bds.BDS05,
}

// Format15 is the definition of a message at Format15
var Format15 = MessageFormat{
	typeCode: 15,
	register: bds.BDS05,
}

// Format16 is the definition of a message at Format16
var Format16 = MessageFormat{
	typeCode: 16,
	register: bds.BDS05,
}

// Format17 is the definition of a message at Format17
var Format17 = MessageFormat{
	typeCode: 17,
	register: bds.BDS05,
}

// Format18 is the definition of a message at Format18
var Format18 = MessageFormat{
	typeCode: 18,
	register: bds.BDS05,
}

// Format19 is the definition of a message at Format19
var Format19 = MessageFormat{
	typeCode: 19,
	register: bds.BDS09,
}

// Format20 is the definition of a message at Format20
var Format20 = MessageFormat{
	typeCode: 20,
	register: bds.BDS05,
}

// Format21 is the definition of a message at Format21
var Format21 = MessageFormat{
	typeCode: 21,
	register: bds.BDS05,
}

// Format22 is the definition of a message at Format22
var Format22 = MessageFormat{
	typeCode: 22,
	register: bds.BDS05,
}

// Format28 is the definition of a message at Format28
var Format28 = MessageFormat{
	typeCode: 28,
	register: bds.BDS61,
}

// Format29 is the definition of a message at Format29
var Format29 = MessageFormat{
	typeCode: 29,
	register: bds.BDS62,
}

// Format31 is the definition of a message at Format31
var Format31 = MessageFormat{
	typeCode: 31,
	register: bds.BDS65,
}
