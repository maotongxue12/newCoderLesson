package tree

import (
	"fmt"
	"math"
	"newCode/sort"
)

//题目一：在二叉树中找到一个节点的后继节点
//该结构比普通二叉树节点结构多了一个指向父节点的parent指针。假设有一棵Node类型的节点组成的二叉树，
//树中每个节点的parent指针都正确地指向自己的父节点，头节点的parent指向null。
//只给一个在二叉树中的某个节点node，请实现返回node的后继节点的函数。
//在二叉树的中序遍历的序列中，node的下一个节点叫作node的后继节点。
func newSpecialTreeNode(val int, left, right, parent *treeNode) *treeNode {
	return &treeNode{
		val:    val,
		left:   left,
		right:  right,
		parent: parent,
	}
}

//方法一：
//（1）若该节点存在右孩子，则其后继节点为为右孩子的最左子节点
//（2）若该节点无右孩子，则寻找其父节点，祖父节点...直至找到该节点为其左子树的第一个先祖节点，该先祖节点即为其后继节点
//若直至寻到根节点，则该节点的后继节点即为根节点
func selectSuccessorNode(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		temp := getLeftNode(node)
		return temp
	} else {
		for node.parent != nil {
			par := node.parent
			if par.left == node {
				return par
			}
			node = par
		}
	}
	return nil
}

//如果node节点存在右孩子，则找到右孩子的最左叶节点
func getLeftNode(node *treeNode) *treeNode {
	temp := node.right
	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

//方法二：利用中序遍历找到其后继节点
func selectSuccessorNodeTraversal(head *treeNode, node *treeNode) *treeNode {
	if head == nil {
		return nil
	}
	//中序遍历
	stack := make([]*treeNode, 0)
	//flag用于标记待寻找后继节点的节点
	flag := false
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.left
		} else {
			head = stack[len(stack)-1]
			if flag == true {
				return head
			}
			if head == node {
				flag = true
			}
			stack = stack[:len(stack)-1]
			head = head.right
		}
	}
	return nil
}

//折纸问题，【题目】 请把一段纸条竖着放在桌子上，然后从纸条的下边向上方对折1次，压出折痕后展开。
//此时 折痕是凹下去的，即折痕突起的方向指向纸条的背面。如果从纸条的下边向上方连续对折 2 次，压出
//折痕后展开，此时有三条折痕，从上到下依次是下折痕、下折痕和上折痕。
//给定一 个输入参数N，代表纸条都从下边向上方连续对折N次，请从上到下打印所有折痕的方向。
//例如：N=1时，打印： down N=2时，打印： down down up
//思路：观察发现，该结构可以理解为根节点为down且每个子树左节点为down，右节点为up的二叉树的中序遍历
//左子树固定为down，右子树固定为up，不需要构建树结构
func paperFolding(allLevel int) {
	printPaperFoldingProcess(1, allLevel, "down")
}

func printPaperFoldingProcess(curLevel, allLevel int, string2 string) {
	if curLevel > allLevel {
		return
	}
	printPaperFoldingProcess(curLevel+1, allLevel, "down")
	fmt.Println(string2)
	printPaperFoldingProcess(curLevel+1, allLevel, "up")
}

//判断一棵树是否是搜索二叉树
//搜索二叉树特征，中序遍历序列递增
func isSearchTree(head *treeNode) bool {
	if head == nil {
		return true
	}
	var preValue int = sort.INT_MIN
	stack := make([]*treeNode, 0)
	for len(stack) != 0 || head != nil {
		if head != nil {
			stack = append(stack, head)
			head = head.left
		} else {
			temp := stack[len(stack)-1]
			if temp.val < preValue {
				return false
			}
			preValue = temp.val
			stack = stack[:len(stack)-1]
			head = temp.right
		}
	}
	return true
}

