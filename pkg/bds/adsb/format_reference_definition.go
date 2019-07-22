package adsb

//go:generate go run gen/gen_formats_reference.go

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

	// GetADSBLevel returns the ADSB version implementing the message
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

// GetADSBLevel returns the ADSB version implementing the format
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
