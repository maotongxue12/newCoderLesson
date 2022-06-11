package main

import "fmt"

func findNumInMatrix(matrix [][]int, key int) bool {
	leftRow := 0
	leftColumn := 0
	rightRow := len(matrix) - 1
	rightColumn := len(matrix[0]) - 1
	startRow := leftRow
	startColumn := rightColumn
	for startRow <= rightRow && startColumn >= leftColumn {
		if matrix[startRow][startColumn] == key {
			return true
		}
		if  key < matrix[startRow][startColumn] {
			startColumn--
		} else {
			startRow++
		}
	}
	return false
}

func main() {
	matrix := [][]int{{0,1,2,5}, {2,3,4,7}, {4,4,4,8}, {5,7,7,9}}
	res := findNumInMatrix(matrix, 8)
	fmt.Println(res)
}