package i3status

import (
	"strconv"
	"time"
)

type Widget struct {
	Output   chan Message
	Refresh  time.Duration
	Instance int
}

var instanceCount int

func NewWidget(output chan Message) *Widget {
	instanceCount++
	w := Widget{
		Output:   output,
		Refresh:  1000,
		Instance: instanceCount,
	}
	return &w
}

func (w *Widget) basicLoop() {
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

func (w *Widget) Start() {
	go w.basicLoop()
}
