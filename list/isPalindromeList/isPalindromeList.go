package main

import (
	"fmt"
)

type node struct {
	Val  int
	Next *node
}

func printList(head *node) {
	for head != nil {
		if head.Next != nil {
			fmt.Print(head.Val, "->")
		} else {
			fmt.Print(head.Val)
		}
		head = head.Next
	}
	fmt.Println()
}

func newNode(val int, next *node) *node {
	return &node{
		Val:  val,
		Next: next,
	}
}

//方法一：可以用栈来判断链表是否为回文结构，栈的长度可以为链表的一半，也可以为整个链表
//将链表的后半部分入栈，然后逐个弹出栈并和链表的前半部分进行比较，判断是否为回文结构。时间复杂度o(n), 空间复杂度o(n)
func isPalindromeList(head *node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//利用快慢指针找到链表的中点，如果链表节点个数为奇数，这应该将后半部分入栈，如果链表中节点个数为偶数，
	//则将中点的后半部分入栈即可（可以不包括中点）
	fast := head
	slow := head.Next
	//确保无论链表节点数为偶数还是奇数，都保证slow指针指向右半部分第一个（奇数时不包括中点）
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//将链表的后半部分压到栈中，go语言可采用切片结构
	temp := make([]int, 0)
	for slow != nil {
		temp = append(temp, slow.Val)
		slow = slow.Next
	}

	//将切片中数据模拟栈的结构弹出，逐一与链表中的前半部分比较
	size := len(temp)
	for i := 0; i < size; i++ {
		if head.Val != temp[len(temp)-1] {
			return false
		} else {
			head = head.Next
			temp = temp[:len(temp)-1]
		}
	}
	return true
}

//方法二：直接将链表的后半部分倒置，然后比较链表的后半部分和链表的前半部分
func isPalindromeList2(head *node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//找到链表的中点右边第一个位置
	fast, slow := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//倒置链表后半部分
	var pre *node
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	//分别从前后两部分比较链表
	res := true
	leftHead := head
	temp := pre		//记录倒置后链表的头元素，用于后面恢复链表原状
	for pre != nil && leftHead != nil {
		if pre.Val != leftHead.Val {
			res = false
			break
		} else {
			pre = pre.Next
			leftHead = leftHead.Next
		}
	}

	//将链表倒置的后半部分恢复原状
	cur = temp
	pre = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return res
}

func main() {
	node3 := newNode(6, nil)
	node2 := newNode(4, node3)
	node1 := newNode(4, node2)
	head1 := newNode(6, node1)

	node03 := newNode(1, nil)
	node02 := newNode(3, node03)
	node01 := newNode(2, node02)
	head2 := newNode(1, node01)

	node002 := newNode(3, nil)
	node001 := newNode(2, node002)
	head3 := newNode(3, node001)

	fmt.Println(isPalindromeList(head1))
	fmt.Println(isPalindromeList(head2))
	fmt.Println(isPalindromeList(head3))

	fmt.Println()

	printList(head1)
	fmt.Println(isPalindromeList2(head1))
	printList(head1)

	printList(head2)
	fmt.Println(isPalindromeList2(head2))
	printList(head2)

	printList(head3)
	fmt.Println(isPalindromeList2(head3))
	printList(head3)
}


