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
			Convey("timer status is started", func() {
				So(w.Status, ShouldContainSubstring, "started")
				msg = <-c
				So(msg.FullText, ShouldContainSubstring, "started")
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
				So(w.Status, ShouldContainSubstring, "paused")
				So(msg.FullText, ShouldContainSubstring, "paused")
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
				So(w.Status, ShouldContainSubstring, "stopped")
				So(msg.FullText, ShouldContainSubstring, "stopped")
			})
		})
	})
}
