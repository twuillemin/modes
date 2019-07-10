package modes

import "errors"

type Message interface {
	GetMessageData() MessageData
}

// -------------------------------------------------------------------------------------
//                                         DF0
// -------------------------------------------------------------------------------------

// MessageDF0 is a message at the format DF0
type MessageDF0 struct {
	MessageData
	VerticalStatus      VerticalStatus
	CrossLinkCapability CrossLinkCompatibility
	SensitivityLevel    SensitivityLevelReport
	ReplyInformation    ReplyInformationAirAir
	AltitudeCode        AltitudeCode
}

func (message *MessageDF0) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF0 parses a message at the DF0 format
func ParseDF0(message MessageData) (*MessageDF0, error) {

	// Format of the message is as follow:
	//
	//     DF  VS CC _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |   AP
	// 0 0 0 0 0 x x _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 24bits

	if message.DownLinkFormat != 0 {
		return nil, errors.New("DF0 message must have a DownLinkFormat of 0")
	}

	if len(message.Payload) != 3 {
		return nil, errors.New("DF0 message must be 7 bytes long")
	}

	return &MessageDF0{
		MessageData:         message,
		VerticalStatus:      readVerticalStatus(message),
		CrossLinkCapability: readCrossLinkCompatibility(message),
		SensitivityLevel:    readSensitivityLevelReport(message),
		ReplyInformation:    readReplyInformationAirAir(message),
		AltitudeCode:        readAltitudeCode(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF4
// -------------------------------------------------------------------------------------

// MessageDF4 is a message at the format DF4
type MessageDF4 struct {
	MessageData
	FlightStatus    FlightStatus
	DownlinkRequest DownlinkRequest
	UtilityMessage  UtilityMessage
	AltitudeCode    AltitudeCode
}

func (message *MessageDF4) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF4 parses a message at the DF4 format
func ParseDF4(message MessageData) (*MessageDF4, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |   AP
	// 0 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 24bits

	if message.DownLinkFormat != 4 {
		return nil, errors.New("DF4 message must have a DownLinkFormat of 4")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF4 message must be 7 bytes long")
	}

	return &MessageDF4{
		MessageData:     message,
		FlightStatus:    readFlightStatus(message),
		DownlinkRequest: readDownlinkRequest(message),
		UtilityMessage:  readUtilityMessage(message),
		AltitudeCode:    readAltitudeCode(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF5
// -------------------------------------------------------------------------------------

// MessageDF5 is a message at the format DF5
type MessageDF5 struct {
	MessageData
	FlightStatus    FlightStatus
	DownlinkRequest DownlinkRequest
	UtilityMessage  UtilityMessage
	Identity        Identity
}

func (message *MessageDF5) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF5 parses a message at the DF5 format
func ParseDF5(message MessageData) (*MessageDF5, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      ID    |        ID       |   AP
	// 0 0 1 0 1 f f f | d d d d d u u u | u u u i i i i i | i i i i i i i i | 24bits

	if message.DownLinkFormat != 5 {
		return nil, errors.New("DF5 message must have a DownLinkFormat of 5")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF5 message must be 7 bytes long")
	}

	return &MessageDF5{
		MessageData:     message,
		FlightStatus:    readFlightStatus(message),
		DownlinkRequest: readDownlinkRequest(message),
		UtilityMessage:  readUtilityMessage(message),
		Identity:        readIdentity(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF11
// -------------------------------------------------------------------------------------

// MessageDF11 is a message at the format DF11
type MessageDF11 struct {
	MessageData
	Capability       Capability
	AddressAnnounced AddressAnnounced
}

func (message *MessageDF11) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF11 parses a message at the DF11 format
func ParseDF11(message MessageData) (*MessageDF11, error) {

	// Format of the message is as follow:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |   PI
	// 0 1 0 1 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 24bits

	if message.DownLinkFormat != 11 {
		return nil, errors.New("DF11 message must have a DownLinkFormat of 11")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF11 message must be 7 bytes long")
	}

	return &MessageDF11{
		MessageData:      message,
		Capability:       readCapability(message),
		AddressAnnounced: readAddressAnnounced(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF16
// -------------------------------------------------------------------------------------

// MessageDF16 is a message at the format DF16
type MessageDF16 struct {
	MessageData
	VerticalStatus   VerticalStatus
	SensitivityLevel SensitivityLevelReport
	ReplyInformation ReplyInformationAirAir
	AltitudeCode     AltitudeCode
	MessageACAS      MessageACAS
}

func (message *MessageDF16) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF16 parses a message at the DF5 format
func ParseDF16(message MessageData) (*MessageDF16, error) {

	// Format of the message is as follow:
	//
	//     DF   VS _ _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |    MV   |  AP
	// 1 0 0 0 0 x _ _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 56 bits |24bits

	if message.DownLinkFormat != 16 {
		return nil, errors.New("DF16 message must have a DownLinkFormat of 16")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF16 message must be 14 bytes long")
	}

	return &MessageDF16{
		MessageData:      message,
		VerticalStatus:   readVerticalStatus(message),
		SensitivityLevel: readSensitivityLevelReport(message),
		ReplyInformation: readReplyInformationAirAir(message),
		AltitudeCode:     readAltitudeCode(message),
		MessageACAS:      readMessageACAS(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF17
// -------------------------------------------------------------------------------------

// MessageDF17 is a message at the format DF17
type MessageDF17 struct {
	MessageData
	Capability              Capability
	AddressAnnounced        AddressAnnounced
	MessageExtendedSquitter MessageExtendedSquitter
}

func (message *MessageDF17) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF17 parses a message at the DF17 format
func ParseDF17(message MessageData) (*MessageDF17, error) {

	// Format of the message is as follow:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 0 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 17 {
		return nil, errors.New("DF17 message must have a DownLinkFormat of 17")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF17 message must be 14 bytes long")
	}

	return &MessageDF17{
		MessageData:             message,
		Capability:              readCapability(message),
		AddressAnnounced:        readAddressAnnounced(message),
		MessageExtendedSquitter: readMessageExtendedSquitter(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF18
// -------------------------------------------------------------------------------------

// MessageDF18 is a message at the format DF18
type MessageDF18 struct {
	MessageData
	ControlField            ControlField
	AddressAnnounced        AddressAnnounced
	MessageExtendedSquitter MessageExtendedSquitter
}

func (message *MessageDF18) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF18 parses a message at the DF17 format
func ParseDF18(message MessageData) (*MessageDF18, error) {

	// Format of the message is as follow:
	//
	//     DF     CF   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 1 0 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 17 {
		return nil, errors.New("DF18 message must have a DownLinkFormat of 18")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF18 message must be 14 bytes long")
	}

	return &MessageDF18{
		MessageData:             message,
		ControlField:            readControlField(message),
		AddressAnnounced:        readAddressAnnounced(message),
		MessageExtendedSquitter: readMessageExtendedSquitter(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF19
// -------------------------------------------------------------------------------------

// MessageDF19 is a message at the format DF19
type MessageDF19 struct {
	MessageData
	ApplicationField ApplicationField
}

func (message *MessageDF19) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF19 parses a message at the DF19 format
func ParseDF19(message MessageData) (*MessageDF19, error) {

	// Format of the message is as follow:
	//
	//     DF     AF   | Military use
	// 1 0 0 1 1 a a a |   104 bits

	if message.DownLinkFormat != 19 {
		return nil, errors.New("DF19 message must have a DownLinkFormat of 19")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF19 message must be 14 bytes long")
	}

	return &MessageDF19{
		MessageData:      message,
		ApplicationField: readApplicationField(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF20
// -------------------------------------------------------------------------------------

// MessageDF20 is a message at the format DF20
type MessageDF20 struct {
	MessageData
	FlightStatus    FlightStatus
	DownlinkRequest DownlinkRequest
	UtilityMessage  UtilityMessage
	AltitudeCode    AltitudeCode
	MessageCommB    MessageCommB
}

func (message *MessageDF20) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF20 parses a message at the DF4 format
func ParseDF20(message MessageData) (*MessageDF20, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |  Comm-B |  AP/DP
	// 1 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 20 {
		return nil, errors.New("DF20 message must have a DownLinkFormat of 20")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF20 message must be 14 bytes long")
	}

	return &MessageDF20{
		MessageData:     message,
		FlightStatus:    readFlightStatus(message),
		DownlinkRequest: readDownlinkRequest(message),
		UtilityMessage:  readUtilityMessage(message),
		AltitudeCode:    readAltitudeCode(message),
		MessageCommB:    readMessageCommB(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF21
// -------------------------------------------------------------------------------------

// MessageDF21 is a message at the format DF21
type MessageDF21 struct {
	MessageData
	FlightStatus    FlightStatus
	DownlinkRequest DownlinkRequest
	UtilityMessage  UtilityMessage
	Identity        Identity
	MessageCommB    MessageCommB
}

func (message *MessageDF21) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF21 parses a message at the DF21 format
func ParseDF21(message MessageData) (*MessageDF21, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      ID    |        ID       |  Comm-B |  AP/DP
	// 1 0 1 0 1 f f f | d d d d d u u u | u u u i i i i i | i i i i i i i i | 56 bits | 24bits

	if message.DownLinkFormat != 21 {
		return nil, errors.New("DF21 message must have a DownLinkFormat of 21")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF21 message must be 14 bytes long")
	}

	return &MessageDF21{
		MessageData:     message,
		FlightStatus:    readFlightStatus(message),
		DownlinkRequest: readDownlinkRequest(message),
		UtilityMessage:  readUtilityMessage(message),
		Identity:        readIdentity(message),
		MessageCommB:    readMessageCommB(message),
	}, nil
}

// -------------------------------------------------------------------------------------
//                                         DF24
// -------------------------------------------------------------------------------------

// MessageDF24 is a message at the format DF24
type MessageDF24 struct {
	MessageData
	ControlELM       ControlELM
	NumberOfDSegment NumberOfDSegment
	MessageCommD     MessageCommD
}

func (message *MessageDF24) GetMessageData() MessageData {
	return message.GetMessageData()
}

// ParseDF24 parses a message at the DF24 format
func ParseDF24(message MessageData) (*MessageDF24, error) {

	// Format of the message is as follow:
	//
	//  DF _ KE   ND   |  Comm-B |   AP
	// 1 0 _ k n n n n | 80 bits | 24bits

	if message.DownLinkFormat&0x18 == 0x18 {
		return nil, errors.New("DF24 message must have a DownLinkFormat of 24")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF24 message must be 14 bytes long")
	}

	return &MessageDF24{
		MessageData:      message,
		ControlELM:       readControlELM(message),
		NumberOfDSegment: readNumberOfDSegment(message),
		MessageCommD:     readMessageCommD(message),
	}, nil
}
