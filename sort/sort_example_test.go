package sort

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func comparator(arr []int) int {
	var res int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				res += arr[i]
			}
		}
	}
	return res
}

func TestSmallSum(t *testing.T) {
	times := 20
	maxLength := 10
	maxValue := 20
	Convey("teset smallSum functions ", t, func() {
		for i := 0; i < times; i++ {
			arrA := generateRandomArr(maxLength, maxValue)
			arrB := make([]int, len(arrA))
			copy(arrB, arrA)
			fmt.Println(arrA, arrB)
			resA := comparator(arrA)
			resB := smallSum(arrB, 0, len(arrB)-1)
			fmt.Println(resA, resB)
			So(resA, ShouldEqual, resB)
		}
	})
}

func comparatorMaxGap(arr []int) int {
	sort.Ints(arr)
	var maxGap int
	for i := 0; i < len(arr)-1; i++ {
		tmp := arr[i+1]-arr[i]
		if maxGap < tmp {
			maxGap = tmp
		}
	}
	return maxGap
}

func TestMaxGap(t *testing.T) {
	times := 30
	maxLength := 10
	maxValue := 20
	Convey("teset smallSum functions ", t, func() {
		for i := 0; i < times; i++ {
			arrA := generateRandomArr(maxLength, maxValue)
			arrB := make([]int, len(arrA))
			copy(arrB, arrA)
			fmt.Println(arrA, arrB)
			resA := comparatorMaxGap(arrA)
			resB := maxGap(arrB)
			fmt.Println(resA, resB)
			So(resA, ShouldEqual, resB)
		}
	})
}
