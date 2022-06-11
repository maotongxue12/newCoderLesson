package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func generateRandomArr(maxLength, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength)
	arr := make([]int, length)
	for i:=0; i < length; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

func comparator(arr []int) int {
	var res int = 0
	sort.Ints(arr)
	for i:=1; i < len(arr); i++ {
		gapValue := arr[i] - arr[i-1]
		if res < gapValue {
			res = gapValue
		}
	}
	return res
}

func maxGap(arr []int) int {
	bucketNums := len(arr) + 1
	var res int
	maxNums := make([]int, bucketNums)
	minNums := make([]int, bucketNums)
	existNums := make([]bool, bucketNums)
	minValue := INT_MAX
	maxValue := INT_MIN
	for i:=0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	if minValue == maxValue {
		return 0
	}
	for i := 0; i < bucketNums-1; i++ {
		bid := bucketBid(arr[i], bucketNums-1, minValue, maxValue)
		if existNums[bid] && maxNums[bid] > arr[i] {
			maxNums[bid] = maxNums[bid]
		} else {
			maxNums[bid] = arr[i]
		}
		if existNums[bid] && minNums[bid] < arr[i] {
			minNums[bid] = minNums[bid]
		} else {
			minNums[bid] = arr[i]
		}
		existNums[bid] = true
	}

	lastMax := maxNums[0]
	for i := 1; i < bucketNums; i++ {
		if existNums[i] {
			maxValue := minNums[i] - lastMax
			lastMax = maxNums[i]
			if res < maxValue {
				res = maxValue
			}
		}
	}
	return res
}

func bucketBid(num, len, min, max int) int {
	bid := int((num - min) * len / (max - min))
	return bid
}

func main() {
	times := 50
	maxLength := 10
	maxValue := 100
	success := true
	for i := 0; i < times; i++ {
		time.Sleep(time.Second*1)
		arr1 := generateRandomArr(maxLength, maxValue)
		//fmt.Println(arr1)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		testRes := maxGap(arr1)
		fmt.Println(testRes)
		trueRes := comparator(arr2)
		fmt.Println(trueRes)
		if trueRes != testRes {
			success = false
		}
	}
	if success == false {
		fmt.Println("The maxGap algorithm is error")
	} else {
		fmt.Println("The maxGap algorithm is correct")
	}
}