package i3status

import (
	"fmt"
	"strconv"
	"time"
)

type DateWidget struct {
	BaseWidget
}

func NewDateWidget() *DateWidget {
	instanceCount++
	w := DateWidget{
		*NewBaseWidget(),
	}
	return &w
}

func (w *DateWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Date"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	const layout = "3:04pm, Jan 2, 2006"
	for {
		msg.FullText = fmt.Sprintf("%s", time.Now().Format(layout))
		w.Output <- *msg
		time.Sleep(w.Refresh * time.Millisecond)
	}
}
func (w *DateWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
