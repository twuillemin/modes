package messages

type MessageModeS interface {
	GetName() string
	GetDownLinkFormat() int
	PrettyPrint()
}
