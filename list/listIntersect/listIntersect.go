package main

import (
	"fmt"
	"math"
)

type node struct {
	val int
	next *node
}
//解题思路：先利用快慢指针法判断是否有环，然后判断两个链表是否相交
//分为有环和有环；有环和无环；无环和无环三种情况
func listIsLoop(head *node) *node {
	//利用双指针判断两链表是否有环并返回入环的第一个节点
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		//fast==slow则存在环，此时fast指向head然后和slow分别进行next遍历，相交点即为入环的第一个节点
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.next
				slow = slow.next
			}
			return fast
		}
	}
	return nil
}

//如果两个链表一个有环一个无环则两链表不可能相交

//两个链表无环时，求出两个链表的长度差，然后长链表减去长度差之后开始遍历，相交点即为两链表的第一个相交点
func getIntersectNodeNoLoop(head1 *node, head2 *node) *node {
	cur1 := head1
	cur2 := head2
	var length int = 0
	for cur1 != nil {
		cur1 = cur1.next
		length++
	}
	for cur2 != nil {
		cur2 = cur2.next
		length--
	}
	if cur1 != cur2 {
		return nil
	}
	cur1 = head1
	cur2 = head2
	if length > 0 {
		for length > 0 {
			cur1 = cur1.next
			length--
		}
	} else {
		for math.Abs(float64(length)) > 0 {
			cur2 = cur2.next
			length--
		}
	}
	for cur1 != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return nil
}

//两个链表有环时，分为三种情况
//1. 两个有环链表相交于入环前的节点，该情景同两个无环链表找到交点的方法相同
//2. 两个有环链表相交于环中一点
//3. 两个有环链表不相交
//情况一
func getIntersectBothLoop(head1 *node, head2 *node, loop1 *node, loop2 *node)  *node {
	//情况1，此时相交于入环前的节点（包括入环点）
	if loop1 == loop2 {
		cur1 := head1
		cur2 := head2
		var length int
		for cur1 != loop1 {
			cur1 = cur1.next
			length++
		}
		for cur2 != loop2 {
			cur2 = cur2.next
			length--
		}
		cur1 = head1
		cur2 = head2
		if length > 0 {
			for length > 0 {
				cur1 = cur1.next
				length--
			}
		} else {
			leng := math.Abs(float64(length))
			for leng > 0 {
				cur2 = cur2.next
				leng--
			}
		}
		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
	//情形2和情形3
		cur1 := loop1.next
		for cur1 != loop1 {
			//情形2
			if cur1 == loop2 {
				return cur1
			}
			cur1 = cur1.next
		}
		//情形3
		return nil
	}
}

func getIntersectNode(head1, head2 *node) (intersectNode *node) {
	loop1 := listIsLoop(head1)
	loop2 := listIsLoop(head2)
	if loop1 != nil || loop2 != nil {
		intersectNode = getIntersectBothLoop(head1, head2, loop1, loop2)
		return intersectNode
	} else if loop1 == nil && loop2 == nil {
		intersectNode = getIntersectNodeNoLoop(head1, head2)
		return intersectNode
	} else {
		return intersectNode
	}
}

func printList(head *node) {
	for head != nil {
		if head.next != nil {
			fmt.Print(head.val, "->")
		} else {
			fmt.Print(head.val)
		}
		head = head.next
	}
	fmt.Println()
}

func newNode(val int, next *node) *node {
	return &node{
		val:  val,
		next: next,
	}
}

func main() {
	//1->2->3->4->5->6->7
	node6 := newNode(7, nil)
	node5 := newNode(6, node6)
	node4 := newNode(5, node5)
	node3 := newNode(4, node4)
	node2 := newNode(3, node3)
	node1 := newNode(2, node2)
	head1 := newNode(1, node1)

	//0->9->8->6->7
	node02 := newNode(8, node5)
	node01 := newNode(9, node02)
	head2 := newNode(0, node01)
	fmt.Println(getIntersectNode(head1, head2))

	//1->2->3->4->5->6->7->4...
	node006 := newNode(7, nil)
	node005 := newNode(6, node006)
	node004 := newNode(5, node005)
	node003 := newNode(4, node004)
	node002 := newNode(3, node003)
	node001 := newNode(2, node002)
	head01 := newNode(1, node001)
	node006.next = node003

	// 0->9->8->2...
	node0003 := newNode(2, nil)
	node0002 := newNode(8, node0003)
	node0001 := newNode(9, node0002)
	head02 := newNode(0, node0001)
	node0003.next = node001
	fmt.Println(getIntersectNode(head01, head02))

	// 0->9->8->6...
	node14 := newNode(8, nil)
	node15 := newNode(9, node14)
	head03:= newNode(0, node15)
	node14.next = node005
	fmt.Println(getIntersectNode(head01, head03))

}