package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

//生成随机数组
func generateRandom(maxlength, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxlength)
	arr := make([]int, length)
	for index := 0; index < length; index++ {
		arr[index] = rand.Intn(maxValue)
	}
	return arr
}

//冒泡排序
func bubbleSort(value []int) {
	for i := 0; i < len(value)-1; i++ {
		for j := 0; j < len(value)-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
			}
		}
	}
}

//插入排序
func insertSort(value []int) {
	if len(value) == 0 || len(value) == 1 {
		return
	}
	for i := 1; i < len(value); i++ {
		temp := value[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value[j] > temp {
				value[j+1] = value[j]
			} else {
				break
			}
		}
		value[j+1] = temp
	}
}

//选择排序
func selectSort(value []int) {
	if len(value) == 0 || len(value) == 1 {
		return
	}
	for i := 0; i < len(value)-1; i++ {
		index := i
		for j := i + 1; j < len(value); j++ {
			if value[j] < value[index] {
				index = j
			}
		}
		if index != i {
			value[i], value[index] = value[index], value[i]
		}
	}
}

//快速排序
//01 填坑法
func quickSort01(value []int, left int, right int) {
	if len(value) == 0 || len(value) == 1 {
		return
	}
	if left < right {
		index := process(value, left, right)
		quickSort01(value, left, index-1)
		quickSort01(value, index+1, right)
	}
}

func process(value []int, left int, right int) int {
	temp := value[left]
	i, j := left, right
	for i < j {
		for i < j && value[j] >= temp {
			j--
		}
		if value[j] < temp {
			value[i] = value[j]
			i++
		}
		for i < j && value[i] <= temp {
			i++
		}
		if value[i] > temp {
			value[j] = value[i]
			j--
		}
	}
	value[i] = temp
	return i
}

//快速排序
//02 左右区间法
func quickSort02(value []int, left, right int) {
	if len(value) == 0 || len(value) == 1 {
		return
	}
	if left < right {
		lessArea, moreArea := process02(value, left, right)
		quickSort02(value, left, lessArea)
		quickSort02(value, moreArea+1, right)
	}
}

//当前值小于temp值时，将当前值和左区的下一个值做交换，left加1，继续比较下一个值
//当前值等于temp值时，left+1，不做交换
//当前值大于temp值时，将当前值和右区的值做交换，left不变，继续比较交换后的值。
func process02(arr []int, left int, right int) (int, int) {
	temp := arr[right]
	minArea, maxArea := left-1, right
	for left <= maxArea {
		if arr[left] < temp {
			minArea++
			arr[minArea], arr[left] = arr[left], arr[minArea]
			left++
		} else if arr[left] == temp {
			left++
		} else {
			arr[left], arr[maxArea] = arr[maxArea], arr[left]
			maxArea--
		}
	}
	return minArea, maxArea
}

//归并排序
func mergeSort(value []int, left, right int) {
	if len(value) == 0 || len(value) == 1 {
		return
	}
	if left == right {
		return
	}

	mid := left + (right-left)>>1
	mergeSort(value, left, mid)
	mergeSort(value, mid+1, right)
	merge(value, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	pLeft := left
	pRight := mid + 1
	i := 0
	temp := make([]int, right-left+1)
	for pLeft <= mid && pRight <= right {
		if arr[pLeft] <= arr[pRight] {
			temp[i] = arr[pLeft]
			i++
			pLeft++
		} else {
			temp[i] = arr[pRight]
			i++
			pRight++
		}
	}
	for pLeft <= mid {
		temp[i] = arr[pLeft]
		i++
		pLeft++
	}
	for pRight <= right {
		temp[i] = arr[pRight]
		pRight++
		i++
	}
	for j := 0; j < len(temp); j++ {
		arr[left+j] = temp[j]
	}
}

//希尔排序
func shellSort(value []int) {
	for gapValue := len(value) / 2; gapValue > 0; gapValue = gapValue / 2 {
		for i := gapValue; i < len(value); i += gapValue {
			temp := value[i]
			j := i - gapValue
			for ; j >= 0 && value[j] > temp; j -= gapValue {
				value[j+gapValue] = value[j]
			}
			value[j+gapValue] = temp
		}
	}
}

//堆排序
func heapInsert(arr []int, index int) {
	for arr[(index-1)/2] < arr[index] {
		arr[(index-1)/2], arr[index] = arr[index], arr[(index-1)/2]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index int, size int) {
	left := index*2 + 1
	var largest int
	for left <= size {
		if left+1 <= size && arr[left] < arr[left+1] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[index] > arr[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = index*2 + 1
	}
}

func heapSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	size := len(arr) - 1
	for size > 0 {
		arr[0], arr[size] = arr[size], arr[0]
		size--
		heapify(arr, 0, size)
	}
}

func main() {
	maxLength := 50
	maxValue := 100
	success := true
	for i := 0; i < 50; i++ {
		time.Sleep(time.Second * 1)
		arr1 := generateRandom(maxLength, maxValue)
		arr2 := make([]int, len(arr1))
		arr3 := make([]int, len(arr1))
		copy(arr2, arr1)
		copy(arr3, arr1)
		sort.Ints(arr1)
		//selectSort(arr2)
		//mergeSort(arr3, 0, len(arr3)-1)
		//shellSort(arr3)
		heapSort(arr3)
		flag := reflect.DeepEqual(arr1, arr3)
		if flag == false {
			success = false
		}
		fmt.Println(arr3)
	}
	if success == false {
		fmt.Println("The sort algorithm is error")
	} else {
		fmt.Println("The sort algorithm is correct")
	}
}
