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