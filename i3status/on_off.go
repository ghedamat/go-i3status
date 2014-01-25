package i3status

import (
	"fmt"
	"strconv"
)

type OnOffWidget struct {
	BaseWidget
	On bool
}

func NewOnOffWidget() *OnOffWidget {
	instanceCount++
	w := OnOffWidget{
		*NewBaseWidget(),
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
	for {
		i := <-w.Input
		if i.Name == "OnOff" {
			if w.On {
				w.On = false
			} else {
				w.On = true
			}
			go w.sendMessage()
		}
	}
}

func (w *OnOffWidget) Start() {
	go w.sendMessage()
	go w.readLoop()
}
