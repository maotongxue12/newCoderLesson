package main

import "fmt"

//数组模拟栈
const maxSize int = 5

type stack struct {
	arr   [maxSize]interface{}
	index int
}

func newStack() *stack {
	return &stack{
		index: 0,
	}
}

func (s *stack) empty() bool {
	return s.index == 0
}

func (s *stack) length() int {
	return s.index
}

func (s *stack) push(data interface{}) error {
	if s.index == maxSize {
		return fmt.Errorf("the stack is full")
	}
	s.arr[s.index] = data
	s.index++
	return nil
}

func (s *stack) pop() error {
	if s.index == 0 {
		return fmt.Errorf("the stack is empty")
	}
	s.index--
	return nil
}

func (s *stack) top() (interface{}, error) {
	if s.index == 0 {
		return nil, fmt.Errorf("the stack is empty")
	}
	num := s.arr[s.index-1]
	return num, nil
}

//两个栈实现队列
//栈1元素倒入栈2元素，当栈2元素不为空时，不能倒数据
//栈1倒数据时，必须一次性倒完
type queue struct {
	normalStack   *stack
	tutorialStack *stack
}

func newQueue() *queue {
	return &queue{
		normalStack:   newStack(),
		tutorialStack: newStack(),
	}
}

func (q *queue) length() int {
	return q.normalStack.length() + q.tutorialStack.length()
}

func (q *queue) push(data interface{}) error {
	if q.normalStack.length() == maxSize {
		return fmt.Errorf("the queue is full")
	}
	q.normalStack.push(data)
	return nil
}

func (q *queue) pop() error {
	if q.tutorialStack.empty() && q.normalStack.empty() {
		return fmt.Errorf("the queue is empty")
	}
	for q.tutorialStack.empty() {
		for !q.normalStack.empty() {
			value, _ := q.normalStack.top()
			q.tutorialStack.push(value)
			q.normalStack.pop()
		}
	}
	q.tutorialStack.pop()
	return nil
}

func (q *queue) front() (interface{}, error) {
	if q.tutorialStack.empty() && q.normalStack.empty() {
		return nil, fmt.Errorf("the queue is empty")
	}
	for q.tutorialStack.empty() {
		for !q.normalStack.empty() {
			value, _ := q.normalStack.top()
			q.tutorialStack.push(value)
			q.normalStack.pop()
		}
	}
	value, err := q.tutorialStack.top()
	return value, err
}

func (q *queue) printQueue() {
	if q.normalStack.empty() && q.tutorialStack.empty() {
		return
	}
	for i := q.tutorialStack.length()-1; i >= 0 ; i-- {
		fmt.Print(i, "->", q.tutorialStack.arr[i])
	}
	for i := 0; i< q.normalStack.length(); i++ {
		fmt.Print(i, "->", q.normalStack.arr[i])
	}
}

func main() {
	queue := newQueue()
	queue.push(5)
	queue.push(4)
	queue.push(3)
	value, _ := queue.front()
	fmt.Println(value)
	queue.pop()
	value, _ = queue.front()
	fmt.Println(value)
	queue.pop()
	value, _ = queue.front()
	fmt.Println(value)
	queue.push(6)
	value, _ = queue.front()
	fmt.Println(value)
	queue.pop()
	value, _ = queue.front()
	fmt.Println(value)
	queue.pop()
	value, _ = queue.front()
	fmt.Println(value)
}
