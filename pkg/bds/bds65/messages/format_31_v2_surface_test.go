package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"testing"
)

func TestReadFormat31V2SurfaceValid(t *testing.T) {

	msg, err := ReadFormat31V2Surface(buildValidFormat31V2SurfaceMessage())
	if err != nil {
		t.Error(err)
	}

	if msg.GetMessageFormat() != adsb.Format31V2 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format31V2.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS65.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS65.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.SubtypeV2Surface {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeV2Surface.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.SurfaceCapabilityClass.ExtendedSquitterIn != fields.ESICapable {
		t.Errorf("Expected Extended Squitter In \"%v\", got \"%v\"",
			fields.ESICapable.ToString(),
			msg.SurfaceCapabilityClass.ExtendedSquitterIn.ToString())
	}

	if msg.SurfaceCapabilityClass.B2Low != fields.B2LLessThan70W {
		t.Errorf("Expected B2Low \"%v\", got \"%v\"",
			fields.B2LLessThan70W.ToString(),
			msg.SurfaceCapabilityClass.B2Low.ToString())
	}

	if msg.SurfaceCapabilityClass.UniversalAccessTransceiverCapability != fields.UATCapable {
		t.Errorf("Expected UAT Capable \"%v\", got \"%v\"",
			fields.UATCapable.ToString(),
			msg.SurfaceCapabilityClass.UniversalAccessTransceiverCapability.ToString())
	}

	if msg.SurfaceCapabilityClass.NavigationAccuracyCategoryVelocity != 5 {
		t.Errorf("Expected Navigation Accuracy Category \"%v\", got \"%v\"",
			5,
			msg.SurfaceCapabilityClass.NavigationAccuracyCategoryVelocity.ToString())
	}

	if msg.SurfaceCapabilityClass.NICSupplementC != fields.NICCSOne {
		t.Errorf("Expected NIC Supplement C \"%v\", got \"%v\"",
			fields.NICCSOne,
			msg.SurfaceCapabilityClass.NICSupplementC.ToString())
	}

	if msg.LengthAndWidth != fields.LWLength65Width59Point5 {
		t.Errorf("Expected Length And Width \"%v\", got \"%v\"",
			fields.LWLength65Width59Point5.ToString(),
			msg.LengthAndWidth.ToString())
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

	if msg.OperationalMode.GPSAntenna.GPSAntennaLateral != fields.GLATRight02Meters {
		t.Errorf("Expected GPS Antenna Lateral \"%v\", got \"%v\"",
			fields.GLATRight02Meters.ToString(),
			msg.OperationalMode.GPSAntenna.GPSAntennaLateral.ToString())
	}

	if msg.OperationalMode.GPSAntenna.GPSAntennaLongitudinal != fields.GLON18Meters {
		t.Errorf("Expected GPS Antenna Lateral \"%v\", got \"%v\"",
			fields.GLON18Meters.ToString(),
			msg.OperationalMode.GPSAntenna.GPSAntennaLongitudinal.ToString())
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

	if msg.SourceIntegrityLevel != fields.SILLevel3 {
		t.Errorf("Expected SIL \"%v\", got \"%v\"",
			fields.SILLevel3.ToString(),
			msg.SourceIntegrityLevel.ToString())
	}

	if msg.TrackAngleHeading != fields.TAHTrackAngleReported {
		t.Errorf("Expected Track Angle Heading \"%v\", got \"%v\"",
			fields.TAHTrackAngleReported.ToString(),
			msg.TrackAngleHeading.ToString())
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
}

func TestReadFormat31V2SurfaceTooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat31V2SurfaceMessage()[:6]

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V2SurfaceBadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat31V2SurfaceMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V2SurfaceBadSubType(t *testing.T) {

	// Change subtype to airborne
	data := buildValidFormat31V2SurfaceMessage()
	data[0] = data[0] & 0xF8

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V2SurfaceBadADSBLevel(t *testing.T) {

	// Set data at ADSB level 0
	data := buildValidFormat31V2SurfaceMessage()
	data[5] = data[5] & 0x1F

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V2SurfaceBadContent(t *testing.T) {

	// Set Content to 1
	data := buildValidFormat31V2SurfaceMessage()
	data[1] = data[1] | 0x40

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V2SurfaceBadOMFormat(t *testing.T) {

	// Set Service Format to 1
	data := buildValidFormat31V2SurfaceMessage()
	data[3] = data[3] | 0x40

	_, err := ReadFormat31V2Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat31V2SurfaceMessage() []byte {
	data := make([]byte, 7)

	// 1111 1000: code 31 (11111) + subtype 0 (001)
	data[0] = 0xF9

	// 0001 0011: Content (00) + Reserved (0) + 1090 In (1) + Reserved (00) + B2 Low (1) + UAT In (1)
	data[1] = 0x13

	// 1011 1010: NACv (101) + NIC Supplement C (1) + UAT In capable (1) +  + Length>65 and Width > 59.5 (1010)
	data[2] = 0xBA

	// 0011 0111: OM Format (00) + ACAS RA Active (1) + Ident active (1) + Reserved (0) +
	// Single Antenna (1) + System Design Level 3 (11)
	data[3] = 0x37

	// 1010 1010:  GPS Antenna Lateral: Right (1) 2 meters (01) Longitudinal: 18 m (01010)
	data[4] = 0xAA

	// 0101 1010: ADSB Version (010) + NIC Supplement A(1) + EPU < 10 m (1010)
	data[5] = 0x5A

	// 0011 1110: Reserved (00) + SIL 3 (11) + Track angle (1) + Magnetic North (1) + SIl By Sample (1) + Reserved (0)
	data[6] = 0x3E

	return data
}
