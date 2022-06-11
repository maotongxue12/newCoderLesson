package main

import (
	"fmt"
	"math/rand"
	"time"
)

//小和问题
func smallSum(arr []int, left, right int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}
	if left == right {
		return 0
	}
	mid := left + (right-left) >> 1
	return smallSum(arr, left, mid) + smallSum(arr, mid+1, right) + smallSumPartition(arr, left, mid, right)
}
func smallSumPartition(arr []int, left, mid, right int) int{
	pLeft := left
	pRight := mid+1
	var sum int
	i := 0
	temp := make([]int, right-left+1)
	for pLeft <= mid && pRight <= right {
		if arr[pLeft] < arr[pRight] {
			sum = sum + arr[pLeft] * (right-pRight+1)
			temp[i] = arr[pLeft]
			pLeft++
			i++
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
	return sum
}

//生成随机数组
func generateRandomArr(maxLength, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}
func comparator(arr []int) int {
	var res int = 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				res = res + arr[i]
			}
		}
	}
	return res
}

func main() {
	times := 20
	maxLength := 10
	maxValue := 100
	success := true
	for i := 0; i < times; i++ {
		time.Sleep(time.Second*1)
		arr1 := generateRandomArr(maxLength, maxValue)
		fmt.Println(arr1)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		testRes := smallSum(arr1, 0, len(arr1)-1)
		fmt.Println(testRes)
		trueRes := comparator(arr2)
		fmt.Println(trueRes)
		if trueRes != testRes {
			success = false
		}
	}
	if success == false {
		fmt.Println("The smallSum algorithm is error")
	} else {
		fmt.Println("The smallSum algorithm is correct")
	}
}