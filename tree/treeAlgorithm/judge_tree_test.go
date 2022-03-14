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
