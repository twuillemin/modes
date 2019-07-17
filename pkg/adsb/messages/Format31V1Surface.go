package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format31V1Surface is a message at the format BDS 6,5 the ADSB V1 / Surface
type Format31V1Surface struct {
	SurfaceCapabilityClass fields.CapabilityClassSurface
	OperationalMode        fields.OperationalMode
	VersionNumber          fields.VersionNumber
}

// GetName returns the name of the message
func (message *Format31V1Surface) GetName() string {
	return bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V1Surface) GetBDS() string {
	return bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V1Surface) GetFormatTypeCode() byte {
	return 31
}

// GetOperationalStatusSubTypeCode returns the code of the Operational Status Sub Type
func (message *Format31V1Surface) GetOperationalStatusSubTypeCode() byte {
	return 1
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Surface) ToString() string {
	return fmt.Sprintf("Message:          %v (%v)\n"+
		"SubType:          1 - Surface\n"+
		"VersionNumber:    %v\n"+
		"SurfaceCapabilityClass:\n%v\n"+
		"OperationalMode:\n%v",
		message.GetBDS(),
		message.GetName(),
		message.VersionNumber.ToString(),
		message.SurfaceCapabilityClass.ToString(),
		message.OperationalMode.ToString())
}

// ReadFormat31V1Surface reads a message at the format Format31V1Surface
func ReadFormat31V1Surface(data []byte) (*Format31V1Surface, error) {

	return &Format31V1Surface{
		SurfaceCapabilityClass: fields.ReadCapabilityClassSurface(data),
		OperationalMode:        fields.ReadOperationalMode(data),
		VersionNumber:          fields.ReadVersionNumber(data),
	}, nil
}
