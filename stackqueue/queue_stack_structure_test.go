package stackqueue

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewQueue(t *testing.T) {
	que := newQueue(5)
	que.push(5)
	que.push(4)
	fmt.Println(que.length())
	que.push(3)
	fmt.Println(que.length())
	que.pop()
	fmt.Println(que.front())
	fmt.Println(que.length())
}

func TestNewStack(t *testing.T) {
	stk := newStack(5)

	Convey("test stack functions", t, func() {
		err := stk.push(1)
		err = stk.push(2)
		err = stk.push(3)
		err = stk.push(4)
		err = stk.push(5)
		So(err, ShouldBeNil)
		fmt.Println(stk.length())
		err = stk.push(6)
		So(err, ShouldNotBeNil)
		err = stk.pop()
		err = stk.pop()
		err = stk.pop()
		err = stk.pop()
		val, err := stk.top()
		fmt.Println(val)
		So(err, ShouldBeNil)
		err = stk.pop()
		So(err, ShouldBeNil)
		err = stk.pop()
		So(err, ShouldNotBeNil)
		val, err = stk.top()
		fmt.Println(val)
		So(err, ShouldNotBeNil)
	})
}
