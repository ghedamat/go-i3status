package i3status

import (
	"fmt"
	"strconv"
	"time"
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
	Status    string
	StartTime time.Time
	Minutes   float64
}

func NewTimerWidget() *TimerWidget {
	instanceCount++
	w := TimerWidget{
		BaseWidget: *NewBaseWidget(),
		Status:     "stopped",
	}
	return &w
}

func (w *TimerWidget) message() string {
	min := 0.0
	if w.Status == "running" {
		min = w.Minutes + time.Since(w.StartTime).Minutes()
	} else {
		min = w.Minutes
	}
	return fmt.Sprintf("Timer %s %.2f", w.Status, min)
}

func (w *TimerWidget) sendMessage() {
	msg := NewMessage()
	msg.Name = "Timer"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	msg.FullText = fmt.Sprint(w.message())
	w.Output <- *msg
}

func (w *TimerWidget) run() {
	w.Status = "running"
	w.StartTime = time.Now()
}
func (w *TimerWidget) pause() {
	w.Status = "paused"
	elapsed := time.Since(w.StartTime)
	w.Minutes = w.Minutes + elapsed.Minutes()
}
func (w *TimerWidget) resume() {
	w.Status = "running"
	w.StartTime = time.Now()
}
func (w *TimerWidget) stop() {
	w.Status = "stopped"
	w.Minutes = 0
}

func (w *TimerWidget) toggleStatus(button int) {
	if button == LeftButton {
		if w.Status == "stopped" {
			w.run()
		} else if w.Status == "paused" {
			w.resume()
		} else if w.Status == "running" {
			w.pause()
		}
	} else if button == RightButton {
		w.stop()
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

func (w *TimerWidget) msgLoop() {
	for {
		time.Sleep(1 * 1000000000)
		go w.sendMessage()
	}
}

func (w *TimerWidget) Start() {
	go w.readLoop()
	go w.msgLoop()
}
