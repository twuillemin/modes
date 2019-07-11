package messages

type Message interface {
	GetName() string
	GetDownLinkFormat() int
	PrettyPrint()
}
