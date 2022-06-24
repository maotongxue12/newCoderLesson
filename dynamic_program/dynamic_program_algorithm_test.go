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
	fmt.Println(cowNumber(1))
	fmt.Println(cowNumber(2))
	fmt.Println(cowNumber(3))
	fmt.Println(cowNumber(4))
	fmt.Println(cowNumber(5))
	fmt.Println(cowNumber(9))
	fmt.Println(cowNumber(10))
	fmt.Println(cowNumber(11))
	fmt.Println(cowNumber(12))

	fmt.Println(cowNumber2(6))
	fmt.Println(cowNumber2(11))
	fmt.Println(cowNumber2(12))
}

func TestProcessMinPath(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(minPath(matrix))
}