package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOnOffWidgetConstructor(t *testing.T) {
	Convey("Given an input and an output channel", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		Convey("When OnOff is created", func() {
			w := i3status.NewOnOffWidget(c, i)
			Convey("input and output channel are available", func() {
				So(w.Input, ShouldEqual, i)
				So(w.Output, ShouldEqual, c)
			})
		})
	})
}

func TestOnOffWidget(t *testing.T) {
	Convey("Given an OnOff Widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewOnOffWidget(c, i)
		w.Start()
		Convey("When and entry is sent", func() {
			i <- i3status.Entry{}
			Convey("widget status goes on", func() {
				So(w.On, ShouldEqual, true)
			})
		})
	})
	Convey("Given an OnOff Widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewOnOffWidget(c, i)
		w.Start()
		Convey("When and entry is sent", func() {
			msg := <-c
			So(msg.FullText, ShouldEqual, "Pants Off")
			i <- i3status.Entry{}
			Convey("widget message is On", func() {
				msg := <-c
				So(msg.FullText, ShouldEqual, "Pants On")
			})
		})
	})
}
