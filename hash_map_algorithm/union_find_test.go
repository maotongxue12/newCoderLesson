package hash_map_algorithm

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
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

func TestIsSameSet(t *testing.T) {
	convey.Convey("test the union_find_test function", t, func() {
		sets := []node{{4}, {2}, {3}, {11}, {7}, {9}}
		initialUnionSet(sets)
		res1 := isSameSet(sets[0], sets[4])
		convey.So(res1, convey.ShouldBeFalse)
		union(sets[0], sets[4])

		res2 := isSameSet(sets[0], sets[4])
		convey.So(res2, convey.ShouldBeTrue)
	})
}
