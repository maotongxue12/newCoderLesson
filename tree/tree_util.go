package tree

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

func print
