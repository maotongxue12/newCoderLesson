package tree

import (
	"strconv"
)

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
		tempRes[length-1] = res[i]
		length--
	}
	return tempRes
}

//非递归后续遍历
func postorderTraversal02(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		//防止左右子树被重复压栈，head=temp.left则左子树已被压入栈，head=temp.right则右子树已被压入栈
		//此时左子树已经被压入栈
		if temp.left != nil && head != temp.left && head != temp.right {
			stack = append(stack, temp.left)
		} else if temp.right != nil && head != temp.right {
			stack = append(stack, temp.right)
		} else {
			stack = stack[:len(stack)-1]
			res = append(res, temp.val)
			head = temp
		}
	}
	return res
}

//层次遍历
func levelTraversal(head *treeNode) (res []int) {
	if head == nil {
		return
	}
	queue := make([]*treeNode, 0)
	queue = append(queue, head)
	for len(queue) != 0 {
		temp := queue[0]
		res = append(res, temp.val)
		queue = queue[1:]
		if temp.left != nil {
			queue = append(queue, temp.left)
		}
		if temp.right != nil {
			queue = append(queue, temp.right)
		}
	}
	return res
}

//题目二：介绍二叉树的序列化和反序列化
//先序遍历序列化
func preTraversalSerialization(head *treeNode) (res string) {
	if head == nil {
		return "#_"
	}
	res = strconv.Itoa(head.val) + "_"
	res += preTraversalSerialization(head.left)
	res += preTraversalSerialization(head.right)
	return res
}

//先序遍历序列化--非递归
func preTraversalSerializationNonRec(head *treeNode) (res string) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, _ := strconv.Atoi("#")
		if temp.val == left {
			res += "#" + "_"
			continue
		}
		res = res + strconv.Itoa(temp.val) + "_"

		if temp.right != nil {
			stack = append(stack, temp.right)
		} else {
			tempVal, _ := strconv.Atoi("#")
			stack = append(stack, &treeNode{val: tempVal})
		}
		if temp.left != nil {
			stack = append(stack, temp.left)
		} else {
			tempVal, _ := strconv.Atoi("#")
			stack = append(stack, &treeNode{val: tempVal})
		}
	}
	return res
}

//中序遍历序列化
func inTraversalSerialization(head *treeNode) (res string) {
	if head == nil {
		return "#_"
	}
	res = inTraversalSerialization(head.left)
	res += strconv.Itoa(head.val) + "_"
	res += inTraversalSerialization(head.right)
	return res
}

//非递归中序遍历序列化
func inTraversalSerializationNonRec(head *treeNode) (res string) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.left
			if head == nil {
				left, _ := strconv.Atoi("#")
				stack = append(stack, &treeNode{val: left})
			}
		} else {
			head = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left, _ := strconv.Atoi("#")
			if head.val == left {
				res += "#" + "_"
				head = nil
				continue
			}
			res += strconv.Itoa(head.val) + "_"
			head = head.right
			if head == nil {
				left, _ := strconv.Atoi("#")
				stack = append(stack, &treeNode{val: left})
			}
		}
	}
	return res
}

//后续遍历序列化
func postTraversalSerialization(head *treeNode) (res string) {
	if head == nil {
		return "#_"
	}
	res = postTraversalSerialization(head.left)
	res += postTraversalSerialization(head.right)
	res += strconv.Itoa(head.val) + "_"
	return res
}

//非递归后续序列化
func postTraversalSerializationNonRec(head *treeNode) (res string) {
	if head == nil {
		return
	}
	stack := make([]*treeNode, 0)
	stack = append(stack, head)

	return res
}
