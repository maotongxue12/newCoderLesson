package treeAlgorithm

import (
	"fmt"
	"newCode/tree/treeUtil"
	"testing"
)

func TestTraverseTree(t *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03,nil)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	fmt.Print("pre-order-recursion:   ")
	preOrderTraverseTree(head)
	fmt.Println()
	fmt.Print("pre-order-unRecursive: ")
	preOrderUnRecursive(head)
	fmt.Println()
	fmt.Print("in-order-recursion:   ")
	inOrderTraverseTree(head)
	fmt.Println()
	fmt.Print("in-order-unRecursion: ")
	inOrderUnRecursive(head)
	fmt.Println()
	fmt.Print("post-order-recursion:   ")
	postOrderTraverseTree(head)
	fmt.Println()
	fmt.Print("post-order-unRecursion: ")
	postOrderUnRecursive(head)
	fmt.Println()
	fmt.Print("post-order-unRecursion: ")
	postOrderUnRecursive02(head)
	fmt.Println()
	fmt.Print("level traversal: ")
	levelTraversal(head)
}