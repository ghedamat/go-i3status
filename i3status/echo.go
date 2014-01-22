package i3status

import (
	"fmt"
	"strconv"
)

type EchoWidget struct {
	BaseWidget
}

func NewEchoWidget() *EchoWidget {
	instanceCount++
	w := EchoWidget{
		*NewBaseWidget(),
	}
	return &w
}

func (w *EchoWidget) sendMessage(e Entry) {
	msg := NewMessage()
	msg.Name = "Echo"
	msg.Color = "#00ffff"
	msg.Instance = strconv.Itoa(w.Instance)
	msg.FullText = fmt.Sprint(e)
	w.Output <- *msg
}

func (w *EchoWidget) readLoop() {
	for {
		e := <-w.Input
		go w.sendMessage(e)
	}
}

func (w *EchoWidget) Start() {
	go w.sendMessage(Entry{})
	go w.readLoop()
}
