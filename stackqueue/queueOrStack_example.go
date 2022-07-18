package stackqueue

import "errors"

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
	count int
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
	index int
}

func (c *catDotQueue)  empty() bool {
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
			count: c.index,
		})
		c.index++
	} else if animal.getPetType() == "dog" {
		c.dogQueue = append(c.dogQueue, pet{
			petType: animal.getPetType(),
			count: c.index,
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

func (c *catDotQueue)  pollDog() (p pet, err error) {
	if len(c.dogQueue) > 0 {
		p = c.dogQueue[0]
		c.dogQueue = c.dogQueue[1:]
	} else {
		return p, errors.New("the dogQueue is empty")
	}
	return p, nil
}

//题目二：转圈打印矩阵