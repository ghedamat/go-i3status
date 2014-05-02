package i3status

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type WlanInfo struct {
	NetworkStatus string
	NetworkName   string
	color         string
}

type WlanWidget struct {
	BaseWidget
}

func NewWlanWidget() *WlanWidget {
	instanceCount++
	w := WlanWidget{
		BaseWidget: *NewBaseWidget(),
	}
	return &w
}

func (w *WlanWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Wlan"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		msg.FullText, msg.Color = w.getStatus()
		w.Output <- *msg
		time.Sleep(5000 * time.Millisecond)
	}
}

func (w *WlanWidget) getStatus() (string, string) {
	info := WlanInfo{}
	cmd := exec.Command("nm-tool")
	str, _ := cmd.Output()
	text := fmt.Sprintf("%s", str)
	lines := strings.Split(text, "\n")
	re := regexp.MustCompile("[a-zA-Z]+\\s*([a-z]+)")
	info.NetworkStatus = re.FindAllString(lines[3], -1)[1]

	if info.NetworkStatus == "connecting" {
		info.color = YELLOW
	}

	if info.NetworkStatus == "connected" {
		info.color = GREEN
	}

	if info.NetworkStatus == "disconnected" {
		info.color = RED
	} else {
		re = regexp.MustCompile("\\[.*\\]")
		matches := re.FindAllString(lines[5], -1)
		if len(matches) > 0 {
			info.NetworkName = matches[0]
		}
	}

	return fmt.Sprintf("WLAN: %s %s", info.NetworkStatus, info.NetworkName), info.color
}

func (w *WlanWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
