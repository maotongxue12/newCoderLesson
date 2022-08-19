package tree

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelectSuccessorNode(t *testing.T) {
	convey.Convey("test the selectSuccessorNode function ", t, func() {
		node07 := newSpecialTreeNode(9, nil, nil, nil)
		node06 := newSpecialTreeNode(7, nil, nil, nil)
		//node05 := newNode(1, nil, nil)
		//node04 := newNode(4, nil, nil)
		node03 := newSpecialTreeNode(2, nil, nil, nil)
		node02 := newSpecialTreeNode(8, nil, node06, nil)
		node01 := newSpecialTreeNode(3, node03, node07, nil)
		head := newSpecialTreeNode(5, node01, node02, nil)
		node01.parent = head
		node02.parent = head
		node03.parent = node01
		node06.parent = node02
		node07.parent = node01
		printTree(head)
		fmt.Println(selectSuccessorNode(node01))
		convey.So(selectSuccessorNode(node01), convey.ShouldEqual, selectSuccessorNodeTraversal(head, node01))

		//node007 := newSpecialTreeNode(1, nil, nil, nil)
		node006 := newSpecialTreeNode(9, nil, nil, nil)
		node005 := newSpecialTreeNode(7, nil, nil, nil)
		//node04 := newSpecialTreeNode(4, nil, nil, nil)
		//node003 := newSpecialTreeNode(2, node007, nil, nil)
		node002 := newSpecialTreeNode(8, node005, node006, nil)
		node001 := newSpecialTreeNode(3, nil, nil, nil)
		head1 := newSpecialTreeNode(5, node001, node002, nil)
		node001.parent = head1
		node002.parent = head1
		node005.parent = node002
		node006.parent = node002
		printTree(head1)
		fmt.Println(selectSuccessorNode(node001))
		convey.So(selectSuccessorNode(node001), convey.ShouldEqual, selectSuccessorNodeTraversal(head1, node001))
	})

}
