package i3status

import (
	"fmt"
	"strconv"
)

type OnOffWidget struct {
	BaseWidget
	Input chan Entry
	On    bool
}

func NewOnOffWidget(output chan Message, input chan Entry) *OnOffWidget {
	instanceCount++
	w := OnOffWidget{
		BaseWidget{
			output,
			1000,
			instanceCount,
		},
		input,
		false,
	}
	return &w
}

func (w *OnOffWidget) message() string {
	if w.On {
		return "Pants On"
	} else {
		return "Pants Off"
	}
}

func (w *OnOffWidget) sendMessage() {
	msg := NewMessage()
	msg.Name = "OnOff"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	msg.FullText = fmt.Sprint(w.message())
	w.Output <- *msg
}

func (w *OnOffWidget) readLoop() {
	<-w.Input
	if w.On {
		w.On = false
	} else {
		w.On = true
	}
	go w.sendMessage()
}

func (w *OnOffWidget) Start() {
	go w.sendMessage()
	go w.readLoop()
}
