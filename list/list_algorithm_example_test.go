package list

import (
	_ "container/list"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverseSingleList(t *testing.T) {
	node1 := new(singleNode)
	node2 := new(singleNode)
	node3 := new(singleNode)
	node4 := new(singleNode)
	node1.Val = 1
	node1.NextNode = node2
	node2.Val = 2
	node2.NextNode = node3
	node3.Val = 4
	node3.NextNode = node4
	node4.Val = 5
	node4.NextNode = nil
	printSingleList(node1)
	head := reverseSingleList(node1)
	printSingleList(head)

	doubleNode01 := newDoubleNode(1, nil, nil)
	doubleNode02 := newDoubleNode(2, doubleNode01, nil)
	doubleNode03 := newDoubleNode(3, doubleNode02, nil)
	doubleNode04 := newDoubleNode(4, doubleNode03, nil)
	doubleNode01.NextNode = doubleNode02
	doubleNode02.NextNode = doubleNode03
	doubleNode03.NextNode = doubleNode04
	printDoubleList(doubleNode01)
	head01 := reverseDoubleList(doubleNode01)
	printDoubleList(head01)
}

func TestPrintCommonList(t *testing.T) {
	node3 := newNode(6, nil)
	node2 := newNode(4, node3)
	node1 := newNode(3, node2)
	head1 := newNode(1, node1)

	node03 := newNode(6, nil)
	node02 := newNode(3, node03)
	node01 := newNode(2, node02)
	head2 := newNode(0, node01)
	printCommonList(head1, head2)
}

func TestIsPalindromeList(t *testing.T) {
	Convey("test the isPalindromeList function", t, func() {
		node3 := newNode(6, nil)
		node2 := newNode(4, node3)
		node1 := newNode(4, node2)
		head1 := newNode(6, node1)
		printSingleList(head1)
		res1 := isPalindromeList(head1)
		So(res1, ShouldBeTrue)
		So(isPalindromeList01(head1), ShouldBeTrue)
		printSingleList(head1)

		node03 := newNode(1, nil)
		node02 := newNode(3, node03)
		node01 := newNode(2, node02)
		head2 := newNode(1, node01)
		printSingleList(head2)
		res2 := isPalindromeList(head2)
		So(res2, ShouldBeFalse)
		So(isPalindromeList01(head2), ShouldBeFalse)
		printSingleList(head2)

		node002 := newNode(3, nil)
		node001 := newNode(2, node002)
		head3 := newNode(3, node001)
		printSingleList(head3)
		res3 := isPalindromeList(head3)
		So(res3, ShouldBeTrue)
		So(isPalindromeList01(head3), ShouldBeTrue)
		printSingleList(head3)

		head4 := newNode(1, nil)
		printSingleList(head4)
		So(isPalindromeList(head4), ShouldBeTrue)
		So(isPalindromeList01(head4), ShouldBeTrue)
		printSingleList(head4)
	})
}

func printSingleList(head *singleNode) {
	for i := head; i != nil; i = i.NextNode {
		if i.NextNode == nil {
			fmt.Print(i.Val, "\n")
		} else {
			fmt.Print(i.Val, "-->")
		}
	}
}

func newNode(val int, next *singleNode) *singleNode {
	return &singleNode{
		Val:      val,
		NextNode: next,
	}
}

func newDoubleNode(key int, preNode *doubleNode, nextNode *doubleNode) *doubleNode {
	return &doubleNode{
		Val:      key,
		PreNode:  preNode,
		NextNode: nextNode,
	}
}

func printDoubleList(head *doubleNode) {
	for i := head; i != nil; i = i.NextNode {
		if i.NextNode == nil {
			fmt.Print(i.Val, "\n")
		} else {
			fmt.Print(i.Val, "<-->")
		}
	}
}
