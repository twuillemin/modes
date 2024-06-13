package bds65

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAircraftOperationalStatusAirborneADSB2(t *testing.T) {

	id1, _ := hex.DecodeString("F8230003004008")
	message, err := ReadAircraftOperationalStatusAirborneADSB2(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS65 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS65)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.AirborneCapabilityClass.ACASOperational != fields.AOOperational {
		t.Errorf("for AirborneCapabilityClass.ACASOperational: got %v, want %v", message.AirborneCapabilityClass.ACASOperational, fields.AOOperational)
	}

	if message.AirborneCapabilityClass.ExtendedSquitterIn != fields.ESINoCapability {
		t.Errorf("for AirborneCapabilityClass.ExtendedSquitterIn: got %v, want %v", message.AirborneCapabilityClass.ExtendedSquitterIn, fields.ESINoCapability)
	}

	if message.AirborneCapabilityClass.AirReferencedVelocityReportCapability != fields.ARVCapable {
		t.Errorf("for AirborneCapabilityClass.AirReferencedVelocityReportCapability: got %v, want %v", message.AirborneCapabilityClass.AirReferencedVelocityReportCapability, fields.ARVCapable)
	}

	if message.AirborneCapabilityClass.TargetStateReportCapability != fields.TSRCCapable {
		t.Errorf("for AirborneCapabilityClass.TargetStateReportCapability: got %v, want %v", message.AirborneCapabilityClass.TargetStateReportCapability, fields.TSRCCapable)
	}

	if message.AirborneCapabilityClass.TargetChangeReportCapability != fields.TCRCNoCapability {
		t.Errorf("for AirborneCapabilityClass.TargetChangeReportCapability: got %v, want %v", message.AirborneCapabilityClass.TargetChangeReportCapability, fields.TCRCNoCapability)
	}

	if message.AirborneCapabilityClass.UniversalAccessTransceiverCapability != fields.UATNoCapability {
		t.Errorf("for AirborneCapabilityClass.UniversalAccessTransceiverCapability: got %v, want %v", message.AirborneCapabilityClass.UniversalAccessTransceiverCapability, fields.UATNoCapability)
	}

	if message.OperationalMode.ACASResolutionAdvisoryActive != fields.ARAV2NotActive {
		t.Errorf("for OperationalMode.ACASResolutionAdvisoryActive: got %v, want %v", message.OperationalMode.ACASResolutionAdvisoryActive, fields.ARAV2NotActive)
	}

	if message.OperationalMode.IdentSwitchActive != fields.ISANotActive {
		t.Errorf("for OperationalMode.IdentSwitchActive: got %v, want %v", message.OperationalMode.IdentSwitchActive, fields.ISANotActive)
	}

	if message.OperationalMode.SingleAntennaFlag != fields.SAFTwoAntennas {
		t.Errorf("for OperationalMode.SingleAntennaFlag: got %v, want %v", message.OperationalMode.SingleAntennaFlag, fields.SAFTwoAntennas)
	}

	if message.OperationalMode.SystemDesignAssurance != fields.SDALevelB {
		t.Errorf("for OperationalMode.SystemDesignAssurance: got %v, want %v", message.OperationalMode.SystemDesignAssurance, fields.SDALevelB)
	}

	if message.NICSupplementA != fields.NICAZero {
		t.Errorf("for NICSupplementA: got %v, want %v", message.NICSupplementA, fields.NICAZero)
	}

	if message.NavigationalAccuracyCategoryPosition != fields.NACV2PEPUGreaterThan18Point52Km {
		t.Errorf("for NavigationalAccuracyCategoryPosition: got %v, want %v", message.NavigationalAccuracyCategoryPosition, fields.NACV2PEPUGreaterThan18Point52Km)
	}

	if message.GeometricVerticalAccuracy != fields.GVAUnknownOrGreaterThan150m {
		t.Errorf("for GeometricVerticalAccuracy: got %v, want %v", message.GeometricVerticalAccuracy, fields.GVAUnknownOrGreaterThan150m)
	}

	if message.SourceIntegrityLevel != fields.SILLevel0 {
		t.Errorf("for SourceIntegrityLevel: got %v, want %v", message.SourceIntegrityLevel, fields.SILLevel0)
	}

	if message.NICBaro != fields.NICBGilhamCrossCheckedOrNonGilham {
		t.Errorf("for NICBaro: got %v, want %v", message.NICBaro, fields.NICBGilhamCrossCheckedOrNonGilham)
	}

	if message.HorizontalReferenceDirection != fields.HRDTrueNorth {
		t.Errorf("for HorizontalReferenceDirection: got %v, want %v", message.HorizontalReferenceDirection, fields.HRDTrueNorth)
	}

	if message.SourceIntegrityLevelSupplement != fields.SILSByHour {
		t.Errorf("for SourceIntegrityLevelSupplement: got %v, want %v", message.SourceIntegrityLevelSupplement, fields.SILSByHour)
	}
}
