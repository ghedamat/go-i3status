package i3status

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

const (
	Charging    = "charging"
	Discharging = "discharging"
	Charged     = "charged"
)

type PowerWidget struct {
	BaseWidget
}

func NewPowerWidget() *PowerWidget {
	instanceCount++
	w := PowerWidget{
		BaseWidget: *NewBaseWidget(),
	}
	return &w
}

func (w *PowerWidget) execCommand() string {
	var out bytes.Buffer
	cmd := exec.Command("acpi", "-b")
	cmd.Stdout = &out
	cmd.Run()
	str, _ := out.ReadString('\n')
	return str
}

func (w *PowerWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Power"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		str := w.execCommand()
		msg.FullText = fmt.Sprintf("%s", str)
		w.Output <- *msg
		time.Sleep(5000 * time.Millisecond)
	}
}

func (w *PowerWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
