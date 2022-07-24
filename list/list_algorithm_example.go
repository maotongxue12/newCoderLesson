package list

import (
	"fmt"
)

//单链表节点
type singleNode struct {
	Val      int
	NextNode *singleNode
}

//双链表节点
type doubleNode struct {
	Val      int
	PreNode  *doubleNode
	NextNode *doubleNode
}

//题目一：反转单向和双向链表 【题目】 分别实现反转单向链表和反转双向链表的函数。
//【要求】 如果链表长度为N，时间复杂度要求为O(N)，额外空间复杂度要求为O(1)

//空间复杂度要求为O(1)，则需要原地反转
func reverseSingleList(head *singleNode) *singleNode {
	var pre *singleNode = nil
	cur := head
	for cur != nil {
		next := cur.NextNode
		cur.NextNode = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseDoubleList(head *doubleNode) *doubleNode {
	var pre *doubleNode = nil
	cur := head
	for cur != nil {
		next := cur.NextNode
		cur.NextNode = pre
		cur.PreNode = next
		pre = cur
		cur = next
	}
	return pre
}

//题目二：打印两个有序链表的公共部分
//【题目】 给定两个有序链表的头指针head1和head2，打印两个链表的公共部分
//思路：两个链表都为有序链表，则可以根据当前值进行判断
func printCommonList(head1, head2 *singleNode) {
	if head1 == nil || head2 == nil {
		return
	}
	cur1 := head1
	cur2 := head2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur1 = cur1.NextNode
		} else if cur1.Val > cur2.Val {
			cur2 = cur2.NextNode
		} else {
			fmt.Print(cur1.Val, " ")
			cur1 = cur1.NextNode
			cur2 = cur2.NextNode
		}
	}
}

//题目三：判断一个链表是否为回文结构
//【题目】 给定一个链表的头节点head，请判断该链表是否为回文结构。
//例如：1->2->1，返回true。1->2->2->1，返回true。15->6->15，返回true。1->2->3，返回false。
//思路：方法一：借助于栈，需要额外空间复杂度，将链表节点依次加入栈，然后再逐个弹出，与链表节点进行比较，
func isPalindromeList(head *singleNode) bool {
	if head == nil {
		return true
	}

	//利用快慢指针找到链表的中点，如果链表节点个数为奇数，这应该将后半部分入栈，如果链表中节点个数为偶数，
	//则将中点的后半部分入栈即可（可以不包括中点）
	//初始点slow在fast之后，可以使得无论链表节点个数为奇数或者偶数，slow最后停在中点位置右侧位置，即为偶数时，停在后半部分起始点
	//为奇数时slow停留在中点后的第一个位置
	fast := head
	slow := head.NextNode
	for fast.NextNode != nil && fast.NextNode.NextNode != nil {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
	}
	//将链表的后半部分入栈
	cur := slow
	stack := make([]*singleNode, 0)
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.NextNode
	}
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		if temp.Val == head.Val {
			head = head.NextNode
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}

//进阶：如果链表长度为N，时间复杂度达到O(N)，额外空间复杂度达到O(1)。
//思路：找到链表的中的，对后半部分进行原地反转链表，空间复杂度为O(1)，即需要实现原链表比较
func isPalindromeList01(head *singleNode) bool {
	if head == nil {
		return true
	}
	//找到链表中点位置
	fast, slow := head, head.NextNode
	for fast.NextNode != nil && fast.NextNode.NextNode != nil {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
	}

	//反转后半部分链表
	cur := slow
	var pre *singleNode = nil
	for cur != nil {
		next := cur.NextNode
		cur.NextNode = pre
		pre = cur
		cur = next
	}

	//比较链表前半部分和后半部分是否相等,当存在节点不相等时，不能直接返回，需要把链表回复原状
	cur = pre
	flag := true
	for cur != nil {
		if cur.Val == head.Val {
			cur = cur.NextNode
			head = head.NextNode
		} else {
			flag = false
			break
		}
	}

	//将链表恢复原状
	cur = pre
	pre = nil
	for cur != nil {
		next := cur.NextNode
		cur.NextNode = pre
		pre = cur
		cur = next
	}

	return flag
}

//题目四：将单向链表按某值划分成左边小、中间相等、右边大的形式
//【题目】 给定一个单向链表的头节点head，节点的值类型是整型，再给定一个 整 数pivot。
//实现一个调整链表的函数，将链表调整为左部分都是值小于 pivot 的节点，中间部分都是值等于pivot的节点，右部分都是值大于 pivot的节点。
//除这个要求外，对调整后的节点顺序没有更多的要求。 例如：链表9->0->4->5- >1，pivot=3。 调整后链表可以是1->0->4->9->5，
//也可以是0->1->9->5->4。总之，满足左部分都是小于3的节点，中间部分都是等于3的节点（本例中这个部分为空）右部分都是大于3的节点即可。
//对某部分内部的节点顺序不做要求。
