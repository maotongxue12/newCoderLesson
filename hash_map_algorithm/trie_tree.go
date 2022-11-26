package hash_map_algorithm

type node01 struct {
	pattern string	//仅叶子节点时，pattern有值
	children []*node	//指向的孩子节点
	part string		//当前节点的值
	end	int			//以当前节点字符结尾的字符串个数
}

type node02 struct {
	path int	//记录当前经过当前节点字符串的个数
	end int		//以此叶子节点为字符串结尾的个数
	next []*node02	//当前节点指向下一个节点的
}

func newTrieTreeNode() *node02 {
	return &node02{
		path: 0,
		end: 0,
		next: make([]*node02, 26),
	}
}

func (n *node02)insert(word string) {
	if word == "" {
		return
	}
	node := n
	for _, ch := range word {
		index := ch - 'a'
		if node.next[index] == nil {
			temp := newTrieTreeNode()
			node.next[index] = temp
		}
		node = node.next[index]
		node.path++
	}
	node.end++
}

func (n *node02)search(word string) int {
	if word == "" {
		return 0
	}
	node := n
	for _, ch := range word {
		index := ch - 'a'
		if node.next[index] == nil {
			return 0
		}
		node = node.next[index]
	}
	return node.end
}

func (n *node02)delete(word string) {
	if word == "" {
		return
	}
	node := n
	//删除路径前需要先判断该是否存在该字符串路径
	if 0 == n.search(word) {
		return
	}
	for _, ch := range word {
		index := ch - 'a'
		node.next[index].path--
		if 0 == node.next[index].path {
			node.next[index] = nil
			return
		}
		node = node.next[index]
	}
	node.end--
}

func (n *node02)prefixNums(pre string) int {
	if pre == "" {
		return 0
	}
	node := n
	for _, ch := range pre {
		index := ch - 'a'
		if node.next[index] == nil {
			return 0
		}
		node = node.next[index]
	}
	return node.path
}