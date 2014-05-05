package i3status

import (
	"fmt"
	forecast "github.com/mlbright/forecast/v2"
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
	"strings"
	"time"
)

type WeatherWidget struct {
	BaseWidget
}

func NewWeatherWidget() *WeatherWidget {
	instanceCount++
	w := WeatherWidget{
		BaseWidget: *NewBaseWidget(),
	}
	return &w
}

func (w *WeatherWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Wlan"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	usr, _ := user.Current()
	keybytes, err := ioutil.ReadFile(usr.HomeDir + "/.forecast.io.rc")
	if err != nil {
		log.Fatal(err)
	}
	key := string(keybytes)
	key = strings.TrimSpace(key)

	for {
		msg.FullText, msg.Color = w.getStatus(key)
		w.Output <- *msg
		time.Sleep(5000 * time.Millisecond)
	}
}

func (w *WeatherWidget) getStatus(key string) (string, string) {

	lat := "43.6595"
	long := "-79.3433"

	f, err := forecast.Get(key, lat, long, "now", forecast.CA)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s: %s %.2f C", f.Timezone, f.Currently.Summary, f.Currently.Temperature), BLUE
}

func (w *WeatherWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}
