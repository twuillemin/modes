package target_state_and_status

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds62"
)

type TargetStateAndStatus struct {
	bds62.TargetStateAndStatus
	adsbVersion adsb.ADSBVersion
}

// GetADSBVersion returns the ADSB level used to read the data
func (message TargetStateAndStatus) GetADSBVersion() adsb.ADSBVersion {
	return message.adsbVersion
}

func (message TargetStateAndStatus) ToString() string {
	return fmt.Sprintf(""+
		"%v",
		message.TargetStateAndStatus.ToString())
}

// ReadTargetStateAndStatus reads a message at the format TargetStateAndStatus
func ReadTargetStateAndStatus(adsbVersion adsb.ADSBVersion, data []byte) (*TargetStateAndStatus, error) {

	bds, err := bds62.ReadTargetStateAndStatus(data)
	if err != nil {
		return nil, err
	}

	return &TargetStateAndStatus{
		TargetStateAndStatus: bds,
		adsbVersion:          adsbVersion,
	}, nil
}
