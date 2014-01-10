package i3status

import (
	"fmt"
	"strconv"
)

type OnOffWidget struct {
	BaseWidget
	Input chan Entry
	On    bool
	Sub   *Subscriber
}

func NewOnOffWidget(output chan Message, sub *Subscriber) *OnOffWidget {
	instanceCount++
	input := make(chan Entry)
	w := OnOffWidget{
		BaseWidget{
			output,
			1000,
			instanceCount,
		},
		input,
		false,
		sub,
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
	for {
		<-w.Input
		if w.On {
			w.On = false
		} else {
			w.On = true
		}
		go w.sendMessage()
	}
}

func (w *OnOffWidget) Start() {
	w.Sub.Subscribe(w.Input)
	go w.sendMessage()
	go w.readLoop()
}
