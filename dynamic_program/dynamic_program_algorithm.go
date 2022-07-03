package dynamic_program

import (
	"fmt"
	"math"
)

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
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
	for i := 1; i <= n; i++ {
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
	printAllSub(str, i+1, res+string(str[i]))
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
		for i := from; i < to; i++ {
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

//动态规划思想：动态规划实质是对暴力递归试法的一种改善，用一个表保留递归过程中计算值，避免递归造成的重复计算
//暴力递归->优化为动态规划，需要为无后效性问题，即当前状态的返回值和之前的轨迹无关，之和问题本身有关
//由暴力递归试法可知，第n年的牛数依赖于n-1年和n-3年，base case为前三年的母牛数
func cowNumberDynamic(year int) int {
	if year == 1 || year == 2 || year == 3 {
		return year
	}
	res := 3
	pre := 2
	prePre := 1
	//可以理解为建立一个deep表，维护当前年份的前三年的母牛数目即可
	for i := 4; i <= year; i++ {
		tmp1 := res
		tmp2 := pre
		//当前年的母牛等于前面一年的母牛数（n-1年）+前面3年的母牛数（n-3年）
		res = res + prePre
		pre = tmp1
		prePre = tmp2
	}
	return res
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

//题目六：给你一个栈，请你逆序这个栈，不能申请额外的数据结构，只能使用递归函数。如何实现？
//暴力递归思路：首先实现一个函数，能够将栈传进这个函数以后，返回栈底的元素，并且栈底上面的元素依次盖下来，即逆序取栈的元素
func removeAndReturnBottomElement(stack *[]int) int {
	res := (*stack)[len(*stack)-1]
	*stack = (*stack)[0 : len(*stack)-1]
	if len(*stack) == 0 {
		return res
	} else {
		//返回栈底元素
		last := removeAndReturnBottomElement(stack)
		//利用递归中栈会保存过程中的变量
		*stack = append(*stack, res)
		return last
	}
}

func reverseStack(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	val := removeAndReturnBottomElement(stack)
	reverseStack(stack)
	*stack = append(*stack, val)
}

//题目七：给你一个二维数组，二维数组中的每个数都是正数，要求从左上角走到右下角，每一步只能向右或者向下。沿途经过的数字要累
//加起来。返回最小的路径和。
//暴力递归：除了最后一列和最后一行外，其余的点都有向右和向下两种走法，则应该取两种路径的最小走法，最后一列只能向下走，最后一行只能向右走
//baseCase为走到右下角
func minPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	return processMinPath(matrix, 0, 0)
}

func processMinPath(matrix [][]int, cow, column int) int {
	res := matrix[cow][column]
	if cow == len(matrix)-1 && column == len(matrix[0])-1 {
		return res
	} else if cow == len(matrix)-1 {
		return res + processMinPath(matrix, cow, column+1)
	} else if column == len(matrix[0])-1 {
		return res + processMinPath(matrix, cow+1, column)
	} else {
		return res + int(math.Min(float64(processMinPath(matrix, cow+1, column)), float64(processMinPath(matrix, cow, column+1))))
	}
}

//动态规划解法：观察暴力递归试法，当前状态与其他状态的规律，当前processMinPath结果仅受参数cow和column影响
//则需要建立一个二维的DP表，横坐标表示column， 纵坐标表示cow，维护processMinPath值的结果
//对应二维表数值的计算规律：当cow == len(matrix)-1 && column == len(matrix[0])-1，DP[cow, column]值为matrix[cow][column]
//当cow == len(matrix)-1时，DP[cow, column]对应值为matrix[cow][column]+processMinPath(matrix, cow, column+1),即matrix[cow][column]+DP表下一列的值
//column == len(matrix[0])-1, DP[cow, column]对应值为matrix[cow][column]+processMinPath(matrix, cow+1, column),即matrix[cow][column]+DP表下一行的值
//否则，DP[cow, column]对应值为min{processMinPath(matrix, cow+1, column), processMinPath(matrix, cow, column+1)}+matrix[cow][column]
//由此，可以计算出整个DP表的值
func minPathDynamic(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	cow := len(matrix)
	column := len(matrix[0])
	dp := make([][]int, cow)
	for i := 0; i < cow; i++ {
		dp[i] = make([]int, column)
	}
	dp[cow-1][column-1] = matrix[cow-1][column-1]
	//计算dp表的最后一列元素值
	for i := cow - 2; i >= 0; i-- {
		dp[i][column-1] = matrix[i][column-1] + dp[i+1][column-1]
	}
	//计算dp表的最后一行元素值
	for i := column - 2; i >= 0; i-- {
		dp[cow-1][i] = matrix[cow-1][i] + dp[cow-1][i+1]
	}
	//计算dp表的中间元素值
	for i := cow - 2; i >= 0; i-- {
		for j := column - 2; j >= 0; j-- {
			dp[i][j] = matrix[i][j] + int(math.Min(float64(dp[i+1][j]), float64(dp[i][j+1])))
		}
	}
	return dp[0][0]
}

//题目八：给你一个数组arr，和一个整数aim。如果可以任意选择arr中的数字，能不能累加得到aim，返回true或者false
//暴力递归：可以参考题目三，打印字符串全部子序列的思想，如数组[2, 5, 6], aim=8，状态值为process(arr, i, preSum, aim)，其中
//i表示当前索引值，preSum表示累加到当前索引的前一个的索引的累加和，如process(arr, 1, 2, aim)表示索引值为1时，累加到索引值为0的和为2
//且每种索引有两种情况：（1）包含当前该索引值(index-1)的累加和（2）不包含当前索引值(index-1)的累加和
func arrSum(arr []int, aim int) bool {
	return processArrSum(arr, 0, 0, aim)
}

func processArrSum(arr []int, index, preSum, aim int) bool {
	//计算累加到数组最后一个数字的累加和时，递归方程中index应该为数组长度
	if index == len(arr) {
		if preSum == aim {
			return true
		} else {
			return false
		}
	}
	//当前索引值的前一个索引的累加和分为包含arr[index]和不包含arr[index]两种情况
	return processArrSum(arr, index+1, preSum, aim) || processArrSum(arr, index+1, preSum+arr[index], aim)
}

//动态规划解法：状态方程即为上述的递归方程式，以index和preSum构建DP表，记录状态方程的状态值,DP表的行表示index[0, len(arr)+1]，列为prSum值[0， preSum]
func arrSumDynamic(arr []int, aim int) bool {
	if len(arr) == 0 {
		return false
	}
	cow := len(arr)
	var column int
	for _, val := range arr {
		column += val
	}
	if aim > column {
		return false
	}
	dp := make([][]bool, cow+1)
	for i := 0; i <= cow; i++ {
		dp[i] = make([]bool, column+1)
	}
	//计算dp表记录对应状态值，由暴力递归的base case可知，当index=len(arr)时，preSum=aim时，即dp[cow][aim]=true
	//由暴力递归的函数可推知状态方程：dp[i][j]=dp[i+1][j] || dp[i+1][j+arr[i]]
	for i := 0; i <= cow; i++ {
		//由dp状态方程可知，dp[cow][aim]为true时，dp[cow-1][aim]=true ... dp[0][aim]=true
		//dp[cow][0]=false, dp[cow][1]=false, ... , dp[cow][aim-1]=false, dp[cow][aim]=true, dp[cow][aim+1]=false, ... dp[cow][column]=false
		dp[i][aim] = true
	}
	//dp[cow][aim] = true
	for i := cow - 1; i >= 0; i-- {
		for j := 0; j <= column; j++ {
			dp[i][j] = dp[i+1][j]
			if j+arr[i] <= column {
				dp[i][j] = dp[i+1][j] || dp[i+1][j+arr[i]]
			}
		}
	}
	return dp[0][0]
}

//题目九：给定两个数组w和v，两个数组长度相等，w[i]表示第i件商品的重量，v[i]表示第i件商品的价值。再给定一个整数bag，要求你挑选商品的重量加起来一定不能超过bag，
//返回满足这个条件下，你能获得的最大价值。
//背包问题，暴力递归法: 解题思路同题目三和题目八：计算价值数组所有可能子序列的和，递归试法时包含两种情况：包含当前索引值和不包含当前索引值
func knapsack(w, v []int, bag int) int {
	//alreadyWeight表示之前所作决定形成的重量
	return processKnapsack(w, v, 0, 0, bag)
}

func processKnapsack(w, v []int, index, alreadyWeight, bag int) int {
	//base case
	//当此时重量超过bag值时，则当前索引值不能被加入该商品集中.该方案舍弃
	//返回值要保证加上v[index]之后不会比不加v[index]的大，否则，return上级函数时还会取这个方案
	if alreadyWeight > bag {
		return INT_MIN
	}
	//当索引超过数组长度时，index之后返回的价值肯定是0
	if index == len(w) {
		return 0
	}
	subRes01 := processKnapsack(w, v, index+1, alreadyWeight, bag)
	alreadyWeight = alreadyWeight + w[index]
	subRes02 := v[index] + processKnapsack(w, v, index+1, alreadyWeight, bag)
	maxRes := math.Max(float64(subRes01), float64(subRes02))
	return int(maxRes)
	/*return int(
		math.Max(float64(processKnapsack(w, v, index+1, alreadyWeight, bag)),
			float64(v[index]+processKnapsack(w, v, index+1, alreadyWeight+w[index], bag))),
	)*/
}

//动态规划思想：dp状态方程：processKnapsack(w, v, index, alreadyWeight, bag)=math.Max(processKnapsack(w, v, index+1, alreadyWeight, bag),
//v[index]+processKnapsack(w, v, index+1, alreadyWeight+w[index], bag))
//则可以以index, totalWeight为行和列构建dp表， 且dp[.][bag+1]=INT_MIN, bag左边的列值都置为最小值，即当alreadyWeight大于bag时，保证不会采用该方案
//dp[cow][.]=0, 即最后一行都置为0
func knapsackDynamic(w, v []int, bag int) int {
	var totalWeight int
	for _, val := range w {
		totalWeight += val
	}
	var totalVal int
	for _, val := range v {
		totalVal += val
	}
	if bag >= totalWeight {
		return totalVal
	}
	cow := len(w)
	//初始化dp表
	dp := make([][]int, cow+1)
	for i := 0; i <= cow; i++ {
		dp[i] = make([]int, totalWeight+1)
	}

	//dp表bag 右列值置为最小值
	for i := 0; i <= cow; i++ {
		for j := bag + 1; j <= totalWeight; j++ {
			dp[i][j] = INT_MIN
		}
	}
	//dp表最后一行值为0，即可以保持初始默认值
	for i := cow - 1; i >= 0; i-- {
		for j := 0; j <= totalWeight; j++ {
			dp[i][j] = dp[i+1][j]
			if j+w[i] <= totalWeight {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(v[i]+dp[i+1][j+w[i]])))
			}
		}
	}
	return dp[0][0]
}
