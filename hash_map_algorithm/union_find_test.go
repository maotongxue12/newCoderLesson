package hash_map_algorithm

import (
	"fmt"
	"testing"
)

func TestIsLandNums(t *testing.T) {
	testCase01 := [][]int{
		{0, 0, 1, 0, 1, 0},
		{1, 1, 1, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	testCase02 := [][]int{
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 1, 1, 1, 0, 1, 1, 1, 0 },
		{ 0, 1, 1, 1, 0, 0, 0, 1, 0 },
		{ 0, 1, 1, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 1, 0, 0 },
		{ 0, 0, 0, 0, 1, 1, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
	}
	fmt.Println(isLandNums(testCase01))
	fmt.Println(isLandNums(testCase02))
}
