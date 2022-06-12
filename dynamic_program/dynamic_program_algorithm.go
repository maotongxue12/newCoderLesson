package dynamic_program

import "fmt"

/*
题目一：求n!的结果
 */
//方法一：暴力递归
func getFibonacci01(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * getFibonacci01(n-1)
}

//方法二：动态规划
func getFibonacci02(n int) int {
	res := 1
	if n == 1 || n == 0 {
		return res
	}
	//由暴力递归可观察，第n个数值的结果仅依赖于第(n-1)的结果，即需要用res记录(n-1)个数值的结果
	for i := 1; i <=n; i++ {
		res *= i
	}
	return res
}

//题目二：汉诺塔：打印n层汉诺塔从最左边移动到最右边的全部过程
func move(disk int, N string, M string) {
	fmt.Println("把 ", disk, " 号圆盘从 ", N, " 移到 ", M)
}

//把n号盘从A移到C，以B为辅助
func hanoi(n int, A string, B string, C string) {
	if n == 1 {
		move(1, A, C)
	} else {
		hanoi(n-1, A, C, B)
		move(n, A, C)
		hanoi(n-1, B, A, C)
	}
}


//题目三：打印一个字符串的全部子序列，包括空字符串
//思路：如abc序列，从头遍历，每个索引都包含两种情况，一种是包含该字母，另一种是不包含该字母，如索引为0时，分为包含a和不包含a两种情况，索引为1时，分为包含b和不包含b两种情况，
//索引为2时，包含包含和不包含c两种情况，当索引为3时，超出索引序列，退出该递归，打印上一级索引序列
var ans []string
func printAllSub(str string, i int, res string) {
	if i == len(str) {
		ans = append(ans, res)
		fmt.Println(res)
		return
	}
	//两种情况，包含该索引的字符和不包含该索引的字符
	printAllSub(str, i+1, res + string(str[i]))
	printAllSub(str, i+1, res)
}

//题目四：打印一个字符串的全部排列
//思路：递归算法，对待全排序的序列中从左选定一个数作为基数，然后对他右边的数进行全排列，以此递归。
//如abcd，可以理解首先固定a，求bcd的全排列，要求bcd的全排列，进行固定b，求cd的全排列，同样，求cd的全排列，需要固定c，求d的全排列，此时，d的全排列就是他本身一种，以此递归
//递归的退出条件就是剩最后一个元素，全排列只有它本身。通俗理解就是首先固定第一个元素，求后面元素的全排列，然后递归固定第二个元素，求后面元素的全排列，直至剩余一个元素，返回当前元素
//的全排列--该唯一元素本身
//在上述全排列递归算法中，要保证遍历所有排列方式，当固定第一个元素时，需要分为四种情况，以a为首元素和分别以b，c，d为首元素，即a分别要与后面的三个元素做交换，并以交换后的第一个元素
//作为基准，向后递归
func printAllPermutations(str string) {
	byteStr := []byte(str)
	processPermutations(byteStr, 0, len(str))
}

func processPermutations(str []byte, from, to int) {
	if from == to {
		fmt.Println(string(str))
	} else {
		for i := from; i < to; i++ {
			str[from], str[i] = str[i], str[from]
			processPermutations(str, from+1, to)
			//当首元素a和第二个元素b交换，递归对后面元素进行全排列之后，要让a元素和第三个元素交换，
			//但此时第一个元素已经换成了b，所以需要交换回来，保证和后面元素交换的还是首元素a
			//str[i], str[from] = str[from], str[i]
		}
	}
}