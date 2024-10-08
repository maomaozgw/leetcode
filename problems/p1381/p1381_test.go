package p1381

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test1381(t *testing.T) {
	Convey("Test 1381", t, func() {
		stack := Constructor(3)
		stack.Push(1)
		stack.Push(2)
		So(stack.Pop(), ShouldEqual, 2)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		stack.Increment(5, 100)
		stack.Increment(2, 100)
		So(stack.Pop(), ShouldEqual, 103)
		So(stack.Pop(), ShouldEqual, 202)
		So(stack.Pop(), ShouldEqual, 201)
		So(stack.Pop(), ShouldEqual, -1)
	})
}
