package recursive

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResolve1st(t *testing.T) {
	Convey("", t, func() {
		val, _ := Resolve1st(1, 2, 1, 1)
		So(val, ShouldEqual, 1)

		val, _ = Resolve1st(1, 2, 1, 5)
		So(val, ShouldEqual, 31)
	})
}
