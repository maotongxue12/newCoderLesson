package hash_map_algorithm

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	root := newTrieTreeNode()
	fmt.Println(root.search("zuo"))
	root.insert("zuo")
	fmt.Println(root.search("zuo"))
	root.insert("zuo")
	root.insert("zuo")
	root.insert("zuo")
	fmt.Println(root.search("zuo"))
	root.delete("zuo")
	fmt.Println(root.search("zuo"))
	root.insert("zuoa")
	root.insert("zuoab")
	root.insert("zuoac")
	root.delete("zuoa")
	fmt.Println(root.search("zuoa"))
	fmt.Println(root.prefixNums("zuo"))
}
