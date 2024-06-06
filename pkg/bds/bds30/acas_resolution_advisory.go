package bds30

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/acas/ra"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// ACASResolutionAdvisory is a message at the format BDS 3,0
//
// Specified in Doc 9871 / Table A-2-48
type ACASResolutionAdvisory struct {
	ResolutionAdvisory ra.ResolutionAdvisory
}

// GetRegister returns the Register the message
func (message ACASResolutionAdvisory) GetRegister() register.Register {
	return register.BDS30
}

// ToString returns a basic, but readable, representation of the message
func (message ACASResolutionAdvisory) ToString() string {
	return fmt.Sprintf(""+
		"Message:                   %v\n"+
		"ACAS Data:                 \n%v",
		message.GetRegister().ToString(),
		message.ResolutionAdvisory.ToString())
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message ACASResolutionAdvisory) CheckCoherency() error {
	return nil
}

// ReadACASResolutionAdvisory reads a message as a ACASResolutionAdvisory
func ReadACASResolutionAdvisory(data []byte) (*ACASResolutionAdvisory, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B ACASResolutionAdvisory message must be 7 bytes long")
	}

	// First byte is simply the BDS format 0011 0000
	if data[0] != 0x30 {
		return nil, errors.New("the first byte of data is not 0x20")
	}

	resolutionAdvisory, err := ra.ReadResolutionAdvisory(data[1:])
	if err != nil {
		return nil, err
	}

	return &ACASResolutionAdvisory{
		ResolutionAdvisory: *resolutionAdvisory,
	}, nil
}
