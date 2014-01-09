package i3status

import (
	"fmt"
	"strconv"
	"time"
)

type OnOffWidget struct {
	BaseWidget
	Input chan Entry
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
	}
	return &w
}

func (w *OnOffWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Date"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		msg.FullText = fmt.Sprintf("%s", time.Now())
		w.Output <- *msg
		time.Sleep(w.Refresh * time.Millisecond)
	}
}
func (w *OnOffWidget) Start() {
	go w.basicLoop()
}
