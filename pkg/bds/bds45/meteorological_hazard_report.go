package bds45

import (
	"fmt"

	"github.com/twuillemin/modes/pkg/bds"
)

type MeteorologicalHazardReport interface {
	bds.Message
	// GetTurbulenceStatus is a simple function to discriminate against other packages
	GetTurbulenceStatus() bool
}

// ReadMeteorologicalHazardReport reads a message at the format MeteorologicalHazardReport
func ReadMeteorologicalHazardReport(data []byte) (MeteorologicalHazardReport, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	messageV1, err := ReadMeteorologicalHazardReportV1(data)
	if err == nil {
		return messageV1, err
	}

	messageV0, err := ReadMeteorologicalHazardReportV0(data)
	if err == nil {
		return messageV0, err
	}

	return nil, err
}
