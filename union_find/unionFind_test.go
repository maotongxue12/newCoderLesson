package union_find

import (
	"fmt"
	"testing"
)

func TestIslandNums(t *testing.T) {
	testCase01 := [][]int{
		{0, 0, 1, 0, 1, 0},
		{1, 1, 1, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	testCase02 := [][]int{{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
	{ 0, 1, 1, 1, 0, 1, 1, 1, 0 },
	{ 0, 1, 1, 1, 0, 0, 0, 1, 0 },
	{ 0, 1, 1, 0, 0, 0, 0, 0, 0 },
	{ 0, 0, 0, 0, 0, 0, 1, 0, 0 },
	{ 0, 0, 0, 0, 1, 1, 0, 0, 0 },
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0 }}
	fmt.Println(islandNums(testCase01))
	fmt.Println(islandNums(testCase02))
}