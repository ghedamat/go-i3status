package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWlanWidgetConstructor(t *testing.T) {
	Convey("Given two channels", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		Convey("When wlan is created", func() {
			w := i3status.NewWlanWidget()
			w.SetChannels(c, i)
			Convey("output channel is available", func() {
				So(w.Output, ShouldEqual, c)
			})
		})
	})
}

func TestWlanWidgetHasMessage(t *testing.T) {
	Convey("Given a widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		w := i3status.NewWlanWidget()
		w.SetChannels(c, i)
		Convey("When wlan is started", func() {
			w.Start()
			Convey("output message is available", func() {
				msg := <-c
				So(msg.FullText, ShouldContainSubstring, "WLAN")
			})
		})
	})
}
