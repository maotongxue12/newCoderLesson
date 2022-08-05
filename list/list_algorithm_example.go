package list

import (
	"fmt"
	"math"
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

//单链表插入数据
func insertSingNode(head, node *singleNode) {
	cur := head
	if head != nil {
		for cur.NextNode != nil {
			cur = cur.NextNode
		}
		cur.NextNode = node
	} else {
		head = node
	}
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
//【题目】 给定一个单向链表的头节点head，节点的值类型是整型，再给定一个整数pivot。
//实现一个调整链表的函数，将链表调整为左部分都是值小于 pivot的节点，中间部分都是值等于pivot的节点，右部分都是值大于 pivot的节点。
//除这个要求外，对调整后的节点顺序没有更多的要求。 例如：链表9->0->4->5- >1，pivot=3。 调整后链表可以是1->0->4->9->5，
//也可以是0->1->9->5->4。总之，满足左部分都是小于3的节点，中间部分都是等于3的节点（本例中这个部分为空）右部分都是大于3的节点即可。
//对某部分内部的节点顺序不做要求。
//思路：可参考快速排序中的左右区间法，不需要递归，仅用一次排序就可以，但该方法需要将链表中的值转为数组，然后在反转为链表，需要额外空间复杂度o(n)
func dividePartitionList(head *singleNode, pivot int) {
	if head == nil {
		return
	}
	cur := head
	//将链表按顺序放在切片中
	temp := make([]int, 0)
	for cur != nil {
		temp = append(temp, cur.Val)
		cur = cur.NextNode
	}
	partition(temp, 0, len(temp)-1, pivot)
	//将排好序的数组恢复为链表
	cur = head
	for i := 0; i < len(temp); i++ {
		cur.Val = temp[i]
		cur = cur.NextNode
	}
}

func partition(arr []int, left, right, pivot int) {
	lessArea := left - 1
	cur := 0
	moreArea := right
	for cur <= moreArea {
		if arr[cur] < pivot {
			lessArea++
			arr[lessArea], arr[cur] = arr[cur], arr[lessArea]
			cur++
		} else if arr[cur] == pivot {
			cur++
		} else {
			arr[cur], arr[moreArea] = arr[moreArea], arr[cur]
			moreArea--
		}
	}
}

//题目四：进阶 在原问题的要求之上再增加如下两个要求。
//在左、中、右三个部分的内部也做顺序要求，要求每部分里的节点从左 到右的
//顺序与原链表中节点的先后次序一致。时间复杂度请达到O(N)，额外空间复杂度请达到O(1)
//思路：利用快排思想时，空间复杂度为o(n), 且快排是不稳定算法，该题目做顺序要求，即一趟排序要实现稳定
//有空间复杂度要求，所以不能使用快速排序相关算法，可以结合链表的特点：
//首先，第一次遍历链表，找到min，equal，max三个值，即找到链表中第一个小于，等于和大于pivot的值
//再进行第二次遍历，分别将小于，等于和大于pivot的值挂到前面的节点之后，最后形成小于pivot值，等于pivot值和大于pivot值的三个链表
//最后将三个链表进行拼接，既可保证各个节点值有序，需要注意的是代码中在对三个链表拼接时，要注意可能存在空链表的异常情况
func divideAdvanceList(head *singleNode, pivot int) *singleNode {
	if head == nil {
		return nil
	}
	cur := head
	var minHead, equalHead, maxHead, minTail, equalTail, maxTail *singleNode

	//两次遍历可以优化为一次遍历既可以将对应节点挂到相应的区间节点上
	for cur != nil {
		if cur.Val < pivot {
			if minHead == nil {
				minHead, minTail = cur, cur
			} else {
				minTail.NextNode = cur
				minTail = cur
			}
		} else if cur.Val == pivot {
			if equalHead == nil {
				equalHead, equalTail = cur, cur
			} else {
				equalTail.NextNode = cur
				equalTail = cur
			}
		} else {
			if maxHead ==nil {
				maxHead, maxTail = cur, cur
			} else {
				maxTail.NextNode = cur
				maxTail = cur
			}
		}
		cur = cur.NextNode
	}

	//要注意把链表最后一个节点的nextNode置为nil，避免出现链表循环
	if minTail != nil {
		minTail.NextNode = nil

	}
	if equalTail != nil {
		equalTail.NextNode = nil
	}
	if maxTail != nil {
		maxTail.NextNode = nil
	}

	//合并三个链表，需要考虑链表为空的异常情况
	if minHead != nil {
		if equalHead != nil {
			minTail.NextNode = equalHead
			if maxHead != nil {
				equalTail.NextNode = maxHead
			}
		} else {
			if maxHead != nil {
				minTail.NextNode = maxHead
			}
		}
		return minHead
	} else {
		if equalHead != nil {
			if maxHead != nil {
				equalTail.NextNode = maxHead
			}
			return equalHead
		} else {
			return maxHead
		}
	}
}

//题目五：复制含有随机指针节点的链表
//【题目】 一种特殊的链表节点类描述如下：
type randomNode struct {
	val int
	next *randomNode
	random *randomNode
}
//Node类中的value是节点值，next指针和正常单链表中next指针的意义一样，都指向下一个节点，rand指针是Node类中新增的指针，这个指
//针可能指向链表中的任意一个节点，也可能指向null。 给定一个由Node节点类型组成的无环单链表的头节点head，请实现一个函数完成
//这个链表中所有结构的复制，并返回复制的新链表的头节点。
//思路：借助hashMap来维护节点之间的关系，其中，map中的key值分别存储原始节点，value分别存储对应的复制节点，根据key值可以匹配对应的复制节点
//根据key值找到对应next节点和random节点，然后由next节点和random的对应value值，即可确定key值复制节点的对应next节点和random节点
func copyRandomList(head *randomNode) *randomNode {
	if head == nil {
		return nil
	}
	cur := head
	tempMap := make(map[*randomNode]*randomNode)
	//第一次遍历链表，将复制节点和初始节点更新到map中
	for cur != nil {
		tempNode := &randomNode{
			val: cur.val,
			next: nil,
			random: nil,
		}
		tempMap[cur] = tempNode
		cur = cur.next
	}
	//第二次遍历链表，根据map中key-value关系初始化copy的链表
	cur = head
	copyHead := tempMap[cur]
	for cur != nil {
		tempMap[cur].next = tempMap[cur.next]
		tempMap[cur].random = tempMap[cur.random]
		cur = cur.next
	}
	return copyHead
}

//题目五：进阶 不使用额外的数据结构，只用有限几个变量，且在时间复杂度为 O(N)内完成原问题要实现的函数。
//思路：要求不能使用额外的数据结构，可以将每个copy节点建立在原节点之后，如1->1'->2->2'->3->3'->4->4'->5->5',遍历链表，
//确定每个链表copy节点的random节点，然后拆分链表
func copyRandomList02(head *randomNode) *randomNode {
	if head == nil {
		return nil
	}
	//第一次遍历，分别将复制节点插入到初始链表对应节点之后
	cur := head
	for cur != nil {
		temp := &randomNode{
			val: cur.val,
			next: nil,
			random: nil,
		}
		next := cur.next
		cur.next = temp
		temp.next = next
		cur = next
	}

	//第二次遍历包含复制节点的整个链表，根据初始节点next值和random值确定复制节点的random值
	cur = head
	for cur != nil {
		if cur.random != nil {
			cur.next.random = cur.random.next
		}
		cur = cur.next.next
	}

	//第三次遍历链表，进行断链，提取出复制节点的链表
	cur = head
	newHead := cur.next
	//因为初始链表的每个节点都进行了复制，所以复制后的链表节点为偶数个节点，
	//即当cur!=nil时，cur.next一定也不等于nil
	for cur != nil {
		copyNode := cur.next
		nextNode := copyNode.next
		cur.next = nextNode
		if nextNode != nil {
			copyNode.next = nextNode.next
		} else {
			copyNode.next = nil
		}
		cur = nextNode
	}
	return newHead
}

//题目六：两个单链表相交的一系列问题
//【题目】 在本题中，单链表可能有环，也可能无环。给定两个单链表的头节点 head1和head2，这两个链表可能相交，也可能
//不相交。请实现一个函数， 如果两个链表相交，请返回相交的第一个节点；如果不相交，返回null 即可。
//要求：如果链表1的长度为N，链表2的长度为M，时间复杂度请达到 O(N+M)，额外空间复杂度请达到O(1)。
//思路：由于单链表可能有环也可能无环，则相交共分为三种情况：（1）无环和无环；（2）有环和无环；（3）有环和有环
//其中由链表特性可知，有环和无环的链表是不可能相交的（相交时相当于两个链表都有环了）
func listIntersect(head1, head2 *singleNode) *singleNode {
	//判断两个链表是否有环
	loop1 := isLoopList(head1)
	loop2 := isLoopList(head2)
	//无环和无环链表的相交情况
	if loop1 == nil && loop2 == nil {
		//分别求出两个链表的长度和end节点位置，如果end节点位置不相同，则两个链表必不相交
		//end节点相同时，计算长度差，然后长链表减去长度差之后开始遍历，短链表直接遍历，直至两个点相遇
		lengthA := 0
		curA := head1
		var endA *singleNode
		for curA != nil {
			lengthA++
			endA = curA
			curA = curA.NextNode
		}

		curB := head2
		lengthB := 0
		var endB *singleNode
		for curB != nil {
			lengthB++
			endB = curB
			curB = curB.NextNode
		}

		if endA != endB {
			return nil
		}
		intersectNode := noLoopListIntersect(head1, head2, lengthA, lengthB)
		return intersectNode
	}
	//有环和无环链表，不可能相交

	//有环和有环链表相交情况，分为三种情形：在入环节点之前相交；在环上相交，此时相交点可以为链表1的入环点，也可以为链表2的入环点；两个链表不相交
	if loop1 != nil && loop2 != nil {
		//情形一：在入环节点前相交，此时相交点的求法等同于无环链表与无环链表的相交
		if loop1 == loop2 {
			length := 0
			curA := head1
			curB := head2
			for curA != loop1 {
				length++
				curA = curA.NextNode
			}
			for curB != loop2 {
				length--
				curB = curB.NextNode
			}
			curA = head1
			curB = head2
			if length > 0 {
				for length > 0 {
					curA = curA.NextNode
					length--
				}
			} else {
				length = int(math.Abs(float64(length)))
				for length > 0 {
					curB = curB.NextNode
					length--
				}
			}
			//该循环返回相交点，包含了入环点相交
			for curA != curB {
				curA = curA.NextNode
				curB = curB.NextNode
			}
			return curA
		} else {
			//情形二
			curA := loop1.NextNode
			for curA != loop1 {
				if curA == loop2 {
					return loop2
				}
				curA = curA.NextNode
			}
			//情形三
			return nil
		}
	}
	return nil
}

//判断链表是否有环，如果有环返回入环的第一个节点，否则返回nil
//思路一：利用setMap(setMap结构即只存储key值，无value值，golang中没有setMap的实现，可以手动实现一个)结构
//首先cur指针从头遍历链表，将每个节点加入到set集合中，如果遍历到cur指针指向nil，则该链表无环，如果遍历到第一个在set集合
//中已经存在的节点，则该节点既是入环的第一个节点

// Set 利用空struct实现一个set结构
type Set map[*singleNode]struct{}

func (s Set) Has(key *singleNode) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key *singleNode) {
	s[key] = struct{}{}
}

func (s Set) Delete(key *singleNode) {
	delete(s, key)
}

func isLoopListForSet(head *singleNode) *singleNode {
	if head == nil {
		return nil
	}
	cur := head
	setMap := make(Set)
	for cur != nil {
		if setMap.Has(cur) {
			return cur
		}
		setMap.Add(cur)
		cur = cur.NextNode
	}
	return nil
}

//思路二 ：利用快慢指针法判断是否是循环链表，并返回第一个入环节点
func isLoopList(head *singleNode) *singleNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.NextNode != nil {
		fast = fast.NextNode.NextNode
		slow = slow.NextNode
		//如果快慢指针相遇，则肯定有环，需要判断入环的第一个节点
		//相遇之后，slow指针不变，fast指针指向头结点，然后步长改为1，分别进行遍历，
		//此后fast和slow指针相遇的第一个节点就是入环的第一个节点（数学证明）
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.NextNode
				slow = slow.NextNode
			}
			return fast
		}
	}
	return nil
}

//判断两个无环链表的相交情况，如果相交返回相交点，否则返回nil
//为了和两个有环链表相交情形一复用代码，传入end1和end2参数
func noLoopListIntersect(head1, head2 *singleNode, lengthA, lengthB int) *singleNode {
	var curA *singleNode
	var curB *singleNode
	lengthDif := int(math.Abs(float64(lengthA)-float64(lengthB)))
	curA = head1
	curB = head2
	if lengthA >= lengthB {
		for i := 0; i < lengthDif; i++ {
			curA = curA.NextNode
		}
	} else {
		for i := 0; i < lengthDif; i++ {
			curB = curB.NextNode
		}
	}
	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.NextNode
		curB = curB.NextNode
	}
	return nil
}