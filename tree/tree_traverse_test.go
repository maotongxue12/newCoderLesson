package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPreorderTraversalR(t *testing.T) {
	Convey("test the preorderTraversalR function", t, func() {
		node07 := newTreeNode(9, nil, nil)
		node06 := newTreeNode(7, nil, nil)
		node05 := newTreeNode(1, nil, nil)
		//node04 := newTreeNode(4, nil, nil)
		node03 := newTreeNode(2, node07, nil)
		node02 := newTreeNode(8, node05, node06)
		node01 := newTreeNode(3, node03, nil)
		head := newTreeNode(5, node01, node02)
		res := preorderTraversalR(head)
		fmt.Print(res)
	})
}
