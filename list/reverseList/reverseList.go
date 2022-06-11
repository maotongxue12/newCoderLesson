package main

import "fmt"

type node struct {
	Val  int
	Next *node
}

type DoubleNode struct {
	Val int
	PreNode *DoubleNode
	NextNode *DoubleNode
}

func reverseList(head *node) *node {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *node = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseDoubleList(head *DoubleNode) *DoubleNode {
	if head == nil || head.NextNode == nil {
		return head
	}
	var pre *DoubleNode
	cur := head
	for cur != nil {
		next := cur.NextNode
		cur.NextNode = pre
		cur.PreNode = next
		pre = cur
		cur = next
	}
	return pre
}

func printList(head *node) {
	for head != nil {
		if head.Next != nil {
			fmt.Print(head.Val, "->")
		} else {
			fmt.Print(head.Val)
		}
		head = head.Next
	}
	fmt.Println()
}

func printDoubleList(head *DoubleNode) {
	for head != nil {
		if head.NextNode != nil {
			fmt.Print(head, "--->")
		} else {
			fmt.Print(head)
		}
		head = head.NextNode
	}
	fmt.Println()
}

func newDoubleNode(key int, preNode *DoubleNode, nextNode *DoubleNode) *DoubleNode {
	return &DoubleNode{
		Val: key,
		PreNode: preNode,
		NextNode: nextNode,
	}
}

func main() {
	node1 := new(node)
	node2 := new(node)
	node3 := new(node)
	node4 := new(node)
	node1.Val = 1
	node1.Next = node2
	node2.Val = 2
	node2.Next = node3
	node3.Val = 4
	node3.Next = node4
	node4.Val = 5
	node4.Next = nil
	printList(node1)
	head := reverseList(node1)
	printList(head)

	doubleNode01 := newDoubleNode(1, nil, nil)
	doubleNode02 := newDoubleNode(2, doubleNode01, nil)
	doubleNode03 := newDoubleNode(3, doubleNode02, nil)
	doubleNode04 := newDoubleNode(4, doubleNode03, nil)
	doubleNode01.NextNode = doubleNode02
	doubleNode02.NextNode = doubleNode03
	doubleNode03.NextNode = doubleNode04
	printDoubleList(doubleNode01)
	doubleHead := reverseDoubleList(doubleNode01)
	printDoubleList(doubleHead)

}