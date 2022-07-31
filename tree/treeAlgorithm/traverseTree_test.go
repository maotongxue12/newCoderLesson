package treeAlgorithm

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
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
	node01 := newNode(3, node03, nil)
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

//先序序列化
func TestPreOrderTraverSerialization(t *testing.T) {
	node10 := newNode(9, nil, nil)
	node09 := newNode(4, nil, nil)
	node08 := newNode(5, nil, node09)

	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, node10, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03, node08)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	res := preOrderTraverSerialization(head)
	fmt.Println(res)
	res2 := preOrderTraverSerializationUnRecursive(head)
	fmt.Println(res2)
	Convey("test tree serial: ", t, func() {
		So(res, ShouldEqual, res2)
	})
}

//中序遍历序列化
func TestInOrderSerialization(t *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03, nil)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	res := inOrderSerialization(head)
	fmt.Println(res)
	res2 := inOrderSerializationUnRecursive(head)
	fmt.Println(res2)
	Convey("test tree serial: ", t, func() {
		So(res, ShouldEqual, res2)
	})
}

//后续遍历序列化
func TestPostOrderSerialization(t *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03, nil)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	res := postOrderSerialization(head)
	fmt.Println(res)
	res2 := postOrderSerializationUnRecursive(head)
	fmt.Println(res2)
	Convey("test tree serial: ", t, func() {
		So(res, ShouldEqual, res2)
	})
}

//后续序列反序列化
func TestPostOrderDeserialization(T *testing.T) {
	strs := "#_#_9_#_2_#_3_#_#_1_#_#_7_8_5"
	head := postOrderDeserialization(strs)
	treeUtil.PrintAvlTree(head)
}

func TestInorderDeserialization(T *testing.T) {
	str := "5_3_2_9_#_#_#_#_8_1_#_#_7_#_#"
	head := preorderDeserialization(str)
	treeUtil.PrintAvlTree(head)
}

//层次遍历序列化
func TestLevelSerialization(t *testing.T) {
	node07 := newNode(9, nil, nil)
	node06 := newNode(7, nil, nil)
	node05 := newNode(1, nil, nil)
	//node04 := newNode(4, nil, nil)
	node03 := newNode(2, node07, nil)
	node02 := newNode(8, node05, node06)
	node01 := newNode(3, node03, nil)
	head := newNode(5, node01, node02)
	treeUtil.PrintAvlTree(head)
	res := levelSerialization(head)
	fmt.Println(res)
	Convey("test tree serial: ", t, func() {
		So(res, ShouldEqual, "5 3 8 2 # 1 7 9 # # # # # # # ")
	})
}

//层次遍历反序列化
func TestLevelDeserialization(t *testing.T) {
	str := "5_3_8_2_#_1_7_9_#_#_#_#_#_#_#"
	head := levelDeserialization(str)
	treeUtil.PrintAvlTree(head)
}
