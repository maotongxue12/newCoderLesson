package main

import (
	"fmt"
)

type sliceStack struct {
	arr       []interface{}
	stackSize int
}

func newSliceStack() *sliceStack {
	return &sliceStack{
		arr: make([]interface{}, 0),
	}
}

func (s *sliceStack) empty() bool {
	return s.stackSize == 0
}

//获取栈的大小
func (s *sliceStack) size() int {
	return s.stackSize
}

//push栈元素
func (s *sliceStack) push(t interface{}) {
	s.arr = append(s.arr, t)
	s.stackSize = s.stackSize + 1
}

//弹出栈顶元素
func (s *sliceStack) pop() {
	if s.stackSize == 0 {
		fmt.Errorf("stack is empty")
		return
	}
	s.stackSize--
	s.arr = s.arr[:s.stackSize]
}

//获取栈顶元素
func (s *sliceStack) top() interface{} {
	if s.stackSize == 0 {
		fmt.Println("statck is empty")
		return nil
	}
	num := s.arr[s.stackSize-1]
	return num
}

//打印栈
func (s *sliceStack) printStack() {
	if s.stackSize == 0 {
		fmt.Println("the stack is empty, please check!")
		return
	}
	for i := 0; i < s.stackSize; i++ {
		fmt.Println(i, "->", s.arr[i])
	}
}

//用数组实现循环队列
const maxSize int = 5

type queue struct {
	arr [maxSize]interface{}
	first int	//指向队首
	last int	//指向队尾元素的下一个位置
}

//初始化队列
func newQueue() *queue{
	return &queue{
		first: 0,
		last: 0,
	}
}

//判断是否为空
func (q *queue) empty() bool {
	return q.first == q.last
}

//获取队列长度
func (q *queue) length() int {
	return (q.last - q.first + maxSize) % maxSize
}

//元素进队列,长度为5的队列只能存四个元素，始终空一个元素
func (q *queue) push(data interface{}) {
	if q.first == (q.last+1)%maxSize {
		fmt.Println("the queue is full!")
		return
	}
	q.arr[q.last] = data
	q.last = (q.last+1) % maxSize
}

//获取队首元素
func (q *queue) front() interface{} {
	if q.first == q.last {
		fmt.Println("the queue is empty, please check!")
		return nil
	}
	element := q.arr[q.first]
	q.first = (q.first+1)%maxSize
	return element
}

//弹出队首元素
func(q *queue) pop() {
	if q.first == q.last {
		fmt.Println("the queue is empty")
	}
	q.first = (q.first+1)%maxSize
}

func main() {
	testStack := newSliceStack()
	testStack.push(5)
	testStack.push(3)
	testStack.push(2)
	testStack.printStack()
	testStack.pop()
	fmt.Println(testStack.top())
	testStack.pop()
	testStack.pop()

	testQueue := newQueue()
	testQueue.push(5)
	testQueue.push(4)
	testQueue.push(3)
	testQueue.push(2)
	testQueue.push(1)
	fmt.Println(testQueue.length())
	fmt.Println(testQueue.first)
	fmt.Println(testQueue.last)
}