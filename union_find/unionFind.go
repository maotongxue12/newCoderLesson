package union_find

//一个矩阵中只有0和1两种值，每个位置都可以和自己的上、下、左、右
//四个位置相连，如果有一片1连在一起，这个部分叫做一个岛，求一个
//矩阵中有多少个岛
//思路：利用递归方法，循环遍历该矩阵，然后对当前遍历的值执行递归操作，寻找上下左右可连接的1，直至
//不能找到1，并将遍历过的值修改为2，防止下个元素进行重复遍历
func islandNums(matrix [][]int) (res int) {
	line := len(matrix)
	conlumn := len(matrix[0])
	for i := 0; i < line; i++ {
		for j := 0; j < conlumn; j++ {
			if matrix[i][j] == 1 {
				//确定当前点岛的范围
				processIsland(matrix, i, j, line, conlumn)
				res++
			}
		}
	}
	return res
}

func processIsland(matrix [][]int, i, j, line, column int){
	if i < 0 || i >= line || j < 0 || j >= column || matrix[i][j] != 1 {
		return
	}
	matrix[i][j] = 2
	processIsland(matrix, i-1, j, line, column)
	processIsland(matrix, i+1, j, line, column)
	processIsland(matrix, i, j-1, line, column)
	processIsland(matrix, i, j+1, line, column)
}

//实现并查集结构，并查集主要包括两个函数（1）is_same_set(_, _) 判断两个元素是否在一个集合里
//（2）union(_, _) 将元素1所在的集合和元素2所在集合合并为一个集合
type node struct {

}
//初始化两个集合，fatherMap集合存储当前节点所对应的根节点
//sizeMap集合存储当前节点所在集合的元素个数，仅该节点是集合的代表节点时，该节点对应value值才有意义
//初始状态，每个节点都分别为一个集合
var fatherMap = make(map[node]node)
var sizeMap = make(map[node]int)
func initUnionFindSet(list []node) {
	for i := 0; i < len(list); i++{
		fatherMap[list[i]] = list[i]
		sizeMap[list[i]] = 1
	}
}

//找到当前节点的代表节点，如果该节点的父节点不是当前集合的代表节点，则将该集合进行扁平化优化，将该节点调整为代表节点的
//子节点
func findHead(Node node) node {
	var stack []node
	curNode := Node
	father := fatherMap[Node]
	//扁平化处理节点，代表节点的父节点指向节点本身
	for curNode != father {
		stack = append(stack, curNode)
		curNode = father
		father = fatherMap[curNode]
	}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fatherMap[cur] = father
	}
	return father
}

//判断两个节点是否在一个集合中，即找到两个节点的表示节点并判断是否是同一个
func isSameSet(a, b node) bool {
	return findHead(a) == findHead(b)
}

//合并两个a, b两个节点为一个集合
func union(a, b node) {
	aHead := findHead(a)
	bHead := findHead(b)
	if aHead != bHead {
		aSize := sizeMap[aHead]
		bSize := sizeMap[bHead]
		//将小的集合进大的集合
		if aSize <= bSize {
			fatherMap[aHead] = bHead
			sizeMap[bHead] = aSize + bSize
		} else {
			fatherMap[bHead] = aHead
			sizeMap[aHead] = aSize + bSize
		}
	}
}