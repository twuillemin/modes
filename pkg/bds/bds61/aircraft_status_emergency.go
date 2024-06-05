package bds61

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// AircraftStatusEmergency is a message at the format BDS 6,1
type AircraftStatusEmergency struct {
	FormatTypeCode          byte
	Subtype                 fields.Subtype
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
	ModeACode               uint16
}

func (message AircraftStatusEmergency) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AircraftStatusEmergency) GetRegister() register.Register {
	return register.BDS61
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftStatusEmergency) CheckCoherency() error {

	if message.EmergencyPriorityStatus > 5 {
		return errors.New("field EmergencyPriorityStatus is a Reserved value")
	}

	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftStatusEmergency) ToString() string {
	return fmt.Sprintf(""+
		"Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergeny State:            %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.EmergencyPriorityStatus.ToString())
}

// ReadAircraftStatusEmergency reads a AircraftStatus / Subtype 1 (Emergency/priority status)
func ReadAircraftStatusEmergency(data []byte) (*AircraftStatusEmergency, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 28, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeEmergencyPriorityStatus {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftStatusEmergency", subType.ToString())
	}

	byte1 := data[1] & 0x1F
	byte2 := data[2]

	for i := uint32(3); i < 7; i++ {
		if (data[1] & 0x1F) != 0 {
			return nil, errors.New("the bits 25 to 56 are expected to be 0")
		}
	}

	return &AircraftStatusEmergency{
		FormatTypeCode:          formatTypeCode,
		Subtype:                 subType,
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
		ModeACode:               bitutils.Pack2Bytes(byte1, byte2),
	}, nil
}
