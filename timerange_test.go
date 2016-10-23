package timerange

import (
	"flag"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimerange(t *testing.T) {
	Convey("", t, func() {
		Convey("It implements flag.Value interface", func() {
			var t Timerange
			flag.Var(&t, "time", "")
		})

		Convey("It returns error if no given value", func() {
			t := NewTimerange()
			So(t.Set(""), ShouldEqual, ErrValueIsEmpty)
		})

		Convey("It returns error if value is not parseable as time.Time", func() {
			t := NewTimerange()
			err := t.Set("foobar")

			So(err, ShouldNotBeNil)
			_, ok := err.(*time.ParseError)
			So(ok, ShouldBeTrue)
		})

		Convey("It parses and stores timestamp when given a timestamp", func() {
			t := NewTimerange()
			So(t.Set("2016/10/22"), ShouldBeNil)

			So(len(t.TimeValues), ShouldEqual, 1)
			So(t.TimeValues[0].Year(), ShouldEqual, 2016)
			So(t.TimeValues[0].Month(), ShouldEqual, time.October)
			So(t.TimeValues[0].Day(), ShouldEqual, 22)
		})

		Convey("It returns error if more range has more than 2 value with seperator", func() {
			t := NewTimerange()
			So(t.Set("2016/10/22..2016/10/23..2016/10/24"), ShouldEqual, ErrInvalidRange)
		})

		Convey("It returns error if any of timestamp in range is invalid", func() {
			t := NewTimerange()
			err := t.Set("2016/10/21..2016/10")

			So(err, ShouldNotBeNil)
			_, ok := err.(*time.ParseError)
			So(ok, ShouldBeTrue)

			err = t.Set("2016/10/111..2016/10/21")
			So(err, ShouldNotBeNil)
			_, ok = err.(*time.ParseError)
			So(ok, ShouldBeTrue)
		})

		Convey("It returns error start date is after end date", func() {
			t := NewTimerange()
			err := t.Set("2016/10/24..2016/10/21")

			So(err, ShouldNotBeNil)
		})

		Convey("It parses and stores all timestamps between given range of timestamps", func() {
			t := NewTimerange()
			So(t.Set("2016/10/22..2016/10/24"), ShouldBeNil)

			So(len(t.TimeValues), ShouldEqual, 3)

			So(t.TimeValues[0].Day(), ShouldEqual, 22)
			So(t.TimeValues[1].Day(), ShouldEqual, 23)
			So(t.TimeValues[2].Day(), ShouldEqual, 24)
		})
	})
}
