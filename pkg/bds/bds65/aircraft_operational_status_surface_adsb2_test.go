package bds65

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAircraftOperationalStatusSurfaceADSB2(t *testing.T) {

	id1, _ := hex.DecodeString("F9002102814A3C")
	message, err := ReadAircraftOperationalStatusSurfaceADSB2(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS65 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS65)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.SurfaceCapabilityClass.ExtendedSquitterIn != fields.ESINoCapability {
		t.Errorf("for SurfaceCapabilityClass.ExtendedSquitterIn: got %v, want %v", message.SurfaceCapabilityClass.ExtendedSquitterIn, fields.ESINoCapability)
	}

	if message.SurfaceCapabilityClass.B2Low != fields.B2LGreaterThan70W {
		t.Errorf("for SurfaceCapabilityClass.B2Low: got %v, want %v", message.SurfaceCapabilityClass.B2Low, fields.B2LGreaterThan70W)
	}

	if message.SurfaceCapabilityClass.UniversalAccessTransceiverCapability != fields.UATNoCapability {
		t.Errorf("for SurfaceCapabilityClass.UniversalAccessTransceiverCapability: got %v, want %v", message.SurfaceCapabilityClass.UniversalAccessTransceiverCapability, fields.UATNoCapability)
	}

	if message.SurfaceCapabilityClass.NavigationAccuracyCategoryVelocity != 0x01 {
		t.Errorf("for SurfaceCapabilityClass.NavigationAccuracyCategoryVelocity: got %v, want %v", message.SurfaceCapabilityClass.NavigationAccuracyCategoryVelocity, 0x01)
	}

	if message.SurfaceCapabilityClass.NICSupplementC != fields.NICCSZero {
		t.Errorf("for SurfaceCapabilityClass.NICSupplementC: got %v, want %v", message.SurfaceCapabilityClass.NICSupplementC, fields.NICCSZero)
	}

	if message.LengthAndWidth != fields.LWLength15Width23 {
		t.Errorf("for LengthAndWidth: got %v, want %v", message.LengthAndWidth, fields.LWLength15Width23)
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

	if message.OperationalMode.GPSAntenna.GPSAntennaLateral != 0x04 {
		t.Errorf("for OperationalMode.GPSAntenna.GPSAntennaLateral: got %v, want %v", message.OperationalMode.GPSAntenna.GPSAntennaLateral, 0x04)
	}

	if message.OperationalMode.GPSAntenna.GPSAntennaLongitudinal != 0x01 {
		t.Errorf("for OperationalMode.GPSAntenna.GPSAntennaLongitudinal: got %v, want %v", message.OperationalMode.GPSAntenna.GPSAntennaLongitudinal, 0x01)
	}

	if message.NICSupplementA != fields.NICAZero {
		t.Errorf("for NICSupplementA: got %v, want %v", message.NICSupplementA, fields.NICAZero)
	}

	if message.NavigationalAccuracyCategoryPosition != fields.NACV2PEPULowerThan10M {
		t.Errorf("for NavigationalAccuracyCategoryPosition: got %v, want %v", message.NavigationalAccuracyCategoryPosition, fields.NACV2PEPULowerThan10M)
	}

	if message.SourceIntegrityLevel != fields.SILLevel3 {
		t.Errorf("for SourceIntegrityLevel: got %v, want %v", message.SourceIntegrityLevel, fields.SILLevel3)
	}

	if message.TrackAngleHeading != fields.TAHTrackAngleReported {
		t.Errorf("for TrackAngleHeading: got %v, want %v", message.TrackAngleHeading, fields.TAHTrackAngleReported)
	}

	if message.HorizontalReferenceDirection != fields.HRDMagneticNorth {
		t.Errorf("for HorizontalReferenceDirection: got %v, want %v", message.HorizontalReferenceDirection, fields.HRDMagneticNorth)
	}

	if message.SourceIntegrityLevelSupplement != fields.SILSByHour {
		t.Errorf("for SourceIntegrityLevelSupplement: got %v, want %v", message.SourceIntegrityLevelSupplement, fields.SILSByHour)
	}
}
