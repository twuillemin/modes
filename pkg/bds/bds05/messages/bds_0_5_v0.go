package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05V0 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB V0 are expected to implement
type MessageBDS05V0 interface {
	MessageBDS05
	GetSingleAntennaFlag() fields.SingleAntennaFlag
	// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
	GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit
	// GetContainmentRadius returns the ContainmentRadius
	GetContainmentRadius() fields.ContainmentRadius
}

func bds05v0ToString(message MessageBDS05V0) string {
	return fmt.Sprintf("Message:                           %v\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v\n"+
		"Single Antenna:                    %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Altitude:                          %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetMessageFormat().ToString(),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalProtectionLimit().ToString(),
		message.GetContainmentRadius().ToString(),
		message.GetSingleAntennaFlag().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

var hplBarometricByFormat = map[byte]fields.HorizontalProtectionLimitBarometric{
	9:  fields.HPLBLowerThan7Dot5M,
	10: fields.HPLBBetween7Dot5MAnd25M,
	11: fields.HPLBBetween25MAnd185Dot2M,
	12: fields.HPLBBetween185Dot2MAnd370Dot4M,
	13: fields.HPLBBetween370Dot4MAnd926M,
	14: fields.HPLBBetween926MAnd1852M,
	15: fields.HPLBBetween1852MAnd3704M,
	16: fields.HPLBBetween3704MAnd18Point52Km,
	17: fields.HPLBBetween18Point52KmAnd37Point04Km,
	18: fields.HPLBLargerThan37Point04Km,
}

var crBarometricByFormat = map[byte]fields.ContainmentRadiusBarometric{
	9:  fields.CRBLowerThan3M,
	10: fields.CRBBetween3MAnd10M,
	11: fields.CRBBetween10MAnd92Dot6M,
	12: fields.CRBBetween92Dot6MAnd185Dot2M,
	13: fields.CRBBetween185Dot2MAnd463M,
	14: fields.CRBBetween463MAnd926M,
	15: fields.CRBBetween926MAnd1852M,
	16: fields.CRBBetween1Point852KmAnd9Point26Km,
	17: fields.CRBBetween9Point26KmAnd18Point52Km,
	18: fields.CRBLargerThan18Point52Km,
}

var hplGNSSByFormat = map[byte]fields.HorizontalProtectionLimitGNSS{
	20: fields.HPLGLowerThan7Dot5M,
	21: fields.HPLGLowerThan25M,
	22: fields.HPLGGreaterThan25M,
}

var crGNSSByFormat = map[byte]fields.ContainmentRadiusGNSS{
	20: fields.CRGHorizontalLowerThan3MAndVerticalLowerThan4M,
	21: fields.CRGHorizontalLowerThan10MAndVerticalLowerThan15M,
	22: fields.CRGHorizontalGreaterThan10MOrVerticalGreaterThan15M,
}
