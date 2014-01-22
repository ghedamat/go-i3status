package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOnOffWidgetConstructor(t *testing.T) {
	Convey("Given an input and an output channel", t, func() {
		bar := i3status.NewBar()
		Convey("When OnOff is created", func() {
			w := i3status.NewOnOffWidget()
			bar.Add(w)
			Convey("output channel is available", func() {
				So(w.Output, ShouldEqual, bar.Input)
			})
		})
	})
}

func TestOnOffWidget(t *testing.T) {
	Convey("Given an OnOff Widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewOnOffWidget()
		w.SetChannels(c, i)
		w.Start()
		Convey("When an entry is sent", func() {
			w.Input <- i3status.Entry{}
			Convey("widget status goes on", func() {
				So(w.On, ShouldEqual, true)
			})
		})
	})
}
func TestOnOffWidgetMessage(t *testing.T) {
	Convey("Given an OnOff Widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewOnOffWidget()
		w.SetChannels(c, i)
		w.Start()
		Convey("When and entry is sent", func() {
			msg := <-c
			So(msg.FullText, ShouldEqual, "Pants Off")
			w.Input <- i3status.Entry{}
			Convey("widget message is On", func() {
				msg = <-c
				So(msg.FullText, ShouldEqual, "Pants On")
			})
		})
	})
}
