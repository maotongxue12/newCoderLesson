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

//可采用设置一个空位置，以此来判断是否栈满；也可以增加计数变量size，直接根据size判断栈是否满
func (q *queue) push(val int) error {
	if (q.last-q.first+len(q.value))%len(q.value) == len(q.value)-1 {
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

//题目三：实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作
//思路：考虑用两个切片模拟两个栈，一个栈作为实际存储数据的栈，另一个栈作为辅助栈，存储当前栈中元素的最小元素
//且与实际栈中元素相对应
type specialStack struct {
	normalStack *stack
	auxiliaryStack *stack
}

func newSpecialStack(size int) *specialStack {
	return &specialStack{
		normalStack: newStack(size),
		auxiliaryStack: newStack(size),
	}
}

func (s *specialStack)push(val int) (err error) {
	data, _ := s.auxiliaryStack.top();
	//保证auxiliaryStack中保存元素为当前normalStack中元素的最小值
	if s.auxiliaryStack.length() == 0 || val < data {
		err = s.auxiliaryStack.push(val)
	} else {
		err = s.auxiliaryStack.push(data)
	}
	if err != nil {
		return errors.New("the stack is rull")
	}
	s.normalStack.push(val)
	return nil
}

func (s *specialStack) pop() error {
	if s.normalStack.length() == 0 {
		return errors.New("the stack is empty")
	}
	s.normalStack.pop()
	s.auxiliaryStack.pop()
	return nil
}

func (s *specialStack) top() (val int, err error) {
	if s.normalStack.length() == 0 {
		return val, errors.New("the stack is empty")
	}
	val, _ = s.normalStack.top()
	return val, nil
}

func (s *specialStack) getMin() (val int, err error) {
	if s.normalStack.length() == 0 {
		return val, errors.New("the stack is empty")
	}
	val, _ = s.auxiliaryStack.top()
	return val, nil
}

//题目四：如何仅用队列结构实现栈结构？
//思路：采用两个队列，一个normal队列，入栈时即正常入队，一个辅助队列，入栈时不操作
//出栈时先把队列中的n-1个元素压入辅助队列，然后弹出normal队列的最后一个值，即实现了后进先出
type queueToStack struct {
	normalQueue *queue
	auxiliaryQueue *queue
}

func newQueueToStack(size int) *queueToStack {
	return &queueToStack{
		normalQueue: newQueue(size),
		auxiliaryQueue: newQueue(size),
	}
}

func (s *queueToStack) isEmpty() bool {
	if s.normalQueue.length() == 0 {
		return true
	}
	return false
}

func (s *queueToStack) length() int {
	return s.normalQueue.length()
}
func (s *queueToStack) push(val int) error {
	if err := s.normalQueue.push(val); err != nil {
		return errors.New("the stack is full")
	}
	return nil
}

func (s *queueToStack) pop() (val int, err error) {
	if s.normalQueue.length() == 0 {
		return val, errors.New("the stack is empty")
	}
	len := s.normalQueue.length()
	for i := 0; i < len-1; i++ {
		val, _ = s.normalQueue.front()
		s.normalQueue.pop()
		s.auxiliaryQueue.push(val)
	}
	val, _ = s.normalQueue.front()
	s.normalQueue.pop()
	len = s.auxiliaryQueue.length()
	for i := 0; i < len; i++ {
		tmp, _ := s.auxiliaryQueue.front()
		s.normalQueue.push(tmp)
		s.auxiliaryQueue.pop()
	}
	return val, nil
}

func (s *queueToStack) front() (val int, err error) {
	if s.normalQueue.length() == 0 {
		return val, errors.New("the stack is empty")
	}
	len := s.normalQueue.length()
	for i := 0; i < len; i++ {
		val, _ = s.normalQueue.front()
		s.normalQueue.pop()
		s.auxiliaryQueue.push(val)
	}
	len = s.auxiliaryQueue.length()
	for i := 0; i < len; i++ {
		tmp, _ := s.auxiliaryQueue.front()
		s.normalQueue.push(tmp)
		s.auxiliaryQueue.pop()
	}
	return val, nil
}

//题目五：如何仅用栈结构实现队列结构？
//思路：normal 栈用于入队，auxiliary栈用于数据弹出时，先将normal栈中数据弹出压入auxiliary栈中
//然后从auxiliary中弹出；原则一：当将normal栈中数据压入auxiliary栈中时，一次性将normal全部弹出
//原则二：当auxiliary栈为空时，才能开始将normal中数据压入auxiliary中
type stackToQueue struct {
	normalStack *stack
	auxiliaryStack *stack
}

func newStackToQueue(size int) *stackToQueue{
	return &stackToQueue{
		normalStack: newStack(size),
		auxiliaryStack: newStack(size),
	}
}

func (s *stackToQueue) empty() bool {
	if s.normalStack.length() == 0 && s.auxiliaryStack.length() == 0 {
		return true
	}
	return false
}

func (s *stackToQueue) length() int {
	return s.normalStack.length() + s.auxiliaryStack.length()
}

func (s *stackToQueue) push(val int) error {
	if err := s.normalStack.push(val); err != nil {
		return errors.New("the stack is full")
	}
	return nil
}

func (s *stackToQueue) pop() error {
	len := s.normalStack.length()
	if s.auxiliaryStack.isEmpty() {
		for i := 0; i < len; i++ {
			tmp, _ := s.normalStack.top()
			s.normalStack.pop()
			s.auxiliaryStack.push(tmp)
		}
	}
	err := s.auxiliaryStack.pop()
	if err != nil {
		return errors.New("the stack is empty")
	}
	return nil
}

func (s *stackToQueue) top() (val int, err error) {
	len := s.normalStack.length()
	if s.auxiliaryStack.isEmpty() {
		for i := 0; i < len; i++ {
			tmp, _ := s.normalStack.top()
			s.normalStack.pop()
			s.auxiliaryStack.push(tmp)
		}
	}
	val, err = s.auxiliaryStack.top()
	if err != nil {
		return val, errors.New("the stack is empty")
	}
	return val, nil
}