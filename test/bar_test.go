package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBarCreate(t *testing.T) {
	Convey("Given a channel", t, func() {

		c := make(chan i3status.Message)
		Convey("When a Bar is created", func() {
			b := i3status.NewBar(c)
			Convey("it has an input channel", func() {
				So(b.Input, ShouldEqual, c)
			})
		})
	})
}

func TestBarStart(t *testing.T) {
	Convey("Given a bar", t, func() {

		c := make(chan i3status.Message)
		w1 := i3status.NewWidget(c)
		w2 := i3status.NewWidget(c)
		w1.Start()
		w2.Start()

		b := i3status.NewBar(c)

		Convey("When a Bar is started", func() {
			b.Start()
			Convey("it gets messages from the widgets", func() {
				time.Sleep(1 * 1e9)
				So(len(b.Messages), ShouldEqual, 2)
			})
		})
	})
}

func TestBarMessage(t *testing.T) {
	Convey("Given a bar", t, func() {
		c := make(chan i3status.Message)
		b := i3status.NewBar(c)

		Convey("When it's just created", func() {
			Convey("it has an empty message", func() {
				So(b.Message(), ShouldEqual, "[]")
			})
		})

		Convey("When it's started and widgets are running", func() {
			w1 := i3status.NewWidget(c)
			w2 := i3status.NewWidget(c)
			w1.Start()
			w2.Start()
			Convey("it has a message", func() {
				b.Start()
				time.Sleep(1 * 1e9)
				json := `[{"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"3","urgent":false,"separator":true,"separator_block_width":10}, {"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"4","urgent":false,"separator":true,"separator_block_width":10}]`
				So(b.Message(), ShouldEqual, json)
			})
		})
	})
}
