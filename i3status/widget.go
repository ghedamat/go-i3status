package i3status

type Widget interface {
	Start()
	SetChannels(chan Message, chan Entry)
}

var instanceCount int
