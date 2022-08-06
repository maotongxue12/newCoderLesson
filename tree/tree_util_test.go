package tree

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSpace(t *testing.T) {
	convey.Convey("test the getSpace function", t, func() {
		leftLen := 5
		rightLen := 5
		val := "10"
		str := getSpace(leftLen) + val + getSpace(rightLen)
		fmt.Println(str)
	})
}

func TestPrintTree(t *testing.T) {
	convey.Convey("test the printTree function", t, func() {
		node07 := newTreeNode(9, nil, nil)
		node06 := newTreeNode(7, nil, nil)
		node05 := newTreeNode(1, nil, nil)
		node04 := newTreeNode(4, nil, nil)
		node03 := newTreeNode(2, node07, nil)
		node02 := newTreeNode(8, node05, node06)
		node01 := newTreeNode(3, node03, node04)
		head := newTreeNode(5, node01, node02)

		printTree(head)
	})
}