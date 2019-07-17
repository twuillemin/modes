package messages

import "fmt"

// ACASUnknown is a generic container for unknown ACAS messages
type ACASUnknown struct {
	VDS1 byte
	VDS2 byte
	Data []byte
}

// ReadMessageACASUnknown reads a message of an unknown message
//
// Params:
//    - message: The content of the message including the field VDS.
//
// Returns an ACASUnknown message
func ReadMessageACASUnknown(message []byte) ACASMessage {

	vds1 := (message[0] & 0xF0) >> 4
	vds2 := message[0] & 0x0F

	return ACASUnknown{
		VDS1: vds1,
		VDS2: vds2,
		Data: message,
	}
}

// GetName returns the name of the message
func (messageACAS ACASUnknown) GetName() string {
	return "Coordination reply"
}

// GetVDS1 returns the VDS1
func (messageACAS ACASUnknown) GetVDS1() byte {
	return messageACAS.VDS1
}

// GetVDS2 returns the VDS1
func (messageACAS ACASUnknown) GetVDS2() byte {
	return messageACAS.VDS2
}

// ToString returns a basic, but readable, representation of the field
func (messageACAS ACASUnknown) ToString() string {
	return fmt.Sprintf("VDS 1:   %02X\n"+
		"VDS 2:   %02X\n"+
		"Content: %v",
		messageACAS.VDS1,
		messageACAS.VDS2,
		messageACAS.Data)
}
