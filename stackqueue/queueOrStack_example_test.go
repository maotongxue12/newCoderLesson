package stackqueue

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetPetType(t *testing.T) {
	Convey("test the catDogQueue function", t, func() {
		cdQue := new(catDotQueue)
		cdQue.add(pet{petType: "dog"})
		p, err := cdQue.pollDog()
		So(p, ShouldResemble, pet{petType: "dog", count: 0})
		So(err, ShouldBeNil)
		cdQue.add(pet{petType: "cat"})
		cdQue.add(pet{petType: "dog"})
		cdQue.add(pet{petType: "dog"})
		p, err = cdQue.pollAll()
		So(p, ShouldResemble, pet{petType: "cat", count: 1})
		p, err = cdQue.pollAll()
		So(p, ShouldResemble, pet{petType: "dog", count: 2})
		p, err = cdQue.pollCat()
		So(err, ShouldNotBeNil)
		p, err = cdQue.pollDog()
		So(err, ShouldBeNil)
	})
}

func TestPrintRotateMatrix(t *testing.T) {
	matrix01 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	printRotateMatrix(matrix01)
	fmt.Println()
	matrix02 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	printRotateMatrix(matrix02)
}

func TestRotateSquire(t *testing.T) {
	Convey("test the rotateSquire function", t, func() {
		matrix01 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
		rotateSquire(matrix01)
		fmt.Println(matrix01)
		matrix02 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 12}}
		rotateSquire(matrix02)
		fmt.Println(matrix02)
	})
}

func TestZigzagPrintMatrix(t *testing.T) {
	Convey("test the zigzagPrintMatrix function", t, func() {
		matrix01 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
		zigzagPrintMatrix(matrix01)
		fmt.Println()
		matrix02 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		zigzagPrintMatrix(matrix02)
		fmt.Println()
		matrix03 := [][]int{{1, 2, 3}}
		zigzagPrintMatrix(matrix03)
	})
}

func TestSelectNumber(t *testing.T) {
	Convey("test the selectNumber function", t, func() {
		matrix := [][]int{{0, 1, 2, 5}, {2, 3, 4, 7}, {4, 4, 4, 8}, {5, 7, 7, 9}}
		res1 := selectNumber(matrix, 7)
		So(res1, ShouldBeTrue)
		res2 := selectNumber(matrix, 6)
		So(res2, ShouldBeFalse)
	})
}
