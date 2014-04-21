package i3status

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	BATTERY_DISCHARGING = iota
	BATTERY_CHARGING
	BATTERY_FULL
)

type BatteryInfo struct {
	PercentRemaining float64
	SecondsRemaining float64
	Consumption      float64
	status           int
}

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

func (w *PowerWidget) getStatus() (string, string) {
	bi, _ := readBatteryInfo(0)

	bar := MakeBar(bi.PercentRemaining, 20)
	remaining := HumanDuration(int64(bi.SecondsRemaining))
	prefix := "BAT"
	color := WHITE
	if bi.IsCharging() {
		prefix = "CHR"
		color = GREEN
	} else if bi.IsFull() {
		prefix = "FULL"
	} else {
		if bi.PercentRemaining < 10 {
			color = RED
		} else if bi.PercentRemaining < 30 {
			color = YELLOW
		}
	}
	return fmt.Sprintf("%s %s %s %0.1f%%", prefix, remaining, bar, bi.PercentRemaining), color
}

func (w *PowerWidget) basicLoop() {
	msg := NewMessage()
	msg.Name = "Power"
	msg.Color = "#ffffff"
	msg.Instance = strconv.Itoa(w.Instance)
	for {
		msg.FullText, msg.Color = w.getStatus()
		w.Output <- *msg
		time.Sleep(5000 * time.Millisecond)
	}
}

func (w *PowerWidget) Start() {
	go w.basicLoop()
	go w.readLoop()
}

// stolen from https://github.com/coleifer/mastodon/blob/master/battery.go
func (batteryInfo *BatteryInfo) IsCharging() bool {
	return batteryInfo.status == BATTERY_CHARGING
}

func (batteryInfo *BatteryInfo) IsFull() bool {
	return batteryInfo.status == BATTERY_FULL
}

func readBatteryInfo(battery int) (*BatteryInfo, error) {
	rawInfo := make(map[string]string)
	batteryInfo := new(BatteryInfo)

	path := fmt.Sprintf("/sys/class/power_supply/BAT%d/uevent", battery)
	if !FileExists(path) {
		return batteryInfo, errors.New("Battery not found")
	}
	callback := func(line string) bool {
		data := strings.Split(string(line), "=")
		rawInfo[data[0]] = data[1]
		return true
	}
	ReadLines(path, callback)

	var remaining, presentRate, voltage, fullDesign float64
	var wattAsUnit bool
	batteryInfo.status = BATTERY_DISCHARGING

	if rawInfo["POWER_SUPPLY_STATUS"] == "Charging" {
		batteryInfo.status = BATTERY_CHARGING
	} else if rawInfo["POWER_SUPPLY_STATUS"] == "Full" {
		batteryInfo.status = BATTERY_FULL
	}

	/* Convert to float shorthand */
	pf := func(keys ...string) float64 {
		for _, key := range keys {
			if _, ok := rawInfo[key]; ok {
				f, _ := strconv.ParseFloat(rawInfo[key], 64)
				return f
			}
		}
		return 0.
	}

	/* Read values from file */
	remaining = pf("POWER_SUPPLY_ENERGY_NOW", "POWER_SUPPLY_CHARGE_NOW")
	presentRate = pf("POWER_SUPPLY_CURRENT_NOW", "POWER_SUPPLY_POWER_NOW")
	voltage = pf("POWER_SUPPLY_VOLTAGE_NOW")
	fullDesign = pf("POWER_SUPPLY_CHARGE_FULL_DESIGN", "POWER_SUPPLY_ENERGY_FULL_DESIGN")
	_, wattAsUnit = rawInfo["POWER_SUPPLY_ENERGY_NOW"]

	if !wattAsUnit {
		presentRate = (voltage / 1000.0) * (presentRate / 1000.0)
		remaining = (voltage / 1000.0) * (remaining / 1000.0)
		fullDesign = (voltage / 1000.0) * (fullDesign / 1000.0)
	}

	if fullDesign == 0 {
		return batteryInfo, errors.New("Battery full design missing")
	}

	batteryInfo.PercentRemaining = (remaining / fullDesign) * 100

	var remainingTime float64
	if presentRate > 0 {
		if batteryInfo.status == BATTERY_CHARGING {
			remainingTime = (fullDesign - remaining) / presentRate
		} else if batteryInfo.status == BATTERY_DISCHARGING {
			remainingTime = remaining / presentRate
		}
		batteryInfo.SecondsRemaining = remainingTime * 3600
	}
	batteryInfo.Consumption = presentRate / 1000000.0

	return batteryInfo, nil
}
