package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTimerWidgetConstructor(t *testing.T) {
	Convey("Given two channels", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		Convey("When timer is created", func() {
			w := i3status.NewTimerWidget()
			w.SetChannels(c, i)
			Convey("output channel is available", func() {
				So(w.Output, ShouldEqual, c)
			})
		})
	})
}

func TestTimerWidgetStartEvent(t *testing.T) {
	Convey("Given A timer widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewTimerWidget()
		w.SetChannels(c, i)
		w.Start()
		Convey("When a left click event is received", func() {
			msg := <-c
			i <- i3status.Entry{"Timer", "0", 1, 0, 0}
			time.Sleep(1)
			Convey("timer status is running", func() {
				So(w.Status, ShouldEqual, "running")
				msg = <-c
				So(msg.FullText, ShouldEqual, "Timer running")
			})
		})
	})
}
func TestTimerWidgetPauseEvent(t *testing.T) {
	Convey("Given A timer widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewTimerWidget()
		w.SetChannels(c, i)
		w.Start()
		Convey("When a second left click event is received", func() {
			msg := <-c
			i <- i3status.Entry{"Timer", "0", 1, 0, 0}
			msg = <-c
			i <- i3status.Entry{"Timer", "0", 1, 0, 0}
			time.Sleep(1)
			Convey("timer status is paused", func() {
				msg = <-c
				So(w.Status, ShouldEqual, "paused")
				So(msg.FullText, ShouldEqual, "Timer paused")
			})
		})
	})
}
func TestTimerWidgetStopEvent(t *testing.T) {
	Convey("Given A timer widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewTimerWidget()
		w.SetChannels(c, i)
		w.Start()
		Convey("When a right click event is received", func() {
			msg := <-c
			i <- i3status.Entry{"Timer", "0", 1, 0, 0}
			msg = <-c
			i <- i3status.Entry{"Timer", "0", 3, 0, 0}
			time.Sleep(1)
			Convey("timer status is stopped", func() {
				msg = <-c
				So(w.Status, ShouldEqual, "stopped")
				So(msg.FullText, ShouldEqual, "Timer stopped")
			})
		})
	})
}
