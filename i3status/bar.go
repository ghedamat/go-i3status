package i3status

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Bar struct {
	Input    chan Message
	Messages map[string]Message
	subs     [](chan Entry)
	In       io.Reader
}

func NewBar(c chan Message) *Bar {
	b := Bar{
		Input:    c,
		Messages: make(map[string]Message),
	}
	b.start()
	return &b
}

func (b *Bar) Add(w Widget) {
	in := make(chan Entry)
	w.SetChannels(b.Input, in)
	b.subs = append(b.subs, in)
	w.Start()
}

func (b *Bar) barLoop() {
	for {
		msg := <-b.Input
		b.Messages[msg.Name+msg.Instance] = msg
	}
}

func (b *Bar) readLoop() {
	var i string
	if len(b.subs) == 0 {
		return
	}
	for {
		fmt.Fscanf(b.In, "%s", &i)
		for _, c := range b.subs {
			c <- *NewEntry(i)
		}
	}
}

func (b *Bar) start() {
	if b.In == nil {
		b.In = os.Stdin
	}
	go b.barLoop()
	go b.readLoop()
}

func (b *Bar) Len() int {
	return len(b.subs)
}

func (b *Bar) Message() string {
	str := "["
	for _, m := range b.Messages {
		str += m.ToJson() + ", "
	}

	return strings.TrimSuffix(str, ", ") + "]"
}
