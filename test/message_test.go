package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testMsg = i3status.Message{
	FullText:            "fullMsg",
	ShortText:           "shortMsg",
	Color:               "#ff00ff",
	MinWidth:            200,
	Align:               "center",
	Name:                "testblock",
	Instance:            "first",
	Urgent:              true,
	Separator:           true,
	SeparatorBlockWidth: 20,
}

func TestHasMessageType(t *testing.T) {
	Convey("The message struct exists", t, func() {
		So(testMsg.FullText, ShouldEqual, "fullMsg")
	})
}

func TestConvertMessageToJson(t *testing.T) {
	Convey("Given a Message", t, func() {
		msg := i3status.Message{FullText: "fullMsg"}
		Convey("When a message is converted to json", func() {
			json := msg.ToJson()
			Convey("valid json is produced", func() {
				res := `{"full_text":"fullMsg","short_text":"","color":"","min_width":0,"align":"","name":"","instance":"","urgent":false,"separator":false,"separator_block_width":0}`
				So(json, ShouldEqual, res)
			})
		})
	})
}

func TestMessageConstructor(t *testing.T) {
	Convey("Given the Message Constructor", t, func() {
		Convey("When a message is created", func() {
			msg := i3status.NewMessage()
			Convey("it has sane defaults", func() {
				So(msg.Separator, ShouldEqual, true)
				So(msg.Align, ShouldEqual, "left")
			})
		})
	})
}