//判断一棵树是否市完全二叉树
//解题思路：（1）如果一个节点仅有右节点，没有左节点，则该二叉树肯定不是完全二叉树
//（2）在层次遍历过程中，如果一个节点仅有左子树，而无右子树或者无子树，则其后遍历的节点都为叶节点，否则不是该树则不是完全二叉树
func isCompleteBinaryTree(head *treeNode) bool {
	queue := make([]*treeNode, 0)
	queue = append(queue, head)
	var flag bool = false
	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]
		//该节点仅有右子树，无左子树
		if temp.left == nil && temp.right != nil {
			return false
		}
		if flag == true && (temp.left != nil || temp.right != nil) {
			return false
		}
		//如果当前节点无左子树或右子树，则其后节点必须为叶节点，否则为非完全二叉树
		if temp.left == nil || temp.right == nil {
			flag = true
		}
		if temp.left != nil {
			queue = append(queue, temp.left)
		}
		if temp.right != nil {
			queue = append(queue, temp.right)
		}
	}
	return true
}

//判断一棵树是否是平衡二叉树
//解题思路： 平衡二叉树特点，对任一节点其左右子树高度差不超过1
//利用递归方法，需要记录当前树的高度并判断是否平衡，
type isBalanceStruct struct {
	isBalance bool
	height    int
}

func isBalanceBinaryTree(head *treeNode) bool {
	return isBalanceProcess(head).isBalance
}

func isBalanceProcess(head *treeNode) *isBalanceStruct {
	if head == nil {
		return &isBalanceStruct{
			isBalance: true,
			height:    0,
		}
	}
	//判断左子树是否为平衡二叉树,需要注意左右子树都为平衡二叉树时，以该节点为根的树才为平衡二叉树
	//但左右子树有一个非平衡二叉树时，以该节点为根节点的树即非平衡二叉树
	leftData := isBalanceProcess(head.left)
	rightData := isBalanceProcess(head.right)
	if false == leftData.isBalance {
		return &isBalanceStruct{
			isBalance: false,
			height:    -1,
		}
	}
	if false == rightData.isBalance {
		return &isBalanceStruct{
			isBalance: false,
			height:    -1,
		}
	}

	high := math.Abs(float64(leftData.height - rightData.height))
	if high > 1 {
		return &isBalanceStruct{
			isBalance: false,
			height:    -1,
		}
	}

	return &isBalanceStruct{
		isBalance: true,
		height:    int(math.Max(float64(leftData.height), float64(rightData.height))) + 1,
	}
}

//已知一棵完全二叉树，求其节点的个数 时间复杂度低于O(N)
//思路一：遍历 时间复杂度 o(n)
//思路二：递归，利用完全二叉树的性质采用递归方法获取节点个数，首先计算二叉树的深度allLevel，然后判断右子树的最左叶子节点的深度level，
//如果level=allLevel, 则表明该完全二叉树的左子树为满二叉树，深度为allLevel，否则该二叉树的右子树为满二叉树，且该满二叉树的深度为allLevel-1
func computeTreeNodeNums(head *treeNode) int {
	if head == nil {
		return 0
	}
	return processComputerNums(head, 1, mostLeftLeaf(head, 1))
}

func processComputerNums(head *treeNode, curLevel, allLevel int) int {
	//当前节点层数等于树的总层数，则该节点为叶子节点，即以该节点为根的数仅有一个节点
	if curLevel == allLevel {
		return 1
	}
	//证明左子树为满二叉树, 需要对右子树进行递归
	if mostLeftLeaf(head.right, curLevel+1) == allLevel {
		return 1<<(allLevel-curLevel) + processComputerNums(head.right, curLevel+1, allLevel)
	} else {
		return processComputerNums(head.left, curLevel+1, allLevel) + 1<<(allLevel-curLevel-1)
	}
}

//计算二叉树的深度，由完全二叉树性质可知，最左叶子节点的深度即为二叉树的高,也可通过获取最左叶子节点计算完全二叉树高度
func computerHeight(head *treeNode) int {
	if head == nil {
		return 0
	}
	leftHeight := computerHeight(head.left)
	rightHeight := computerHeight(head.right)
	return int(math.Max(float64(leftHeight), float64(rightHeight)) + 1)
}

//level表示当前节点所在树的深度，当level为1时，返回值为树的总高度
func mostLeftLeaf(head *treeNode, level int) int {
	for head != nil {
		head = head.left
		level++
	}
	return level - 1
}