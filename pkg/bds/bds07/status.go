package bds07

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds07/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// Status is a message at the format BDS 0,7
//
// Specified in Doc 9871 / D.2.4.1
type Status struct {
	TransmissionRateSubfield fields.TransmissionRateSubfield
	AltitudeTypeSubfield     fields.AltitudeTypeSubfield
}

// GetRegister returns the Register the message
func (message Status) GetRegister() register.Register {
	return register.BDS07
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message Status) CheckCoherency() error {
	if message.TransmissionRateSubfield >= 3 {
		return errors.New("field TransmissionRateSubfield is a Reserved value")
	}

	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message Status) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Transmission Rate Subfield:              %v\n"+
		"Altitude Type Subfield                   %v\n",
		message.GetRegister().ToString(),
		message.TransmissionRateSubfield.ToString(),
		message.AltitudeTypeSubfield.ToString())
}

// ReadStatus reads a message as a Status
func ReadStatus(data []byte) (*Status, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B Status message must be 7 bytes long")
	}

	// Bits 4 to 8 are reserved and must be 0
	if data[0]&0x1C != 0 {
		return nil, errors.New("the bits 4 to 8 are reserved and must be 0")
	}

	for i := uint(1); i < 7; i++ {
		if data[i] != 0 {
			return nil, errors.New("the bits 9 to 56 are reserved and must be 0")
		}
	}

	return &Status{
		TransmissionRateSubfield: fields.ReadTransmissionRateSubfield(data),
		AltitudeTypeSubfield:     fields.ReadAltitudeTypeSubfield(data),
	}, nil
}
