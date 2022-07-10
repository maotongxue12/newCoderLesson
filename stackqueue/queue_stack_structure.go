package stackqueue

import (
	"errors"
	"fmt"
)

//题目一：用数组结构实现大小固定的队列
//解题思路：队列是先进先出的一种，go语言中存在一种切片结构，可以直接当作队列结构来用，
//利用切片封装一个struct结构体，来实现队列的相关函数
type queue struct {
	value []int
	first int
	last  int
}

func newQueue(maxSize int) *queue {
	return &queue{
		value: make([]int, maxSize+1),
		first: 0,
		last:  0,
	}
}

func (q *queue) length() int {
	return (q.last - q.first + len(q.value)) % len(q.value)
}

func (q *queue) isEmpty() bool {
	if q.first == q.last {
		return true
	}
	return false
}

func (q *queue) push(val int) error {
	if (q.last-q.first+len(q.value))%len(q.value) == len(q.value) {
		return errors.New("the queue is full")
	}
	q.value[q.last] = val
	q.last = (q.last + 1) % len(q.value)
	return nil
}

func (q *queue) pop() error {
	if q.last == q.first {
		err := errors.New("the queue is empty")
		return err
	}
	q.first = (q.first+1) % len(q.value)
	return nil
}

func (q *queue) front() (val int, err error) {
	if q.last == q.first {
		err := errors.New("the queue is empty")
		return val, err
	}
	val = q.value[q.first]
	return val, err
}


//题目二：用数组实现栈结构
//解题思路：栈，先进后出
type stack struct {
	value []int
	index int
}

func newStack(maxSize int) *stack {
	return &stack{
		value: make([]int, maxSize+1),
		index: 0,
	}
}

func (s *stack) isEmpty() bool {
	if s.index == 0 {
		return true
	}
	return false
}

func (s *stack) length() int {
	return s.index
}

func (s *stack) push(val int) error {
	if s.index == len(s.value)-1 {
		return fmt.Errorf("the stack is full")
		//return errors.New("the stack is full")
	}
	s.value[s.index] = val
	s.index++
	return nil
}

func (s *stack) pop() error {
	if s.index == 0 {
		return errors.New("the stack is empty")
	}
	s.index--
	return nil
}

func (s *stack) top() (val int, err error) {
	if s.index == 0 {
		return val, errors.New("the stack is empty")
	}
	val = s.value[s.index-1]
	return val, nil
}