package sortalgorithm

//冒泡排序
func bubbleSort(arr []int) {
	//比较趟数
	for i := 0; i < len(arr)-1; i++ {
		//每趟的比较次数
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
}

//希尔排序
func shellSort(arr []int) {
	for gap := len(arr) / 2; gap >= 1; gap /= 2 {
		for i := gap; i < len(arr); i += gap {
			temp := arr[i]
			j := i - gap
			for ; j >= 0; j -= gap {
				if arr[j] > temp {
					arr[j+gap] = arr[j]
				} else {
					break
				}
			}
			arr[j+gap] = temp
		}
	}
}

//选择排序
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		if index != i {
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
}

//快速排序: 填坑法
func quickSort01(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := process01(arr, left, right)
	quickSort01(arr, left, index-1)
	quickSort01(arr, index+1, right)
}

func process01(arr []int, left, right int) int {
	standardVal := arr[left]
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= standardVal {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= standardVal {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = standardVal
	return i
}

//快速排序：左右区间法
func quickSort02(arr []int, left, right int) {
	if left >= right {
		return
	}
	minArea, maxArea := process02(arr, left, right)
	quickSort02(arr, left, minArea)
	//maxArea区间包含了等于temp的值，所以递归时应该去掉这个值，如果包含了唯一的temp值，会进入无限递归中
	quickSort02(arr, maxArea+1, right)
}

//当前值小于temp值时，将当前值和左区的下一个值做交换，left加1，继续比较下一个值，这样和temp相同的值可以保证在一起
//当前值等于temp值时，left+1，不做交换，左区不包含和temp相等的值，右区包含一个和temp相等的值
//当前值大于temp值时，将当前值和右区的值做交换，left不变，继续比较交换后的值。
func process02(arr []int, left, right int) (int, int) {
	temp := arr[left]
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
func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)>>2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	mergeProcess(arr, left, mid, right)
}

func mergeProcess(arr []int, left, mid, right int) {
	var temp []int
	i := left
	j := mid + 1
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}
	for j <= right {
		temp = append(temp, arr[j])
		j++
	}
	for k := 0; k < len(temp); k++ {
		arr[left+k] = temp[k]
	}
}

//堆排序--大根堆
func maxHeap(arr []int) {
	size := len(arr)
	//构建大根堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	for size > 0 {
		size--
		arr[0], arr[size] = arr[size], arr[0]
		heapify(arr, 0, size)
	}
}

//构建大根堆
func heapInsert(arr []int, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if arr[index] > arr[parent] {
			arr[index], arr[parent] = arr[parent], arr[index]
		}
		index = parent
	}
}

func heapify(arr []int, index, size int) {
	left := index*2 + 1
	largest := index
	for left < size {
		right := left + 1
		if right < size && arr[right] > arr[left] {
			largest = right
		} else {
			largest = left
		}
		if arr[index] > arr[largest] {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = index*2 + 1
	}
}
