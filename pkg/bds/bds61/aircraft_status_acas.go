package bds61

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/acas/ra"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AircraftStatusACAS is a message at the format BDS 6,1
type AircraftStatusACAS struct {
	FormatTypeCode     byte
	Subtype            fields.Subtype
	ResolutionAdvisory ra.ResolutionAdvisory
}

func (message AircraftStatusACAS) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AircraftStatusACAS) GetRegister() register.Register {
	return register.BDS61
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftStatusACAS) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftStatusACAS) ToString() string {
	return fmt.Sprintf(""+
		"Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"ACAS Data:                 \n%v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.ResolutionAdvisory.ToString())
}

// ReadAircraftStatusACAS reads a AircraftStatus / Subtype 2 (ACAS RA Broadcast)
func ReadAircraftStatusACAS(data []byte) (*AircraftStatusACAS, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 28, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeRABroadcast {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftStatusACAS", subType.ToString())
	}

	resolutionAdvisory, err := ra.ReadResolutionAdvisory(data[1:])
	if err != nil {
		return nil, err
	}

	return &AircraftStatusACAS{
		FormatTypeCode:     formatTypeCode,
		Subtype:            subType,
		ResolutionAdvisory: *resolutionAdvisory,
	}, nil
}
