package hash_map_algorithm

func isLandNums(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	nums := 0
	rows := len(matrix)
	columns := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == 1 {
				processIsLand(matrix, i, j, rows, columns)
				nums++
			}
		}
	}
	return nums
}

func processIsLand(matrix [][]int, i, j, rows, columns int) {
	if i < 0 || i >= rows || j < 0 || j >= columns || matrix[i][j] != 1  {
		return
	}
	matrix[i][j] = 2
	processIsLand(matrix, i-1, j, rows, columns)
	processIsLand(matrix, i+1, j, rows, columns)
	processIsLand(matrix, i, j-1, rows, columns)
	processIsLand(matrix, i, j+1, rows, columns)
}

//union_find并查集结构可以理解为指向父节点的树，其中子节点都指向父节点，最后根结点指向自己
//主要包括两个函数：（1）is_same_set(_, _) 判断两个元素是否在一个集合里
//（2）union(_, _) 将元素1所在的集合和元素2所在集合合并为一个集合，虽然该函数是合并两个集合，
//但是函数中的参数为两个集合中的元素
//node 可以任意定义节点元素类型
type node struct{
	element int
}

//为了表示维护节点之间的关系，初始化两个map结构：fatherMap维护节点所对应的父节点
//sizeMap表示该节点所在集合的个数，同时，仅该节点为代表节点时，该节点对应value值才有意义
var fatherMap = make(map[node]node, 0)
var sizeMap = make(map[node]int, 0)

//初始化并查集
func initialUnionSet(sets []node) {
	fatherMap = make(map[node]node)
	sizeMap = make(map[node]int)
	for _, val := range sets {
		fatherMap[val] = val
		sizeMap[val] = 1
	}
}

//获取当前节点的代表节点，同时，在获取代表节点时会做一个优化，对于链型的集合进行扁平化处理
//将该节点的所有父节点都指向代表节点
func findFather(n node) node {
	fatherNode := fatherMap[n]
	curNode := n
	stack := make([]node, 0)
	for curNode != fatherNode {
		stack = append(stack, curNode)
		curNode = fatherNode
		fatherNode = fatherMap[fatherNode]
	}

	for len(stack) != 0 {
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fatherMap[val] = fatherNode
	}
	return fatherNode
}

//判断两个元素是否处于同一个集合中: 首先分别找到两个节点所属的根节点，判断是否属于同一个集合
func isSameSet(a, b node) bool {
	return findFather(a) == findFather(b)
}


//合并两个元素所在的集合
func union(a, b node) {
	fatherNodeA := findFather(a)
	fatherNodeB := findFather(b)
	if fatherNodeA != fatherNodeB {
		aSize := sizeMap[fatherNodeA]
		bSize := sizeMap[fatherNodeB]
		if aSize <= bSize {
			fatherMap[fatherNodeA] = fatherNodeB
			sizeMap[fatherNodeB] = aSize + bSize
		} else {
			fatherMap[fatherNodeB] = fatherNodeA
			sizeMap[fatherNodeA] = aSize + bSize
		}
	}
}