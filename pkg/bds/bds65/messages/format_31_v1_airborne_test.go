package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"testing"
)

func TestReadFormat31V1Airborne(t *testing.T) {

	msg, err := ReadFormat31V1Airborne(buildValidFormat31V1AirborneMessage())
	if err != nil {
		t.Error(err)
	}

	if msg.GetMessageFormat() != adsb.Format31V1 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format31V1.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS65.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS65.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.SubtypeV1Airborne {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeV1Airborne.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.AirborneCapabilityClass.NotACASStatus != fields.NotACASStatusNotInstalledNotOperational {
		t.Errorf("Expected Not ACAS Status \"%v\", got \"%v\"",
			fields.NotACASStatusNotInstalledNotOperational.ToString(),
			msg.AirborneCapabilityClass.NotACASStatus.ToString())
	}

	if msg.AirborneCapabilityClass.CockpitDisplayOfTrafficInformationStatus != fields.CDTIOperational {
		t.Errorf("Expected CDTI \"%v\", got \"%v\"",
			fields.CDTIOperational.ToString(),
			msg.AirborneCapabilityClass.CockpitDisplayOfTrafficInformationStatus.ToString())
	}

	if msg.AirborneCapabilityClass.AirReferencedVelocityReportCapability != fields.ARVCapable {
		t.Errorf("Expected ARV \"%v\", got \"%v\"",
			fields.ARVCapable.ToString(),
			msg.AirborneCapabilityClass.AirReferencedVelocityReportCapability.ToString())
	}

	if msg.AirborneCapabilityClass.TargetStateReportCapability != fields.TSRCCapable {
		t.Errorf("Expected TSR \"%v\", got \"%v\"",
			fields.TSRCCapable.ToString(),
			msg.AirborneCapabilityClass.TargetStateReportCapability.ToString())
	}

	if msg.AirborneCapabilityClass.TargetChangeReportCapability != fields.TCRCCapableMultipleTC {
		t.Errorf("Expected TCR \"%v\", got \"%v\"",
			fields.TCRCCapableMultipleTC.ToString(),
			msg.AirborneCapabilityClass.TargetChangeReportCapability.ToString())
	}

	if msg.OperationalMode.ACASResolutionAdvisoryActive != fields.ARAV1Active {
		t.Errorf("Expected ACAS RA Active \"%v\", got \"%v\"",
			fields.ARAV1Active.ToString(),
			msg.OperationalMode.ACASResolutionAdvisoryActive.ToString())
	}

	if msg.OperationalMode.IdentSwitchActive != fields.ISAActive {
		t.Errorf("Expected Ident Switch Active \"%v\", got \"%v\"",
			fields.ISAActive.ToString(),
			msg.OperationalMode.IdentSwitchActive.ToString())
	}

	if msg.OperationalMode.ReceivingATCServices != fields.RASReceivingATC {
		t.Errorf("Expected Receiving ATC Active \"%v\", got \"%v\"",
			fields.RASReceivingATC.ToString(),
			msg.OperationalMode.ReceivingATCServices.ToString())
	}

	if msg.NICSupplement != fields.NICAOne {
		t.Errorf("Expected NIC Supplement \"%v\", got \"%v\"",
			fields.NICAOne.ToString(),
			msg.NICSupplement.ToString())
	}

	if msg.NavigationalAccuracyCategoryPosition != fields.NACV1PEPULowerThan10MAndVEPULowerThan15M {
		t.Errorf("Expected NACp \"%v\", got \"%v\"",
			fields.NACV1PEPULowerThan10MAndVEPULowerThan15M.ToString(),
			msg.NavigationalAccuracyCategoryPosition.ToString())
	}

	if msg.SurveillanceIntegrityLevel != fields.SUILLevel3 {
		t.Errorf("Expected SIL \"%v\", got \"%v\"",
			fields.SUILLevel3.ToString(),
			msg.NavigationalAccuracyCategoryPosition.ToString())
	}

	if msg.NICBaro != fields.NICBGilhamCrossCheckedOrNonGilham {
		t.Errorf("Expected NICbaro \"%v\", got \"%v\"",
			fields.NICBGilhamCrossCheckedOrNonGilham.ToString(),
			msg.NICBaro.ToString())
	}

	if msg.HorizontalReferenceDirection != fields.HRDMagneticNorth {
		t.Errorf("Expected Horizontal Reference Direction \"%v\", got \"%v\"",
			fields.HRDMagneticNorth.ToString(),
			msg.HorizontalReferenceDirection.ToString())
	}
}

func buildValidFormat31V1AirborneMessage() []byte {
	data := make([]byte, 7)

	// 1111 1000: code 31 (11111) + subtype 0 (000)
	data[0] = 0xF8

	// 0011 0011: ServiceLevel (00) + ACAS not installed (1) + Traffic display operational (1) + ServiceLevel (00) +
	// Capability of sending velocity messages (1) + Capability of sending targetState messages (1)
	data[1] = 0x33

	// 1000 0000: multiple TC Reports (10)+Reserved (000000)
	data[2] = 0x80

	// 0011 1000: OM Format (00) + ACAS RA Active (1) + Ident active (1) + Receiving ATC (1) + Reserved (000)
	data[3] = 0x38

	// 0000 0000:  Reserved (0000 0000)
	data[4] = 0x00

	// 0011 1010: ADSB Version (001) + NIC Supplement (1) + EPU < 10 m and VEPU < 15 m (1010)
	data[5] = 0x3A

	// 0011 1100: BAQ (00) + SIL 3 (11) + NICbaro verified (1) + Magnetic North (1) + Reserved (00)
	data[6] = 0x3C

	return data
}
