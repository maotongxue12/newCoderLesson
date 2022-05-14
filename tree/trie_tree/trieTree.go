package trie_tree

//前缀树数据结构，前缀树路径记录当前字符,如{abc, ace, bcd}
/******************************
			 O
       a/         \b
       O           O
    b/  \c       c/
    O    O       O
  c/      \e   d/
  O        O   O
 *******************************/
type trieTreeNode struct {
	path int	//记录当前经过当前节点字符串的个数
	end int		//以此叶子节点为字符串结尾的个数
	next []*trieTreeNode	//当前节点指向下一个节点的
}

func newTrieTreeMode() *trieTreeNode {
	return &trieTreeNode{
		path: 0,
		end: 0,
		next: make([]*trieTreeNode,26),	//26个字母，默认初始化26条路径
	}
}

func (t *trieTreeNode) insert(word string) {
	if word == "" {
		return
	}
	node := t
	for _, v := range word {
		index := v - 'a'
		if node.next[index] == nil {
			node.next[index] = newTrieTreeMode()
		}
		node = node.next[index]
		node.path++
	}
	node.end++
}

func (t *trieTreeNode) delete(word string) {
	if t.search(word) != 0 {
		node := t
		for _, val := range word {
			index := val - 'a'
			node.next[index].path--
			if 0 == node.next[index].path {
				node.next[index] = nil
				return
			}
			node = node.next[index]
		}
		node.end--
	}
}

func (t *trieTreeNode) search(word string) int {
	if word == "" {
		return 0
	}
	node := t
	for _, val := range word {
		index := val - 'a'
		if node.next[index] == nil {
			return 0
		}
		node = node.next[index]
	}
	return node.end
}

func (t *trieTreeNode) prefixNumber(pre string) int {
	if "" == pre {
		return 0
	}
	node := t
	for _, val := range pre {
		index := val - 'a'
		if nil == node.next[index] {
			return 0
		}
		node = node.next[index]
	}
	return node.path
}