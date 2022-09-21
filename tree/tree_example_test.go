package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelectSuccessorNode(t *testing.T) {
	Convey("test the selectSuccessorNode function ", t, func() {
		node07 := newSpecialTreeNode(9, nil, nil, nil)
		node06 := newSpecialTreeNode(7, nil, nil, nil)
		//node05 := newTreeNode(1, nil, nil)
		//node04 := newTreeNode(4, nil, nil)
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
		So(selectSuccessorNode(node01), ShouldEqual, selectSuccessorNodeTraversal(head, node01))

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
		So(selectSuccessorNode(node001), ShouldEqual, selectSuccessorNodeTraversal(head1, node001))
	})

}

func TestPaperFolding(t *testing.T) {
	Convey("test the paperFolding function", t, func() {
		Convey("test one case", func() {
			paperFolding(2)
			fmt.Println()
		})
		Convey("test two case", func() {
			paperFolding(1)
			fmt.Println()
			paperFolding(3)
		})

	})

}

func TestIsSearchTree(t *testing.T)  {
	node07 := newTreeNode(9, nil, nil)
	node06 := newTreeNode(7, nil, nil)
	node05 := newTreeNode(1, nil, nil)
	//node04 := newTreeNode(4, nil, nil)
	node03 := newTreeNode(2, node07, nil)
	node02 := newTreeNode(8, node05, node06)
	node01 := newTreeNode(3, node03,nil)
	head := newTreeNode(5, node01, node02)
	printTree(head)
	fmt.Println(isSearchTree(head))

	node007 := newTreeNode(1, nil, nil)
	node006 := newTreeNode(9, nil, nil)
	node005 := newTreeNode(7, nil, nil)
	//node04 := newTreeNode(4, nil, nil)
	node003 := newTreeNode(2, node007, nil)
	node002 := newTreeNode(8, node005, node006)
	node001 := newTreeNode(3, node003,nil)
	head1 := newTreeNode(5, node001, node002)
	printTree(head1)
	fmt.Println(isSearchTree(head1))
}

func TestIsCompleteBinaryTree(t *testing.T) {
	Convey("test the isCompleteBinaryTree function...", t, func() {
		node07 := newTreeNode(9, nil, nil)
		node06 := newTreeNode(7, nil, nil)
		//node05 := newTreeNode(1, nil, nil)
		//node04 := newTreeNode(4, nil, nil)
		node03 := newTreeNode(2, nil, nil)
		node02 := newTreeNode(8, nil, node06)
		node01 := newTreeNode(3, node03,node07)
		head := newTreeNode(5, node01, node02)
		printTree(head)
		fmt.Println(isCompleteBinaryTree(head))
		So(isCompleteBinaryTree(head), ShouldBeFalse)

		//node007 := newTreeNode(1, nil, nil)
		node006 := newTreeNode(9, nil, nil)
		node005 := newTreeNode(7, nil, nil)
		//node04 := newTreeNode(4, nil, nil)
		//node003 := newTreeNode(2, node007, nil)
		node002 := newTreeNode(8, node005, node006)
		node001 := newTreeNode(3, nil,nil)
		head1 := newTreeNode(5, node001, node002)
		printTree(head1)
		So(isCompleteBinaryTree(head1), ShouldBeFalse)
	})
}

func TestIsBalanceBinaryTree(t *testing.T) {
	node07 := newTreeNode(9, nil, nil)
	node06 := newTreeNode(7, nil, nil)
	node05 := newTreeNode(1, nil, nil)
	//node04 := newTreeNode(4, nil, nil)
	node03 := newTreeNode(2, nil, nil)
	node02 := newTreeNode(8, node05, node06)
	node01 := newTreeNode(3, node03,node07)
	head := newTreeNode(5, node01, node02)
	printTree(head)
	fmt.Println(isBalanceBinaryTree(head))
	fmt.Println(computerHeight(head))

	node007 := newTreeNode(1, nil, nil)
	node006 := newTreeNode(9, nil, nil)
	node005 := newTreeNode(7, nil, nil)
	//node04 := newTreeNode(4, nil, nil)
	node003 := newTreeNode(2, node007, nil)
	node002 := newTreeNode(8, node005, node006)
	node001 := newTreeNode(3, node003,nil)
	head1 := newTreeNode(5, node001, node002)
	printTree(head1)
	fmt.Println(isBalanceBinaryTree(head1))
	fmt.Println(computerHeight(head1))
}

func TestComputeTreeNodeNums(t *testing.T) {
	node07 := newTreeNode(9, nil, nil)
	node06 := newTreeNode(7, nil, nil)
	node05 := newTreeNode(1, nil, nil)
	//node04 := newTreeNode(4, nil, nil)
	node03 := newTreeNode(2, nil, nil)
	node02 := newTreeNode(8, node05, node06)
	node01 := newTreeNode(3, node03,node07)
	head := newTreeNode(5, node01, node02)
	printTree(head)
	Convey("test computer node nums: ", t, func() {
		fmt.Println(computeTreeNodeNums(head))
		//So(computerNodeLevelOrder(head), ShouldEqual, computeTreeNodeNums(head))
	})

	node007 := newTreeNode(9, nil, nil)
	node006 := newTreeNode(7, nil, nil)
	node005 := newTreeNode(1, nil, nil)
	node004 := newTreeNode(5, nil, nil)
	node008 := newTreeNode(4, nil, nil)
	node003 := newTreeNode(2, node004, node008)
	node002 := newTreeNode(8, node005, node006)
	node001 := newTreeNode(3, node003,node007)
	head0 := newTreeNode(5, node001, node002)
	printTree(head0)
	Convey("test computer node nums: ", t, func() {
		fmt.Println(computeTreeNodeNums(head0))
		//So(computerNodeLevelOrder(head0), ShouldEqual, computerNodeNums(head0))
	})
}