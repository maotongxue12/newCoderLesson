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

//前序遍历序列化--递归
func preOrderTraverSerialization(head *treeUtil.Node) string {
	if head == nil {
		return "# "
	}
	res := fmt.Sprintf("%d", head.Value) + " "
	res += preOrderTraverSerialization(head.Left)
	res += preOrderTraverSerialization(head.Right)
	return res
}

//前序遍历序列化--非递归
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
func preorderDeserialization(str string) (head *treeUtil.Node){
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

//中序遍历序列化--递归
func inOrderSerialization(head *treeUtil.Node) string {
	if head == nil {
		return "#"
	}
	var res string
	res += inOrderSerialization(head.Left)
	res = res + " " + fmt.Sprintf("%d", head.Value) + " "
	res += inOrderSerialization(head.Right)
	return res
}

//中序遍历序列化--非递归
func inOrderSerializationUnRecursive(head *treeUtil.Node) string {
	if head == nil {
		return ""
	}
	var res string
	stack := make([]*treeUtil.Node, 0)

	for head != nil || len(stack) != 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left
			if head == nil {
				val, _ := strconv.Atoi("#")
				stack = append(stack, &treeUtil.Node{Value: val})
			}
		} else {
			head = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leaf, _ := strconv.Atoi("#")
			if head.Value == leaf {
				res = res + "#" + " "
				head = nil
				continue
			} else {
				res = res + fmt.Sprintf("%d", head.Value) + " "
			}
			head = head.Right
			if head == nil {
				val, _ := strconv.Atoi("#")
				stack = append(stack, &treeUtil.Node{Value: val})
			}
		}
	}
	//返回res时需要删除res中最后一个多余的空格
	return res[:len(res)-1]
}

//中序遍历反序列化,因无法确定二叉树的根节点，因此无法反序列化
var i int = -1
func inOrderDeserialization(str string) (head *treeUtil.Node){
	tempStr := strings.Split(str, "_")
	i++
	value := tempStr[i]
	if value == "#" {
		return nil
	}

	//val,_ := strconv.Atoi(value)
	return head
}

//后续遍历反序列化--递归
var j int
func postOrderDeserialization(str string) (head *treeUtil.Node) {
	strs := strings.Split(str, "_")
	j = len(strs)
	head = postDeserialization(strs)
	return head
}

func postDeserialization(strs []string) (head *treeUtil.Node) {
	j--
	if strs[j] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(strs[j])
	head = newNode(val, nil, nil)
	head.Right = postDeserialization(strs)
	head.Left = postDeserialization(strs)
	return head
}

//后续遍历序列化--递归
func postOrderSerialization(head *treeUtil.Node) string {
	if head == nil {
		return "#"
	}
	var res string
	res += postOrderSerialization(head.Left) + " "
	res += postOrderSerialization(head.Right) + " "
	res += fmt.Sprintf("%d", head.Value)
	return res
}

//后续遍历序列化--非递归
func postOrderSerializationUnRecursive(head *treeUtil.Node) string {
	if head == nil {
		return ""
	}
	stack := make([]*treeUtil.Node, 0)
	tempStack := make([]*treeUtil.Node, 0)
	stack = append(stack, head)
	var res string
	//var result string
	for len(stack) != 0 {
		head = stack[len(stack)-1]
		tempStack = append(tempStack, head)
		stack = stack[:len(stack)-1]
		temp, _ := strconv.Atoi("#")
		if head.Value == temp {
			res += "#" + " "
			continue
		}
		res += fmt.Sprintf("%d", head.Value) + " "
		if head.Left != nil {
			stack = append(stack, head.Left)
		} else {
			val, _ := strconv.Atoi("#")
			stack = append(stack, &treeUtil.Node{Value: val})
		}
		if head.Right != nil {
			stack = append(stack, head.Right)
		} else {
			val, _ := strconv.Atoi("#")
			stack = append(stack, &treeUtil.Node{Value: val})
		}
	}
	res = res[:len(res)-1]
	return reverseString(res)
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//层次遍历序列化--非递归
func levelSerialization(head *treeUtil.Node) string {
	if head == nil {
		return ""
	}
	var res string
	queue := make([]*treeUtil.Node, 0)
	res += fmt.Sprintf("%d", head.Value) + " "
	queue = append(queue, head)
	for len(queue) != 0 {
		value := queue[0]
		queue = queue[1:]
		if value.Left != nil {
			res += fmt.Sprintf("%d", value.Left.Value) + " "
			queue = append(queue, value.Left)
		} else {
			res += "# "
		}
		if value.Right != nil {
			res += fmt.Sprintf("%d", value.Right.Value) + " "
			queue = append(queue, value.Right)
		} else {
			res += "# "
		}
	}
	return res
}

//层次遍历反序列化
func levelDeserialization(string2 string) (head *treeUtil.Node) {
	strs := strings.Split(string2, "_")
	var index int = 0
	queue := make([]*treeUtil.Node, 0)
	head = getChildNode(strs[index])
	queue = append(queue, head)
	index++
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left = getChildNode(strs[index])
		index++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		node.Right = getChildNode(strs[index])
		index++
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return head
}

func getChildNode(val string) *treeUtil.Node {
	if "#" == val {
		return nil
	} else {
		nodeVal, _ := strconv.Atoi(val)
		return newNode(nodeVal, nil, nil)
	}
}