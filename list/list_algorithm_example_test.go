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

func TestDividePartitionList(t *testing.T) {
	node6 := newNode(4, nil)
	node5 := newNode(4, node6)
	node4 := newNode(1, node5)
	node3 := newNode(5, node4)
	node2 := newNode(2, node3)
	node1 := newNode(3, node2)
	head1 := newNode(9, node1)
	printSingleList(head1)
	dividePartitionList(head1, 4)
	printSingleList(head1)

	node06 := newNode(4, nil)
	node05 := newNode(4, node06)
	node04 := newNode(1, node05)
	node03 := newNode(5, node04)
	node02 := newNode(2, node03)
	node01 := newNode(3, node02)
	head01 := newNode(9, node01)
	printSingleList(head01)
	head := divideAdvanceList(head01, 5)
	printSingleList(head)
}

func TestCopyRandomList(t *testing.T) {
	node3 := newRandomNode(7, nil, nil)
	node2 := newRandomNode(5, node3, nil)
	node1 := newRandomNode(3, node2, nil)
	head := newRandomNode(1, node1, nil)
	head.random = node2
	node2.random = node1
	printRandomList(head)
	copyHead := copyRandomList(head)
	printRandomList(copyHead)
	newHead := copyRandomList02(head)
	printRandomList(newHead)
}

func TestListIntersect(t *testing.T) {
	Convey("test listIntersect function", t, func() {
		//1->2->3->4->5->6->7
		node6 := newNode(7, nil)
		node5 := newNode(6, node6)
		node4 := newNode(5, node5)
		node3 := newNode(4, node4)
		node2 := newNode(3, node3)
		node1 := newNode(2, node2)
		head1 := newNode(1, node1)

		//0->9->8->6->7
		node02 := newNode(8, node5)
		node01 := newNode(9, node02)
		head2 := newNode(0, node01)
		fmt.Println(listIntersect(head1, head2))

		//1->2->3->4->5->6->7->4...
		node006 := newNode(7, nil)
		node005 := newNode(6, node006)
		node004 := newNode(5, node005)
		node003 := newNode(4, node004)
		node002 := newNode(3, node003)
		node001 := newNode(2, node002)
		head01 := newNode(1, node001)
		node006.NextNode = node003

		// 0->9->8->2...
		node0003 := newNode(12, nil)
		node0002 := newNode(8, node0003)
		node0001 := newNode(9, node0002)
		head02 := newNode(0, node0001)
		node0003.NextNode = node001
		res := listIntersect(head01, head02)
		fmt.Println(res)

		// 0->9->8->6...
		node14 := newNode(8, nil)
		node15 := newNode(9, node14)
		head03:= newNode(0, node15)
		node14.NextNode = node005
		fmt.Println(listIntersect(head01, head03))
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

func newRandomNode(val int, next, random *randomNode) *randomNode {
	return &randomNode{
		val: val,
		next: next,
		random: random,
	}
}

func printRandomList(head *randomNode) {
	for head != nil {
		if head.next != nil {
			fmt.Print(head.val)
			if head.random != nil {
				fmt.Print("*", head.random.val, "->")
			} else {
				fmt.Print("->")
			}

		} else {
			fmt.Print(head.val)
		}
		head = head.next
	}
	fmt.Println()
}