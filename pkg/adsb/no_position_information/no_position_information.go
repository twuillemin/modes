package no_position_information

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds05"
)

type NoPositionInformation struct {
	adsbVersion                 adsb.ADSBVersion
	FormatTypeCode              byte
	AltitudeBarometric          int
	NavigationIntegrityCategory byte
}

// GetADSBVersion returns the ADSB level used to read the data
func (message NoPositionInformation) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message NoPositionInformation) CheckCoherency() error {
	return nil
}

func (message NoPositionInformation) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           No Position Information\n"+
		"Format Type Code:                  %v\n"+
		"Altitude Baro (feet):              %v\n"+
		"Navigation Integrity Category:     %v",
		message.FormatTypeCode,
		message.AltitudeBarometric,
		message.NavigationIntegrityCategory)
}

// ReadNoPositionInformation reads a message at the format NoPositionInformation
func ReadNoPositionInformation(adsbVersion adsb.ADSBVersion, data []byte) (*NoPositionInformation, error) {

	// Update the first byte to have a standard BDS code 0,5 (Airborne position with baro altitude) - format 15
	dataTmp := []byte{0x78, data[1], data[2], 0, 0, 0, 0}

	bds, err := bds05.ReadAirbornePosition(dataTmp)
	if err != nil {
		return nil, err
	}

	return &NoPositionInformation{
		FormatTypeCode:              0,
		AltitudeBarometric:          bds.AltitudeInFeet,
		NavigationIntegrityCategory: 0,
		adsbVersion:                 adsbVersion,
	}, nil
}
