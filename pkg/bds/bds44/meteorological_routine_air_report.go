package bds44

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds44/fields"
)

type MeteorologicalRoutineAirReport interface {
	bds.Message
	GetSource() fields.Source
}

// ReadMeteorologicalRoutineAirReport reads a message at the format MeteorologicalRoutineAirReport
func ReadMeteorologicalRoutineAirReport(data []byte) (MeteorologicalRoutineAirReport, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	source := fields.ReadSource(data)
	switch source {
	case fields.SourceINS, fields.SourceGNSS, fields.SourceDMEDME, fields.SourceVORDME:
		return ReadMeteorologicalRoutineAirReportV2(data)
	case fields.SourceInvalid:
		return ReadMeteorologicalRoutineAirReportV1(data)

	default:
		return nil, fmt.Errorf("the field SubType must be comprised between 0 and 1 included, got %v", source)
	}
}
