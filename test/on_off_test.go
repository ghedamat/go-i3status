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
