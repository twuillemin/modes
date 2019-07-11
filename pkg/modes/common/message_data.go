package common

// MessageData is the basic message structure that is applicable to all messages, except DF24
type MessageData struct {
	DownLinkFormat uint8
	FirstField     uint8
	Payload        []uint8
	Parity         []uint8
}
