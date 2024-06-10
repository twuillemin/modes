package bds06

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// SurfacePosition is a message at the format BDS 0,6
type SurfacePosition struct {
	FormatTypeCode    byte
	Movement          fields.Movement
	GroundTrackStatus bool
	GroundTrack       float32
	Time              fields.Time
	CPRFormat         fields.CompactPositionReportingFormat
	EncodedLatitude   fields.EncodedLatitude
	EncodedLongitude  fields.EncodedLongitude
}

// GetRegister returns the Register the message
func (message SurfacePosition) GetRegister() register.Register {
	return register.BDS06
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message SurfacePosition) CheckCoherency() error {
	return nil
}

func (message SurfacePosition) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Format Type Code:                  %v\n"+
		"Mevement:                          %v\n"+
		"Ground Track Status:               %v\n"+
		"Ground Track (degrees):            %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetRegister().ToString(),
		message.FormatTypeCode,
		message.Movement.ToString(),
		message.GroundTrackStatus,
		message.GroundTrack,
		message.Time.ToString(),
		message.CPRFormat.ToString(),
		message.EncodedLatitude,
		message.EncodedLongitude)
}

// ReadSurfacePosition reads a message at the format Format09V1
func ReadSurfacePosition(data []byte) (*SurfacePosition, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode < 5 || formatTypeCode > 8 {
		return nil, fmt.Errorf("the field FormatTypeCode must be comprised between 5 and 8 included, got %v", formatTypeCode)
	}

	groundTrack, groundTrackStatus := fields.ReadGroundTrack(data)

	return &SurfacePosition{
		FormatTypeCode:    formatTypeCode,
		Movement:          fields.ReadMovement(data),
		GroundTrackStatus: groundTrackStatus,
		GroundTrack:       groundTrack,
		Time:              fields.ReadTime(data),
		CPRFormat:         fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:   fields.ReadEncodedLatitude(data),
		EncodedLongitude:  fields.ReadEncodedLongitude(data),
	}, nil
}
