package stackqueue

import (
	"errors"
	"fmt"
)

//题目一：实现一种狗猫队列的结构，要求如下： 用户可以调用add方法将cat类或dog类的 实例放入队列中；
//用户可以调用pollAll方法，将队列中所有的实例按照进队列 的先后顺序依次弹出； 用户可以调用pollDog方法，
//将队列中dog类的实例按照进队列的先后顺序依次弹出；用户可以调用pollCat方法，将队列中cat类的实例按照进队
//列的先后顺序依次弹出；用户可以调用isEmpty方法，检查队列中是否还有dog或cat的实例；用户可以调用isDogEmpty方法，
//检查队列中是否有dog类的实例；用户可以调用isCatEmpty方法，检查队列中是否有cat类的实例。

//思路：首先考虑可以分别依次弹出dog和cat，则需要用两个队列来维护；
//同时要保证可以不分cat和dog，可以依次弹出，则可以给cat和dog添加一个时间戳之类的索引，当调用poolAll时，
//比较cat队列和dog队列的队首时间戳，哪个时间戳靠前就弹出哪个，以此来时间部分cat和dog依次弹出

//pet结构体，封装宠物类型,并增加入队索引（时间戳）,记录当前cat或dog在整体pet中的索引
type pet struct {
	petType string
	count   int
}

func (p *pet) getPetType() string {
	return p.petType
}

func (p *pet) setPetType(kind string) {
	p.petType = kind
}

//猫狗队列结构体，分别存储cat和dog, 记录当前pet的索引，并将该索引值赋值pet中count字段，表示cat或dog的整体入队先后
type catDotQueue struct {
	catQueue []pet
	dogQueue []pet
	index    int
}

func (c *catDotQueue) empty() bool {
	if len(c.catQueue) == 0 && len(c.dogQueue) == 0 {
		return true
	}
	return false
}

func (c *catDotQueue) isDogEmpty() bool {
	if len(c.dogQueue) == 0 {
		return true
	}
	return false
}

func (c *catDotQueue) isCatEmpty() bool {
	if len(c.catQueue) == 0 {
		return true
	}
	return false
}

func (c *catDotQueue) add(animal pet) error {
	if animal.getPetType() == "cat" {
		c.catQueue = append(c.catQueue, pet{
			petType: animal.getPetType(),
			count:   c.index,
		})
		c.index++
	} else if animal.getPetType() == "dog" {
		c.dogQueue = append(c.dogQueue, pet{
			petType: animal.getPetType(),
			count:   c.index,
		})
		c.index++
	} else {
		return errors.New("the pet is not cat or dog")
	}
	return nil
}

func (c *catDotQueue) pollAll() (p pet, err error) {
	if len(c.catQueue) > 0 && len(c.dogQueue) > 0 {
		if c.catQueue[0].count < c.dogQueue[0].count {
			p = c.catQueue[0]
			c.catQueue = c.catQueue[1:len(c.catQueue)]
		} else {
			p = c.dogQueue[0]
			c.dogQueue = c.dogQueue[1:len(c.dogQueue)]
		}
	} else if len(c.catQueue) > 0 {
		p = c.catQueue[0]
		c.catQueue = c.catQueue[1:len(c.catQueue)]
	} else if len(c.dogQueue) > 0 {
		p = c.dogQueue[0]
		c.dogQueue = c.dogQueue[1:len(c.dogQueue)]
	} else {
		return p, errors.New("the queue is empty")
	}
	return p, nil
}

func (c *catDotQueue) pollCat() (p pet, err error) {
	if len(c.catQueue) > 0 {
		p = c.catQueue[0]
		c.catQueue = c.catQueue[1:len(c.catQueue)]
	} else {
		return p, errors.New("the catQueue is empty")
	}

	return p, nil
}

func (c *catDotQueue) pollDog() (p pet, err error) {
	if len(c.dogQueue) > 0 {
		p = c.dogQueue[0]
		c.dogQueue = c.dogQueue[1:]
	} else {
		return p, errors.New("the dogQueue is empty")
	}
	return p, nil
}

//题目二：转圈打印矩阵
//思路：处理此类从宏观角度考虑，不能局限于某个点，转圈打印即可以先打印外围，然后打印内圈
func printRotateMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	leftRow := 0
	leftColumn := 0
	rightRow := len(matrix) - 1
	rightColumn := len(matrix[0]) - 1
	for leftRow <= rightRow && leftColumn <= rightColumn {
		printEdgeMatrix(matrix, leftRow, leftColumn, rightRow, rightColumn)
		leftRow++
		leftColumn++
		rightRow--
		rightColumn--
	}
}

