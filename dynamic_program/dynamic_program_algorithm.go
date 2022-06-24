package dynamic_program

import (
	"fmt"
	"math"
)

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
			fmt.Println("递归前str字符串：", string(str))
			processPermutations(str, from+1, to)
			fmt.Println("递归后str字符串：", string(str))
			//当首元素a和第二个元素b交换，递归对后面元素进行全排列之后，要让a元素和第三个元素交换，
			//但此时第一个元素已经换成了b，所以需要交换回来，保证和后面元素交换的还是首元素a
			str[i], str[from] = str[from], str[i]
		}
	}
}

//题目五：打印一个字符串的全部排列，要求不要出现重复的排列
//思路：去重的全排列就是从第一个数字起，每个数分别与它后面非重复出现的数字交换。
func printPermutationsNoRepeat(str string) {
	byteStr := []byte(str)
	processPermutationsNoRepeat(byteStr, 0, len(byteStr))
}

//在[i, j)区间是否有字符与下标为j的字符相等
func isSwap(str []byte, i, j int) bool {
	for index := i; index < j; index++ {
		if str[index] == str[j] {
			return false
		}
	}
	return true
}

func processPermutationsNoRepeat(b []byte, from, to int) {
	if from == to {
		fmt.Println(string(b))
	} else {
		for i := from ; i < to; i++ {
			if isSwap(b, from, i) {
				b[i], b[from] = b[from], b[i]
				processPermutationsNoRepeat(b, from+1, to)
				b[i], b[from] = b[from], b[i]
			}
		}
	}
}

//题目五：母牛每年生一只母牛，新出生的母牛成长三年后也能每年生一只母牛，假设不会死。求N年后，母牛的数量。
//暴力递归：总结可发现递归方程式为f(n)=f(n-1)+f(n-3)，即前一年的牛加上(n-3)年的牛的数量（n-3的牛经过三年都可以生小牛，f(n-3)即为当前可以生小牛的母牛数量==新出生小牛数量）
func cowNumber(year int) int {
	if year < 1 {
		return 0
	}
	if year == 1 || year == 2 || year == 3 {
		return year
	}
	return cowNumber(year-1) + cowNumber(year-3)
}

//进阶：如果每只母牛只能活10年，求N年后，母牛的数量。
func cowNumber2(year int) int {
	if year < 1 {
		return 0
	}
	if year == 1 || year == 2 || year == 3 {
		return year
	}
	if year > 3 && year <= 10 {
		return cowNumber2(year-1) + cowNumber2(year-3)
	} else {
		return cowNumber2(year-1) + cowNumber2(year-3) - cowNumber2(year-10)
	}
}

//题目六：

//题目七：给你一个二维数组，二维数组中的每个数都是正数，要求从左上角走到右下角，每一步只能向右或者向下。沿途经过的数字要累
//加起来。返回最小的路径和。
//暴力递归：除了最后一列和最后一行外，其余的点都有向右和向下两种走法，则应该取两种路径的最小走法，最后一列只能向下走，最后一行只能向右走
//baseCase为走到右下角
func minPath(matrix [][]int) int {
	return processMinPath(matrix, 0, 0)
}

func processMinPath(matrix [][]int, cow, column int) int {
	res := matrix[cow][column]
	if cow == len(matrix)-1 && column == len(matrix[0])-1 {
		return res
	} else if cow == len(matrix)-1 {
		return res + processMinPath(matrix, cow, column+1)
	} else if column == len(matrix[0])-1 {
		return res + processMinPath(matrix, cow + 1, column)
	} else {
		return res + int(math.Min(float64(processMinPath(matrix, cow+1, column)), float64(processMinPath(matrix, cow, column+1))))
	}
}

