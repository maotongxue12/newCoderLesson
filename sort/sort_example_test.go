package sort

import (
	"sort"
	"testing"
)

func comp

func TestSmallSum(t *testing.T) {
	times := 10
	maxLength := 10
	maxValue := 20
	for i := 0; i < times; i++ {
		arrA := generateRandomArr(maxLength, maxValue)
		arrB := make([]int, len(arrA))
		copy(arrB, arrA)
		sort.Ints(arrA)

	}

}

