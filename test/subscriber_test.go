package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestSubscribe(t *testing.T) {
	Convey("Given a subscriber and an Entry channel", t, func() {
		sub := new(i3status.Subscriber)
		c := make(chan i3status.Entry)
		Convey("when a channel c subscribed", func() {
			sub.Subscribe(c)
			Convey("length is 1", func() {
				So(sub.Len(), ShouldEqual, 1)
			})
		})
	})
	Convey("Given a subscriber with a channel", t, func() {
		sub := new(i3status.Subscriber)
		c := make(chan i3status.Entry)
		sub.Subscribe(c)
		Convey("when a message arrives", func() {
			input := `{"name":"test","instance":"eth0","button":1,"x":1320,"y":1400}`
			sub.In = strings.NewReader(input)
			sub.Start()

			Convey("an Entry is sent on the channel", func() {
				en := <-c
				So(en.Name, ShouldEqual, "test")
			})
		})
	})
}
