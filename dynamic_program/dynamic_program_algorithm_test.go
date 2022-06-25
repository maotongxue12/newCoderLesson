package dynamic_program

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetFibonacci01(t *testing.T) {
	res01 := getFibonacci01(10)
	res02 := getFibonacci02(10)
	Convey("test fibonacci nums: ", t, func() {
		So(res01, ShouldEqual, res02)
	})
}

func TestHanoi(t *testing.T) {
	hanoi(3, "A", "B", "C")
	fmt.Println()
	hanoi(5, "A", "B", "C")
}

func TestPrintAllSub(t *testing.T) {
	printAllSub("abc", 0, "")
	fmt.Println(ans)
	fmt.Println(len(ans))
}

func TestPrintAllPermutations(t *testing.T) {
	str := "abcd"
	printAllPermutations(str)
	fmt.Println()
	str = "abc"
	printAllPermutations(str)
}

func TestPermutationsNoRepeat(t *testing.T) {
	str := "abb"
	printPermutationsNoRepeat(str)
}

func TestCowNumber(t *testing.T) {
	Convey("test the cow number:", t, func() {
		So(cowNumber(1), ShouldEqual, cowNumberDynamic(1))
		So(cowNumber(3), ShouldEqual, cowNumberDynamic(3))
		So(cowNumber(4), ShouldEqual, cowNumberDynamic(4))
		So(cowNumber(9), ShouldEqual, cowNumberDynamic(9))
		So(cowNumber(12), ShouldEqual, cowNumberDynamic(12))
	})

	fmt.Println()
	fmt.Println(cowNumber2(6))
	fmt.Println(cowNumber2(11))
	fmt.Println(cowNumber2(12))
}

func TestMinPath(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(minPath(matrix))
	matrix2 := [][]int{{ 1, 3, 5, 9 }, { 8, 1, 3, 4 }, { 5, 0, 6, 1 }, { 8, 8, 4, 0 }}
	fmt.Println(minPath(matrix2))
}

func TestMinPathDynamic(t *testing.T) {
	matrix2 := [][]int{{ 1, 3, 5, 9 }, { 8, 1, 3, 4 }, { 5, 0, 6, 1 }, { 8, 8, 4, 0 }}
	fmt.Println(minPath(matrix2))
	Convey("test minPath function: ", t, func() {
		So(minPath(matrix2), ShouldEqual, minPathDynamic(matrix2))
	})
}