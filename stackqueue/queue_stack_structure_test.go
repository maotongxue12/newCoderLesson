package stackqueue

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewQueue(t *testing.T) {
	Convey("test the arrToQueue function", t, func() {
		que := newQueue(5)
		que.push(5)
		que.push(4)
		fmt.Println(que.length())
		que.push(3)
		fmt.Println(que.length())
		que.pop()
		fmt.Println(que.front())
		fmt.Println(que.length())
		que.push(2)
		que.push(1)
		que.push(0)
		err := que.push(8)
		So(err, ShouldNotBeNil)
		fmt.Println(que.front())
	})
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

func TestNewSpecialStack(t *testing.T) {
	sStack := newSpecialStack(5)
	Convey("test the specialStack function", t, func() {
		sStack.push(5)
		sStack.push(4)
		sStack.push(8)
		sStack.push(2)
		sStack.push(8)
		minValue, _ := sStack.getMin()
		So(minValue, ShouldEqual, 2)
		sStack.pop()
		minValue, _ = sStack.getMin()
		So(minValue, ShouldEqual, 2)
		sStack.pop()
		minValue, _ = sStack.getMin()
		So(minValue, ShouldEqual, 4)
		sStack.pop()
		sStack.pop()
		minValue, _ = sStack.getMin()
		So(minValue, ShouldEqual, 5)
	})
}

func TestNewQueueToStack(t *testing.T) {
	stack := newQueueToStack(5)
	Convey("tesh the queueToStack function", t, func() {
		stack.push(5)
		stack.push(4)
		stack.push(3)
		val, _ := stack.front()
		So(3, ShouldEqual, val)
		stack.pop()
		val, _ = stack.front()
		So(4, ShouldEqual, val)
		stack.pop()
		val, _ = stack.front()
		So(5, ShouldEqual, val)
		stack.pop()
		_, err := stack.front()
		So(err, ShouldNotBeNil)
		stack.push(7)
		val, _ = stack.front()
		So(7, ShouldEqual, val)
	})
}

func TestNewStackToQueue(t *testing.T) {
	stq := newStackToQueue(5)
	Convey("test the stackToQueue function", t, func() {
		stq.push(5)
		stq.push(4)
		stq.push(3)
		stq.push(2)
		val, _ := stq.top()
		So(val, ShouldEqual, 5)
		stq.pop()
		val, _ = stq.top()
		So(val, ShouldEqual, 4)
		err := stq.pop()
		So(err, ShouldBeNil)
		err = stq.pop()
		So(err, ShouldBeNil)
		val, err = stq.top()
		So(val, ShouldEqual,2)
		err = stq.pop()
		So(err, ShouldBeNil)
		val, err = stq.top()
		So(err, ShouldNotBeNil)
	})
}