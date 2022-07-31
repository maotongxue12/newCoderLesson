package tree

//题目一：实现二叉树的先序、中序、后序遍历，包括递归方式和非递归方式
//递归先序遍历算法
func preorderTraversalR(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	res = append(res, head.val)
	res = append(res, preorderTraversalR(head.left)...)
	res = append(res, preorderTraversalR(head.right)...)
	return res
}

//递归中序遍历算法
func inorderTraversalR(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	inorderTraversalR(head.left)
	res = append(res, head.val)
	inorderTraversalR(head.right)
	return res
}

//递归后续遍历算法
func postorderTraversalR(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	postorderTraversalR(head.left)
	postorderTraversalR(head.right)
	res = append(res, head.val)
	return res
}
