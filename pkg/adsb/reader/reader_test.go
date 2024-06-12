package reader

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/airborne_position"
	"github.com/twuillemin/modes/pkg/adsb/no_position_information"
	"testing"
)

func TestReadADSB(t *testing.T) {

	noPositionMessage, _ := hex.DecodeString("00676000000000")
	msg, err := ReadADSBMessage(adsb.ADSBV0, false, false, noPositionMessage)
	if err != nil {
		t.Errorf("Unable to read a no position information message")
	} else {
		_, ok := msg.(*no_position_information.NoPositionInformation)
		if !ok {
			t.Errorf("the no position information message was not read as a NoPositionInformation")
		}
	}

	airbornePositionType0, _ := hex.DecodeString("00A3C22B7776D7")
	msg, err = ReadADSBMessage(adsb.ADSBV0, false, false, airbornePositionType0)
	if err != nil {
		t.Errorf("Unable to read an airborne position type 0")
	} else {
		_, ok := msg.(*airborne_position.AirbornePositionType0)
		if !ok {
			t.Errorf("the airborne position type 0 message was not read as a AirbornePositionType0")
		}
	}
}
