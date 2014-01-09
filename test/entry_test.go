package i3status_test

import (
	"github.com/ghedamat/go-i3status/i3status"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEntry(t *testing.T) {
	Convey("Given a json string", t, func() {
		str := `{
      "name": "ethernet",
      "instance": "eth0",
      "button": 1,
      "x": 1320,
      "y": 1400
    }`
		Convey("When NewEntry is called", func() {
			e := i3status.NewEntry(str)
			Convey("a new Entry is created", func() {
				So(e.Name, ShouldEqual, "ethernet")
				So(e.Instance, ShouldEqual, "eth0")
				So(e.Button, ShouldEqual, 1)
			})
		})
	})
}
