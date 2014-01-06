package i3status

import (
	"strings"
)

type Bar struct {
	Input    chan Message
	Messages map[string]Message
}

func NewBar(c chan Message) *Bar {
	b := Bar{
		Input:    c,
		Messages: make(map[string]Message),
	}
	return &b
}

func (b *Bar) barLoop() {
	for {
		msg := <-b.Input
		b.Messages[msg.Name+msg.Instance] = msg
	}
}

func (b *Bar) Start() {
	go b.barLoop()
}

func (b *Bar) Message() string {
	str := "["
	for _, m := range b.Messages {
		str += m.ToJson() + ", "
	}

	return strings.TrimSuffix(str, ", ") + "]"
}
