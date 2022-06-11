package main

import "fmt"

func printEdge(matrix [][]int, leftRow, leftList, rightRow, rightList int) {
	if leftRow == rightRow {
		for i := leftList; i <= rightList; i++ {
			fmt.Print(matrix[leftRow][i], " ")
		}
	} else if leftList == rightList {
		for i := leftRow; i <= rightRow; i++ {
			fmt.Print(matrix[i][leftList], " ")
		}
	} else {
		for i := leftList; i < rightList; i++ {
			fmt.Print(matrix[leftRow][i], " ")
		}
		for i := leftRow; i < rightRow; i++ {
			fmt.Print(matrix[i][rightList], " ")
		}
		for i := rightList; i > leftList; i-- {
			fmt.Print(matrix[rightRow][i], " ")
		}
		for i := rightRow; i > leftRow; i-- {
			fmt.Print(matrix[i][leftList], " ")
		}
	}
}

func printMatrix(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}
	leftRow := 0
	leftList := 0
	rightRow := len(matrix)-1
	rightList := len(matrix[0])-1
	for  leftRow <= rightRow && leftList <= rightList {
		printEdge(matrix, leftRow, leftList, rightRow, rightList)
		leftRow++
		leftList++
		rightList--
		rightRow--
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	/*var matrix [][]int
	for i := 0; i < row; i++ {
		submatrix := []int{1, 2, 3, 4}
		matrix = append(matrix, submatrix)
	}*/
	printMatrix(matrix)
}
