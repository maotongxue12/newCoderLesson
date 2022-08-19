package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPreorderTraversalR(t *testing.T) {
	Convey("test the preorderTraversalR function", t, func() {
		node07 := newTreeNode(9, nil, nil)
		node06 := newTreeNode(7, nil, nil)
		node05 := newTreeNode(1, nil, nil)
		node04 := newTreeNode(4, nil, nil)
		node03 := newTreeNode(2, node07, nil)
		node02 := newTreeNode(8, node05, node06)
		node01 := newTreeNode(3, node03, node04)
		head := newTreeNode(5, node01, node02)
		printTree(head)
		resRec := preorderTraversalR(head)
		res := preorderTraversal(head)
		fmt.Println("先序遍历: ", resRec, res)
		So(resRec, ShouldResemble, res)
		res01Rec := inorderTraversalR(head)
		res01 := inorderTraversal(head)
		fmt.Println("中序遍历: ", res01Rec, res01)
		So(res01Rec, ShouldResemble, res01)
		res02Rec := postorderTraversalR(head)
		res02 := postorderTraversal(head)
		res03 := postorderTraversal02(head)
		fmt.Println("后序遍历: ", res02Rec, res02, res03)
		So(res02Rec, ShouldResemble, res02)
		So(res02, ShouldResemble, res03)

		resLevel := levelTraversal(head)
		fmt.Println("层次遍历: ", resLevel)

	})
}

func TestPreTraversalSerialization(t *testing.T) {
	Convey("test the tree traversalSerialization", t, func() {
		node07 := newTreeNode(9, nil, nil)
		node06 := newTreeNode(7, nil, nil)
		node05 := newTreeNode(1, nil, nil)
		node04 := newTreeNode(4, nil, nil)
		node03 := newTreeNode(2, node07, nil)
		node02 := newTreeNode(8, node05, node06)
		node01 := newTreeNode(3, node03, node04)
		head := newTreeNode(5, node01, node02)
		printTree(head)
		resPre := preTraversalSerialization(head)
		fmt.Println(resPre)
		resPreUnRec := preTraversalSerializationNonRec(head)
		So(resPre, ShouldResemble, resPreUnRec)
		resIn := inTraversalSerialization(head)
		resInUnRec := inTraversalSerializationNonRec(head)
		So(resIn, ShouldResemble, resInUnRec)
		fmt.Println(resIn)
		resPost := postTraversalSerialization(head)
		fmt.Println(resPost)
	})
}
