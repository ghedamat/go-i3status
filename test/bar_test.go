package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
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

func TestBarAdd(t *testing.T) {
	Convey("Given a bar", t, func() {
		c := make(chan i3status.Message)
		b := i3status.NewBar(c)

		w1 := i3status.NewBaseWidget()
		w2 := i3status.NewBaseWidget()
		Convey("When widgets are added", func() {
			b.Add(w1)
			b.Add(w2)
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
			w1 := i3status.NewBaseWidget()
			w2 := i3status.NewBaseWidget()
			Convey("it has a message", func() {
				b.Add(w1)
				b.Add(w2)
				time.Sleep(1 * 1e9)
				json := `[{"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"3","urgent":false,"separator":true,"separator_block_width":10}, {"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"4","urgent":false,"separator":true,"separator_block_width":10}]`
				So(b.Message(), ShouldEqual, json)
			})
		})
	})
}

func TestBarOrder(t *testing.T) {
	Convey("Given some widgets", t, func() {
		//w1 := i3status.NewWidget(c)
		//w2 := i3status.NewWidget(c)
		Convey("When a bar is created", func() {
			Convey("widget are rendered in order", func() {
			})
		})
	})
}

func TestBarAddLength(t *testing.T) {
	Convey("Given a Bar", t, func() {
		c := make(chan i3status.Message)
		b := i3status.NewBar(c)
		Convey("when a widget is added", func() {
			w1 := i3status.NewBaseWidget()
			b.Add(w1)
			Convey("length is 1", func() {
				So(b.Len(), ShouldEqual, 1)
			})
		})
	})
}
func TestBarSendsEntries(t *testing.T) {
	Convey("Given a subscriber with a channel", t, func() {
		c := make(chan i3status.Message)
		b := i3status.NewBar(c)
		w1 := i3status.NewBaseWidget()
		b.Add(w1)
		Convey("when a message arrives", func() {
			input := `{"name":"test","instance":"eth0","button":1,"x":1320,"y":1400}`
			b.In = strings.NewReader(input)

			Convey("an Entry is sent on the channel", func() {
				en := <-w1.Input
				So(en.Name, ShouldEqual, "test")
			})
		})
	})
}
