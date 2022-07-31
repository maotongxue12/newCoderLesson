package tree

import (
	"fmt"
	"math"
	"strconv"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func newTreeNode(value int, leftNode, rightNode *treeNode) *treeNode {
	return &treeNode{
		val:   value,
		left:  leftNode,
		right: rightNode,
	}
}

//获取树的高度
func getTreeHeight(root *treeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(getTreeHeight(root.left)), float64(getTreeHeight(root.left))))
}

//递归填充数组
func writeArray(root *treeNode, row, treeHeight int, resArray [][]string) {
	if root == nil {
		return
	}

}

func printTree(root *treeNode) {
	height := getTreeHeight(root)
	fmt.Printf("height: %v\n", height)
	//包含树枝符号的总高度
	totalHeight := height*2 - 1
	// 最大宽度为3 * 2^(n - 1) + 1，公式如下：
	/**
	   父亲节点占1, 两个孩子空间各占1, 连接线各占1, 每个父子关系共占5, 每个关系之间空1, 最左最右各空1
	  第2行： 5 + 2 （1个父子结构占位+左右两个空格分割）
	  第3行：2 * 5 + (1 + 2) （2个父子结构占位+1个相邻父子结构间空格+左右两个空格分割）
	  第4行：4 * 5 + (3 + 2) （4个父子结构占位+3个相邻父子结构间空格+左右两个空格分割）
	  第5行：8 * 5 + (7 + 2)
	  第n行：5 * 2 ^ (n - 2) + (2 ^ (n - 2) - 1) + 2 = 6 * 2 ^ (n-2) + 1 = 3 * 2 ^ (n - 1) + 1
	*/
	var totalWidth int
	if height == 0 {
		totalWidth = 1
	} else {
		totalWidth = (2 << (height - 2) * 3) + 1
	}
	//初始化数打印组
	printArray := make([][]string, totalHeight)
	for i := range printArray {
		printArray[i] = make([]string, totalWidth)
		for j := range printArray[i] {
			printArray[i][j] = " "
		}
	}

	//填充打印数组
	writeArray(root, 0, totalWidth/2, height, printArray)

	//打印
	for i := range printArray {
		var res string
		for j := range printArray[i] {
			res = res + printArray[i][j]
		}
		fmt.Println(res)
	}
}
