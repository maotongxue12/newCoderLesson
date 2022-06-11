package main

import "fmt"

type node struct {
	Val int
	Next *node
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

func newNode(val int, next *node) *node {
	return &node{
		Val:  val,
		Next: next,
	}
}

func divedePartition(head *node, pivot int) *node {
	if head == nil || head.Next == nil {
		return head
	}
	var minHead, equalHead, maxHead *node
	var curMin, curEqual, curMax *node
	for head != nil {
		if head.Val < pivot {
			if minHead == nil {
				minHead = head
				curMin = minHead
			} else {
				curMin.Next = head
				curMin= curMin.Next
			}
		} else if head.Val > pivot {
			if maxHead == nil {
				maxHead = head
				curMax = maxHead
			} else {
				curMax.Next = head
				curMax = curMax.Next
			}
		} else {
			if equalHead == nil {
				equalHead = head
				curEqual = head
			} else {
				curEqual.Next = head
				curEqual = curEqual.Next
			}
		}
		head = head.Next
	}

	if curMin != nil {
		curMin.Next = equalHead
		if curEqual != nil {
			curEqual.Next = maxHead
			if curMax != nil {
				curMax.Next = nil
			}
		} else {
			curMin.Next = maxHead
			if curMax != nil {
				curMax.Next = nil
			}
		}
		return minHead
	} else {
		if curEqual != nil {
			curEqual.Next = maxHead
			if curMax != nil {
				curMax.Next = nil
			}
			return equalHead
		} else {
			return maxHead
		}
	}
}

func main() {
	node4 := newNode(1, nil)
	node3 := newNode(5, node4)
	node2 := newNode(2, node3)
	node1 := newNode(0, node2)
	head1 := newNode(9, node1)
	printList(head1)

	head := divedePartition(head1, 3)
	printList(head)
}