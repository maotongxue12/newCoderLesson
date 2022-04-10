package treeAlgorithm

import (
	"fmt"
	"newCode/tree/treeUtil"
	"strconv"
	"strings"
)

func newNode(Value int, lf, rt *treeUtil.Node) *treeUtil.Node {
	return &treeUtil.Node{
		Value: Value,
		Left:  lf,
		Right: rt,
	}
}

func newNodeParent(Value int, lf, rt, par *treeUtil.Node) *treeUtil.Node {
	return &treeUtil.Node{
		Value:  Value,
		Left:   lf,
		Right:  rt,
		Parent: par,
	}
}

//递归遍历 先序、中序、后序
func preOrderTraverseTree(head *treeUtil.Node) {
	if head == nil {
		return
	}
	fmt.Print(head.Value, " ")
	preOrderTraverseTree(head.Left)
	preOrderTraverseTree(head.Right)
}

func inOrderTraverseTree(head *treeUtil.Node) {
	if head == nil {
		return
	}
	inOrderTraverseTree(head.Left)
	fmt.Print(head.Value, " ")
	inOrderTraverseTree(head.Right)
}

func postOrderTraverseTree(head *treeUtil.Node) {
	if head == nil {
		return
	}
	postOrderTraverseTree(head.Left)
	postOrderTraverseTree(head.Right)
	fmt.Print(head.Value, " ")
}

//非递归遍历先序、中序、后序
func preOrderUnRecursive(head *treeUtil.Node) {
	if head == nil {
		return
	}
	stack := make([]*treeUtil.Node, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		fmt.Print(stack[len(stack)-1].Value, " ")
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
}

func inOrderUnRecursive(head *treeUtil.Node) {
	if head == nil {
		return
	}
	stack := make([]*treeUtil.Node, 0)
	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
		} else {
			head = stack[len(stack)-1]
			fmt.Print(head.Value, " ")
			stack = stack[:len(stack)-1]
			head = head.Right
		}
	}
}

func postOrderUnRecursive(head *treeUtil.Node) {
	if head == nil {
		return
	}
	stack01 := make([]*treeUtil.Node, 0)
	stack02 := make([]*treeUtil.Node, 0)
	stack01 = append(stack01, head)
	for len(stack01) != 0 {
		value := stack01[len(stack01)-1]
		stack02 = append(stack02, value)
		stack01 = stack01[:len(stack01)-1]
		if value.Left != nil {
			stack01 = append(stack01, value.Left)
		}
		if value.Right != nil {
			stack01 = append(stack01, value.Right)
		}
	}
	for len(stack02) != 0 {
		value := stack02[len(stack02)-1]
		fmt.Print(value.Value, " ")
		stack02 = stack02[:len(stack02)-1]
	}
}

func postOrderUnRecursive02(head *treeUtil.Node) {
	if head == nil {
		return
	}
	stack := make([]*treeUtil.Node, 0)
	stack = append(stack, head)
	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		//保证左右子树不被重复压栈
		if temp.Left != nil && head != temp.Left && head != temp.Right {
			stack = append(stack, temp.Left)
		} else if temp.Right != nil && head != temp.Right {
			stack = append(stack, temp.Right)
		} else {
			fmt.Print(temp.Value, " ")
			stack = stack[:len(stack)-1]
			head = temp
		}
	}
}

//层次遍历, 利用队列
func levelTraversal(head *treeUtil.Node) {
	if head == nil {
		return
	}
	list := make([]*treeUtil.Node, 0)
	list = append(list, head)
	for len(list) != 0 {
		value := list[0]
		fmt.Print(value.Value, " ")
		/*if len(list) > 1 {
			list = list[1:len(list)]
		} else {
			list = make([]*treeUtil.Node, 0)
		}*/
		list = list[1:]
		if value.Left != nil {
			list = append(list, value.Left)
		}
		if value.Right != nil {
			list = append(list, value.Right)
		}
	}
}

//前序遍历树的序列化
func preOrderTraverSerialization(head *treeUtil.Node) string {
	if head == nil {
		return "# "
	}
	res := fmt.Sprintf("%d", head.Value) + " "
	res += preOrderTraverSerialization(head.Left)
	res += preOrderTraverSerialization(head.Right)
	return res
}

//前序遍历非递归序列化
func preOrderTraverSerializationUnRecursive(head *treeUtil.Node) string {
	if head == nil {
		return ""
	}
	stack := make([]*treeUtil.Node, 0)
	stack = append(stack, head)
	var res string
	for len(stack) != 0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		leaf, _ := strconv.Atoi("#")
		if value.Value == leaf {
			res = res + "#" + " "
			continue
		}
		res = res + fmt.Sprintf("%d", value.Value) + " "

		if value.Right != nil {
			stack = append(stack, value.Right)
		} else {
			val, _ := strconv.Atoi("#")
			stack = append(stack, &treeUtil.Node{Value: val})
		}
		if value.Left != nil {
			stack = append(stack, value.Left)
		} else {
			val, _ := strconv.Atoi("#")
			stack = append(stack, &treeUtil.Node{Value: val})
		}
	}
	return res
}

//前序反序列化
func inorderDeserialization(str string) (head *treeUtil.Node){
	string := strings.Split(str, "_")
	head = deserialization(string)
	fmt.Println(string)
	return head
}

var index int = -1
func deserialization(str []string) (head *treeUtil.Node) {
	index++
	value := str[index]
	if value == "#" {
		return nil
	}
	val, _ := strconv.Atoi(value)
	head = newNode(val, nil, nil)
	head.Left = deserialization(str)
	head.Right = deserialization(str)
	return head
}