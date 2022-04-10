package treeAlgorithm

import (
	"fmt"
	"newCode/tree/treeUtil"
	"testing"
)

func TestJudgeBSTTree(T *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03,nil)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	fmt.Println(judgeBSTTree(head))

	node007 := newNode(1, nil, nil)
	node006 := newNode(9, nil, nil)
	node005 := newNode(7, nil, nil)
	//node04 := newNode(4, nil, nil)
	node003 := newNode(2, node007, nil)
	node002 := newNode(8, node005, node006)
	node001 := newNode(3, node003,nil)
	head1 := newNode(5, node001, node002)
	treeUtil.PrintAvlTree(head1)
	fmt.Println(judgeBSTTree(head1))
}

func TestJudgeIsBalanceTree(T *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, nil, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03,node07)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	fmt.Println(judgeIsBalanceTree(head))

	node007 := newNode(1, nil, nil)
	node006 := newNode(9, nil, nil)
	node005 := newNode(7, nil, nil)
	//node04 := newNode(4, nil, nil)
	node003 := newNode(2, node007, nil)
	node002 := newNode(8, node005, node006)
	node001 := newNode(3, node003,nil)
	head1 := newNode(5, node001, node002)
	treeUtil.PrintAvlTree(head1)
	fmt.Println(judgeIsBalanceTree(head1))
}

func TestJudgeIsCompleteBinaryTree(T *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	//node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, nil, nil)
	node02 := newNode(8, nil, node06)
	node01 := newNode(3, node03,node07)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	fmt.Println(judgeIsCompleteBinaryTree(head))

	//node007 := newNode(1, nil, nil)
	node006 := newNode(9, nil, nil)
	node005 := newNode(7, nil, nil)
	//node04 := newNode(4, nil, nil)
	//node003 := newNode(2, node007, nil)
	node002 := newNode(8, node005, node006)
	node001 := newNode(3, nil,nil)
	head1 := newNode(5, node001, node002)
	treeUtil.PrintAvlTree(head1)
	fmt.Println(judgeIsCompleteBinaryTree(head1))
}

func TestGetDescendantNode02(T *testing.T) {
	node07 := newNodeParent(9, nil, nil, nil)
	node06 := newNodeParent(7, nil, nil, nil)
	//node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNodeParent(2, nil, nil, nil)
	node02 := newNodeParent(8, nil, node06, nil)
	node01 := newNodeParent(3, node03, node07, nil)
	head := newNodeParent(5, node01, node02, nil)
	node01.Parent = head
	node02.Parent = head
	node03.Parent = node01
	node06.Parent = node02
	node07.Parent = node01
	treeUtil.PrintAvlTree(head)
	fmt.Println(getDescendantNode02(node01))

	//node007 := newNodeParent(1, nil, nil, nil)
	node006 := newNodeParent(9, nil, nil, nil)
	node005 := newNodeParent(7, nil, nil, nil)
	//node04 := newNodeParent(4, nil, nil, nil)
	//node003 := newNodeParent(2, node007, nil, nil)
	node002 := newNodeParent(8, node005, node006, nil)
	node001 := newNodeParent(3, nil,nil, nil)
	head1 := newNodeParent(5, node001, node002, nil)
	node001.Parent = head1
	node002.Parent = head1
	node005.Parent = node002
	node006.Parent = node002
	treeUtil.PrintAvlTree(head1)
	fmt.Println(getDescendantNode02(head1))
	fmt.Println(getDescendantNode02(head1))
}