//按圈打印
func printEdgeMatrix(matrix [][]int, leftRow, leftColumn, rightRow, rightColumn int) {
	if leftRow == rightRow {
		for i := leftColumn; i <= rightColumn; i++ {
			fmt.Print(matrix[leftRow][i], " ")
		}
	} else if leftColumn == rightColumn {
		for i := leftRow; i <= rightRow; i++ {
			fmt.Print(matrix[i][leftColumn], " ")
		}
	} else {
		for i := leftColumn; i < rightColumn; i++ {
			fmt.Print(matrix[leftRow][i], " ")
		}
		for i := leftRow; i < rightRow; i++ {
			fmt.Print(matrix[i][rightColumn], " ")
		}
		for i := rightColumn; i > leftColumn; i-- {
			fmt.Print(matrix[rightRow][i], " ")
		}
		for i := rightRow; i > leftRow; i-- {
			fmt.Print(matrix[i][leftColumn], " ")
		}
	}
}

//题目三：旋转正方形矩阵--给定一个整型正方形矩阵matrix，请把该矩阵调整成 顺时针旋转90度的样子
//思路:同题目二思路，首先旋转矩阵的外围，然后从外到内，依次按圈旋转90°
func rotateSquire(matrix [][]int) {
	leftRow, leftColumn := 0, 0
	rightRow := len(matrix) - 1
	if rightRow == 0 {
		return
	}
	rightColumn := len(matrix[0]) - 1
	for leftRow <= rightRow {
		rotateEdge(matrix, leftRow, leftColumn, rightRow, rightColumn)
		leftRow++
		leftColumn++
		rightRow--
		rightColumn--
	}
}

func rotateEdge(matrix [][]int, leftCow, leftColumn, rightRow, rightColumn int) {
	times := rightColumn - leftColumn
	for i := 0; i < times; i++ {
		temp := matrix[leftCow][leftColumn+i]
		matrix[leftCow][leftColumn+i] = matrix[rightRow-i][leftColumn]
		matrix[rightRow-i][leftColumn] = matrix[rightRow][rightColumn-i]
		matrix[rightRow][rightColumn-i] = matrix[leftCow+i][rightColumn]
		matrix[leftCow+i][rightColumn] = temp
	}
}

//题目四：“之”字形打印矩阵 【题目】 给定一个矩阵matrix，按照“之”字形的方式打印这 个矩阵，
//例如： 1 2 3 4 5 6 7 8 9 10 11 12 “之”字形打印的结果为：1，2，5，9，6，3，4，7，10，11， 8，12
//【要求】 额外空间复杂度为O(1)。
func zigzagPrintMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	endRow := len(matrix) - 1
	endColumn := len(matrix[0]) - 1
	leftRow, leftColumn := 0, 0
	rightRow, rightColumn := 0, 0
	flag := true
	for leftColumn <= endColumn {
		printZigzag(matrix, leftRow, leftColumn, rightRow, rightColumn, flag)
		if leftRow == endRow {
			leftColumn++
			leftRow = leftRow
		} else {
			leftRow++
			leftColumn = leftColumn
		}
		if rightColumn == endColumn {
			rightColumn = rightColumn
			rightRow++
		} else {
			rightColumn++
			rightRow = rightRow
		}
		flag = !flag
	}
}

func printZigzag(matrix [][]int, leftRow, leftColumn, rightRow, rightColumn int, flag bool) {
	//如果flag == true，则从左下向右上打印，反之从右上向左下打印
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

//题目五：在行列都排好序的矩阵中找数
//【题目】 给定一个有N*M的整型矩阵matrix和一个整数K， matrix的每一行和每一 列都是排好序的。
//实现一个函数，判断K 是否在matrix中。 例如： 0 1 2 5 2 3 4 7 4 4 4 8 5 7 7 9 如果K为7，返回true；
//如果K为6，返 回false。
//【要求】 时间复杂度为O(N+M)，额外空间复杂度为O(1)。
//思路：注意题目中标注了每一行和每一列都是排好序的，即同一行中右边的值肯定比左边的大，同一列中下边值肯定比上边大
//则可以从（0，M）开始寻找，如果key值大于该值，则向下寻找，如果key小于该值，则向左寻找，直至找到（N, 0）点
func selectNumber(matrix [][]int, key int) bool {
	if len(matrix) == 0 {
		return false
	}
	N := len(matrix) - 1
	M := len(matrix[0]) - 1
	row, column := 0, M
	for row <= N && column >= 0 {
		if matrix[row][column] == key {
			return true
		} else if matrix[row][column] < key {
			row++
		} else {
			column--
		}
	}
	return false
}
