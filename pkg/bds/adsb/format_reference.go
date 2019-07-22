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

	// GetADSBLevel() returns the ADSB version implementing the message
	GetADSBLevel() Level
}

// -------------------------------------------------------------------------
//                         INTERNAL STRUCTURE
// -------------------------------------------------------------------------

// The basic structure for keeping information about known adsb messages
type adsbFormatReferenceDefinition struct {
	typeCode  byte
	register  bds.Register
	adsbLevel Level
}

// GetTypeCode returns the id of the ADSB message
func (adsb adsbFormatReferenceDefinition) GetTypeCode() byte {
	return adsb.typeCode
}

// GetRegister returns the BDS implementing the message
func (adsb adsbFormatReferenceDefinition) GetRegister() bds.Register {
	return adsb.register
}

// GetADSBLevel() returns the ADSB version implementing the format
func (adsb adsbFormatReferenceDefinition) GetADSBLevel() Level {
	return adsb.adsbLevel
}

// ToString returns a basic, but readable, representation of the message
func (adsb adsbFormatReferenceDefinition) ToString() string {
	return fmt.Sprintf("%v - %v (%v) [%v]",
		adsb.typeCode,
		adsb.register.GetDescription(),
		adsb.register.GetId(),
		adsb.adsbLevel.ToString())
}

// -------------------------------------------------------------------------
//                              INSTANCES
// -------------------------------------------------------------------------

// -----------------------------  FORMAT 01  -------------------------------

// Format01V0OrMore is the definition of a message at at MessageFormat 01 for ADSB V0 or higher
var Format01V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  1,
	register:  bds.BDS08,
	adsbLevel: Level0OrMore,
}

// -----------------------------  FORMAT 02  -------------------------------

// Format02V0OrMore is the definition of a message at at MessageFormat 02 for ADSB V0 or higher
var Format02V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  2,
	register:  bds.BDS08,
	adsbLevel: Level0OrMore,
}

// -----------------------------  FORMAT 03  -------------------------------

// Format03V0OrMore is the definition of a message at at MessageFormat 03 for ADSB V0 or higher
var Format03V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  3,
	register:  bds.BDS08,
	adsbLevel: Level0OrMore,
}

// -----------------------------  FORMAT 04  -------------------------------

// Format04V0OrMore is the definition of a message at at MessageFormat 04 for ADSB V0 or higher
var Format04V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  4,
	register:  bds.BDS08,
	adsbLevel: Level0OrMore,
}

// -----------------------------  FORMAT 05  -------------------------------

// Format05V0 is the definition of a message at at MessageFormat 05 for ADSB V0
var Format05V0 = adsbFormatReferenceDefinition{
	typeCode:  5,
	register:  bds.BDS06,
	adsbLevel: Level0Exactly,
}

// Format05V1 is the definition of a message at at MessageFormat 05 for ADSB V1
var Format05V1 = adsbFormatReferenceDefinition{
	typeCode:  5,
	register:  bds.BDS06,
	adsbLevel: Level1Exactly,
}

