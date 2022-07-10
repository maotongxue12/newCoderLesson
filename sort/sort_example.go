package sort

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
