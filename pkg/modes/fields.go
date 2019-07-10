package modes

import "fmt"

// -----------------------------------------------------------------------------------------
//
//                                 VerticalStatus (VS)
//
// -----------------------------------------------------------------------------------------

// VerticalStatus (VS) downlink field shall indicate the status of the aircraft
//
// Defined at 3.1.2.8.2.1
type VerticalStatus int

const (
	// VerticalStatusAirborne signifies that the aircraft is airborne
	VerticalStatusAirborne VerticalStatus = 0
	// VerticalStatusOnTheGround signifies that the aircraft is on the ground
	VerticalStatusOnTheGround VerticalStatus = 1
)

// readVerticalStatus reads the VS field from a message
func readVerticalStatus(message MessageData) VerticalStatus {
	if message.FirstField&0x04 != 0 {
		return VerticalStatusOnTheGround
	}
	return VerticalStatusAirborne
}

// -----------------------------------------------------------------------------------------
//
//                                 CrossLinkCompatibility (CC)
//
// -----------------------------------------------------------------------------------------

// CrossLinkCompatibility (CC) downlink field shall indicate the ability of the transponder to support the cross-link
// capability, i.e. decode the contents of the DS field in an interrogation with UF equals 0 and respond with the
// contents of the specified GICB register in the corresponding reply with DF equals 16.
//
// Defined at 3.1.2.8.2.3
type CrossLinkCompatibility int

const (
	// CrossLinkCompatibilityNotSupported signifies that the transponder cannot support the cross-link capability.
	CrossLinkCompatibilityNotSupported CrossLinkCompatibility = 0
	// CrossLinkCompatibilitySupported signifies that the transponder supports the cross-link capability.
	CrossLinkCompatibilitySupported CrossLinkCompatibility = 1
)

// readCrossLinkCompatibility reads the CC field from a message
func readCrossLinkCompatibility(message MessageData) CrossLinkCompatibility {
	if message.FirstField&0x02 != 0 {
		return CrossLinkCompatibilitySupported
	}
	return CrossLinkCompatibilityNotSupported
}

// -----------------------------------------------------------------------------------------
//
//                                 SensitivityLevelReport (SL)
//
// -----------------------------------------------------------------------------------------

// SensitivityLevelReport (SL) downlink field shall be included in both short and long air-air reply formats
// (DF = 0 and 16). This field shall denote the sensitivity level at which ACAS is currently operating.
//
// Defined at 4.3.8.4.2.5
type SensitivityLevelReport int

const (
	// SensitivityLevelACASInoperative signifies that ACAS is inoperative
	SensitivityLevelACASInoperative SensitivityLevelReport = 0
	// SensitivityLevel1 signifies that ACAS is operating at sensitivity level 1
	SensitivityLevel1 SensitivityLevelReport = 1
	// SensitivityLevel2 signifies that ACAS is operating at sensitivity level 2
	SensitivityLevel2 SensitivityLevelReport = 2
	// SensitivityLevel3 signifies that ACAS is operating at sensitivity level 3
	SensitivityLevel3 SensitivityLevelReport = 3
	// SensitivityLevel4 signifies that ACAS is operating at sensitivity level 4
	SensitivityLevel4 SensitivityLevelReport = 4
	// SensitivityLevel5 signifies that ACAS is operating at sensitivity level 5
	SensitivityLevel5 SensitivityLevelReport = 5
	// SensitivityLevel6 signifies that ACAS is operating at sensitivity level 6
	SensitivityLevel6 SensitivityLevelReport = 6
	// SensitivityLevel7 signifies that ACAS is operating at sensitivity level 7
	SensitivityLevel7 SensitivityLevelReport = 7
)

// readSensitivityLevelReport reads the SL field from a message
func readSensitivityLevelReport(message MessageData) SensitivityLevelReport {

	// The 3 first bits of the message
	sensitivity := message.Payload[0] >> 5

	return SensitivityLevelReport(sensitivity)
}

// -----------------------------------------------------------------------------------------
//
//                                 ReplyInformationAirAir (RI)
//
// -----------------------------------------------------------------------------------------