// Format05V2 is the definition of a message at at MessageFormat 05 for ADSB V2
var Format05V2 = adsbFormatReferenceDefinition{
	typeCode:  5,
	register:  bds.BDS06,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 06  -------------------------------

// Format06V0 is the definition of a message at at MessageFormat 06 for ADSB V0
var Format06V0 = adsbFormatReferenceDefinition{
	typeCode:  6,
	register:  bds.BDS06,
	adsbLevel: Level0Exactly,
}

// Format06V1 is the definition of a message at at MessageFormat 06 for ADSB V1
var Format06V1 = adsbFormatReferenceDefinition{
	typeCode:  6,
	register:  bds.BDS06,
	adsbLevel: Level1Exactly,
}

// Format06V2 is the definition of a message at at MessageFormat 06 for ADSB V2
var Format06V2 = adsbFormatReferenceDefinition{
	typeCode:  6,
	register:  bds.BDS06,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 07  -------------------------------

// Format07V0 is the definition of a message at at MessageFormat 07 for ADSB V0
var Format07V0 = adsbFormatReferenceDefinition{
	typeCode:  7,
	register:  bds.BDS06,
	adsbLevel: Level0Exactly,
}

// Format07V1 is the definition of a message at at MessageFormat 07 for ADSB V1
var Format07V1 = adsbFormatReferenceDefinition{
	typeCode:  7,
	register:  bds.BDS06,
	adsbLevel: Level1Exactly,
}

// Format07V2 is the definition of a message at at MessageFormat 07 for ADSB V2
var Format07V2 = adsbFormatReferenceDefinition{
	typeCode:  7,
	register:  bds.BDS06,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 08  -------------------------------

// Format08V0 is the definition of a message at at MessageFormat 08 for ADSB V0
var Format08V0 = adsbFormatReferenceDefinition{
	typeCode:  8,
	register:  bds.BDS06,
	adsbLevel: Level0Exactly,
}

// Format08V1 is the definition of a message at at MessageFormat 08 for ADSB V1
var Format08V1 = adsbFormatReferenceDefinition{
	typeCode:  8,
	register:  bds.BDS06,
	adsbLevel: Level1Exactly,
}

// Format08V2 is the definition of a message at at MessageFormat 08 for ADSB V2
var Format08V2 = adsbFormatReferenceDefinition{
	typeCode:  8,
	register:  bds.BDS06,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 09  -------------------------------

// Format09V0 is the definition of a message at at MessageFormat 09 for ADSB V0
var Format09V0 = adsbFormatReferenceDefinition{
	typeCode:  9,
	register:  bds.BDS05,
	adsbLevel: Level0Exactly,
}

// Format09V1 is the definition of a message at at MessageFormat 09 for ADSB V1
var Format09V1 = adsbFormatReferenceDefinition{
	typeCode:  9,
	register:  bds.BDS05,
	adsbLevel: Level1Exactly,
}

// Format09V2 is the definition of a message at at MessageFormat 09 for ADSB V2
var Format09V2 = adsbFormatReferenceDefinition{
	typeCode:  9,
	register:  bds.BDS05,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 19  -------------------------------

// Format19V0OrMore is the definition of a message at at MessageFormat 19 for ADSB V0 or higher
var Format19V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  19,
	register:  bds.BDS09,
	adsbLevel: Level0OrMore,
}

// -----------------------------  FORMAT 28  -------------------------------

// Format28V0OrMore is the definition of a message at at MessageFormat 28 for ADSB V0 or higher
var Format28V0OrMore = adsbFormatReferenceDefinition{
	typeCode:  28,
	register:  bds.BDS61,
	adsbLevel: Level0OrMore,
}

// Format28V0 is the definition of a message at at MessageFormat 28 for ADSB V0
var Format28V0 = adsbFormatReferenceDefinition{
	typeCode:  28,
	register:  bds.BDS61,
	adsbLevel: Level0Exactly,
}

// Format28V1 is the definition of a message at at MessageFormat 28 for ADSB V1
var Format28V1 = adsbFormatReferenceDefinition{
	typeCode:  28,
	register:  bds.BDS61,
	adsbLevel: Level1Exactly,
}

// Format28V2 is the definition of a message at at MessageFormat 28 for ADSB V2
var Format28V2 = adsbFormatReferenceDefinition{
	typeCode:  9,
	register:  bds.BDS61,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 29  -------------------------------

// Format29V1OrMore is the definition of a message at at MessageFormat 29 for ADSB V1 or higher
var Format29V1OrMore = adsbFormatReferenceDefinition{
	typeCode:  29,
	register:  bds.BDS62,
	adsbLevel: Level1OrMore,
}

// Format29V2 is the definition of a message at at MessageFormat 29 for ADSB V2
var Format29V2 = adsbFormatReferenceDefinition{
	typeCode:  29,
	register:  bds.BDS62,
	adsbLevel: Level2,
}

// -----------------------------  FORMAT 31  -------------------------------

// Format31V0 is the definition of a message at at MessageFormat 31 for ADSB V0
var Format31V0 = adsbFormatReferenceDefinition{
	typeCode:  31,
	register:  bds.BDS65,
	adsbLevel: Level0Exactly,
}

// Format31V1 is the definition of a message at at MessageFormat 31 for ADSB V1
var Format31V1 = adsbFormatReferenceDefinition{
	typeCode:  31,
	register:  bds.BDS65,
	adsbLevel: Level1Exactly,
}

// Format31V2 is the definition of a message at at MessageFormat 31 for ADSB V2
var Format31V2 = adsbFormatReferenceDefinition{
	typeCode:  9,
	register:  bds.BDS65,
	adsbLevel: Level2,
}
