package tree

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
