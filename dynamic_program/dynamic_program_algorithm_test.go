package dynamic_program

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
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
	str := "abcb"
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

func TestReverseStack(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(10)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(20)
	}
	fmt.Println(arr)
	reverseStack(&arr)
	fmt.Println(arr)
}

func TestMinPath(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(minPath(matrix))
	matrix2 := [][]int{{1, 3, 5, 9}, {8, 1, 3, 4}, {5, 0, 6, 1}, {8, 8, 4, 0}}
	fmt.Println(minPath(matrix2))
}

func TestMinPathDynamic(t *testing.T) {
	times := 10
	Convey("test minPath function: ", t, func() {
		for i := 0; i < times; i++ {
			time.Sleep(1 * time.Second)
			matrix := generateRandomArr(10, 10, 20)
			res := minPath(matrix)
			resDynmic := minPathDynamic(matrix)
			fmt.Println(matrix)
			fmt.Println(res, resDynmic)
			So(res, ShouldEqual, resDynmic)
		}
	})
}

func TestArrSum(t *testing.T) {
	times := 10
	aim := 20
	arr := make([]int, 6)
	for i := 0; i < times; i++ {
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < len(arr); j++ {
			arr[j] = rand.Intn(20)
		}
		Convey("test arrSum function: ", t, func() {
			res := arrSum(arr, aim)
			resDynamic := arrSumDynamic(arr, aim)
			fmt.Println(res, resDynamic)
			fmt.Println(arr)
			So(res, ShouldEqual, resDynamic)
		})
		time.Sleep(time.Second)
	}
}

func TestKnapsack(t *testing.T) {
	times := 10
	arrA := make([]int, 6)
	arrB := make([]int, 6)
	bag := 100
	for i := 0; i < times; i++ {
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < len(arrA); j++ {
			arrA[j] = rand.Intn(20)
		}
		time.Sleep(500 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < len(arrB); j++ {
			arrB[j] = rand.Intn(20)
		}
		fmt.Println(arrA, arrB)
		Convey("test knapsack function: ", t, func() {
			res := knapsack(arrA, arrB, bag)
			resDynamic := knapsackDynamic(arrA, arrB, bag)
			fmt.Println(res, resDynamic)
			So(res, ShouldEqual, resDynamic)
		})
		time.Sleep(500 * time.Millisecond)
	}
}

func generateRandomArr(maxCow, maxColumn, maxValue int) [][]int {
	rand.Seed(time.Now().UnixNano())
	cow := rand.Intn(maxCow)
	column := rand.Intn(maxColumn)
	matrix := make([][]int, cow)
	for i := 0; i < cow; i++ {
		matrix[i] = make([]int, column)
	}
	for i := 0; i < cow; i++ {
		for j := 0; j < column; j++ {
			matrix[i][j] = rand.Intn(maxValue)
		}
	}
	return matrix
}
