package adsb

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/common"
)

// -------------------------------------------------------------------------
//                         INTERFACE DEFINITION
// -------------------------------------------------------------------------

// MessageFormat is the definition of a format for an ADSB message
type MessageFormat interface {
	common.Printable

	// GetTypeCode returns the id of the MessageFormat message
	GetTypeCode() byte

	// GetRegister returns the BDS implementing the message
	GetRegister() bds.Register
}

// -------------------------------------------------------------------------
//                         INTERNAL STRUCTURE
// -------------------------------------------------------------------------

// The basic structure for keeping information about known adsb messages
type adsbFormatReferenceDefinition struct {
	typeCode byte
	register bds.Register
}

// GetTypeCode returns the id of the ADSB message
func (adsb adsbFormatReferenceDefinition) GetTypeCode() byte {
	return adsb.typeCode
}

// GetRegister returns the BDS implementing the message
func (adsb adsbFormatReferenceDefinition) GetRegister() bds.Register {
	return adsb.register
}

// ToString returns a basic, but readable, representation of the message
func (adsb adsbFormatReferenceDefinition) ToString() string {
	return fmt.Sprintf("%v - %v (%v)",
		adsb.typeCode,
		adsb.register.GetDescription(),
		adsb.register.GetId())
}

// -------------------------------------------------------------------------
//                             THE FORMATS
// -------------------------------------------------------------------------

// Format01 is the definition of a message at Format01
var Format01 = adsbFormatReferenceDefinition{
	typeCode: 1,
	register: bds.BDS08,
}

// Format02 is the definition of a message at Format02
var Format02 = adsbFormatReferenceDefinition{
	typeCode: 2,
	register: bds.BDS08,
}

// Format03 is the definition of a message at Format03
var Format03 = adsbFormatReferenceDefinition{
	typeCode: 3,
	register: bds.BDS08,
}

// Format04 is the definition of a message at Format04
var Format04 = adsbFormatReferenceDefinition{
	typeCode: 4,
	register: bds.BDS08,
}

// Format05 is the definition of a message at Format05
var Format05 = adsbFormatReferenceDefinition{
	typeCode: 5,
	register: bds.BDS06,
}

// Format06 is the definition of a message at Format06
var Format06 = adsbFormatReferenceDefinition{
	typeCode: 6,
	register: bds.BDS06,
}

// Format07 is the definition of a message at Format07
var Format07 = adsbFormatReferenceDefinition{
	typeCode: 7,
	register: bds.BDS06,
}

// Format08 is the definition of a message at Format08
var Format08 = adsbFormatReferenceDefinition{
	typeCode: 8,
	register: bds.BDS06,
}

// Format09 is the definition of a message at Format09
var Format09 = adsbFormatReferenceDefinition{
	typeCode: 9,
	register: bds.BDS05,
}

// Format10 is the definition of a message at Format10
var Format10 = adsbFormatReferenceDefinition{
	typeCode: 10,
	register: bds.BDS05,
}

// Format11 is the definition of a message at Format11
var Format11 = adsbFormatReferenceDefinition{
	typeCode: 11,
	register: bds.BDS05,
}

// Format12 is the definition of a message at Format12
var Format12 = adsbFormatReferenceDefinition{
	typeCode: 12,
	register: bds.BDS05,
}

// Format13 is the definition of a message at Format13
var Format13 = adsbFormatReferenceDefinition{
	typeCode: 13,
	register: bds.BDS05,
}

// Format14 is the definition of a message at Format14
var Format14 = adsbFormatReferenceDefinition{
	typeCode: 14,
	register: bds.BDS05,
}

// Format15 is the definition of a message at Format15
var Format15 = adsbFormatReferenceDefinition{
	typeCode: 15,
	register: bds.BDS05,
}

// Format16 is the definition of a message at Format16
var Format16 = adsbFormatReferenceDefinition{
	typeCode: 16,
	register: bds.BDS05,
}

// Format17 is the definition of a message at Format17
var Format17 = adsbFormatReferenceDefinition{
	typeCode: 17,
	register: bds.BDS05,
}

// Format18 is the definition of a message at Format18
var Format18 = adsbFormatReferenceDefinition{
	typeCode: 18,
	register: bds.BDS05,
}

// Format19 is the definition of a message at Format19
var Format19 = adsbFormatReferenceDefinition{
	typeCode: 19,
	register: bds.BDS09,
}

// Format20 is the definition of a message at Format20
var Format20 = adsbFormatReferenceDefinition{
	typeCode: 20,
	register: bds.BDS05,
}

// Format21 is the definition of a message at Format21
var Format21 = adsbFormatReferenceDefinition{
	typeCode: 21,
	register: bds.BDS05,
}

// Format22 is the definition of a message at Format22
var Format22 = adsbFormatReferenceDefinition{
	typeCode: 22,
	register: bds.BDS05,
}

// Format28 is the definition of a message at Format28
var Format28 = adsbFormatReferenceDefinition{
	typeCode: 28,
	register: bds.BDS61,
}

// Format29 is the definition of a message at Format29
var Format29 = adsbFormatReferenceDefinition{
	typeCode: 29,
	register: bds.BDS62,
}

// Format31 is the definition of a message at Format31
var Format31 = adsbFormatReferenceDefinition{
	typeCode: 31,
	register: bds.BDS65,
}
