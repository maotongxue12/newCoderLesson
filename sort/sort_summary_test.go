package sort

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	times := 5
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		bubbleSort(arr2)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestInsertSort(t *testing.T) {
	times := 5
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		insertSort(arr2)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestShellSort(t *testing.T) {
	times := 20
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		shellSort(arr2)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestSelectSort(t *testing.T) {
	times := 5
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		selectSort(arr2)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestQuickSort01(t *testing.T) {
	times := 10
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		quickSort01(arr2, 0, len(arr2)-1)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestQuickSort02(t *testing.T) {
	times := 10
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		quickSort02(arr2, 0, len(arr2)-1)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestMergeSort(t *testing.T) {
	times := 10
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		mergeSort(arr2, 0, len(arr2)-1)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

func TestMaxHeap(t *testing.T) {
	times := 20
	for i := 0; i < times; i++ {
		arr1 := generateRandomArr(10, 20)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		maxHeap(arr2)
		fmt.Println(arr1, arr2)
		Convey("test bubbleSort algorithm", t, func() {
			So(arr1, ShouldResemble, arr2)
		})
	}
}

//生成随机数组
func generateRandomArr(maxlength, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxlength)
	arr := make([]int, length)
	for index := 0; index < length; index++ {
		arr[index] = rand.Intn(maxValue)
	}
	time.Sleep(1 * time.Second)
	return arr
}
