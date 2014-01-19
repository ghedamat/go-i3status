package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWidgetConstructor(t *testing.T) {
	Convey("Given a Message channel", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		Convey("When a widget is created and a channels are set", func() {
			widget := i3status.NewBaseWidget()
			widget.SetChannels(c, i)
			Convey("the widget has an output Message channel", func() {
				So(widget.Output, ShouldEqual, c)
			})
			Convey("the widget has an instance identifier", func() {
				So(widget.Instance, ShouldNotEqual, 0)
			})
		})
	})

	Convey("Given two Widgets", t, func() {
		w1 := i3status.NewBaseWidget()
		w2 := i3status.NewBaseWidget()
		Convey("When they are created", func() {
			Convey("they have different instance identifiers", func() {
				So(w1.Instance, ShouldNotEqual, w2.Instance)
			})
		})
	})

}

func TestWidgetSendMessage(t *testing.T) {
	Convey("Given a Widget", t, func() {
		c := make(chan i3status.Message)
		i := make(chan i3status.Entry)
		widget := i3status.NewBaseWidget()
		widget.SetChannels(c, i)
		Convey("When a widget is started", func() {
			widget.Start()
			Convey("it sends a Message to the channel", func() {
				msg := <-c
				So(msg.FullText, ShouldEqual, "Basic Widget")
			})
		})
	})
}

func TestWidgetInterface(t *testing.T) {
	Convey("Given two different Widgets", t, func() {
		Convey("they have the same interface", func() {
			w1 := i3status.NewBaseWidget()
			w2 := i3status.NewDateWidget()

			arr := make([]i3status.Widget, 2)
			arr[0] = w1
			arr[1] = w2
		})
	})
}
