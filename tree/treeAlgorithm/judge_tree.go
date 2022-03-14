package treeAlgorithm

import (
	"math"
	"newCode/tree/treeUtil"
)

//判断是否是二叉搜索树结构
//由二叉搜索树的特性可知二叉搜索树的中序遍历非递减
func judgeBSTTree(head *treeUtil.Node) bool {
	if head == nil || (head.Left == nil && head.Right == nil) {
		return true
	}
	var preValue *treeUtil.Node
	stack := make([]*treeUtil.Node, 0)
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		} else {
			head = stack[len(stack)-1]
			if preValue != nil && (head.Value < preValue.Value) {
				return false
			}
			preValue = head
			stack = stack[0:len(stack)-1]
			head = head.Right
		}
	}
	return true
}

//判断一棵树是否为平衡二叉树
//完全二叉树左右子树的高度差最大为1
type balanceStruct struct {
	isBalance bool
	height int
}

func judgeIsBalanceTree(head *treeUtil.Node) bool {
	return process(head).isBalance
}

func newBalanceStruct(isBalance bool, height int) *balanceStruct {
	return &balanceStruct{
		isBalance: isBalance,
		height: height,
	}
}

func process(head *treeUtil.Node) *balanceStruct {
	if head == nil {
		return newBalanceStruct(true, 0)
	}
	leftData := process(head.Left)
	righData := process(head.Right)
	if false == leftData.isBalance {
		return newBalanceStruct(false, -1)
	}
	if false == righData.isBalance {
		return newBalanceStruct(false, -1)
	}
	height := math.Abs(float64(leftData.height - righData.height))
	if height > 1 {
		return newBalanceStruct(false, -1)
	}
	return newBalanceStruct(true, int(math.Max(float64(leftData.height), float64(righData.height)) + 1))
}

