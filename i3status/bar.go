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
	Output   chan string
}

func NewBar() *Bar {
	b := Bar{
		Input:    make(chan Message),
		Messages: make(map[string]Message),
		Output:   make(chan string),
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
		b.Output <- b.Message()
	}
}

func (b *Bar) readLoop() {
	var s string
	if len(b.subs) == 0 {
		return
	}
	for {
		fmt.Fscanln(b.In, &s)
		s = strings.TrimPrefix(s, ",")
		for _, c := range b.subs {
			c <- NewEntry(s)
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
