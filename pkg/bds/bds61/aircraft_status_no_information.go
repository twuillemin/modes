package bds61

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AircraftStatusNoInformation is a message at the format BDS 6,1
type AircraftStatusNoInformation struct {
	FormatTypeCode          byte
	Subtype                 fields.Subtype
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

func (message AircraftStatusNoInformation) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AircraftStatusNoInformation) GetRegister() register.Register {
	return register.BDS61
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftStatusNoInformation) CheckCoherency() error {

	if message.EmergencyPriorityStatus != fields.EPSNoEmergency {
		return errors.New("field EmergencyPriorityStatus is expected to be NoEmergency")
	}

	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftStatusNoInformation) ToString() string {
	return fmt.Sprintf(""+
		"Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency State:           %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.EmergencyPriorityStatus.ToString())
}

// ReadAircraftStatusNoInformation reads a AircraftStatus / subtype 0 (No information)
func ReadAircraftStatusNoInformation(data []byte) (*AircraftStatusNoInformation, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 28, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeNoInformation {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftStatusNoInformation", subType.ToString())
	}

	if (data[1] & 0x1F) != 0 {
		return nil, errors.New("the bits 12 to 16 are expected to be 0")
	}

	for i := uint32(2); i < 7; i++ {
		if (data[1] & 0x1F) != 0 {
			return nil, errors.New("the bits 17 to 56 are expected to be 0")
		}
	}

	return &AircraftStatusNoInformation{
		FormatTypeCode:          formatTypeCode,
		Subtype:                 subType,
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
