package sort

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

//题目八：小和问题 在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组的小和。
//[1,3,4,2,5] 1左边比1小的数，没有； 3左边比3小的数，1； 4左边比4小的数，1、3； 2左边比2小的数，1；
//5左边比5小的数，1、3、4、2； 所以小和为1+1+3+1+1+3+4+2=16
//解题思路：采用归并排序，归并排序特点：当前存在的两个序列已经内部有序，例如A[3,5,8]和B[1,3,5]两个序列归并，
//如果计算A组中比B组中5小的数的个数，只需要比较找到A组第一个比5大的数字，则根据数组长度就可以推算A组有几个比5小的数字
//且计算之后合并为[1,3,3,5,5,8]内部有序，执行下次归并已经不需要处理当前数组内部问题
func smallSum(arr []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)>>1
	leftRes := smallSum(arr, left, mid)
	rightRes := smallSum(arr, mid+1, right)
	return processSmallSum(arr, left, mid, right) + leftRes + rightRes
}

func processSmallSum(arr []int, left, mid, right int) int {
	pLeft := left
	pRight := mid + 1
	var subRes int
	temp := make([]int, 0)
	for pLeft <= mid && pRight <= right {
		if arr[pLeft] < arr[pRight] {
			subRes = subRes + arr[pLeft]*(right-pRight+1)
			temp = append(temp, arr[pLeft])
			pLeft++
		} else {
			temp = append(temp, arr[pRight])
			pRight++
		}
	}
	for pLeft <= mid {
		temp = append(temp, arr[pLeft])
		pLeft++
	}
	for pRight <= right {
		temp = append(temp, arr[pRight])
		pRight++
	}
	for i := 0; i < len(temp); i++ {
		arr[left+i] = temp[i]
	}
	return subRes
}

//题目九： 给定一个数组，求如果排序之后，相邻两数的最大差值，要求时 间复杂度O(N)，且要求不能用非基于比较的排序。
//解题思路：采用桶排序思想，共n个数，设置n+1个桶，则能保证必有一个桶为空，当存在一个空桶时，则空桶左边的最大值和空桶右边的最小值
//之间的差值必然大于桶内之间的差值，但同时，该差值不一定是所有相邻两数的最大差值；空桶保证了在比较差值时可以忽略桶内数值的比较，只
//需要比较相邻桶之间后一个桶的最小值和前一个桶的最小值之间的差值
func maxGap(arr []int) int {
	var res int
	//首先将arr数组的数放入到对应的桶中, 其中只需要放入桶中最大值和最小边界值，所以桶中的中间数值可以舍去
	//即桶中需存放最大值，最小值,其中isExist记录桶中是否有值
	length := len(arr) + 1
	maxNums := make([]int, length)
	minNums := make([]int, length)
	isExist := make([]bool, length)
	maxValue := INT_MIN
	minValue := INT_MAX
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	if maxValue == minValue {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		bucketNum := selectBucket(arr[i], minValue, maxValue, length-1)
		if isExist[bucketNum] && maxNums[bucketNum] >= arr[i] {
			maxNums[bucketNum] = maxNums[bucketNum]
		} else {
			maxNums[bucketNum] = arr[i]
		}
		if isExist[bucketNum] && minNums[bucketNum] <= arr[i] {
			minNums[bucketNum] = minNums[bucketNum]
		} else {
			minNums[bucketNum] = arr[i]
		}
		isExist[bucketNum] = true
	}

	//比较前一个桶的最大值和后一个桶的最小值的差，并获取差的最大值
	lastMax := maxNums[0]
	for i := 1; i < length; i++ {
		if isExist[i] {
			maxGap := minNums[i] - lastMax
			if res < maxGap {
				res = maxGap
			}
			lastMax = maxNums[i]
		}
	}
	return res
}

func selectBucket(value, minValue, maxValue, buckets int) int {
	bucketNum := (value - minValue) * buckets / (maxValue - minValue)
	return bucketNum
}
