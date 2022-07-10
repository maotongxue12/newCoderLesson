package main

import "fmt"

func rotateEdge(matrix [][]int, leftRow, leftColumn, rightRow, rightColumn int) {
	times := rightColumn - leftColumn
	for i := 0; i < times; i++ {
		temp := matrix[leftRow][leftColumn+i]
		matrix[leftRow][leftColumn+i] = matrix[rightRow-i][leftColumn]
		matrix[rightRow-i][leftColumn] = matrix[rightRow][rightColumn-i]
		matrix[rightRow][rightColumn-i] = matrix[leftRow+i][rightColumn]
		matrix[leftRow+i][rightColumn] = temp
	}
}

func rotate(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}
	leftRow, leftList := 0, 0
	rightRow := len(matrix) - 1
	rightList := len(matrix[0]) - 1
	for leftRow <= rightRow && leftList <= rightList {
		rotateEdge(matrix, leftRow, leftList, rightRow, rightList)
		leftRow++
		leftList++
		rightList--
		rightRow--
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 12}}
	rotate(matrix)
	fmt.Println(matrix)
}
