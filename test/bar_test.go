package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"time"
)

func TestBarAdd(t *testing.T) {
	Convey("Given a bar", t, func() {
		b := i3status.NewBar()

		w1 := i3status.NewBaseWidget()
		w2 := i3status.NewBaseWidget()
		Convey("When widgets are added", func() {
			b.Add(w1)
			<-b.Output
			b.Add(w2)
			<-b.Output
			Convey("it gets messages from the widgets", func() {
				time.Sleep(1)
				So(len(b.Messages), ShouldEqual, 2)
			})
		})
	})
}

func TestBarMessage(t *testing.T) {
	Convey("Given a bar", t, func() {
		b := i3status.NewBar()

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
				<-b.Output
				b.Add(w2)
				<-b.Output
				time.Sleep(1)
				json := `[{"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"3","urgent":false,"separator":true,"separator_block_width":10}, {"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"4","urgent":false,"separator":true,"separator_block_width":10}]`
				So(b.Message(), ShouldEqual, json)
			})
		})
	})
}

func TestBarOutput(t *testing.T) {
	Convey("Given a bar", t, func() {
		b := i3status.NewBar()

		Convey("When it's started and widgets are running", func() {
			w1 := i3status.NewBaseWidget()
			w2 := i3status.NewBaseWidget()
			Convey("it has a message", func() {
				b.Add(w1)
				msg := <-b.Output
				b.Add(w2)
				msg = <-b.Output
				time.Sleep(1)
				json := `[{"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"5","urgent":false,"separator":true,"separator_block_width":10}, {"full_text":"Basic Widget","short_text":"","color":"#ffffff","min_width":0,"align":"left","name":"Basic","instance":"6","urgent":false,"separator":true,"separator_block_width":10}]`
				So(msg, ShouldEqual, json)
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
		b := i3status.NewBar()
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
		b := i3status.NewBar()
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

func TestBarSendsMultipleEntries(t *testing.T) {
	Convey("Given a subscriber with a channel", t, func() {
		b := i3status.NewBar()
		w1 := i3status.NewBaseWidget()
		b.Add(w1)
		Convey("when a multiple messages are sent", func() {

			Convey("an Entry is sent on the channel", func() {
				input := `{"name":"test","instance":"eth0","button":1,"x":1320,"y":1400}`
				b.In = strings.NewReader(input)
				time.Sleep(1)
				en := <-w1.Input
				So(en.Name, ShouldEqual, "test")
				input = `{"name":"test1","instance":"eth0","button":1,"x":1320,"y":1400}`
				b.In = strings.NewReader(input)
				time.Sleep(1)
				en = <-w1.Input
				So(en.Name, ShouldEqual, "test1")
				input = `{"name":"test2","instance":"eth0","button":1,"x":1320,"y":1400}`
				b.In = strings.NewReader(input)
				time.Sleep(1)
				en = <-w1.Input
				So(en.Name, ShouldEqual, "test2")
			})
		})
	})
}
