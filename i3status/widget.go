package i3status

import (
	"strconv"
	"time"
)

type Widget interface {
	Start()
}

type BaseWidget struct {
	Output   chan Message
	Refresh  time.Duration
	Instance int
}

var instanceCount int

func NewBaseWidget(output chan Message) *BaseWidget {
	instanceCount++
	w := BaseWidget{
		Output:   output,
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
