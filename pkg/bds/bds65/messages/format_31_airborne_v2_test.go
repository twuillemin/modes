package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"testing"
)

func TestReadFormat31AirborneV2Valid(t *testing.T) {

	msg, err := ReadFormat31AirborneV2(buildValidFormat31AirborneV2Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format31 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format31.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS65.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS65.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.SubtypeAirborne {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeAirborne.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.AirborneCapabilityClass.ACASOperational != fields.AOOperational {
		t.Errorf("Expected ACAS Operational \"%v\", got \"%v\"",
			fields.AOOperational.ToString(),
			msg.AirborneCapabilityClass.ACASOperational.ToString())
	}

	if msg.AirborneCapabilityClass.ExtendedSquitterIn != fields.ESICapable {
		t.Errorf("Expected Extended Squitter In \"%v\", got \"%v\"",
			fields.ESICapable.ToString(),
			msg.AirborneCapabilityClass.ExtendedSquitterIn.ToString())
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

	if msg.AirborneCapabilityClass.UniversalAccessTransceiverCapability != fields.UATCapable {
		t.Errorf("Expected UAT Capable \"%v\", got \"%v\"",
			fields.UATCapable.ToString(),
			msg.AirborneCapabilityClass.UniversalAccessTransceiverCapability.ToString())
	}

	if msg.OperationalMode.ACASResolutionAdvisoryActive != fields.ARAV2Active {
		t.Errorf("Expected ACAS RA Active \"%v\", got \"%v\"",
			fields.ARAV2Active.ToString(),
			msg.OperationalMode.ACASResolutionAdvisoryActive.ToString())
	}

	if msg.OperationalMode.IdentSwitchActive != fields.ISAActive {
		t.Errorf("Expected Ident Switch Active \"%v\", got \"%v\"",
			fields.ISAActive.ToString(),
			msg.OperationalMode.IdentSwitchActive.ToString())
	}

	if msg.OperationalMode.SingleAntennaFlag != fields.SAFOneAntenna {
		t.Errorf("Expected Single Antenna Flag \"%v\", got \"%v\"",
			fields.SAFOneAntenna.ToString(),
			msg.OperationalMode.SingleAntennaFlag.ToString())
	}

	if msg.OperationalMode.SystemDesignAssurance != fields.SDALevelB {
		t.Errorf("Expected System Design Assurance \"%v\", got \"%v\"",
			fields.SDALevelB.ToString(),
			msg.OperationalMode.SystemDesignAssurance.ToString())
	}

	if msg.NICSupplementA != fields.NICAOne {
		t.Errorf("Expected NIC Supplement \"%v\", got \"%v\"",
			fields.NICAOne.ToString(),
			msg.NICSupplementA.ToString())
	}

	if msg.NavigationalAccuracyCategoryPosition != fields.NACV2PEPULowerThan10M {
		t.Errorf("Expected NACp \"%v\", got \"%v\"",
			fields.NACV2PEPULowerThan10M.ToString(),
			msg.NavigationalAccuracyCategoryPosition.ToString())
	}

	if msg.GeometricVerticalAccuracy != fields.GVALowerThan45m {
		t.Errorf("Expected Geometric Vertical Accuracy \"%v\", got \"%v\"",
			fields.GVALowerThan45m.ToString(),
			msg.GeometricVerticalAccuracy.ToString())
	}

	if msg.SourceIntegrityLevel != fields.SILLevel3 {
		t.Errorf("Expected SIL \"%v\", got \"%v\"",
			fields.SILLevel3.ToString(),
			msg.SourceIntegrityLevel.ToString())
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

	if msg.SourceIntegrityLevelSupplement != fields.SILSBySample {
		t.Errorf("Expected SIL Supplement \"%v\", got \"%v\"",
			fields.SILSBySample.ToString(),
			msg.SourceIntegrityLevelSupplement.ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat31AirborneV2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat31AirborneV2Message()[:6]

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31AirborneV2BadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat31AirborneV2Message()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31AirborneV2BadSubType(t *testing.T) {

	// Change subtype to surface
	data := buildValidFormat31AirborneV2Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31AirborneV2BadADSBLevel(t *testing.T) {

	// Set data at ADSB level 0
	data := buildValidFormat31AirborneV2Message()
	data[5] = data[5] & 0x1F

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31AirborneV2BadContent(t *testing.T) {

	// Set Content to 1
	data := buildValidFormat31AirborneV2Message()
	data[1] = data[1] | 0x40

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31AirborneV2BadOMFormat(t *testing.T) {

	// Set Service Format to 1
	data := buildValidFormat31AirborneV2Message()
	data[3] = data[3] | 0x40

	_, err := ReadFormat31AirborneV2(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat31AirborneV2Message() []byte {
	data := make([]byte, 7)

	// 1111 1000: code 31 (11111) + subtype 0 (000)
	data[0] = 0xF8

	// 0011 0011: Content (00) + ACAS operational (1) + 1090 In (1) + Reserved (00) +
	// Capability of sending velocity messages (1) + Capability of sending targetState messages (1)
	data[1] = 0x33

	// 1010 0000: multiple TC Reports (10) + UAT In capable (1) + Reserved (00000)
	data[2] = 0xA0

	// 0011 1111: OM Format (00) + ACAS RA Active (1) + Ident active (1) + Receiving ATC (1) +
	// Single Antenna (1) + System Design ReaderLevel 3 (11)
	data[3] = 0x3F

	// 0000 0000:  Reserved (0000 0000)
	data[4] = 0x00

	// 0101 1010: ADSB Version (010) + NIC Supplement A(1) + EPU < 10 m (1010)
	data[5] = 0x5A

	// 1011 1110: GVA<45m (10) + SIL 3 (11) + NICbaro verified (1) + Magnetic North (1) + SIl By Sample (1) + Reserved (0)
	data[6] = 0xBE

	return data
}