// ReplyInformationAirAir (RI) downlink field shall report the aircraft’s maximum cruising true
// airspeed capability and type of reply to interrogating aircraft.
//   - Value of 0 signifies a reply to an air-air interrogation UF = 0 with AQ = 0, no operating ACAS.
//   - Values from 1 to 7 are reserved for ACAS
//   - Values from 8 to 715 signifies a reply to an air-air interrogation UF = 0 with AQ = 1 and that the
//   maximum airspeed is defined as detailed in the constant values.
//
// Defined at 3.1.2.8.2.2
type ReplyInformationAirAir int

const (
	// ReplyInformationNoOperatingACAS signifies no operating ACAS.
	ReplyInformationNoOperatingACAS VerticalStatus = 0
	// ReplyInformationNReservedACAS1 is reserved for ACAS (1).
	ReplyInformationNReservedACAS1 VerticalStatus = 1
	// ReplyInformationNReservedACAS2 is reserved for ACAS (2).
	ReplyInformationNReservedACAS2 VerticalStatus = 2
	// ReplyInformationNReservedACAS3 is reserved for ACAS (3).
	ReplyInformationNReservedACAS3 VerticalStatus = 3
	// ReplyInformationNReservedACAS4 is reserved for ACAS (4).
	ReplyInformationNReservedACAS4 VerticalStatus = 4
	// ReplyInformationNReservedACAS5 is reserved for ACAS (5).
	ReplyInformationNReservedACAS5 VerticalStatus = 5
	// ReplyInformationNReservedACAS6 is reserved for ACAS (6).
	ReplyInformationNReservedACAS6 VerticalStatus = 6
	// ReplyInformationNReservedACAS7 is reserved for ACAS (7).
	ReplyInformationNReservedACAS7 VerticalStatus = 7
	// ReplyInformationMaximumAirSpeedNotAvailable signifies no maximum airspeed data available
	ReplyInformationMaximumAirSpeedNotAvailable VerticalStatus = 8
	// ReplyInformationMaximumAirSpeed0To140 signifies maximum airspeed is less than or equal to 140 km/h (75 kt)
	ReplyInformationMaximumAirSpeed0To140 VerticalStatus = 9
	// ReplyInformationMaximumAirSpeed140To280 signifies maximum airspeed is greater than 140 and less than or equal
	// to 280 km/h (75 and 150 kt)
	ReplyInformationMaximumAirSpeed140To280 VerticalStatus = 10
	// ReplyInformationMaximumAirSpeed280To560 signifies maximum airspeed is greater than 280 and less than or equal
	// to 560 km/h (150 and 300 kt)
	ReplyInformationMaximumAirSpeed280To560 VerticalStatus = 11
	// ReplyInformationMaximumAirSpeed560To1110 signifies maximum airspeed is greater than 560 and less than or equal
	// to 1 110 km/h (300 and 600 kt)
	ReplyInformationMaximumAirSpeed560To1110 VerticalStatus = 12
	// ReplyInformationMaximumAirSpeed1110To2220 signifies maximum airspeed is greater than 1 110 and less than or
	// equal to 2 220 km/h (600 and 1 200 kt)
	ReplyInformationMaximumAirSpeed1110To2220 VerticalStatus = 13
	// ReplyInformationMaximumAirSpeedOver2220 signifies maximum airspeed is more than 2 220 km/h (1 200 kt)
	ReplyInformationMaximumAirSpeedOver2220 VerticalStatus = 14
	// ReplyInformationNotAssigned is not assigned.
	ReplyInformationNotAssigned VerticalStatus = 15
)

func readReplyInformationAirAir(message MessageData) ReplyInformationAirAir {

	replyInformation := ((message.Payload[0] & 0x03) << 2) | ((message.Payload[1] & 0x080) >> 7)

	return ReplyInformationAirAir(replyInformation)
}

// -----------------------------------------------------------------------------------------
//
//                                 AltitudeCode (AC)
//
// -----------------------------------------------------------------------------------------

// AltitudeCodeReportMethod details how the altitude of the AltitudeCode (AC) field is reported. It corresponds to AC
// fields bit 26 (M-bit) and bit 28 (Q-bit):
//   - M = 0 and Q = 0 => 100 foot increments
//   - M = 0 and Q = 1 => 25 foot increments. For this method, the coding method is only able to provide values between
//   minus 1 000 ft and plus 50 175 ft.
//   - M = 1 => metric units
//
// Defined at 3.1.2.6.5.4
type AltitudeCodeReportMethod int

