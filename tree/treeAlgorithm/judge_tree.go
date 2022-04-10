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

//判断一颗二叉树是否是完全二叉树
//思路：结合完全二叉树特点，可采用层序遍历方法实现该算法，当进行层序遍历时：
//（1）如果当前节点有右子树没有左子树则必然不是完全二叉树；（2）如果当前节点仅有一个子树，则在他之后遍历的节点全部为叶节点，否则此二叉树必然不是完全二叉树
func judgeIsCompleteBinaryTree(head *treeUtil.Node) bool {
	if head == nil {
		return true
	}
	queue := make([]*treeUtil.Node, 0)
	queue = append(queue, head)
	flag := false
	for len(queue) != 0 {
		head = queue[0]
		queue = queue[1:]
		if head.Left == nil && head.Right != nil {
			return false
		}
		//如果当前节点有左子树而无右子树，则需要判断后续遍历的节点是否为叶子节点
		//需要设置一个标志，来表示当前遍历的节点只有左子树没有右子树
		if head.Left == nil || head.Right == nil {
			flag = true
		}
		if flag == true && (head.Left != nil || head.Right != nil) {
			return false
		}
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
	return true
}


//问题：在比普通二叉树节点结构多了一个指向父节点的parent指针的二叉树中找到一个节点的后继节点, 后继节点即使中序遍历的下一个节点
//方法一：可采用中序遍历法判断该节点的下一个节点
func getDescendantNode(head *treeUtil.Node, key *treeUtil.Node) *treeUtil.Node {
	if head == nil {
		return nil
	}
	//用以标记当前节点是否是需要寻找后继节点的节点
	flag := false
	stack := make([]*treeUtil.Node, 0)
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		} else {
			head = stack[len(stack)-1]
			if true == flag {
				return head
			}
			if head == key {
				flag = true
			}
			stack = stack[:len(stack)-1]
			head = head.Right
		}
	}
	return nil
}

//方法二：利用该二叉树带有父节点的特点，不需要完整遍历整棵树
//当该节点有右子树时，则该节点的后继节点为右子树根节点的左子树的最左叶节点
//当该节点没有右子树时，则从该节点向上索引父节点，找到该节点为其左子树的第一个先祖节点即为该节点的后继节点
func getDescendantNode02(key *treeUtil.Node) *treeUtil.Node {
	if key == nil {
		return nil
	}
	if key.Right != nil {
		return getMostLeftLeafNode(key)
	} else {
		cur := key
		parent := cur.Parent
		for parent != nil {
			if parent.Left == cur {
				return parent
			}
			cur = parent
			parent = cur.Parent
		}
	}
	return nil
}

//获取当前节点右子树的最左叶子节点
func getMostLeftLeafNode(node *treeUtil.Node) *treeUtil.Node {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}