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

	// There are three formats possible for BDS 4,4:
	// - the V0 (defined in annex A) which has a Source defined in the first bits and a specific structure
	// - the V1 (defined annex E) which has a 0 for the first bits (so invalid source as of V0) and another structure
	// - the V2 which is an undocumented mix comprising the Source from V0 with the structure of V1. However, this is
	//   the format that was found in real world data.

	source := fields.ReadSource(data)
	switch source {
	// As a Source is defined, it may be a V2 or V0 format. So try the V2 first and if it cannot be read, try the V0.
	case fields.SourceINS, fields.SourceGNSS, fields.SourceDMEDME, fields.SourceVORDME:
		message, err := ReadMeteorologicalRoutineAirReportV2(data)
		if err == nil {
			return message, nil
		}
		return ReadMeteorologicalRoutineAirReportV0(data)

	// The Source is zeroed, so try the V1 format.
	case fields.SourceInvalid:
		return ReadMeteorologicalRoutineAirReportV1(data)

	default:
		return nil, fmt.Errorf("the field SubType must be comprised between 0 and 1 included, got %v", source)
	}
}
