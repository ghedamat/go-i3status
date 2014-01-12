package i3status

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type I3statusWidget struct {
	BaseWidget
}

func NewI3statusWidget() *I3statusWidget {
	instanceCount++
	w := I3statusWidget{
		*NewBaseWidget(),
	}
	return &w
}

func (w *I3statusWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "I3status"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	cmd := exec.Command("i3status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Start()
	for {
		str, _ := out.ReadString('\n')
		msg.FullText = fmt.Sprintf("%s", str)
		w.Output <- *msg
		time.Sleep(5000 * time.Millisecond)
	}
}
func (w *I3statusWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