const (
	// AltitudeCodeReportNotAvailable signifies that altitude information is not available or that the altitude
	// has been determined invalid.
	AltitudeCodeReportNotAvailable AltitudeCodeReportMethod = 1
	// AltitudeCodeReportMetricUnits signifies that altitude is reported in metric units.
	AltitudeCodeReportMetricUnits AltitudeCodeReportMethod = 2
	// AltitudeCodeReport100FootIncrements signifies that altitude is reported in 100-foot increments.
	AltitudeCodeReport100FootIncrements AltitudeCodeReportMethod = 3
	// AltitudeCodeReport25FootIncrements signifies that altitude is reported in 25-foot increments.
	AltitudeCodeReport25FootIncrements AltitudeCodeReportMethod = 4
)

// AltitudeCode (AC) field shall contain altitude coded in feet or metric units.
//
// Defined at 3.1.2.6.5.4
type AltitudeCode struct {
	ReportMethod AltitudeCodeReportMethod

	AltitudeInFeet int
}

// readAltitudeCode reads the altitude code from a message
func readAltitudeCode(message MessageData) AltitudeCode {

	// Altitude code is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// message     |_  _  _  x  x  x  x  x |x  M  x  Q  x  x  x  x
	// 100 foot    |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 0  B2 D2 B4 D4

	// Get the raw altitude code
	altitudeCode := (uint16(message.Payload[1])<<8 | uint16(message.Payload[2])) & 0x1fff

	if altitudeCode == 0 {
		return AltitudeCode{
			AltitudeCodeReportNotAvailable,
			-1,
		}
	}

	// Get the M bit
	mBit := (altitudeCode & 0x0040) != 0

	// If altitude reported in metric units
	if mBit {

		// Not specified for now
		return AltitudeCode{
			AltitudeCodeReportMetricUnits,
			0,
		}
	}

	// Get the Q bit
	qBit := (altitudeCode & 0x0010) != 0

	// If altitude reported 25 foot increments
	if qBit {

		// If the M bit equals 0 and the Q bit equals 1, the 11-bit field represented by bits 20 to 25, 27 and 29 to 32
		// shall represent a binary coded field with a least significant bit (LSB) of 25 ft. The binary value of the
		// positive decimal integer “N” shall be encoded to report pressure-altitude in the range [(25 N – 1 000)
		// plus or minus 12.5 ft]. The coding of 3.1.2.6.5.4 c) shall be used to report pressure-altitude
		// above 50 187.5 ft.
		n := uint16(0)
		n |= (altitudeCode & 0x1F80) >> 2
		n |= (altitudeCode & 0x0020) >> 1
		n |= altitudeCode & 0x000F

		return AltitudeCode{
			AltitudeCodeReport25FootIncrements,
			25*int(n) - 1000,
		}
	}

	// Otherwise, altitude is reported in 100 foot increment

	// The altitude shall be coded according to the pattern for Mode C replies of 3.1.1.7.12.2.3.
	// Starting with bit 20 the sequence shall be C1, A1, C2, A2, C4, A4, ZERO, B1, ZERO, B2, D2, B4, D4.
	c1 := (altitudeCode & 0x1000) != 0
	a1 := (altitudeCode & 0x0800) != 0
	c2 := (altitudeCode & 0x0400) != 0
	a2 := (altitudeCode & 0x0200) != 0
	c4 := (altitudeCode & 0x0100) != 0
	a4 := (altitudeCode & 0x0080) != 0
	b1 := (altitudeCode & 0x0020) != 0
	b2 := (altitudeCode & 0x0008) != 0
	d2 := (altitudeCode & 0x0004) != 0
	b4 := (altitudeCode & 0x0002) != 0
	d4 := (altitudeCode & 0x0001) != 0

	increment500 := grayToBinary(d2, d4, a1, a2, a4, b1, b2, b4)
	// subIncrement is given from 1 to 5 (so there is always one bit in c1,c2,c4), but it is from 0 to 4
	subIncrement := grayToBinary(false, false, false, false, false, c1, c2, c4)
	increment100 := subIncrement - 1
	// And increment is reversed alternatively
	if increment500%2 != 0 {
		increment100 = 4 - increment100
	}

	// Compute the altitude
	altitudeFeet := -1200 + int(increment500)*500 + int(increment100)*100

	return AltitudeCode{
		AltitudeCodeReport100FootIncrements,
		altitudeFeet,
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Flight Status (FS)
//
// -----------------------------------------------------------------------------------------

// FlightStatus (FS) field shall contain information about the flight status
//
// Defined at 3.1.2.6.5.1
type FlightStatus int

const (
	// FlightStatusNoAlertNoSPIAirborne signifies no alert and no SPI, aircraft is airborne
	FlightStatusNoAlertNoSPIAirborne FlightStatus = 0
	// FlightStatusNoAlertNoSPIOnTheGround signifies no alert and no SPI, aircraft is on the ground
	FlightStatusNoAlertNoSPIOnTheGround FlightStatus = 1
	// FlightStatusAlertNoSPIAirborne signifies alert and no SPI, aircraft is airborne
	FlightStatusAlertNoSPIAirborne FlightStatus = 2
	// FlightStatusAlertNoSPIOnTheGround signifies alert and no SPI, aircraft is on the ground
	FlightStatusAlertNoSPIOnTheGround FlightStatus = 3
	// FlightStatusAlertSPIAirborne signifies alert and SPI, aircraft is airborne
	FlightStatusAlertSPIAirborne FlightStatus = 4
	// FlightStatusAlertSPIOnTheGround signifies alert and SPI, aircraft is on the ground
	FlightStatusAlertSPIOnTheGround FlightStatus = 5
	// FlightStatusReserved is reserved
	FlightStatusReserved FlightStatus = 6
	// FlightStatusNotAssigned is not assigned
	FlightStatusNotAssigned FlightStatus = 7
)

// readFlightStatus reads the FS field from a message
func readFlightStatus(message MessageData) FlightStatus {

	return FlightStatus(message.FirstField)
}

// -----------------------------------------------------------------------------------------
//
//                                 Downlink Request (DR)
//
// -----------------------------------------------------------------------------------------

// DownlinkRequest (DR) field shall contain requests to downlink information.
//
// Defined at 3.1.2.6.5.2 for codes 0 to 15 and 3.1.2.7.7.1 for codes 16 to 31
type DownlinkRequest int

const (
	// DownlinkRequestNoDownlinkRequest signifies no downlink request
	DownlinkRequestNoDownlinkRequest DownlinkRequest = 0
	// DownlinkRequestToSendCommBMessage signifies request to send Comm-B message
	DownlinkRequestToSendCommBMessage DownlinkRequest = 1
	// DownlinkRequestReservedACAS2 is reserved for ACAS
	DownlinkRequestReservedACAS2 DownlinkRequest = 2
	// DownlinkRequestReservedACAS3 is reserved for ACAS
	DownlinkRequestReservedACAS3 DownlinkRequest = 3
	// DownlinkRequestCommBMessage1Available signifies Comm-B broadcast message 1 available
	DownlinkRequestCommBMessage1Available DownlinkRequest = 4
	// DownlinkRequestCommBMessage2Available signifies Comm-B broadcast message 2 available
	DownlinkRequestCommBMessage2Available DownlinkRequest = 5
	// DownlinkRequestReservedACAS6 is reserved for ACAS
	DownlinkRequestReservedACAS6 DownlinkRequest = 6
	// DownlinkRequestReservedACAS7 is reserved for ACAS
	DownlinkRequestReservedACAS7 DownlinkRequest = 7
	// DownlinkRequestNotAssigned8 is not assigned
	DownlinkRequestNotAssigned8 DownlinkRequest = 8
	// DownlinkRequestNotAssigned9 is not assigned
	DownlinkRequestNotAssigned9 DownlinkRequest = 9
	// DownlinkRequestNotAssigned10 is not assigned
	DownlinkRequestNotAssigned10 DownlinkRequest = 10
	// DownlinkRequestNotAssigned11 is not assigned
	DownlinkRequestNotAssigned11 DownlinkRequest = 11
	// DownlinkRequestNotAssigned12 is not assigned
	DownlinkRequestNotAssigned12 DownlinkRequest = 12
	// DownlinkRequestNotAssigned13 is not assigned
	DownlinkRequestNotAssigned13 DownlinkRequest = 13
	// DownlinkRequestNotAssigned14 is not assigned
	DownlinkRequestNotAssigned14 DownlinkRequest = 14
	// DownlinkRequestNotAssigned15 is not assigned
	DownlinkRequestNotAssigned15 DownlinkRequest = 15
	// DownlinkRequestELMAvailable1Segments announces the presence of a downlink ELM of 1 segments
	DownlinkRequestELMAvailable1Segments DownlinkRequest = 16
	// DownlinkRequestELMAvailable2Segments announces the presence of a downlink ELM of 2 segments
	DownlinkRequestELMAvailable2Segments DownlinkRequest = 17
	// DownlinkRequestELMAvailable3Segments announces the presence of a downlink ELM of 3 segments
	DownlinkRequestELMAvailable3Segments DownlinkRequest = 18
	// DownlinkRequestELMAvailable4Segments announces the presence of a downlink ELM of 4 segments
	DownlinkRequestELMAvailable4Segments DownlinkRequest = 19
	// DownlinkRequestELMAvailable5Segments announces the presence of a downlink ELM of 5 segments
	DownlinkRequestELMAvailable5Segments DownlinkRequest = 20
	// DownlinkRequestELMAvailable6Segments announces the presence of a downlink ELM of 6 segments
	DownlinkRequestELMAvailable6Segments DownlinkRequest = 21
	// DownlinkRequestELMAvailable7Segments announces the presence of a downlink ELM of 7 segments
	DownlinkRequestELMAvailable7Segments DownlinkRequest = 22
	// DownlinkRequestELMAvailable8Segments announces the presence of a downlink ELM of 8 segments
	DownlinkRequestELMAvailable8Segments DownlinkRequest = 23
	// DownlinkRequestELMAvailable9Segments announces the presence of a downlink ELM of 9 segments
	DownlinkRequestELMAvailable9Segments DownlinkRequest = 24
	// DownlinkRequestELMAvailable10Segments announces the presence of a downlink ELM of 10 segments
	DownlinkRequestELMAvailable10Segments DownlinkRequest = 25
	// DownlinkRequestELMAvailable11Segments announces the presence of a downlink ELM of 11 segments
	DownlinkRequestELMAvailable11Segments DownlinkRequest = 26
	// DownlinkRequestELMAvailable12Segments announces the presence of a downlink ELM of 12 segments
	DownlinkRequestELMAvailable12Segments DownlinkRequest = 27
	// DownlinkRequestELMAvailable13Segments announces the presence of a downlink ELM of 13 segments
	DownlinkRequestELMAvailable13Segments DownlinkRequest = 28
	// DownlinkRequestELMAvailable14Segments announces the presence of a downlink ELM of 14 segments
	DownlinkRequestELMAvailable14Segments DownlinkRequest = 29
	// DownlinkRequestELMAvailable15Segments announces the presence of a downlink ELM of 15 segments
	DownlinkRequestELMAvailable15Segments DownlinkRequest = 30
	// DownlinkRequestELMAvailable16Segments announces the presence of a downlink ELM of 16 segments
	DownlinkRequestELMAvailable16Segments DownlinkRequest = 31
)

// readSensitivityLevelReport reads the DR field from a message
func readDownlinkRequest(message MessageData) DownlinkRequest {

	return DownlinkRequest(message.Payload[0] >> 3)
}

// -----------------------------------------------------------------------------------------
//
//                                 Utility Message (UM)
//
// -----------------------------------------------------------------------------------------

// UtilityMessage (UM) field shall contain transponder communications status information
//
// Defined at 3.1.2.6.5.3
type UtilityMessage struct {
	// InterrogatorIdentifier subfield reports the identifier of the interrogator that is reserved for
	// multisite communications.
	InterrogatorIdentifier uint8
	// IdentifierDesignator subfield reports the type of reservation made by the interrogator identified in IIS
	IdentifierDesignator UtilityMessageIdentifierDesignator
}

// UtilityMessageIdentifierDesignator subfield reports the type of reservation made by the interrogator
// identified in IIS.
//
// Defined at 3.1.2.6.5.3.1
type UtilityMessageIdentifierDesignator int

const (
	// UtilityMessageIdentifierDesignatorNoInformation signifies no information
	UtilityMessageIdentifierDesignatorNoInformation UtilityMessageIdentifierDesignator = 0
	// UtilityMessageIdentifierDesignatorCommB signifies IIS contains Comm-B II code
	UtilityMessageIdentifierDesignatorCommB UtilityMessageIdentifierDesignator = 1
	// UtilityMessageIdentifierDesignatorCommC signifies IIS contains Comm-C II code
	UtilityMessageIdentifierDesignatorCommC UtilityMessageIdentifierDesignator = 2
	// UtilityMessageIdentifierDesignatorCommD signifies IIS contains Comm-D II code
	UtilityMessageIdentifierDesignatorCommD UtilityMessageIdentifierDesignator = 3
)

// readUtilityMessage reads the UM field from a message
func readUtilityMessage(message MessageData) UtilityMessage {

	ii := ((message.Payload[0] & 0x07) << 1) | ((message.Payload[1] & 0x80) >> 7)
	id := (message.Payload[1] & 0x60) >> 5

	return UtilityMessage{
		InterrogatorIdentifier: ii,
		IdentifierDesignator:   UtilityMessageIdentifierDesignator(id),
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Identity (ID)
//
// -----------------------------------------------------------------------------------------

// Identity (ID) field shall contain aircraft identity code, in accordance with the pattern for Mode A replies
//
// Defined at 3.1.2.6.7.1
type Identity struct {
	// Identity is the identity, which is always a 4 digit chars.
	Identity string
}

// IsSpecialEmergency informs that the identity is a special code for an aircraft with
// radiocommunication failure. Identity string is "7700".
func (identity Identity) IsSpecialEmergency() bool {
	return identity.Identity == "7700"
}

// IsSpecialRadiocommunicationFailure informs that the identity is a special code for an aircraft with
// radiocommunication failure. Identity string is "7600".
func (identity Identity) IsSpecialRadiocommunicationFailure() bool {
	return identity.Identity == "7600"
}

// IsSpecialUnlawfulInterference informs that the identity is a special code for an aircraft t which is being
// subjected to unlawful interference. Identity string is "7500".
func (identity Identity) IsSpecialUnlawfulInterference() bool {
	return identity.Identity == "7500"
}

// IsSpecialInstructionNotReceived informs that the identity is a special code for an aircraft which has not
// received any instructions from air traffic control units to operate the transponder. Identity string is "2000".
func (identity Identity) IsSpecialInstructionNotReceived() bool {
	return identity.Identity == "2000"
}

// readIdentity reads the identity from a message
func readIdentity(message MessageData) Identity {

	// Identity is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// id bits     |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 D1  B2 D2 B4 D4

	// Get the raw identity code
	identity := (uint16(message.Payload[1])<<8 | uint16(message.Payload[2])) & 0x1fff

	// Starting with bit 20 the sequence shall be C1, A1, C2, A2, C4, A4, ZERO, B1, D1, B2, D2, B4, D4.
	c1 := (identity & 0x1000) >> 12
	a1 := (identity & 0x0800) >> 11
	c2 := (identity & 0x0400) >> 10
	a2 := (identity & 0x0200) >> 9
	c4 := (identity & 0x0100) >> 8
	a4 := (identity & 0x0080) >> 7
	b1 := (identity & 0x0020) >> 5
	d1 := (identity & 0x0010) >> 4
	b2 := (identity & 0x0008) >> 3
	d2 := (identity & 0x0004) >> 2
	b4 := (identity & 0x0002) >> 1
	d4 := (identity & 0x0021)

	// Defined at 3.1.1.6
	a := a4<<2 | a2<<1 | a1
	b := b4<<2 | b2<<1 | b1
	c := c4<<2 | c2<<1 | c1
	d := d4<<2 | d2<<1 | d1

	return Identity{
		Identity: fmt.Sprintf("%v%v%v%v", a, b, c, d),
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Capability (CA)
//
// -----------------------------------------------------------------------------------------

// Capability (CA) field shall convey information on the transponder level, the additional information below,
// and shall be used in formats DF = 11 and DF = 17.
//
// Defined at 3.1.2.5.2.2.1
type Capability int

const (
	// CapabilityLevel1Transponder signifies Level 1 transponder (surveillance only), and no ability to set CA
	// code 7 and either airborne or on the ground
	CapabilityLevel1Transponder Capability = 0
	// CapabilityReserved1 is reserved
	CapabilityReserved1 Capability = 1
	// CapabilityReserved2 is reserved
	CapabilityReserved2 Capability = 2
	// CapabilityReserved3 is reserved
	CapabilityReserved3 Capability = 3
	// CapabilityLevel2OnTheGround signifies Level 2 or above transponder and ability to set CA code 7 and on the ground
	CapabilityLevel2OnTheGround Capability = 4
	// CapabilityLevel2Airborne signifies Level 2 or above transponder and ability to set CA code 7 and airborne
	CapabilityLevel2Airborne Capability = 5
	// CapabilityLevelOnTheGroundOrAirborne signifies Level 2 or above transponder and ability to set CA code 7 and
	// either airborne or on the ground
	CapabilityLevelOnTheGroundOrAirborne Capability = 6
	// CapabilityFSOrDR signifies the DR field is not equal to 0 or the FS field equals 2, 3, 4 or 5, and either
	//airborne or on the ground
	CapabilityFSOrDR Capability = 7
)

// readFlightStatus reads the FS field from a message
func readCapability(message MessageData) Capability {

	return Capability(message.FirstField)
}

// -----------------------------------------------------------------------------------------
//
//                                 Address Announced (AA)
//
// -----------------------------------------------------------------------------------------

// AddressAnnounced (AA) field shall contain the aircraft address which provides unambiguous identification of
// the aircraft.
//
// Defined at 3.1.2.5.2.2.2
type AddressAnnounced struct {
	Address uint32
}

// readAddressAnnounced reads the AA field from a message
func readAddressAnnounced(message MessageData) AddressAnnounced {

	return AddressAnnounced{
		Address: uint32(message.Payload[0])<<16 | uint32(message.Payload[1])<<8 | uint32(message.Payload[2]),
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Message ACAS (MV)
//
// -----------------------------------------------------------------------------------------

// MessageACAS (MV) field shall contain the aircraft address which provides unambiguous identification of
// the aircraft.
//
// Defined at 3.1.2.8.3.1
type MessageACAS struct {
	Data []byte
}

// readMessageACAS reads the MV field from a message
func readMessageACAS(message MessageData) MessageACAS {

	return MessageACAS{
		Data: message.Payload[3:10],
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Message Extended Squitter (ME)
//
// -----------------------------------------------------------------------------------------

// MessageExtendedSquitter field in DF = 17 shall be used to transmit broadcast messages. Extended squitter shall be
// supported by registers 05, 06, 07, 08, 09, 0A {HEX} and 61-6F {HEX} and shall conform to either version 0, version 1
// or version 2 message formats as described below:
//
// Defined at 3.1.2.8.6.2
type MessageExtendedSquitter struct {
	Data []byte
}

// readAddressAnnounced reads the AA field from a message
func readMessageExtendedSquitter(message MessageData) MessageExtendedSquitter {

	return MessageExtendedSquitter{
		Data: message.Payload[3:10],
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Control Field (CF)
//
// -----------------------------------------------------------------------------------------

// ControlField (CF) field in DF = 18 shall be used to define the format of the 112-bit transmission as follows.
//
// Defined at 3.1.2.8.7.2
type ControlField int

const (
	// ControlFieldADSB signifies ADS-B ES/NT devices that report the ICAO 24-bit address in the AA field
	ControlFieldADSB ControlField = 0
	// ControlFieldADSBReserved signifies reserved for ADS-B for ES/NT devices that use other addressing techniques in the AA field
	ControlFieldADSBReserved ControlField = 1
	// ControlFieldTISBFineFormat signifies fine format TIS-B message
	ControlFieldTISBFineFormat ControlField = 2
	// ControlFieldTISBCoarseFormat signifies coarse format TIS-B message
	ControlFieldTISBCoarseFormat ControlField = 3
	// ControlFieldTISBReservedManagement signifies reserved for TIS-B management messages
	ControlFieldTISBReservedManagement ControlField = 4
	// ControlFieldTISBRelayADSB signifies TIS-B messages that relay ADS-B messages that use other addressing techniques in the AA field
	ControlFieldTISBRelayADSB ControlField = 5
	// ControlFieldADSBRebroadcast signifies  ADS-B rebroadcast using the same type codes and message formats as defined for DF = 17 ADS-B messages
	ControlFieldADSBRebroadcast ControlField = 6
	// ControlFieldReserved is reserved
	ControlFieldReserved ControlField = 7
)

// readControlField reads the CF field from a message
func readControlField(message MessageData) ControlField {

	return ControlField(message.FirstField)
}

// -----------------------------------------------------------------------------------------
//
//                                 Application Field (AF)
//
// -----------------------------------------------------------------------------------------

// ApplicationField (AF) field in DF = 19 shall be used to define the format of the 112-bit transmission as follows.
//
// Defined at 3.1.2.8.8.2
type ApplicationField int

const (
	// ApplicationFieldReserved0 is reserved
	ApplicationFieldReserved0 ApplicationField = 0
	// ApplicationFieldReserved1 is reserved
	ApplicationFieldReserved1 ApplicationField = 1
	// ApplicationFieldReserved2 is reserved
	ApplicationFieldReserved2 ApplicationField = 2
	// ApplicationFieldReserved3 is reserved
	ApplicationFieldReserved3 ApplicationField = 3
	// ApplicationFieldReserved4 is reserved
	ApplicationFieldReserved4 ApplicationField = 4
	// ApplicationFieldReserved5 is reserved
	ApplicationFieldReserved5 ApplicationField = 5
	// ApplicationFieldReserved6 is reserved
	ApplicationFieldReserved6 ApplicationField = 6
	// ApplicationFieldReserved7 is reserved
	ApplicationFieldReserved7 ApplicationField = 7
)

// readApplicationField reads the AF field from a message
func readApplicationField(message MessageData) ApplicationField {

	return ApplicationField(message.FirstField)
}

// -----------------------------------------------------------------------------------------
//
//                                 Message Comm-B (MB)
//
// -----------------------------------------------------------------------------------------

// MessageCommB field shall be used to transmit data link messages to the ground
//
// Defined at 3.1.2.6.6.1
type MessageCommB struct {
	Data []byte
}

// readMessageCommB reads the MB field from a message
func readMessageCommB(message MessageData) MessageCommB {

	return MessageCommB{
		Data: message.Payload[3:10],
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 ControlELM (KE)
//
// -----------------------------------------------------------------------------------------

// ControlELM (KE) downlink field shall define the content of the ND and MD fields.
//
// Defined at 3.1.2.7.3.1
type ControlELM int

const (
	// DownlinkELMTransmission signifies that the aircraft is airborne
	DownlinkELMTransmission ControlELM = 0
	// UplinkELMAcknowledgement signifies that the aircraft is on the ground
	UplinkELMAcknowledgement ControlELM = 1
)

// readControlELM reads the KE field from a message
func readControlELM(message MessageData) ControlELM {
	if message.DownLinkFormat&0x02 != 0 {
		return UplinkELMAcknowledgement
	}
	return DownlinkELMTransmission
}

// -----------------------------------------------------------------------------------------
//
//                                 Number Of D-Segment (ND)
//
// -----------------------------------------------------------------------------------------

// NumberOfDSegment (ND) downlink field shall designate the number of the message segment contained in MD
//
// Defined at 3.1.2.7.3.2
type NumberOfDSegment uint8

// readNumberOfDSegment reads the ND field from a message
func readNumberOfDSegment(message MessageData) NumberOfDSegment {

	return NumberOfDSegment((message.DownLinkFormat&0x01)<<3 | message.FirstField)
}

// -----------------------------------------------------------------------------------------
//
//                                 Message Comm-D (MD)
//
// -----------------------------------------------------------------------------------------

// MessageCommD field shall contain:
//   a) one of the segments of a sequence used to transmit a downlink ELM to the interrogator; or
//   b) control codes for an uplink ELM.
//
// Defined at 3.1.2.7.3.3
type MessageCommD struct {
	Data []byte
}

// readMessageCommD reads the MB field from a message
func readMessageCommD(message MessageData) MessageCommD {

	return MessageCommD{
		Data: message.Payload,
	}
}

// -----------------------------------------------------------------------------------------
//
//                                 Miscellaneous functions
//
// -----------------------------------------------------------------------------------------

// grayToBinary convert the given bits (b0 being the MSB and b7 the LSB) from gray code to "classical" binary
func grayToBinary(b0, b1, b2, b3, b4, b5, b6, b7 bool) uint8 {

	num := uint8(0)
	if b0 {
		num |= 0x80
	}
	if b1 {
		num |= 0x40
	}
	if b2 {
		num |= 0x20
	}
	if b3 {
		num |= 0x10
	}
	if b4 {
		num |= 0x08
	}
	if b5 {
		num |= 0x04
	}
	if b6 {
		num |= 0x02
	}
	if b7 {
		num |= 0x01
	}

	mask := num >> 1

	for mask != 0 {
		num = num ^ mask
		mask = num >> 1
	}

	return num
}
