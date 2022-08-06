package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

//二叉树可以用常规的三种遍历结果来描述其结构，但是不够直观，尤其是二叉树中有重复 值的时候，
//仅通过三种遍历的结果来构造二叉树的真实结构更是难上加难，有时则根本不可能。 给定一棵二叉树的头节点 head，
//已知二叉树节点值的类型为 32 位整型，请实现一个打印二叉 树的函数，可以直观地展示树的形状，也便于画出真实的结构。
//思路分析：第一步要解决的问题是用什么样的方式可以无歧义的打印二叉树，也就是设计打印的样式。
//第二个要解决的问题：规定节点打印时占用的统一长度。整形值(INT32_MIN)占用的长度是11；
//加上前缀和后缀(H、v、^)之后占用长度为13，前后两个空格，固定长度设为17
func newTreeNode(value int, leftNode, rightNode *treeNode) *treeNode {
	return &treeNode{
		val:   value,
		left:  leftNode,
		right: rightNode,
	}
}

func printTree(head *treeNode) {
	if head == nil {
		return
	}
	printProcess(head, 0, "H", 17)
}

//打印的整体过程是：右子树–>根节点–>左子树
func printProcess(head *treeNode, height int, to string, length int) {
	if head == nil {
		return
	}
	printProcess(head.right, height+1, "v", length)
	var builder strings.Builder
	builder.WriteString(to)
	builder.WriteString(strconv.Itoa(head.val))
	builder.WriteString(to)
	val := builder.String()
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	var b strings.Builder
	b.WriteString(getSpace(lenL))
	b.WriteString(val)
	b.WriteString(getSpace(lenR))
	val = b.String()
	fmt.Println(getSpace(length*height) + val)
	printProcess(head.left, height+1, "^", length)
}

func getSpace(int2 int) string {
	str := make([]byte, 0)
	for i := 0; i < int2; i++ {
		str = append(str, ' ')
	}
	return string(str)
}
