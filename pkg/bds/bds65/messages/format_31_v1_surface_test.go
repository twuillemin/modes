package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"testing"
)

func TestReadFormat31V1Surface(t *testing.T) {

	msg, err := ReadFormat31V1Surface(buildValidFormat31V1SurfaceMessage())
	if err != nil {
		t.Fatal(err)
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

	if msg.GetSubtype() != fields.SubtypeV1Surface {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeV1Surface.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.SurfaceCapabilityClass.PositionOffsetApplied != fields.POAIsADSB {
		t.Errorf("Expected POA \"%v\", got \"%v\"",
			fields.POAIsADSB.ToString(),
			msg.SurfaceCapabilityClass.PositionOffsetApplied.ToString())
	}

	if msg.SurfaceCapabilityClass.CockpitDisplayOfTrafficInformationStatus != fields.CDTIOperational {
		t.Errorf("Expected CDTI \"%v\", got \"%v\"",
			fields.CDTIOperational.ToString(),
			msg.SurfaceCapabilityClass.CockpitDisplayOfTrafficInformationStatus.ToString())
	}

	if msg.SurfaceCapabilityClass.B2Low != fields.B2LLessThan70W {
		t.Errorf("Expected B2Low \"%v\", got \"%v\"",
			fields.B2LLessThan70W.ToString(),
			msg.SurfaceCapabilityClass.B2Low.ToString())
	}

	if msg.LengthAndWidth != fields.LWLength65Width59Point5 {
		t.Errorf("Expected Length And Width \"%v\", got \"%v\"",
			fields.LWLength65Width59Point5.ToString(),
			msg.LengthAndWidth.ToString())
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
		t.Errorf("Expected ATC Services \"%v\", got \"%v\"",
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

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat31V1SurfaceTooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat31V1SurfaceMessage()[:6]

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V1SurfaceBadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat31V1SurfaceMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V1SurfaceBadSubType(t *testing.T) {

	// Change subtype to airborne
	data := buildValidFormat31V1SurfaceMessage()
	data[0] = data[0] & 0xF8

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V1SurfaceBadADSBLevel(t *testing.T) {

	// Set data at ADSB level 0
	data := buildValidFormat31V1SurfaceMessage()
	data[5] = data[5] & 0x1F

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V1SurfaceBadServiceLevel(t *testing.T) {

	// Set Service Level to 1
	data := buildValidFormat31V1SurfaceMessage()
	data[1] = data[1] | 0x04

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31V1SurfaceBadOMFormat(t *testing.T) {

	// Set Service Format to 1
	data := buildValidFormat31V1SurfaceMessage()
	data[3] = data[3] | 0x40

	_, err := ReadFormat31V1Surface(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat31V1SurfaceMessage() []byte {
	data := make([]byte, 7)

	// 1111 1000: code 31 (11111) + subtype 0 (001)
	data[0] = 0xF9

	// 0011 0010: ServiceLevel (00) + POA is ADSB (1) + Traffic display operational (1) + ServiceLevel (00) +
	// B2Low Lower than 70 W (1) + Reserved (0)
	data[1] = 0x32

	// 0000 1010: Reserved (0000) + Length>65 and Width > 59.5 (1010)
	data[2] = 0x0A

	// 0011 1000: OM Format (00) + ACAS RA Active (1) + Ident active (1) + Receiving ATC (1) + Reserved (000)
	data[3] = 0x38

	// 0000 0000:  Reserved (0000 0000)
	data[4] = 0x00

	// 0011 1010: ADSB Version (001) + NIC Supplement (1) + EPU < 10 m and VEPU < 15 m (1010)
	data[5] = 0x3A

	// 0011 1100: BAQ (00) + SIL 3 (11) + Track Angle (1) + Magnetic North (1) + Reserved (00)
	data[6] = 0x3C

	return data
}
