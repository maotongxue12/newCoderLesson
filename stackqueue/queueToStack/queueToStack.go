package main

import (
	"fmt"
)

//用数组实现队列结构
const maxSize int = 5

type queue struct {
	arr   [maxSize]interface{}
	last  int
	first int
}

func newQueue() *queue {
	return &queue{
		last:  0,
		first: 0,
	}
}

func (q *queue) empty() bool {
	return q.last == q.first
}

func (q *queue) length() int {
	return (q.last - q.first + maxSize) % maxSize
}

func (q *queue) push(data interface{}) error {
	if q.first == (q.last+1)%maxSize {
		return fmt.Errorf("the queue is full")
	}
	q.arr[q.last] = data
	q.last = (q.last + 1) % maxSize
	return nil
}

func (q *queue) pop() error {
	if q.first == q.last {
		return fmt.Errorf("the queue is empty")
	}
	q.first = (q.first + 1) % maxSize
	return nil
}

func (q *queue) front() (value interface{}, err error) {
	if q.first == q.last {
		return value, fmt.Errorf("the queue is empty")
	}
	value = q.arr[q.first]
	return value, nil
}

//用两个队列结构实现栈结构
type stack struct {
	tutorialQueue *queue
	normalQueue   *queue
}

func newStack() *stack {
	return &stack{
		tutorialQueue: newQueue(),
		normalQueue:   newQueue(),
	}
}

func (s *stack) length() int {
	return s.normalQueue.length()
}

func (s *stack) push(data int) error {
	if s.normalQueue.length() == maxSize-1 {
		return fmt.Errorf("the stack is full")
	}
	s.normalQueue.push(data)
	return nil
}

func (s *stack) pop() error {
	if 0 == s.normalQueue.length() {
		return fmt.Errorf("the stack is empty")
	}
	len := s.normalQueue.length()-1
	for i := 0; i < len; i++ {
		value, _ := s.normalQueue.front()
		s.tutorialQueue.push(value)
		s.normalQueue.pop()
	}
	s.normalQueue.pop()
	len = s.tutorialQueue.length()
	for i := 0; i < len; i++ {
		value, _ := s.tutorialQueue.front()
		s.normalQueue.push(value)
		s.tutorialQueue.pop()
	}
	return nil
}

func (s *stack) top() (value interface{}, err error) {
	if 0 == s.normalQueue.length() {
		return nil, fmt.Errorf("the stack is empty")
	}
	len := s.normalQueue.length()
	for i := 0; i < len; i++ {
		value, _ = s.normalQueue.front()
		s.tutorialQueue.push(value)
		s.normalQueue.pop()
	}
	s.normalQueue.pop()
	len = s.tutorialQueue.length()
	for i := 0; i < len; i++ {
		num, _ := s.tutorialQueue.front()
		s.normalQueue.push(num)
		s.tutorialQueue.pop()
	}
	return value, nil
}

//打印栈
func (s *stack) printStack() {
	if s.length() == 0 {
		fmt.Println("the stack is empty, please check!")
		return
	}
	for i := 0; i < s.length(); i++ {
		fmt.Println(i, "->", s.normalQueue.arr[i])
	}
}

func main() {
	queueStack := newStack()
	queueStack.push(5)
	queueStack.push(4)
	queueStack.push(3)
	queueStack.printStack()
	value, _ := queueStack.top()
	fmt.Println(value)
	queueStack.pop()
	value, _ = queueStack.top()
	fmt.Println(value)
	queueStack.pop()
	value, _ = queueStack.top()
	fmt.Println(value)
	queueStack.pop()
	queueStack.printStack()
}
