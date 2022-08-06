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

//先序遍历非递归算法
//先序遍历从根节点开始遍历，则可采用栈先进后出的思想，出栈时访问该节点
//存放栈的作用是访问该节点的右节点
func preorderTraversal(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		res = append(res, temp.val)
		stack = stack[:len(stack)-1]
		if temp.right != nil {
			stack = append(stack, temp.right)
		}
		if temp.left != nil {
			stack = append(stack, temp.left)
		}
	}
	return res
}

//递归中序遍历算法
func inorderTraversalR(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	//1
	subResL := inorderTraversalR(head.left)
	//2
	res = append(subResL, head.val)
	subResR := inorderTraversalR(head.right)
	res = append(res, subResR...)
	//3
	return res
}

//非递归中序遍历算法
func inorderTraversal(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.left
		} else {
			head = stack[len(stack)-1]
			res = append(res, head.val)
			stack = stack[:len(stack)-1]
			head = head.right
		}
	}
	return res
}

//递归后续遍历算法
func postorderTraversalR(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	res = append(res, postorderTraversalR(head.left)...)
	res = append(res, postorderTraversalR(head.right)...)
	res = append(res, head.val)
	return res
}

//非递归后续遍历算法
//参考先序遍历：中-》左-》右，将先序遍历修改为：中-》右-》左，让偶将该顺序入栈，再反向弹出，即变为后续
//遍历：左-》右-》中
func postorderTraversal(head *treeNode) (res []int) {
	if head == nil {
		return res
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		res = append(res, temp.val)
		stack = stack[:len(stack)-1]
		if temp.left != nil {
			stack = append(stack, temp.left)
		}
		if temp.right != nil {
			stack = append(stack, temp.right)
		}
	}
	length := len(res)
	tempRes := make([]int, length)
	for i := 0; i < len(res); i++ {
		tempRes[length-1] =  res[i]
		length--
	}
	return tempRes
}