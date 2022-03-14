package treeUtil

import "testing"

func newNode(val int, lf , rt *Node) *Node {
	return &Node{
		Value: val,
		Left:  lf,
		Right: rt,
	}
}

func TestPrintAvlTree(t *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03, node04)
	head := newNode(5, node01, node02)

	PrintAvlTree(head)
}
