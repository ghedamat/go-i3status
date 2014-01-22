package i3status

import (
	"fmt"
	"strconv"
)

const (
	LeftButton   = 1
	MiddleButton = 2
	RightButton  = 3
	Started      = "started"
	Stopped      = "stopped"
	Paused       = "paused"
)

type TimerWidget struct {
	BaseWidget
	Status string
}

func NewTimerWidget() *TimerWidget {
	instanceCount++
	w := TimerWidget{
		*NewBaseWidget(),
		"stopped",
	}
	return &w
}

func (w *TimerWidget) message() string {
	return "Timer " + w.Status
}

func (w *TimerWidget) sendMessage() {
	msg := NewMessage()
	msg.Name = "Timer"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	msg.FullText = fmt.Sprint(w.message())
	w.Output <- *msg
}

func (w *TimerWidget) toggleStatus(button int) {
	if button == LeftButton {
		if w.Status == "stopped" || w.Status == "paused" {
			w.Status = "running"
		} else if w.Status == "running" {
			w.Status = "paused"
		}
	} else if button == RightButton {
		w.Status = "stopped"
	}

}

func (w *TimerWidget) readLoop() {
	for {
		i := <-w.Input
		if i.Name == "Timer" {
			w.toggleStatus(i.Button)
			go w.sendMessage()
		}
	}
}

func (w *TimerWidget) Start() {
	go w.sendMessage()
	go w.readLoop()
}
