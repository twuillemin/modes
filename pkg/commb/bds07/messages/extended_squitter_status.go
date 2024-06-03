package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/commb"
	"github.com/twuillemin/modes/pkg/commb/bds07/fields"
)

// ExtendedSquitterStatus is a message at the format BDS 0,7
//
// Specified in Doc 9871 / D.2.4.1
type ExtendedSquitterStatus struct {
	TransmissionRateSubfield fields.TransmissionRateSubfield
	AltitudeTypeSubfield     fields.AltitudeTypeSubfield
}

// GetMessageFormat returns the register of the message
func (message ExtendedSquitterStatus) GetMessageFormat() commb.MessageFormat {
	return commb.FormatExtendedSquitterStatus
}

// ToString returns a basic, but readable, representation of the message
func (message ExtendedSquitterStatus) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Transmission Rate Subfield:              %v\n"+
		"Altitude Type Subfield                   %v\n",
		commb.GetMessageFormatInformation(&message),
		message.TransmissionRateSubfield.ToString(),
		message.AltitudeTypeSubfield.ToString())
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message ExtendedSquitterStatus) CheckCoherency() error {
	if message.TransmissionRateSubfield >= 3 {
		return errors.New("field TransmissionRateSubfield is a Reserved value")
	}

	return nil
}

// ReadExtendedSquitterStatus reads a message as a ExtendedSquitterStatus
func ReadExtendedSquitterStatus(data []byte) (*ExtendedSquitterStatus, error) {
	err := CheckIfDataReadable(data)
	if err != nil {
		return nil, err
	}

	return &ExtendedSquitterStatus{
		TransmissionRateSubfield: fields.ReadTransmissionRateSubfield(data),
		AltitudeTypeSubfield:     fields.ReadAltitudeTypeSubfield(data),
	}, nil
}

// CheckIfDataReadable checks if the data can be read as a ExtendedSquitterStatus
func CheckIfDataReadable(data []byte) error {
	if len(data) != 7 {
		return errors.New("the data for Comm-B ExtendedSquitterStatus message must be 7 bytes long")
	}

	// First byte is simply the BDS format 0001 0000
	if data[0] != 0x10 {
		return errors.New("the first byte of data is not 0x10")
	}

	// Bits 4 to 8 are reserved and must be 0
	if data[0]&0x1C != 0 {
		return errors.New("the bits 4 to 8 are reserved and must be 0")
	}

	for i := uint32(1); i < 7; i++ {
		if data[i] != 0 {
			return errors.New("the bits 9 to 56 are reserved and must be 0")
		}
	}

	return nil
}
