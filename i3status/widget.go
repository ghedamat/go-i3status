package i3status

import (
	"strconv"
	"time"
)

type BaseWidget struct {
	Output   chan Message
	Input    chan Entry
	Refresh  time.Duration
	Instance int
}

var instanceCount int

func (w *BaseWidget) SetChannels(out chan Message, in chan Entry) {
	w.Output = out
	w.Input = in
}

func NewBaseWidget() *BaseWidget {
	instanceCount++
	w := BaseWidget{
		Output:   nil,
		Input:    nil,
		Refresh:  1000,
		Instance: instanceCount,
	}
	return &w
}

func (w *BaseWidget) basicLoop() {
	msg := NewMessage()
	msg.FullText = "Basic Widget"
	msg.Name = "Basic"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		w.Output <- *msg
		time.Sleep(w.Refresh * time.Millisecond)
	}
}

func (w *BaseWidget) Start() {
	go w.basicLoop()
}
