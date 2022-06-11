package main

import "fmt"

//copy链表中的节点结构同原链表的节点结构是相同的，即copy链表中节点的next节点和random节点和原链表一一对应
//如：1->2->3->4->5其中1的random节点为3，4的random节点为2，假设copy链表为1'->2'->3'->4'->5'则1’的random节点为3‘，4’的random节点为2‘
//本题的难点在于如何根据原链表节点的random节点找到对应copy链表节点的random节点，故，假设能够找到节点1和1’，2和2‘。。。5和5’对应关系，
//就能找到每个copy节点对应的random节点，比如1的random节点为3，如果能够根据1和3分别找到1’节点和3‘节点，则就能确定1’节点的random指向
//解决思路：需要建立原链表与copy链表的对应关系，这样才能根据原链表中节点的random节点找到copy链表的random节点
//方法一：用一个map维护原链表与copy链表每个节点的对应关系
//方法二：将每个copy节点建立在原节点之后，如1->1'->2->2'->3->3'->4->4'->5->5',遍历链表，确定每个链表copy节点的random节点，然后拆分链表

type node struct {
	val int
	next *node
	random *node
}

//使用map结构
func copyRandomListMap(head *node) *node {
	if head == nil {
		return nil
	}
	copyListMap := make(map[*node]*node)
	cur := head
	for cur != nil {
		copyNode := &node{
			val: cur.val,
			next: nil,
			random: nil,
		}
		copyListMap[cur] = copyNode
		cur = cur.next
	}
	cur = head
	for cur != nil {
		copyListMap[cur].next = cur.next
		copyListMap[cur].random = cur.random
		cur = cur.next
	}
	return copyListMap[head]
}

//不使用额外的空间
func copyList(head *node) *node{
	cur := head
	for cur != nil {
		copyNode := &node{
			val: cur.val,
			next: nil,
			random: nil,
		}
		next := cur.next
		cur.next = copyNode
		copyNode.next = next
		cur = next
	}

	cur = head
	for cur != nil && cur.next != nil {
		cur.next.random = cur.random
		cur = cur.next.next
	}

	cur = head
	newHead := cur.next
	copyCur := newHead
	for cur != nil && cur.next != nil {
		copyCur = cur.next
		next := copyCur.next
		cur.next = copyCur.next
		if next == nil {
			copyCur.next = nil
		} else {
			copyCur.next = next.next
		}
		cur = next
	}

	return newHead
}

func newNode(val int, next *node, random *node) *node {
	return &node{
		val:  val,
		next: next,
		random: random,
	}
}

func printList(head *node) {
	for head != nil {
		if head.next != nil {
			fmt.Print(head.val)
			if head.random != nil {
				fmt.Print("*", head.random.val, "->")
			} else {
				fmt.Print("->")
			}

		} else {
			fmt.Print(head.val)
		}
		head = head.next
	}
	fmt.Println()
}

func main(){
	node3 := newNode(7, nil, nil)
	node2 := newNode(5, node3, nil)
	node1 := newNode(3, node2, nil)
	head := newNode(1, node1, nil)
	head.random = node2
	node2.random = node1
	printList(head)
	copyHead := copyRandomListMap(head)
	copyHead2 := copyList(head)
	printList(copyHead)
	printList(copyHead2)
}