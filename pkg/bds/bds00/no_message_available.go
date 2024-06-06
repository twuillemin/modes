package bds00

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/register"
)

// NoMessageAvailable is a message at the format BDS 0,0
//
// From Spec: If a command to transmit an air-initiated Comm-B message is received while no message is waiting to be
// transmitted, the reply shall contain all ZEROs in the MB field
type NoMessageAvailable struct {
}

// GetRegister returns the Register the message
func (message NoMessageAvailable) GetRegister() register.Register {
	return register.BDS00
}

// ToString returns a basic, but readable, representation of the message
func (message NoMessageAvailable) ToString() string {
	return fmt.Sprintf(""+
		"Message:                   %v",
		message.GetRegister().ToString())
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message NoMessageAvailable) CheckCoherency() error {
	return nil
}

// ReadNoMessageAvailable reads a message as a NoMessageAvailable
func ReadNoMessageAvailable(data []byte) (*NoMessageAvailable, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B NoMessageAvailable message must be 7 bytes long")
	}

	for i := uint(0); i < 7; i++ {
		if data[i] != 0 {
			return nil, errors.New("bits 1 to 56 must be zero")
		}
	}

	return &NoMessageAvailable{}, nil
}
