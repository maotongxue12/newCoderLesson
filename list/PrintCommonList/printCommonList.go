package main

import "fmt"

type node struct {
	Val int
	Next *node
}

func printCommonElement(head1 *node, head2 *node){
	if head1 == nil || head2 == nil {
		return
	}
	for head1 != nil  && head2 != nil {
		if head1.Val < head2.Val {
			head1 = head1.Next
		} else if head1.Val > head2.Val {
			head2 = head2.Next
		} else {
			fmt.Print(head1.Val, " ")
			head1 = head1.Next
			head2 = head2.Next
		}
	}
}

func newNode(val int, next *node) *node {
	return &node{
		Val: val,
		Next: next,
	}
}

func main() {
	node3 := newNode(6, nil)
	node2 := newNode(4, node3)
	node1 := newNode(3, node2)
	head1 := newNode(1, node1)

	node03 := newNode(4, nil)
	node02 := newNode(3, node03)
	node01 := newNode(2, node02)
	head2 := newNode(0, node01)

	printCommonElement(head1, head2)
}