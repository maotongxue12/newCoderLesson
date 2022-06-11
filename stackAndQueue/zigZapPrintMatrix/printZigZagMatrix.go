package main

import "fmt"

func printLevel(matrix [][]int, leftRow, leftColumn, rightRow, rightColumn int, flag bool) {
	//如果flag为true，则从左下角向右上角打印，否则反向打印
	if flag {
		for leftColumn <= rightColumn {
			fmt.Print(matrix[leftRow][leftColumn], " ")
			leftRow--
			leftColumn++
		}
	} else {
		for rightRow <= leftRow {
			fmt.Print(matrix[rightRow][rightColumn], " ")
			rightRow++
			rightColumn--
		}
	}
}

func printZigZagMatrix(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}
	leftRow, leftColum := 0, 0
	rightRow, rightColum := 0, 0
	endRow := len(matrix) - 1
	endColum := len(matrix[0]) - 1
	flag := true
	for rightRow <= endRow {
		printLevel(matrix, leftRow, leftColum, rightRow, rightColum, flag)
		if leftRow == endRow {
			leftColum = leftColum + 1
			leftRow = leftRow
		} else {
			leftColum = leftColum
			leftRow = leftRow + 1
		}
		if rightColum == endColum {
			rightColum = rightColum
			rightRow = rightRow + 1
		} else {
			rightColum = rightColum + 1
			rightRow = rightRow
		}
		flag = !flag
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	printZigZagMatrix(matrix)
}
