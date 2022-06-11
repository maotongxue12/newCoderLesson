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
	return n*getFibonacci01(n-1)
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

//汉诺塔：打印n层汉诺塔从最左边移动到最右边的全部过程
func move(disk int, N byte, M byte) {
	fmt.Println("把 ", disk, " 号圆盘从 ", N, " 移到 ", M)
}

//把n号盘从A移到C，以b为辅助
func hanoi(n int, A byte, B byte, c byte) {
	if n == 1 {
		move(1, 'A', 'C')
	}
	hanoi(n-1, 'A', 'C', 'B')
	move(n, 'A','C')
	hanoi(n-1, 'B', 'A', 'C')
}