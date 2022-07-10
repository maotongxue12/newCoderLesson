package main

import (
	"fmt"
)

//实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返
//回栈中最小元素的操作。
type specialStack struct {
	minStack []int
	normalStack []int
}

const maxSize int = 5

func newSpecialStack() *specialStack {
	return &specialStack{
		minStack: make([]int, 0),
		normalStack: make([]int, 0),
	}
}

func (s *specialStack) push(data int) {
	if len(s.minStack) == 0 || s.minStack[len(s.minStack)-1] > data{
		s.minStack = append(s.minStack, data)
	} else {
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
	s.normalStack = append(s.normalStack, data)
}

func (s *specialStack) pop() error {
	if len(s.normalStack) == 0 {
		return fmt.Errorf("the stack is empty")
	}
	s.normalStack = s.normalStack[:len(s.normalStack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
	return nil
}

func (s *specialStack) minValue() (value int,err error) {
	if 0 == len(s.normalStack) {
		return value, fmt.Errorf("the stack is empty")
	}
	value = s.minStack[len(s.minStack)-1]
	return  value,nil
}

func main() {
	specialStack := newSpecialStack()
	specialStack.push(5)
	specialStack.push(4)
	specialStack.push(1)
	specialStack.push(2)
	fmt.Println(specialStack.minValue())
}